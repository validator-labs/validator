// Package validationresult contains functions for handling ValidationResult objects.
package validationresult

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/cluster-api/util/patch"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/validator-labs/validator/api/v1alpha1"
	"github.com/validator-labs/validator/pkg/constants"
	"github.com/validator-labs/validator/pkg/types"
	"github.com/validator-labs/validator/pkg/util"
)

const validationErrorMsg = "Validation failed with an unexpected error"

// Patcher is an interface for patching objects.
type Patcher interface {
	Patch(ctx context.Context, obj client.Object, opts ...patch.Option) error
}

// ValidationRule is an interface for validation rules.
type ValidationRule interface {
	client.Object
	Kind() string
	PluginCode() string
	ResultCount() int
}

// Build creates a new ValidationResult for a specific ValidationRule.
func Build(r ValidationRule) *v1alpha1.ValidationResult {
	return &v1alpha1.ValidationResult{
		TypeMeta: metav1.TypeMeta{
			APIVersion: v1alpha1.APIVersion,
			Kind:       r.Kind(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      Name(r),
			Namespace: r.GetNamespace(),
			OwnerReferences: []metav1.OwnerReference{
				{
					APIVersion: r.GetObjectKind().GroupVersionKind().Version,
					Kind:       r.GetObjectKind().GroupVersionKind().Kind,
					Name:       r.GetName(),
					UID:        r.GetUID(),
					Controller: util.Ptr(true),
				},
			},
		},
		Spec: v1alpha1.ValidationResultSpec{
			Plugin:          r.PluginCode(),
			ExpectedResults: r.ResultCount(),
		},
	}
}

// Name returns the name of a ValidationResult for a specific ValidationRule.
func Name(r ValidationRule) string {
	name := fmt.Sprintf("validator-plugin-%s-%s", r.PluginCode(), r.GetName())
	return util.Sanitize(name)
}

// HandleExisting processes a preexisting validation result for the active validator.
func HandleExisting(vr *v1alpha1.ValidationResult, l logr.Logger) {
	l = l.WithValues("name", vr.Name, "namespace", vr.Namespace, "state", vr.Status.State)

	switch vr.Status.State {
	case v1alpha1.ValidationInProgress:
		// validations are only left in progress if an unexpected error occurred
		l.V(0).Info("Previous validation failed with unexpected error")
	case v1alpha1.ValidationFailed:
		// log validation failure, but continue and retry
		cs := getInvalidConditions(vr.Status.ValidationConditions)
		if len(cs) > 0 {
			for _, c := range cs {
				l.V(0).Info("Validation failed. Retrying.",
					"validation", c.ValidationRule, "error", c.Message, "details", c.Details, "failures", c.Failures,
				)
			}
		}
	case v1alpha1.ValidationSucceeded:
		// log validation success, continue to re-validate
		l.V(0).Info("Previous validation succeeded. Re-validating.")
	}
}

// HandleNew creates a new validation result for the active validator.
func HandleNew(ctx context.Context, c client.Client, p Patcher, vr *v1alpha1.ValidationResult, l logr.Logger) error {
	l = l.WithValues("name", vr.Name, "namespace", vr.Namespace)

	// Create the ValidationResult
	if err := c.Create(ctx, vr, &client.CreateOptions{}); err != nil {
		l.V(0).Error(err, "failed to create ValidationResult")
		return err
	}

	// Update the ValidationResult's status
	vr.Status = v1alpha1.ValidationResultStatus{
		State: v1alpha1.ValidationInProgress,
	}
	l = l.WithValues("state", vr.Status.State)

	l.V(0).Info("Preparing to patch ValidationResult")
	if err := patchValidationResult(ctx, p, vr); err != nil {
		l.Error(err, "failed to patch ValidationResult")
		return err
	}

	l.V(0).Info("Successfully patched ValidationResult")
	return nil
}

// SafeUpdate processes and patches a ValidationResult with retry logic.
func SafeUpdate(ctx context.Context, p Patcher, vr *v1alpha1.ValidationResult, vrr types.ValidationResponse, l logr.Logger) error {
	l = l.WithValues("name", vr.Name, "namespace", vr.Namespace)

	if err := Finalize(vr, vrr, l); err != nil {
		l.Error(err, "failed to finalize ValidationResult")
		return err
	}
	if err := patchValidationResult(ctx, p, vr); err != nil {
		l.Error(err, "failed to patch ValidationResult")
		return err
	}

	l.V(0).Info("Successfully patched ValidationResult", "state", vr.Status.State)
	return nil
}

// Finalize finalizes a ValidationResult, ensuring
// that the overall validation status remains failed if a single rule fails.
func Finalize(vr *v1alpha1.ValidationResult, vrr types.ValidationResponse, l logr.Logger) error {
	for i, r := range vrr.ValidationRuleResults {
		// Handle nil ValidationRuleResult
		if r == nil {
			r = &types.ValidationRuleResult{
				Condition: &v1alpha1.ValidationCondition{
					LastValidationTime: metav1.Time{Time: time.Now()},
				},
			}
		}
		// Update overall ValidationResult status
		updateStatus(vr, r, vrr.ValidationRuleErrors[i], l)
	}
	return nil
}

// updateStatus updates a ValidationResult's status with the result of a single validation rule.
func updateStatus(vr *v1alpha1.ValidationResult, vrr *types.ValidationRuleResult, vrrErr error, l logr.Logger) {

	// Finalize result State and Condition in the event of an unexpected error
	if vrrErr != nil {
		vrr.State = util.Ptr(v1alpha1.ValidationFailed)
		vrr.Condition.Status = corev1.ConditionFalse
		vrr.Condition.Message = validationErrorMsg
		vrr.Condition.Failures = append(vrr.Condition.Failures, vrrErr.Error())
	}

	// Update and/or insert the ValidationResult's Conditions with the latest Condition
	idx := getConditionIndexByValidationRule(vr.Status.ValidationConditions, vrr.Condition.ValidationRule)
	if idx == -1 {
		vr.Status.ValidationConditions = append(vr.Status.ValidationConditions, *vrr.Condition)
	} else {
		vr.Status.ValidationConditions[idx] = *vrr.Condition
	}

	// Set State to:
	// - ValidationFailed if ANY condition failed
	// - ValidationSucceeded if ALL conditions succeeded
	// - ValidationInProgress otherwise
	vr.Status.State = *vrr.State
	for _, c := range vr.Status.ValidationConditions {
		if c.Status == corev1.ConditionTrue {
			vr.Status.State = v1alpha1.ValidationSucceeded
		}
		if c.Status == corev1.ConditionFalse {
			vr.Status.State = v1alpha1.ValidationFailed
			break
		}
	}

	l.V(0).Info("Updated ValidationResult status", "overallState", vr.Status.State, "validationRuleState", vrr.State,
		"validationRuleReason", vrr.Condition.ValidationRule, "validationRuleMessage", vrr.Condition.Message,
	)
}

// getInvalidConditions filters a ValidationCondition array and returns all conditions corresponding to a failed validation.
func getInvalidConditions(conditions []v1alpha1.ValidationCondition) []v1alpha1.ValidationCondition {
	invalidConditions := make([]v1alpha1.ValidationCondition, 0)
	for _, c := range conditions {
		if strings.HasPrefix(c.ValidationRule, constants.ValidationRulePrefix) && c.Status == corev1.ConditionFalse {
			invalidConditions = append(invalidConditions, c)
		}
	}
	return invalidConditions
}

// getConditionIndexByValidationRule retrieves the index of a condition from a ValidationCondition array matching a specific reason.
func getConditionIndexByValidationRule(conditions []v1alpha1.ValidationCondition, validationRule string) int {
	for i, c := range conditions {
		if c.ValidationRule == validationRule {
			return i
		}
	}
	return -1
}

// patchValidationResult patches a ValidationResult.
func patchValidationResult(ctx context.Context, p Patcher, vr *v1alpha1.ValidationResult) error {
	return p.Patch(ctx, vr, patch.WithStatusObservedGeneration{})
}
