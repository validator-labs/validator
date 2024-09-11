// Package validationruleresult includes code to help work with ValidationRuleResults.
package validationruleresult

import (
	"fmt"

	"github.com/validator-labs/validator/api/v1alpha1"
	"github.com/validator-labs/validator/pkg/constants"
	"github.com/validator-labs/validator/pkg/types"
	"github.com/validator-labs/validator/pkg/util"
)

// BuildDefault builds a default ValidationResult for a given validation type. The rule name is
// sanitized and then used as part of the one-word validation rule description in the latest
// condition. The latest condition message should be one that conveys that validation succeeded. The
// validation type should be unique for each combination of plugin and rule (e.g.
// "aws-iam-role-policy").
func BuildDefault(ruleName, latestConditionMsg, validationType string) *types.ValidationRuleResult {
	state := v1alpha1.ValidationSucceeded
	latestCondition := v1alpha1.DefaultValidationCondition()
	latestCondition.Message = latestConditionMsg
	latestCondition.ValidationRule = fmt.Sprintf(
		"%s-%s",
		constants.ValidationRulePrefix, util.Sanitize(ruleName),
	)
	latestCondition.ValidationType = validationType
	return &types.ValidationRuleResult{
		Condition: &latestCondition,
		State:     &state,
	}
}
