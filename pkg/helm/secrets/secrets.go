/*
Copyright The Helm Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package secrets includes a client for fetching Helm releases from Kubernetes secrets.
package secrets

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"

	corev1 "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	klabels "k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	klog "k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var magicGzip = []byte{0x1f, 0x8b, 0x08}

// Client is a subset of the kubernetes SecretsInterface, geared towards handling Helm secrets.
type Client interface {
	List(context.Context, klabels.Selector, string) ([]*Release, error)
	Get(context.Context, string, string) (*Release, error)
}

// HelmSecretsClient implements the Client interface.
type HelmSecretsClient struct {
	kubeClient client.Client
}

// NewSecretsClient initializes a new HelmSecretsClient.
func NewSecretsClient(client client.Client) *HelmSecretsClient {
	return &HelmSecretsClient{
		kubeClient: client,
	}
}

// List fetches all Helm releases in a namespace, filtered by a label selector.
func (s *HelmSecretsClient) List(ctx context.Context, labels klabels.Selector, namespace string) ([]*Release, error) {
	// ensure the label selector includes the 'owner: helm' label
	req, err := klabels.NewRequirement("owner", selection.Equals, []string{"helm"})
	if err != nil {
		return nil, err
	}
	if labels == nil {
		labels = klabels.Everything()
	}
	labels = labels.Add(*req)

	// list the Helm secrets
	list := &corev1.SecretList{}
	if err := s.kubeClient.List(ctx, list, &client.ListOptions{
		LabelSelector: labels,
		Namespace:     namespace,
	}); err != nil {
		return nil, err
	}

	// iterate over the Helm secrets and decode each release
	releases := make([]*Release, 0, len(list.Items))
	for _, item := range list.Items {
		cpy := item
		rls, err := decodeRelease(&cpy, string(item.Data["release"]))
		if err != nil {
			klog.Errorf("failed to decode release: %v", err)
			continue
		} else if rls.Chart == nil || rls.Chart.Metadata == nil || rls.Info == nil {
			klog.Warningf("skipping release with empty metadata: %s", rls.Name)
			continue
		}
		releases = append(releases, rls)
	}
	return releases, nil
}

// Get fetches the latest Helm release by name and namespace.
func (s *HelmSecretsClient) Get(ctx context.Context, name string, namespace string) (*Release, error) {
	ls := klabels.Set{}
	ls["name"] = name
	releaseList, err := s.List(ctx, ls.AsSelector(), namespace)
	if err != nil {
		return nil, err
	} else if len(releaseList) == 0 {
		return nil, kerrors.NewNotFound(corev1.Resource("Secret"), name)
	}

	var latest *Release
	for _, rls := range releaseList {
		if latest == nil || latest.Version < rls.Version {
			latest = rls
		}
	}
	return latest, nil
}

// decodeRelease decodes secret data into a Helm release.
func decodeRelease(secret *corev1.Secret, data string) (*Release, error) {
	b, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	} else if len(b) < 3 {
		return nil, fmt.Errorf("unexpected secret content: %s", data)
	}

	// For backwards compatibility with releases that were stored before compression
	// was introduced we skip decompression if the gzip magic header is not found
	if bytes.Equal(b[0:3], magicGzip) {
		r, err := gzip.NewReader(bytes.NewReader(b))
		if err != nil {
			return nil, err
		}
		b2, err := io.ReadAll(r)
		if err != nil {
			return nil, err
		}
		b = b2
	}

	var rls Release
	if err := json.Unmarshal(b, &rls); err != nil {
		return nil, fmt.Errorf("error decoding Helm release %s: %v", string(b), err)
	}

	rls.Secret = secret
	return &rls, nil
}
