// Package validationruleresult includes code to help work with ValidationRuleResults.
package validationruleresult

import (
	"encoding/json"
	"reflect"
	"testing"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/validator-labs/validator/api/v1alpha1"
	"github.com/validator-labs/validator/pkg/types"
	"github.com/validator-labs/validator/pkg/util"
)

func TestBuildDefault(t *testing.T) {
	type args struct {
		latestConditionMsg string
		validationType     string
		validationRule     string
		ruleName           string
	}
	tests := []struct {
		name string
		args args
		want *types.ValidationRuleResult
	}{
		{
			name: "uses validation rule description if provided",
			args: args{
				latestConditionMsg: "a",
				validationType:     "b",
				validationRule:     "c",
				ruleName:           "",
			},
			want: &types.ValidationRuleResult{
				Condition: &v1alpha1.ValidationCondition{
					ValidationType: "b",
					ValidationRule: "c",
					Message:        "a",
					Status:         corev1.ConditionTrue,
				},
				State: util.Ptr(v1alpha1.ValidationSucceeded),
			},
		},
		{
			name: "correctly generates a validation rule description based on rule name if validation rule not provided",
			args: args{
				latestConditionMsg: "a",
				validationType:     "b",
				validationRule:     "",
				ruleName:           "d",
			},
			want: &types.ValidationRuleResult{
				Condition: &v1alpha1.ValidationCondition{
					ValidationType: "b",
					ValidationRule: "validation-d",
					Message:        "a",
					Status:         corev1.ConditionTrue,
				},
				State: util.Ptr(v1alpha1.ValidationSucceeded),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BuildDefault(tt.args.latestConditionMsg, tt.args.validationType, tt.args.validationRule, tt.args.ruleName)

			// Remove the time. It will be different each run.
			got.Condition.LastValidationTime = metav1.Time{}

			gotJSON, _ := json.MarshalIndent(got, "", "  ")
			wantJSON, _ := json.MarshalIndent(tt.want, "", "  ")

			if !reflect.DeepEqual(gotJSON, wantJSON) {
				t.Errorf("BuildDefault() = %v, want %v", string(gotJSON), string(wantJSON))
			}
		})
	}
}
