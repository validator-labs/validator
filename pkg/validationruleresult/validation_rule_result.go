// Package validationruleresult includes code to help work with ValidationRuleResults.
package validationruleresult

import (
	"fmt"

	"github.com/validator-labs/validator/api/v1alpha1"
	"github.com/validator-labs/validator/pkg/constants"
	"github.com/validator-labs/validator/pkg/types"
	"github.com/validator-labs/validator/pkg/util"
)

// BuildDefault builds a default ValidationResult for a given validation type.
//
// The latest condition message param should be one that conveys that validation succeeded. The
// validation type param should be unique for each combination of plugin and rule (e.g.
// "aws-iam-role-policy").
//
// One of the validation rule param or rule name params must be provided. If validation rule is
// provided, it is used as the validation rule description of the condition. If it isn't, a
// description is generated based on the rule name and used instead.
func BuildDefault(latestConditionMsg, validationType, validationRule, ruleName string) *types.ValidationRuleResult {
	state := v1alpha1.ValidationSucceeded
	latestCondition := v1alpha1.DefaultValidationCondition()
	latestCondition.Message = latestConditionMsg
	if validationRule != "" {
		latestCondition.ValidationRule = validationRule
	} else {
		latestCondition.ValidationRule = fmt.Sprintf(
			"%s-%s",
			constants.ValidationRulePrefix, util.Sanitize(ruleName),
		)
	}
	latestCondition.ValidationType = validationType
	return &types.ValidationRuleResult{
		Condition: &latestCondition,
		State:     &state,
	}
}
