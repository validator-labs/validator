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
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ktypes "k8s.io/apimachinery/pkg/types"
	clusterv1beta1 "sigs.k8s.io/cluster-api/api/v1beta1"
	"sigs.k8s.io/cluster-api/util/patch"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	v1alpha1 "github.com/validator-labs/validator/api/v1alpha1"
	"github.com/validator-labs/validator/pkg/constants"
	"github.com/validator-labs/validator/pkg/sinks"
	"github.com/validator-labs/validator/pkg/types"
)

// ValidationResultHash is used to determine whether to re-emit updates to a validation result sink.
const ValidationResultHash = "validator/validation-result-hash"

// ValidationResultReconciler reconciles a ValidationResult object
type ValidationResultReconciler struct {
	client.Client
	Log        logr.Logger
	Namespace  string
	Scheme     *runtime.Scheme
	SinkClient *sinks.Client
}

// +kubebuilder:rbac:groups=validation.spectrocloud.labs,resources=validationresults,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=validation.spectrocloud.labs,resources=validationresults/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=validation.spectrocloud.labs,resources=validationresults/finalizers,verbs=update

// Reconcile reconciles a ValidationResult.
func (r *ValidationResultReconciler) Reconcile(ctx context.Context, req ctrl.Request) (_ ctrl.Result, reterr error) {
	r.Log.V(0).Info("Reconciling ValidationResult", "name", req.Name, "namespace", req.Namespace)

	vc := &v1alpha1.ValidatorConfig{}
	vcKey := ktypes.NamespacedName{Namespace: r.Namespace, Name: constants.ValidatorConfig}
	if err := r.Get(ctx, vcKey, vc); err != nil {
		if !apierrs.IsNotFound(err) {
			r.Log.Error(err, "failed to fetch ValidatorConfig", "key", req)
		}
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	vr := &v1alpha1.ValidationResult{}
	if err := r.Get(ctx, req.NamespacedName, vr); err != nil {
		if !apierrs.IsNotFound(err) {
			r.Log.Error(err, "failed to fetch ValidationResult", "key", req)
		}
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	patcher, err := patch.NewHelper(vr, r.Client)
	if err != nil {
		return ctrl.Result{}, errors.Wrapf(err, "failed to create patch helper for ValidationResult %s", vr.Name)
	}

	currHash := vr.Hash()
	prevHash, ok := vr.ObjectMeta.Annotations[ValidationResultHash]

	sinkConditionIdx, sinkCondition := getConditionByType(vr.Status.Conditions, v1alpha1.SinkEmission)
	if sinkConditionIdx == -1 {
		sinkCondition.Type = v1alpha1.SinkEmission
		sinkCondition.Reason = string(v1alpha1.SinkEmitNA)
		sinkCondition.Status = corev1.ConditionTrue
		sinkCondition.LastTransitionTime = metav1.Now()
	}

	defer func() {
		r.Log.V(1).Info("Preparing to patch ValidationResult", "validationResult", vr.Name)
		if err := patchValidationResult(ctx, patcher, vr); err != nil && reterr == nil {
			reterr = err
			r.Log.Error(err, "failed to patch ValidationResult", "validationResult", vr.Name)
			return
		}
		r.Log.V(1).Info("Successfully patched ValidationResult", "validationResult", vr.Name)
	}()

	r.Log.V(0).Info("Plugin", "name", vr.Spec.Plugin)

	if vr.Status.State == v1alpha1.ValidationFailed || vr.Status.State == v1alpha1.ValidationSucceeded {
		r.Log.V(0).Info("ValidationResult complete", "name", vr.Name, "state", vr.Status.State)

		for _, c := range vr.Status.ValidationConditions {
			r.Log.V(0).Info("ValidationResult metadata", "type", c.ValidationType,
				"rule", c.ValidationRule, "status", c.Status,
				"message", c.Message, "details", c.Details,
				"failures", c.Failures, "time", c.LastValidationTime,
			)
		}

		// Do not emit until the number of conditions matches the expected number of results. Otherwise, N
		// emissions will occur during the 1st reconciliation, where N is the number of rules in the validator.
		if vc.Spec.Sink != nil && len(vr.Status.ValidationConditions) == vr.Spec.ExpectedResults && (!ok || prevHash != currHash) {
			sinkState, err := r.emitToSink(ctx, vc, vr)
			if err != nil {
				sinkCondition.Message = err.Error()
			}
			sinkCondition.Reason = string(sinkState)
		}
	}

	// update ValidationResult annotations
	if prevHash != currHash {
		if vr.ObjectMeta.Annotations == nil {
			vr.ObjectMeta.Annotations = make(map[string]string, 0)
		}
		vr.ObjectMeta.Annotations[ValidationResultHash] = currHash
	}

	// update SinkEmission condition
	if sinkConditionIdx == -1 {
		vr.Status.Conditions = append(vr.Status.Conditions, sinkCondition)
	} else {
		vr.Status.Conditions[sinkConditionIdx] = sinkCondition
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

// emitToSink emits a ValidationResult to a sink - either upon the completion of the 1st reconciliation, or if its hash changes.
func (r *ValidationResultReconciler) emitToSink(ctx context.Context, vc *v1alpha1.ValidatorConfig, vr *v1alpha1.ValidationResult) (v1alpha1.SinkEmitState, error) {
	sink := sinks.NewSink(types.SinkType(vc.Spec.Sink.Type), r.Log)

	var sinkConfig map[string][]byte
	if vc.Spec.Sink.SecretName != "" {
		sinkSecret := &corev1.Secret{}
		sinkConfigKey := ktypes.NamespacedName{Namespace: r.Namespace, Name: vc.Spec.Sink.SecretName}
		if err := r.Client.Get(ctx, sinkConfigKey, sinkSecret); err != nil {
			r.Log.Error(err, "failed to fetch sink configuration secret")
			return v1alpha1.SinkEmitFailed, err
		}
		sinkConfig = sinkSecret.Data
	}

	if err := sink.Configure(*r.SinkClient, sinkConfig); err != nil {
		r.Log.Error(err, "failed to configure sink")
		return v1alpha1.SinkEmitFailed, err
	}
	if err := sink.Emit(*vr); err != nil {
		r.Log.Error(err, "failed to emit ValidationResult to sink", "sinkType", vc.Spec.Sink.Type)
		return v1alpha1.SinkEmitFailed, err
	}

	return v1alpha1.SinkEmitSucceeded, nil
}

// getConditionIndexByType retrieves the index of a condition from a Condition array matching a specific condition type.
func getConditionByType(conditions []clusterv1beta1.Condition, conditionType clusterv1beta1.ConditionType) (int, clusterv1beta1.Condition) {
	for i, c := range conditions {
		if c.Type == conditionType {
			return i, c
		}
	}
	return -1, clusterv1beta1.Condition{}
}

// patchValidationResult patches a ValidationResult, ignoring conflicts on the conditions owned by this controller.
func patchValidationResult(ctx context.Context, patchHelper *patch.Helper, vr *v1alpha1.ValidationResult) error {
	return patchHelper.Patch(
		ctx,
		vr,
		patch.WithOwnedConditions{Conditions: []clusterv1beta1.ConditionType{
			v1alpha1.SinkEmission,
		}},
		patch.WithStatusObservedGeneration{},
	)
}
