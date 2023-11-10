package types

import "github.com/spectrocloud-labs/validator/api/v1alpha1"

// ValidationResult is the result of the execution of a validation rule by a validator
type ValidationResult struct {
	Condition *v1alpha1.ValidationCondition
	State     *v1alpha1.ValidationState
}
