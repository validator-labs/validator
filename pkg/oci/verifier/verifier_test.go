package verifier

import (
	"context"
	"errors"
	"net/http"
	"reflect"
	"testing"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/v1/remote"
)

func TestOptions(t *testing.T) {
	tests := []struct {
		name string
		opts []Options
		want *options
	}{{
		name: "no options",
		want: &options{},
	}, {
		name: "signature option",
		opts: []Options{WithPublicKey([]byte("foo"))},
		want: &options{
			PublicKey: []byte("foo"),
			ROpt:      nil,
		},
	}, {
		name: "keychain option",
		opts: []Options{WithRemoteOptions(remote.WithAuthFromKeychain(authn.DefaultKeychain))},
		want: &options{
			PublicKey: nil,
			ROpt:      []remote.Option{remote.WithAuthFromKeychain(authn.DefaultKeychain)},
		},
	}, {
		name: "keychain and authenticator option",
		opts: []Options{WithRemoteOptions(
			remote.WithAuth(&authn.Basic{Username: "foo", Password: "bar"}),
			remote.WithAuthFromKeychain(authn.DefaultKeychain),
		)},
		want: &options{
			PublicKey: nil,
			ROpt: []remote.Option{
				remote.WithAuth(&authn.Basic{Username: "foo", Password: "bar"}),
				remote.WithAuthFromKeychain(authn.DefaultKeychain),
			},
		},
	}, {
		name: "keychain, authenticator and transport option",
		opts: []Options{WithRemoteOptions(
			remote.WithAuth(&authn.Basic{Username: "foo", Password: "bar"}),
			remote.WithAuthFromKeychain(authn.DefaultKeychain),
			remote.WithTransport(http.DefaultTransport),
		)},
		want: &options{
			PublicKey: nil,
			ROpt: []remote.Option{
				remote.WithAuth(&authn.Basic{Username: "foo", Password: "bar"}),
				remote.WithAuthFromKeychain(authn.DefaultKeychain),
				remote.WithTransport(http.DefaultTransport),
			},
		},
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			o := options{}
			for _, opt := range test.opts {
				opt(&o)
			}
			if !reflect.DeepEqual(o.PublicKey, test.want.PublicKey) {
				t.Errorf("got %#v, want %#v", &o.PublicKey, test.want.PublicKey)
			}

			if test.want.ROpt != nil {
				if len(o.ROpt) != len(test.want.ROpt) {
					t.Errorf("got %d remote options, want %d", len(o.ROpt), len(test.want.ROpt))
				}
				return
			}

			if test.want.ROpt == nil {
				if len(o.ROpt) != 0 {
					t.Errorf("got %d remote options, want %d", len(o.ROpt), 0)
				}
			}
		})
	}
}

func TestNewCosignVerifier(t *testing.T) {
	testCases := []struct {
		name        string
		publicKey   []byte
		remoteOpts  []remote.Option
		expectError error
	}{
		{
			name: "Successful initialization with public key and remote options",
			publicKey: []byte(`-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEKPuCo9AmJCpqGWhefjbhkFcr1GA3
iNa765seE3jYC3MGUe5h52393Dhy7B5bXGsg6EfPpNYamlAEWjxCpHF3Lg==
-----END PUBLIC KEY-----`),
			remoteOpts:  []remote.Option{remote.WithAuthFromKeychain(authn.DefaultKeychain)},
			expectError: nil,
		},
		{
			name:        "Failure due to missing public key",
			publicKey:   nil,
			remoteOpts:  nil,
			expectError: errors.New("public key required for cosign verifier"),
		},
		{
			name:        "Failure due to invalid public key",
			publicKey:   []byte("invalid public key"),
			remoteOpts:  nil,
			expectError: errors.New("PEM decoding failed"),
		},
	}

	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewCosignVerifier(ctx, WithPublicKey(tc.publicKey), WithRemoteOptions(tc.remoteOpts...))

			if err == nil && tc.expectError != nil {
				t.Errorf("Expected error for %s, got nil", tc.name)
			} else if err != nil && tc.expectError == nil {
				t.Errorf("Unexpected error for %s: %v", tc.name, err)
			} else if err != nil && tc.expectError != nil && err.Error() != tc.expectError.Error() {
				t.Errorf("Expected error '%v', got '%v'", tc.expectError, err)
			}
		})
	}
}
