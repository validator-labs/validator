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
	"time"

	apierrs "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/go-logr/logr"
	validationv1alpha1 "github.com/spectrocloud-labs/validator/api/v1alpha1"
)

// ValidationResultReconciler reconciles a ValidationResult object
type ValidationResultReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=validation.spectrocloud.labs,resources=validationresults,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=validation.spectrocloud.labs,resources=validationresults/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=validation.spectrocloud.labs,resources=validationresults/finalizers,verbs=update

func (r *ValidationResultReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	r.Log.V(0).Info("Reconciling ValidationResult", "name", req.Name, "namespace", req.Namespace)

	vr := &validationv1alpha1.ValidationResult{}
	if err := r.Get(ctx, req.NamespacedName, vr); err != nil {
		// ignore not-found errors, since they can't be fixed by an immediate requeue
		if apierrs.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		r.Log.Error(err, "failed to fetch ValidationResult")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	r.Log.V(0).Info("Plugin", "name", vr.Spec.Plugin)

	if vr.Status.State == validationv1alpha1.ValidationFailed || vr.Status.State == validationv1alpha1.ValidationSucceeded {
		r.Log.V(0).Info("ValidationResult complete", "name", vr.Name, "state", vr.Status.State)

		for _, c := range vr.Status.Conditions {
			r.Log.V(0).Info("ValidationResult metadata", "type", c.ValidationType,
				"rule", c.ValidationRule, "status", c.Status,
				"message", c.Message, "details", c.Details,
				"failures", c.Failures, "time", c.LastValidationTime,
			)
		}

		// TODO: send result to a sink
	}

	r.Log.V(0).Info("Validation in progress. Requeuing in 30s.", "name", req.Name, "namespace", req.Namespace)
	return ctrl.Result{RequeueAfter: time.Second * 30}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ValidationResultReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&validationv1alpha1.ValidationResult{}).
		Complete(r)
}
