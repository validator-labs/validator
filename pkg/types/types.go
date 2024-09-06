// Package types contains structs used by to construct ValidationResults.
package types

import (
	corev1 "k8s.io/api/core/v1"

	"github.com/validator-labs/validator/api/v1alpha1"
	"github.com/validator-labs/validator/pkg/util"
)

// ErrValidationFailed is the error message returned when a validation rule fails.
const ErrValidationFailed = "Validation failed with an unexpected error"

// ValidationRuleResult is the result of the execution of a validation rule by a validator.
type ValidationRuleResult struct {
	Condition *v1alpha1.ValidationCondition
	State     *v1alpha1.ValidationState
}

// Finalize sets the ValidationRuleResult state to ValidationFailed if an non-nil error is provided.
func (vrr *ValidationRuleResult) Finalize(err error) {
	if err != nil {
		vrr.State = util.Ptr(v1alpha1.ValidationFailed)
		vrr.Condition.Status = corev1.ConditionFalse
		vrr.Condition.Message = ErrValidationFailed
		vrr.Condition.Failures = append(vrr.Condition.Failures, err.Error())
	}
}

// ValidationResponse is the reconciliation output of one or more validation rules by a validator.
type ValidationResponse struct {
	ValidationRuleResults []*ValidationRuleResult
	ValidationRuleErrors  []error
}

// AddResult adds a ValidationRuleResult and associated error to a ValidationResponse.
func (v *ValidationResponse) AddResult(vrr *ValidationRuleResult, err error) {
	v.ValidationRuleResults = append(v.ValidationRuleResults, vrr)
	v.ValidationRuleErrors = append(v.ValidationRuleErrors, err)
}

// SinkType is the type of sink to which a notification should be sent.
type SinkType string

const (
	// SinkTypeAlertmanager is an Alertmanager sink.
	SinkTypeAlertmanager SinkType = "alertmanager"

	// SinkTypeSlack is a Slack sink.
	SinkTypeSlack SinkType = "slack"
)
