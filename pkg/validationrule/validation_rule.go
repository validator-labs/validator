// Package validationrule describes validation rules in the Validator ecosystem.
package validationrule

// Interface defines validation rule behavior.
type Interface interface {
	// Name returns the name of the rule.
	Name() string
	// SetName sets the name of the rule if it is a rule that supports manually specified names.
	SetName(string)
	// AllowsManualName returns whether the validation rule allows manually specifying its name.
	// Code processing rules should check this first before trying to set a manual name.
	AllowsManualName() bool
}

// ManuallyNamed can be embedded into a rule struct to indicate that the rule allows its name to be
// manually specified.
type ManuallyNamed struct{}

// AllowsManualName returns true, indicating that the rule supports manually specifying its name.
func (ManuallyNamed) AllowsManualName() bool {
	return true
}

// AutomaticallyNamed can be embedded into a rule struct to indicate that the rule does not allow its
// name to be manually specified.
type AutomaticallyNamed struct{}

// AllowsManualName returns false, indicating that the rule does not support manually specifying its
// name.
func (AutomaticallyNamed) AllowsManualName() bool {
	return false
}

// SetName is a no-op because the rule does not support manually specifying its name.
func (AutomaticallyNamed) SetName(string) {
	// no-op
}
