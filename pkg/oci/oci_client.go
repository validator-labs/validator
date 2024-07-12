// Package oci contains the OCI client interface and implementation.
package oci

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/awslabs/amazon-ecr-credential-helper/ecr-login"
	"github.com/awslabs/amazon-ecr-credential-helper/ecr-login/api"
	acr "github.com/chrismellard/docker-credential-acr-env/pkg/credhelper"
	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/google"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	klog "k8s.io/klog/v2"
)

// Scheme is the URI scheme for an OCI registry.
const Scheme = "oci://"

// Client is an interface for interacting with an OCI registry.
type Client struct {
	auth authn.Keychain
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
func NewOCIClient(opts ...Option) *Client {
	c := &Client{}
	for _, o := range opts {
		o(c)
	}
	return c
}

// WithMultiAuth configures the OCI client with multiple authentication keychains.
func WithMultiAuth() Option {
	return func(c *Client) {
		c.auth = authn.NewMultiKeychain(
			authn.DefaultKeychain,
			google.Keychain,
			authn.NewKeychainFromHelper(ecr.NewECRHelper(ecr.WithClientFactory(api.DefaultClientFactory{}))),
			authn.NewKeychainFromHelper(acr.ACRCredHelper{}),
		)
	}
}

// PullChart pulls a Helm chart from the given ImageOptions.
func (c Client) PullChart(opts ImageOptions) error {
	ref, err := name.ParseReference(opts.Ref)
	if err != nil {
		return fmt.Errorf("failed to parse chart reference: %w", err)
	}
	// Assume the chart is in the first layer & extract it
	layers, err := c.PullImage(ref)
	if err != nil {
		return fmt.Errorf("failed to pull chart: %w", err)
	}
	return c.WriteLayer(opts, layers[0])
}

// PullImage pulls an image from the given name.Reference.
func (c Client) PullImage(ref name.Reference) ([]v1.Layer, error) {
	img, err := remote.Image(ref, remote.WithAuthFromKeychain(c.auth))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch image from registry: %w", err)
	}
	return img.Layers()
}

// WriteLayer writes a layer to the filesystem.
func (c Client) WriteLayer(opts ImageOptions, layer v1.Layer) error {
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

	path := filepath.Join(opts.OutDir, fmt.Sprintf("%s.tgz", opts.OutFile))

	if err := os.WriteFile(path, content, os.ModePerm); err != nil {
		return fmt.Errorf("failed to write layer file: %w", err)
	}

	klog.Infof("Layer saved successfully to %s\n", path)
	return nil
}
