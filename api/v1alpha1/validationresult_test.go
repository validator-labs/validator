package v1alpha1

import (
	"reflect"
	"testing"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestHash(t *testing.T) {
	cs := []struct {
		name             string
		validationResult ValidationResult
		expectedHash     string
	}{
		{
			name: "Pass",
			validationResult: ValidationResult{
				ObjectMeta: metav1.ObjectMeta{
					UID: "1",
				},
				Spec: ValidationResultSpec{
					Plugin: "AWS",
				},
				Status: ValidationResultStatus{
					State:     ValidationSucceeded,
					SinkState: SinkEmitSucceeded,
					Conditions: []ValidationCondition{
						{
							ValidationType:     "foo",
							ValidationRule:     "bar",
							Message:            "baz",
							Details:            []string{"detail"},
							Failures:           []string{"failure"},
							Status:             corev1.ConditionTrue,
							LastValidationTime: metav1.Now(),
						},
					},
				},
			},
			expectedHash: "mCyJwAeP5yOG82mDw8Yy1Q==",
		},
	}
	for _, c := range cs {
		hash := c.validationResult.Hash()
		if !reflect.DeepEqual(hash, c.expectedHash) {
			t.Errorf("expected (%s), got (%s)", c.expectedHash, hash)
		}
	}
}
