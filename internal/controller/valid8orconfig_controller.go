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

	v1alpha1 "github.com/spectrocloud-labs/valid8or/api/v1alpha1"
	"github.com/spectrocloud-labs/valid8or/internal/helm"
)

const (
	// A finalizer that is added to the Valid8orConfig CR to ensure that helm delete is executed.
	CleanupFinalizer = "valid8or/cleanup"
)

// Valid8orConfigReconciler reconciles a Valid8orConfig object
type Valid8orConfigReconciler struct {
	client.Client
	HelmClient        helm.HelmClient
	HelmSecretsClient helm.SecretsClient
	Log               logr.Logger
	Scheme            *runtime.Scheme
}

//+kubebuilder:rbac:groups=validation.spectrocloud.labs,resources=valid8orconfigs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=validation.spectrocloud.labs,resources=valid8orconfigs/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=validation.spectrocloud.labs,resources=valid8orconfigs/finalizers,verbs=update

func (r *Valid8orConfigReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	r.Log.V(0).Info("Reconciling Valid8orConfig", "name", req.Name, "namespace", req.Namespace)

	vc := &v1alpha1.Valid8orConfig{}
	if err := r.Get(ctx, req.NamespacedName, vc); err != nil {
		// ignore not-found errors, since they can't be fixed by an immediate requeue
		if apierrs.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		r.Log.Error(err, "failed to fetch Valid8orConfig")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// handle Valid8orConfig deletion
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
	r.Log.V(0).Info("Ensured Valid8orConfig finalizer")

	// deploy/redeploy plugins as required
	if err := r.redeployIfNeeded(ctx, vc); err != nil {
		r.Log.V(0).Error(err, "Valid8orConfig plugin deployment failed", "namespace", vc.Namespace, "name", vc.Name)
		return ctrl.Result{RequeueAfter: time.Second * 5}, err
	}

	return ctrl.Result{RequeueAfter: time.Second * 30}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *Valid8orConfigReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.Valid8orConfig{}).
		Complete(r)
}

// redeployIfNeeded deploys/redeploys each valid8or plugin in a Valid8orConfig
func (r *Valid8orConfigReconciler) redeployIfNeeded(ctx context.Context, vc *v1alpha1.Valid8orConfig) error {
	for _, p := range vc.Spec.Plugins {
		// skip plugin if already deployed
		if isConditionTrue(vc, p.Chart.Name, v1alpha1.HelmChartDeployedCondition) {
			return nil
		}

		r.Log.V(0).Info("Installed/upgraded plugin Helm chart", "namespace", vc.Namespace, "name", p.Chart.Name)

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
			return fmt.Errorf("error installing / upgrading Valid8orConfig: %v", err)
		}

		if err := r.markConditionTrue(ctx, vc, p.Chart.Name); err != nil {
			return fmt.Errorf("error updating Valid8orConfig: %v", err)
		}
	}

	return nil
}

// deletePlugins deletes each valid8or plugin's Helm release
func (r *Valid8orConfigReconciler) deletePlugins(ctx context.Context, vc *v1alpha1.Valid8orConfig) error {
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
		if err := r.HelmClient.Delete(p.Chart.Name, vc.Namespace); err != nil {
			r.Log.V(0).Error(err, "failed to delete valid8or plugin", "namespace", vc.Namespace, "name", p.Chart.Name)
		}
		r.Log.V(0).Info("Deleted Helm release for valid8or plugin", "namespace", vc.Namespace, "name", p.Chart.Name)
	}
	return nil
}

// markConditionTrue marks a Valid8orPluginCondition as true
func (r *Valid8orConfigReconciler) markConditionTrue(ctx context.Context, vc *v1alpha1.Valid8orConfig, chartName string) error {
	condition := v1alpha1.Valid8orPluginCondition{
		Type:               v1alpha1.HelmChartDeployedCondition,
		PluginName:         chartName,
		Status:             corev1.ConditionTrue,
		Message:            fmt.Sprintf("Plugin %s is healthy", chartName),
		LastTransitionTime: metav1.Time{Time: time.Now()},
	}
	idx := conditionIndex(vc, chartName, v1alpha1.HelmChartDeployedCondition)
	if idx == -1 {
		vc.Status.Conditions = append(vc.Status.Conditions, condition)
	} else {
		vc.Status.Conditions[idx] = condition
	}
	return r.Client.Update(ctx, vc)
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

// conditionIndex retrieves the index of a Valid8orPluginCondition from a Valid8orConfig's status
func conditionIndex(vc *v1alpha1.Valid8orConfig, chartName string, conditionType v1alpha1.ConditionType) int {
	for i, c := range vc.Status.Conditions {
		if c.Type == conditionType && c.PluginName == chartName {
			return i
		}
	}
	return -1
}

// isConditionTrue checks whether a Valid8orPluginCondition is true
func isConditionTrue(vc *v1alpha1.Valid8orConfig, chartName string, conditionType v1alpha1.ConditionType) bool {
	idx := conditionIndex(vc, chartName, conditionType)
	if idx == -1 {
		return false
	}
	return vc.Status.Conditions[idx].Status == corev1.ConditionTrue
}
