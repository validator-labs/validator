package types

import "github.com/spectrocloud-labs/validator/api/v1alpha1"

// ValidationRuleResult is the result of the execution of a validation rule by a validator.
type ValidationRuleResult struct {
	Condition *v1alpha1.ValidationCondition
	State     *v1alpha1.ValidationState
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

type SinkType string

const (
	SinkTypeAlertmanager SinkType = "alertmanager"
	SinkTypeSlack        SinkType = "slack"
)
