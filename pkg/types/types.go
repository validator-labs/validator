package types

import "github.com/spectrocloud-labs/valid8or/api/v1alpha1"

// ValidationResult is the result of the execution of a validation rule by a validator
type ValidationResult struct {
	Condition *v1alpha1.ValidationCondition
	State     *v1alpha1.ValidationState
}

// MonotonicBool starts off false and remains true permanently if updated to true
type MonotonicBool struct {
	Ok bool
}

// Update updates the status of a monotonic bool. If the monotonic bool is already true, Update() is a noop.
func (m *MonotonicBool) Update(ok bool) {
	m.Ok = ok || m.Ok
}
