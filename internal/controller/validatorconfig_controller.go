/*
Copyright 2023.

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

package controller

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	connect "connectrpc.com/connect"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"golang.org/x/exp/slices"
	corev1 "k8s.io/api/core/v1"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/cluster-api/util/patch"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"buf.build/gen/go/spectrocloud/spectro-cleanup/connectrpc/go/cleanup/v1/cleanupv1connect"
	cleanv1 "buf.build/gen/go/spectrocloud/spectro-cleanup/protocolbuffers/go/cleanup/v1"
	ociauth "github.com/validator-labs/validator-plugin-oci/pkg/auth"
	ocic "github.com/validator-labs/validator-plugin-oci/pkg/ociclient"

	v1alpha1 "github.com/validator-labs/validator/api/v1alpha1"
	"github.com/validator-labs/validator/pkg/helm"
	helmrelease "github.com/validator-labs/validator/pkg/helm/release"
)

const (
	// CleanupFinalizer ensures that plugin Helm releases are properly garbage collected.
	CleanupFinalizer = "validator/cleanup"

	// PluginValuesHash is an annotation key added to a ValidatorConfig to determine whether to update a plugin's Helm release.
	PluginValuesHash = "validator/plugin-values"
)

// ValidatorConfigReconciler reconciles a ValidatorConfig object
type ValidatorConfigReconciler struct {
	client.Client
	HelmClient        helm.Client
	HelmReleaseClient helmrelease.Client
	Log               logr.Logger
	Scheme            *runtime.Scheme
}

// +kubebuilder:rbac:groups=validation.spectrocloud.labs,resources=validatorconfigs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=validation.spectrocloud.labs,resources=validatorconfigs/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=validation.spectrocloud.labs,resources=validatorconfigs/finalizers,verbs=update

// Reconcile reconciles a ValidatorConfig.
func (r *ValidatorConfigReconciler) Reconcile(ctx context.Context, req ctrl.Request) (_ ctrl.Result, reterr error) {
	r.Log.V(0).Info("Reconciling ValidatorConfig", "name", req.Name, "namespace", req.Namespace)

	vc := &v1alpha1.ValidatorConfig{}

	if err := r.Get(ctx, req.NamespacedName, vc); err != nil {
		if !apierrs.IsNotFound(err) {
			r.Log.Error(err, "failed to fetch ValidatorConfig", "key", req)
		}
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	patcher, err := patch.NewHelper(vc, r.Client)
	if err != nil {
		return ctrl.Result{}, errors.Wrapf(err, "failed to create patch helper for ValidatorConfig %s", vc.Name)
	}

	if vc.Annotations == nil {
		vc.Annotations = make(map[string]string)
	}

	// handle ValidatorConfig deletion
	if vc.DeletionTimestamp != nil {
		// if namespace is deleting, remove finalizer & the rest will follow
		namespace := &corev1.Namespace{}
		err := r.Client.Get(ctx, types.NamespacedName{Name: req.Namespace}, namespace)
		if err != nil {
			return ctrl.Result{}, nil
		} else if namespace.DeletionTimestamp != nil {
			return ctrl.Result{}, removeFinalizer(ctx, r.Client, vc, CleanupFinalizer)
		}

		// otherwise, just delete the plugins
		if err := r.deletePlugins(ctx, vc); err != nil {
			return ctrl.Result{}, err
		}

		err = removeFinalizer(ctx, r.Client, vc, CleanupFinalizer)

		if emitErr := r.emitFinalizeCleanup(ctx); emitErr != nil {
			r.Log.Error(emitErr, "Failed to emit FinalizeCleanup request")
		}

		return ctrl.Result{}, err
	}

	defer func() {
		r.Log.V(1).Info("Preparing to patch ValidatorConfig", "validatorConfig", vc.Name)
		if err := patchValidatorConfig(ctx, patcher, vc); err != nil && reterr == nil {
			reterr = err
			r.Log.Error(err, "failed to patch ValidatorConfig", "validatorConfig", vc.Name)
			return
		}
		r.Log.V(1).Info("Successfully patched ValidatorConfig", "validatorConfig", vc.Name)
	}()

	// deploy/redeploy plugins as required
	r.redeployIfNeeded(ctx, vc)

	// ensure cleanup finalizer
	ensureFinalizer(vc, CleanupFinalizer)
	r.Log.V(0).Info("Ensured ValidatorConfig finalizer")

	return ctrl.Result{RequeueAfter: time.Second * 30}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ValidatorConfigReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.ValidatorConfig{}).
		Complete(r)
}

// redeployIfNeeded deploys/redeploys each validator plugin in a ValidatorConfig and deletes plugins that have been removed
func (r *ValidatorConfigReconciler) redeployIfNeeded(ctx context.Context, vc *v1alpha1.ValidatorConfig) {
	specPlugins := make(map[string]bool)
	conditions := make([]v1alpha1.ValidatorPluginCondition, len(vc.Spec.Plugins))

	helmConfig := vc.Spec.HelmConfig
	for i, p := range vc.Spec.Plugins {
		specPlugins[p.Chart.Name] = true

		// update plugin's values hash
		valuesUnchanged := r.updatePluginHash(vc, p)

		// skip plugin if already deployed & no change in values
		condition, ok := isConditionTrue(vc, p.Chart.Name, v1alpha1.HelmChartDeployedCondition)
		if ok && valuesUnchanged {
			r.Log.V(0).Info("Values unchanged. Skipping upgrade for plugin Helm chart", "namespace", vc.Namespace, "name", p.Chart.Name)
			conditions[i] = condition
			continue
		}

		opts := &helm.Options{
			Chart:                 p.Chart.Name,
			Repo:                  p.Chart.Repository,
			Registry:              helmConfig.Registry,
			Version:               p.Chart.Version,
			Values:                p.Values,
			InsecureSkipTLSVerify: helmConfig.InsecureSkipTLSVerify,
		}

		if helmConfig.AuthSecretName != "" {
			nn := types.NamespacedName{Name: helmConfig.AuthSecretName, Namespace: vc.Namespace}
			if err := r.configureHelmOpts(ctx, nn, opts); err != nil {
				r.Log.V(0).Error(err, "failed to configure basic auth for Helm upgrade")
				conditions[i] = r.buildHelmChartCondition(p.Chart.Name, err)
				continue
			}
		}

		var cleanupLocalChart bool
		if strings.HasPrefix(helmConfig.Registry, ocic.Scheme) {
			r.Log.V(0).Info("Pulling plugin Helm chart", "name", p.Chart.Name)

			opts.Path = fmt.Sprintf("/charts/%s", opts.Chart)
			opts.Version = strings.TrimPrefix(opts.Version, "v")

			// use OCI client instead of Helm client due to https://github.com/helm/helm/issues/12810
			ociClient, err := ocic.NewOCIClient(
				ocic.WithBasicAuth(opts.Username, opts.Password),
				ocic.WithMultiAuth(ociauth.GetKeychain(opts.Registry)),
				ocic.WithTLSConfig(opts.InsecureSkipTLSVerify, "", opts.CaFile),
			)
			if err != nil {
				r.Log.V(0).Error(err, "failed to create OCI client")
				conditions[i] = r.buildHelmChartCondition(p.Chart.Name, err)
				continue
			}
			ociOpts := ocic.ImageOptions{
				Ref:     fmt.Sprintf("%s/%s:%s", strings.TrimPrefix(opts.Registry, ocic.Scheme), opts.Repo, opts.Version),
				OutDir:  opts.Path,
				OutFile: opts.Chart,
			}
			if err := ociClient.PullChart(ociOpts); err != nil {
				r.Log.V(0).Error(err, "failed to pull Helm chart from OCI registry")
				conditions[i] = r.buildHelmChartCondition(p.Chart.Name, err)
				continue
			}

			r.Log.V(0).Info("Reconfiguring Helm options to deploy local chart", "name", p.Chart.Name)
			opts.Path = fmt.Sprintf("%s/%s.tgz", opts.Path, opts.Chart)
			opts.Chart = ""
			cleanupLocalChart = true
		}

		r.Log.V(0).Info("Installing/upgrading plugin Helm chart", "namespace", vc.Namespace, "name", p.Chart.Name)
		err := r.HelmClient.Upgrade(p.Chart.Name, vc.Namespace, *opts)
		if err != nil {
			// if Helm install/upgrade failed, delete the release so installation is reattempted each iteration
			if strings.Contains(err.Error(), "has no deployed releases") {
				if err := r.HelmClient.Delete(p.Chart.Name, vc.Namespace); err != nil {
					r.Log.V(0).Error(err, "failed to delete Helm release")
				}
			}
		}
		conditions[i] = r.buildHelmChartCondition(p.Chart.Name, err)

		if cleanupLocalChart {
			r.Log.V(0).Info("Cleaning up local chart directory", "path", opts.Path)
			if err := os.RemoveAll(opts.Path); err != nil {
				r.Log.V(0).Error(err, "failed to remove local chart directory")
			}
		}
	}

	// delete any plugins that have been removed
	for _, c := range vc.Status.Conditions {
		_, ok := specPlugins[c.PluginName]
		if !ok && c.Type == v1alpha1.HelmChartDeployedCondition && c.Status == corev1.ConditionTrue {
			r.Log.V(0).Info("Deleting plugin Helm chart", "namespace", vc.Namespace, "name", c.PluginName)
			r.deletePlugin(vc, c.PluginName)
			delete(vc.Annotations, getPluginHashKey(c.PluginName))
		}
	}

	// update plugin conditions
	vc.Status.Conditions = conditions
}

func (r *ValidatorConfigReconciler) configureHelmOpts(ctx context.Context, nn types.NamespacedName, opts *helm.Options) error {
	secret := &corev1.Secret{}
	if err := r.Get(ctx, nn, secret); err != nil {
		return fmt.Errorf(
			"failed to get auth secret %s in namespace %s for chart %s in repo %s: %v",
			nn.Name, nn.Namespace, opts.Chart, opts.Repo, err,
		)
	}

	username, ok := secret.Data["username"]
	if !ok {
		return fmt.Errorf("auth secret for chart %s in repo %s missing required key: 'username'", opts.Chart, opts.Repo)
	}
	opts.Username = string(username)

	password, ok := secret.Data["password"]
	if !ok {
		return fmt.Errorf("auth secret for chart %s in repo %s missing required key: 'password'", opts.Chart, opts.Repo)
	}
	opts.Password = string(password)

	caCert, ok := secret.Data["caCert"]
	if ok {
		caFile := fmt.Sprintf("/etc/ssl/certs/%s-ca.crt", opts.Chart)
		if err := os.WriteFile(caFile, caCert, 0600); err != nil {
			return errors.Wrap(err, "failed to write Helm CA file")
		}
		opts.CaFile = caFile
	}

	return nil
}

// updatePluginHash compares the current plugin's values hash annotation to a hash of its current values,
// updates the values hash annotation on the ValidatorConfig for the current plugin, and returns a flag
// indicating whether the values have changed or not since the last reconciliation
func (r *ValidatorConfigReconciler) updatePluginHash(vc *v1alpha1.ValidatorConfig, p v1alpha1.HelmRelease) bool {
	valuesUnchanged := false
	pluginValuesHashLatest := sha256.Sum256([]byte(p.Values))
	pluginValuesHashLatestB64 := base64.StdEncoding.EncodeToString(pluginValuesHashLatest[:])
	key := getPluginHashKey(p.Chart.Name)

	pluginValuesHash, ok := vc.Annotations[key]
	if ok {
		valuesUnchanged = pluginValuesHash == pluginValuesHashLatestB64
	}
	vc.Annotations[key] = pluginValuesHashLatestB64

	return valuesUnchanged
}

// getPluginHashKey generates an annotation key used to retrieve a plugin's values hash
func getPluginHashKey(pluginName string) string {
	return fmt.Sprintf("%s-%s", PluginValuesHash, pluginName)
}

// deletePlugins deletes each validator plugin's Helm release
func (r *ValidatorConfigReconciler) deletePlugins(ctx context.Context, vc *v1alpha1.ValidatorConfig) error {
	var wg sync.WaitGroup
	for _, p := range vc.Spec.Plugins {
		release, err := r.HelmReleaseClient.Get(ctx, p.Chart.Name, vc.Namespace)
		if err != nil {
			if !apierrs.IsNotFound(err) {
				return err
			}
			return nil
		}
		if release.Secret.Labels == nil || release.Secret.Labels["owner"] != "helm" {
			return nil
		}

		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			r.deletePlugin(vc, name)
		}(p.Chart.Name)
	}

	wg.Wait()
	return nil
}

// deletePlugin deletes the Helm release associated with a ValidatorConfig plugin
func (r *ValidatorConfigReconciler) deletePlugin(vc *v1alpha1.ValidatorConfig, name string) {
	if err := r.HelmClient.Delete(name, vc.Namespace); err != nil {
		r.Log.V(0).Error(err, "failed to delete validator plugin", "namespace", vc.Namespace, "name", name)
	}
	r.Log.V(0).Info("Deleted Helm release for validator plugin", "namespace", vc.Namespace, "name", name)
}

// buildHelmChartCondition builds a ValidatorPluginCondition for a plugin
func (r *ValidatorConfigReconciler) buildHelmChartCondition(chartName string, err error) v1alpha1.ValidatorPluginCondition {
	c := v1alpha1.ValidatorPluginCondition{
		Type:               v1alpha1.HelmChartDeployedCondition,
		PluginName:         chartName,
		Status:             corev1.ConditionTrue,
		Message:            fmt.Sprintf("Plugin chart %s is installed/upgraded", chartName),
		LastTransitionTime: metav1.Time{Time: time.Now()},
	}
	if err != nil {
		c.Status = corev1.ConditionFalse
		c.Message = err.Error()
	}
	r.Log.V(0).Info("Latest ValidatorConfig plugin condition", "name", c.PluginName, "type", c.Type, "status", c.Status, "message", c.Message)
	return c
}

// ensureFinalizer ensures that an object's finalizers include a certain finalizer
func ensureFinalizer(obj client.Object, finalizer string) {
	currentFinalizers := obj.GetFinalizers()
	if !slices.Contains(currentFinalizers, finalizer) {
		newFinalizers := []string{}
		newFinalizers = append(newFinalizers, currentFinalizers...)
		newFinalizers = append(newFinalizers, finalizer)
		obj.SetFinalizers(newFinalizers)
	}
}

// removeFinalizer removes a finalizer from an object's finalizer's (if found)
func removeFinalizer(ctx context.Context, client client.Client, obj client.Object, finalizer string) error {
	finalizers := obj.GetFinalizers()
	if len(finalizers) > 0 {
		newFinalizers := []string{}
		for _, f := range finalizers {
			if f == finalizer {
				continue
			}
			newFinalizers = append(newFinalizers, f)
		}
		if len(newFinalizers) != len(finalizers) {
			obj.SetFinalizers(newFinalizers)
			if err := client.Update(ctx, obj); err != nil {
				return err
			}
		}
	}
	return nil
}

// conditionIndex retrieves the index of a ValidatorPluginCondition from a ValidatorConfig's status
func conditionIndex(vc *v1alpha1.ValidatorConfig, chartName string, conditionType v1alpha1.ConditionType) int {
	for i, c := range vc.Status.Conditions {
		if c.Type == conditionType && c.PluginName == chartName {
			return i
		}
	}
	return -1
}

// isConditionTrue checks whether a ValidatorPluginCondition is true
func isConditionTrue(vc *v1alpha1.ValidatorConfig, chartName string, conditionType v1alpha1.ConditionType) (v1alpha1.ValidatorPluginCondition, bool) {
	idx := conditionIndex(vc, chartName, conditionType)
	if idx == -1 {
		return v1alpha1.ValidatorPluginCondition{}, false
	}
	return vc.Status.Conditions[idx], vc.Status.Conditions[idx].Status == corev1.ConditionTrue
}

func (r *ValidatorConfigReconciler) emitFinalizeCleanup(ctx context.Context) error {
	grpcEnabled := os.Getenv("CLEANUP_GRPC_SERVER_ENABLED")
	if grpcEnabled != "true" {
		r.Log.V(0).Info("Cleanup job gRPC server is not enabled")
		return nil
	}

	host := os.Getenv("CLEANUP_GRPC_SERVER_HOST")
	if host == "" {
		return errors.New("CLEANUP_GRPC_SERVER_HOST is invalid")
	}

	port := os.Getenv("CLEANUP_GRPC_SERVER_PORT")
	_, err := strconv.Atoi(port)
	if err != nil {
		return fmt.Errorf("CLEANUP_GRPC_SERVER_PORT is invalid: %w", err)
	}

	url := fmt.Sprintf("http://%s:%s", host, port)
	client := cleanupv1connect.NewCleanupServiceClient(
		http.DefaultClient,
		url,
	)
	_, err = client.FinalizeCleanup(ctx, connect.NewRequest(&cleanv1.FinalizeCleanupRequest{}))
	if err != nil {
		return fmt.Errorf("FinalizeCleanup request to %s failed: %w", url, err)
	}
	return nil
}

// patchValidatorConfig patches a ValidatorConfig.
func patchValidatorConfig(ctx context.Context, patchHelper *patch.Helper, vc *v1alpha1.ValidatorConfig) error {
	return patchHelper.Patch(ctx, vc, patch.WithStatusObservedGeneration{})
}
