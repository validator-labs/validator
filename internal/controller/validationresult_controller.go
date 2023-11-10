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

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	v1alpha1 "github.com/spectrocloud-labs/validator/api/v1alpha1"
	"github.com/spectrocloud-labs/validator/internal/sinks"
	"github.com/spectrocloud-labs/validator/pkg/constants"
)

// ValidationResultHash is used to determine whether to re-emit updates to a validation result sink.
const ValidationResultHash = "validator/validation-result-hash"

var (
	vr        *v1alpha1.ValidationResult
	vrKey     types.NamespacedName
	sinkState v1alpha1.SinkState
)

// ValidationResultReconciler reconciles a ValidationResult object
type ValidationResultReconciler struct {
	client.Client
	Log        logr.Logger
	Namespace  string
	Scheme     *runtime.Scheme
	SinkClient *sinks.Client
}

//+kubebuilder:rbac:groups=validation.spectrocloud.labs,resources=validationresults,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=validation.spectrocloud.labs,resources=validationresults/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=validation.spectrocloud.labs,resources=validationresults/finalizers,verbs=update

func (r *ValidationResultReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	r.Log.V(0).Info("Reconciling ValidationResult", "name", req.Name, "namespace", req.Namespace)

	vc := &v1alpha1.ValidatorConfig{}
	vcKey := types.NamespacedName{Namespace: r.Namespace, Name: constants.ValidatorConfig}
	if err := r.Get(ctx, vcKey, vc); err != nil {
		// ignore not-found errors, since they can't be fixed by an immediate requeue
		if apierrs.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		r.Log.Error(err, "failed to fetch ValidatorConfig")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	vr = &v1alpha1.ValidationResult{}
	vrKey = req.NamespacedName

	if err := r.Get(ctx, req.NamespacedName, vr); err != nil {
		// ignore not-found errors, since they can't be fixed by an immediate requeue
		if apierrs.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		r.Log.Error(err, "failed to fetch ValidationResult")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	currHash := vr.Hash()
	prevHash, ok := vr.ObjectMeta.Annotations[ValidationResultHash]
	sinkState = v1alpha1.SinkEmitNone

	// always update the ValidationResult's Status
	defer func() {
		r.updateStatus(ctx)
	}()

	r.Log.V(0).Info("Plugin", "name", vr.Spec.Plugin)

	if vr.Status.State == v1alpha1.ValidationFailed || vr.Status.State == v1alpha1.ValidationSucceeded {
		r.Log.V(0).Info("ValidationResult complete", "name", vr.Name, "state", vr.Status.State)

		for _, c := range vr.Status.Conditions {
			r.Log.V(0).Info("ValidationResult metadata", "type", c.ValidationType,
				"rule", c.ValidationRule, "status", c.Status,
				"message", c.Message, "details", c.Details,
				"failures", c.Failures, "time", c.LastValidationTime,
			)
		}

		// Emit ValidationResult to a sink - either upon the completion of the 1st reconciliation, or if its hash changes.
		// Do not emit until the number of conditions matches the expected number of results, otherwise N
		// emissions will occur during the 1st reconciliation, where N is the number of rules in the validator.
		if vc.Spec.Sink != nil && len(vr.Status.Conditions) == vr.Spec.ExpectedResults && (!ok || prevHash != currHash) {

			sink := sinks.NewSink(vc.Spec.Sink.Type, r.Log)
			sinkState = v1alpha1.SinkEmitFailed

			var sinkConfig map[string][]byte
			if vc.Spec.Sink.SecretName != "" {
				sinkSecret := &corev1.Secret{}
				sinkConfigKey := types.NamespacedName{Namespace: r.Namespace, Name: vc.Spec.Sink.SecretName}
				if err := r.Client.Get(ctx, sinkConfigKey, sinkSecret); err != nil {
					r.Log.Error(err, "failed to fetch sink configuration secret")
					return ctrl.Result{}, err
				}
				sinkConfig = sinkSecret.Data
			}

			if err := sink.Configure(*r.SinkClient, *vc, sinkConfig); err != nil {
				r.Log.Error(err, "failed to configure sink")
				return ctrl.Result{}, err
			}
			if err := sink.Emit(*vr); err != nil {
				r.Log.Error(err, "failed to emit ValidationResult to sink", "sinkType", vc.Spec.Sink.Type)
			}

			sinkState = v1alpha1.SinkEmitSucceeded
		}
	}

	// update ValidationResult annotations
	if prevHash != currHash {
		if vr.ObjectMeta.Annotations == nil {
			vr.ObjectMeta.Annotations = make(map[string]string, 0)
		}
		vr.ObjectMeta.Annotations[ValidationResultHash] = currHash
		if err := r.Client.Update(ctx, vr); err != nil {
			return ctrl.Result{}, err
		}
	}

	r.Log.V(0).Info("Validation in progress. Requeuing in 30s.", "name", req.Name, "namespace", req.Namespace)
	return ctrl.Result{RequeueAfter: time.Second * 30}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ValidationResultReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.ValidationResult{}).
		Complete(r)
}

// updateStatus updates the ValidatorResult's status subresource
func (r *ValidationResultReconciler) updateStatus(ctx context.Context) {
	if err := r.Get(ctx, vrKey, vr); err != nil {
		r.Log.V(0).Error(err, "failed to get ValidationResult")
	}

	// all status modifications must happen after r.Client.Update
	vr.Status.SinkState = sinkState

	if err := r.Status().Update(context.Background(), vr); err != nil {
		r.Log.V(0).Error(err, "failed to update ValidationResult status")
	}
	r.Log.V(0).Info("Updated ValidationResult", "conditions", vr.Status.Conditions, "time", time.Now())
}
