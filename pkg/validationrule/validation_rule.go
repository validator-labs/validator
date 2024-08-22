// Package validationrule describes validation rules.
package validationrule

// Interface defines validation rule behavior.
type Interface interface {
	// Name returns the name of the rule.
	Name() string
	// SetName sets the name of the rule if it is a rule that requires manually specifying its name.
	// This should be a no-op for rules that automatically generate their name.
	SetName(string)
	// RequiresName returns whether the validation rule requires its name to be manually specified.
	// This should return false for rules that automatically generate their name.
	RequiresName() bool
}

// ManuallyNamed can be embedded into a rule struct to indicate that the rule requires its name to
// be manually specified.
type ManuallyNamed struct{}

// RequiresName returns true.
func (ManuallyNamed) RequiresName() bool {
	return true
}

// AutomaticallyNamed can be embedded into a rule struct to indicate that the rule does not require
// its name to be manually specified because it is automatically generated from other data in the
// rule.
type AutomaticallyNamed struct{}

// RequiresName returns false.
func (AutomaticallyNamed) RequiresName() bool {
	return false
}

// SetName is a no-op because the rule does not support manually specifying its name.
func (AutomaticallyNamed) SetName(string) {
	// no-op
}
