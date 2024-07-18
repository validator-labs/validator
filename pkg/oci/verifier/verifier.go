// Package verifier contains the verifier interface and the cosign verifier implementation.
package verifier

import (
	"context"
	"crypto"
	"fmt"

	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	coptions "github.com/sigstore/cosign/v2/cmd/cosign/cli/options"
	"github.com/sigstore/cosign/v2/pkg/cosign"
	ociremote "github.com/sigstore/cosign/v2/pkg/oci/remote"
	"github.com/sigstore/sigstore/pkg/cryptoutils"
	"github.com/sigstore/sigstore/pkg/signature"
)

// Verifier is an interface for verifying the authenticity of an OCI image.
type Verifier interface {
	Verify(ctx context.Context, ref name.Reference) (bool, error)
}

// options is a struct that holds options for verifier.
type options struct {
	PublicKey []byte
	ROpt      []remote.Option
}

// Options is a function that configures the options applied to a Verifier.
type Options func(opts *options)

// WithPublicKey sets the public key.
func WithPublicKey(publicKey []byte) Options {
	return func(opts *options) {
		opts.PublicKey = publicKey
	}
}

// WithRemoteOptions is a functional option for overriding the default
// remote options used by the verifier.
func WithRemoteOptions(opts ...remote.Option) Options {
	return func(o *options) {
		o.ROpt = opts
	}
}

// CosignVerifier is a struct which is responsible for executing verification logic.
type CosignVerifier struct {
	opts *cosign.CheckOpts
}

// NewCosignVerifier initializes a new CosignVerifier.
func NewCosignVerifier(ctx context.Context, opts ...Options) (*CosignVerifier, error) {
	o := options{}
	for _, opt := range opts {
		opt(&o)
	}

	if len(o.PublicKey) == 0 {
		return nil, fmt.Errorf("public key required for cosign verifier")
	}

	checkOpts := &cosign.CheckOpts{}

	ro := coptions.RegistryOptions{}
	co, err := ro.ClientOpts(ctx)
	if err != nil {
		return nil, err
	}

	if o.ROpt != nil {
		co = append(co, ociremote.WithRemoteOptions(o.ROpt...))
	}

	checkOpts.RegistryClientOpts = co

	checkOpts.Offline = true
	// TODO(hidde): this is an oversight in our implementation. As it is
	//  theoretically possible to have a custom PK, without disabling tlog.
	checkOpts.IgnoreTlog = true

	pubKeyRaw, err := cryptoutils.UnmarshalPEMToPublicKey(o.PublicKey)
	if err != nil {
		return nil, err
	}

	checkOpts.SigVerifier, err = signature.LoadVerifier(pubKeyRaw, crypto.SHA256)
	if err != nil {
		return nil, err
	}

	return &CosignVerifier{
		opts: checkOpts,
	}, nil
}

// Verify verifies the authenticity of the given ref OCI image.
// It returns a boolean indicating if the verification was successful.
// It returns an error if the verification fails, nil otherwise.
func (v *CosignVerifier) Verify(ctx context.Context, ref name.Reference) (bool, error) {
	signatures, _, err := cosign.VerifyImageSignatures(ctx, ref, v.opts)
	if err != nil {
		return false, err
	}

	if len(signatures) == 0 {
		return false, nil
	}

	return true, nil
}
