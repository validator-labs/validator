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
	"strings"
	"time"

	"github.com/go-logr/logr"
	"golang.org/x/exp/slices"
	corev1 "k8s.io/api/core/v1"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	v1alpha1 "github.com/spectrocloud-labs/validator/api/v1alpha1"
	"github.com/spectrocloud-labs/validator/pkg/helm"
)

const (
	// A finalizer added to a ValidatorConfig to ensure that plugin Helm releases are properly garbage collected
	CleanupFinalizer = "validator/cleanup"

	// An annotation added to a ValidatorConfig to determine whether or not to update a plugin's Helm release
	PluginValuesHash = "validator/plugin-values"
)

// ValidatorConfigReconciler reconciles a ValidatorConfig object
type ValidatorConfigReconciler struct {
	client.Client
	HelmClient        helm.HelmClient
	HelmSecretsClient helm.SecretsClient
	Log               logr.Logger
	Scheme            *runtime.Scheme
}

//+kubebuilder:rbac:groups=validation.spectrocloud.labs,resources=validatorconfigs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=validation.spectrocloud.labs,resources=validatorconfigs/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=validation.spectrocloud.labs,resources=validatorconfigs/finalizers,verbs=update

func (r *ValidatorConfigReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	r.Log.V(0).Info("Reconciling ValidatorConfig", "name", req.Name, "namespace", req.Namespace)

	vc := &v1alpha1.ValidatorConfig{}
	if err := r.Get(ctx, req.NamespacedName, vc); err != nil {
		r.Log.Error(err, "failed to fetch ValidatorConfig", "key", req)
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// Always update the ValidatorConfig's Status
	defer func() {
		r.updateStatus(ctx, vc)
	}()

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
		return ctrl.Result{}, removeFinalizer(ctx, r.Client, vc, CleanupFinalizer)
	}

	// ensure cleanup finalizer
	if err := ensureFinalizer(ctx, r.Client, vc, CleanupFinalizer); err != nil {
		r.Log.Error(err, "Error ensuring finalizer")
		return ctrl.Result{}, err
	}
	r.Log.V(0).Info("Ensured ValidatorConfig finalizer")

	// deploy/redeploy plugins as required
	if err := r.redeployIfNeeded(ctx, vc); err != nil {
		r.Log.V(0).Error(err, "ValidatorConfig plugin deployment failed", "namespace", vc.Namespace, "name", vc.Name)
		return ctrl.Result{}, err
	}

	return ctrl.Result{RequeueAfter: time.Second * 30}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ValidatorConfigReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.ValidatorConfig{}).
		Complete(r)
}

// updateStatus updates the ValidatorConfig's status subresource
func (r *ValidatorConfigReconciler) updateStatus(ctx context.Context, vc *v1alpha1.ValidatorConfig) {
	if err := r.Status().Update(context.Background(), vc); err != nil {
		r.Log.V(0).Error(err, "failed to update ValidatorConfig status")
	}
	r.Log.V(0).Info("Updated ValidatorConfig", "conditions", vc.Status.Conditions, "time", time.Now())
}

// redeployIfNeeded deploys/redeploys each validator plugin in a ValidatorConfig and deletes plugins that have been removed
func (r *ValidatorConfigReconciler) redeployIfNeeded(ctx context.Context, vc *v1alpha1.ValidatorConfig) error {
	specPlugins := make(map[string]bool)
	conditions := make([]v1alpha1.ValidatorPluginCondition, len(vc.Spec.Plugins))

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

		r.Log.V(0).Info("Installing/upgrading plugin Helm chart", "namespace", vc.Namespace, "name", p.Chart.Name)

		err := r.HelmClient.Upgrade(p.Chart.Name, vc.Namespace, helm.UpgradeOptions{
			Chart:   p.Chart.Name,
			Repo:    p.Chart.Repository,
			Version: p.Chart.Version,
			Values:  p.Values,
		})
		if err != nil {
			// if Helm install/upgrade failed, delete the release so installation is reattempted each iteration
			if strings.Contains(err.Error(), "has no deployed releases") {
				if err := r.HelmClient.Delete(p.Chart.Name, vc.Namespace); err != nil {
					r.Log.V(0).Error(err, "failed to delete Helm release")
				}
			}
			return fmt.Errorf("error installing / upgrading ValidatorConfig: %v", err)
		}
		conditions[i] = getHelmChartCondition(p.Chart.Name, true)
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

	// update ValidatorConfig annotations
	if err := r.Client.Update(ctx, vc); err != nil {
		return err
	}

	// all status modifications must happen after r.Client.Update
	vc.Status.Conditions = conditions

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

	if vc.Annotations == nil {
		vc.Annotations = make(map[string]string)
	}
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
	for _, p := range vc.Spec.Plugins {
		release, err := r.HelmSecretsClient.Get(ctx, p.Chart.Name, vc.Namespace)
		if err != nil {
			if !apierrs.IsNotFound(err) {
				return err
			}
			return nil
		}
		if release.Secret.Labels == nil || release.Secret.Labels["owner"] != "helm" {
			return nil
		}
		r.deletePlugin(vc, p.Chart.Name)
	}
	return nil
}

// deletePlugin deletes the Helm release associated with a ValidatorConfig plugin
func (r *ValidatorConfigReconciler) deletePlugin(vc *v1alpha1.ValidatorConfig, name string) {
	if err := r.HelmClient.Delete(name, vc.Namespace); err != nil {
		r.Log.V(0).Error(err, "failed to delete validator plugin", "namespace", vc.Namespace, "name", name)
	}
	r.Log.V(0).Info("Deleted Helm release for validator plugin", "namespace", vc.Namespace, "name", name)
}

// ensureFinalizer ensures that an object's finalizers include a certain finalizer
func ensureFinalizer(ctx context.Context, client client.Client, obj client.Object, finalizer string) error {
	currentFinalizers := obj.GetFinalizers()
	if !slices.Contains(currentFinalizers, finalizer) {
		newFinalizers := []string{}
		newFinalizers = append(newFinalizers, currentFinalizers...)
		newFinalizers = append(newFinalizers, finalizer)
		obj.SetFinalizers(newFinalizers)
		return client.Update(ctx, obj)
	}
	return nil
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

// getHelmChartCondition builds a ValidatorPluginCondition for a plugin
func getHelmChartCondition(chartName string, installed bool) v1alpha1.ValidatorPluginCondition {
	condition := v1alpha1.ValidatorPluginCondition{
		Type:               v1alpha1.HelmChartDeployedCondition,
		PluginName:         chartName,
		Status:             corev1.ConditionTrue,
		Message:            fmt.Sprintf("Plugin %s is installed", chartName),
		LastTransitionTime: metav1.Time{Time: time.Now()},
	}
	if !installed {
		condition.Status = corev1.ConditionFalse
		condition.Message = fmt.Sprintf("Plugin %s was uninstalled", chartName)
	}
	return condition
}
