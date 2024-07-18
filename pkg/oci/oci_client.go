// Package oci contains the OCI client interface and implementation.
package oci

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/awslabs/amazon-ecr-credential-helper/ecr-login"
	"github.com/awslabs/amazon-ecr-credential-helper/ecr-login/api"
	acr "github.com/chrismellard/docker-credential-acr-env/pkg/credhelper"
	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/google"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"github.com/google/go-containerregistry/pkg/v1/validate"
	"github.com/validator-labs/validator/pkg/util"
	klog "k8s.io/klog/v2"

	"github.com/validator-labs/validator/pkg/oci/verifier"
)

const (
	// Scheme is the URI scheme for an OCI registry.
	Scheme = "oci://"

	// VerificationTimeout is the timeout for verifying the authenticity of an image.
	VerificationTimeout = 60 * time.Second
)

var (
	// ErrNoPublicKeys is the error message for when no public keys are provided for signature verification.
	ErrNoPublicKeys = errors.New("no public keys provided for signature verification")
)

// Client is an interface for interacting with an OCI registry.
type Client struct {
	authenticator          authn.Authenticator
	keychain               authn.Keychain
	caCert                 string
	caFile                 string
	insecureSkipTLSVerify  bool
	opts                   []remote.Option
	verificationTimeout    time.Duration
	verificationPublicKeys [][]byte
}

// ImageOptions defines the options for pulling an image.
type ImageOptions struct {
	Ref     string
	OutDir  string
	OutFile string
}

// Option is a functional option for configuring the OCI client.
type Option func(*Client)

// NewOCIClient creates a new OCI client with the given options.
func NewOCIClient(opts ...Option) (*Client, error) {
	c := &Client{}
	for _, o := range opts {
		o(c)
	}
	c.setAuth()
	if err := c.setTransport(); err != nil {
		return nil, err
	}
	return c, nil
}

// WithTLSConfig configures the OCI client with the given TLS options.
func WithTLSConfig(insecureSkipTLSVerify bool, caCert, caFile string) Option {
	return func(c *Client) {
		c.caCert = caCert
		c.caFile = caFile
		c.insecureSkipTLSVerify = insecureSkipTLSVerify
	}
}

// WithAnonymousAuth configures the OCI client with anonymous authentication.
func WithAnonymousAuth() Option {
	return func(c *Client) {
		c.authenticator = authn.Anonymous
	}
}

// WithBasicAuth configures the OCI client with basic authentication.
func WithBasicAuth(username, password string) Option {
	return func(c *Client) {
		if username != "" && password != "" {
			c.authenticator = &authn.Basic{Username: username, Password: password}
		}
	}
}

// WithMultiAuth configures the OCI client with multiple authentication keychains.
func WithMultiAuth() Option {
	return func(c *Client) {
		c.keychain = authn.NewMultiKeychain(
			authn.DefaultKeychain,
			google.Keychain,
			authn.NewKeychainFromHelper(ecr.NewECRHelper(ecr.WithClientFactory(api.DefaultClientFactory{}))),
			authn.NewKeychainFromHelper(acr.ACRCredHelper{}),
		)
	}
}

// WithVerificationPublicKeys configures the OCI client with the given public keys for signature verification.
func WithVerificationPublicKeys(publicKeys [][]byte) Option {
	return func(c *Client) {
		c.verificationPublicKeys = publicKeys
	}
}

// WithVerificationTimeout configures the OCI client with the given verification timeout.
func WithVerificationTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.verificationTimeout = timeout
	}
}

// Catalog returns the list of repositories in the registry.
func (c *Client) Catalog(ctx context.Context, reg name.Registry) ([]string, error) {
	return remote.Catalog(ctx, reg, c.opts...)
}

// List returns a list of tags for the given repository.
func (c *Client) List(ref name.Repository) ([]string, error) {
	return remote.List(ref, c.opts...)
}

// Head checks if the given artifact exists in the registry.
func (c *Client) Head(ref name.Reference) (*v1.Descriptor, error) {
	return remote.Head(ref, c.opts...)
}

// PullChart pulls a Helm chart from the given ImageOptions.
func (c *Client) PullChart(opts ImageOptions) error {
	ref, err := name.ParseReference(opts.Ref)
	if err != nil {
		return fmt.Errorf("failed to parse chart reference: %w", err)
	}

	path := filepath.Join(opts.OutDir, opts.OutFile)

	// Assume the chart is in the first layer & extract it
	img, err := c.PullImage(ref)
	if err != nil {
		return fmt.Errorf("failed to pull chart: %w", err)
	}
	layers, err := img.Layers()
	if err != nil {
		return fmt.Errorf("failed to get image layers: %w", err)
	}
	if err := c.WriteLayer(layers[0], path, opts); err != nil {
		return fmt.Errorf("failed to write chart layer: %w", err)
	}

	return util.Gzip(path, fmt.Sprintf("%s.tgz", path))
}

// PullImage pulls an image from the given name.Reference.
func (c *Client) PullImage(ref name.Reference) (v1.Image, error) {
	img, err := remote.Image(ref, c.opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch image from registry: %w", err)
	}
	return img, nil
}

// ValidateImage validates the given image.
func (c *Client) ValidateImage(image v1.Image, fullLayerValidation bool) error {
	var validateOpts []validate.Option
	if !fullLayerValidation {
		validateOpts = append(validateOpts, validate.Fast)
	}
	return validate.Image(image, validateOpts...)
}

// WriteLayer writes a layer to the filesystem.
func (c *Client) WriteLayer(layer v1.Layer, path string, opts ImageOptions) error {
	r, err := layer.Uncompressed()
	if err != nil {
		return fmt.Errorf("failed to uncompress layer: %w", err)
	}
	defer func() {
		closeErr := r.Close()
		if err == nil {
			err = closeErr
		} else {
			klog.Errorf("failed to close layer reader: %v", closeErr)
		}
	}()

	content, err := io.ReadAll(r)
	if err != nil {
		return fmt.Errorf("failed to read layer content: %w", err)
	}
	if err := os.MkdirAll(opts.OutDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}
	if err := os.WriteFile(path, content, os.ModePerm); err != nil {
		return fmt.Errorf("failed to write layer file: %w", err)
	}

	klog.Infof("Layer saved successfully to %s\n", path)
	return nil
}

// VerifySignature verifies the authenticity of the given image reference URL using the provided public keys.
func (c *Client) VerifySignature(ctx context.Context, ref name.Reference) ([]string, []error) {
	details := make([]string, 0)
	errs := make([]error, 0)

	if len(c.verificationPublicKeys) == 0 {
		errs = append(errs, ErrNoPublicKeys)
		return details, errs
	}

	ctxTimeout, cancel := context.WithTimeout(ctx, c.verificationTimeout)
	defer cancel()

	defaultCosignOciOpts := []verifier.Options{
		verifier.WithRemoteOptions(c.opts...),
	}

	for _, key := range c.verificationPublicKeys {
		verifier, err := verifier.NewCosignVerifier(ctxTimeout, append(defaultCosignOciOpts, verifier.WithPublicKey(key))...)
		if err != nil {
			details = append(details, fmt.Sprintf("failed to create verifier with public key %s", key))
			errs = append(errs, err)
			return details, errs
		}

		hasValidSignature, err := verifier.Verify(ctxTimeout, ref)
		if err != nil {
			details = append(details, fmt.Sprintf("failed to verify signature of %s with public key %s", ref, key))
			errs = append(errs, err)
			continue
		}

		if hasValidSignature {
			details = nil
			errs = nil
			return details, errs
		}
	}

	details = append(details, fmt.Sprintf("no matching signatures were found for '%s'", ref))
	errs = append(errs, fmt.Errorf("failed to verify signature for '%s'", ref))
	return details, errs
}

// setAuth configures a remote.Option for authenticating with the OCI registry.
// If an authenticator was configured, it is used; otherwise the default keychains are used.
func (c *Client) setAuth() {
	if c.authenticator != nil {
		c.opts = append(c.opts, remote.WithAuth(c.authenticator))
		return
	}
	c.opts = append(c.opts, remote.WithAuthFromKeychain(c.keychain))
}

// setTransport configures the HTTP transport for the OCI client.
func (c *Client) setTransport() error {
	// Determine CA pool
	caCertPool, err := x509.SystemCertPool()
	if err != nil {
		return fmt.Errorf("failed to load system cert pool: %w", err)
	}
	if c.caCert != "" {
		if ok := caCertPool.AppendCertsFromPEM([]byte(c.caCert)); !ok {
			return fmt.Errorf("failed to append raw CA cert")
		}
	}
	if c.caFile != "" {
		bs, err := os.ReadFile(c.caFile)
		if err != nil {
			return fmt.Errorf("failed to read CA cert file: %w", err)
		}
		if ok := caCertPool.AppendCertsFromPEM(bs); !ok {
			return fmt.Errorf("failed to append CA cert from %s", c.caFile)
		}
	}

	// Configure transport
	transport := remote.DefaultTransport.(*http.Transport)
	transport.Proxy = http.ProxyFromEnvironment
	transport.TLSClientConfig = &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    caCertPool,
	}
	if c.insecureSkipTLSVerify {
		transport.TLSClientConfig.InsecureSkipVerify = true
	}

	c.opts = append(c.opts, remote.WithTransport(transport))
	return nil
}
