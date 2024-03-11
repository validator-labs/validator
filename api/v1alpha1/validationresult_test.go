package v1alpha1

import (
	"reflect"
	"testing"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clusterv1beta1 "sigs.k8s.io/cluster-api/api/v1beta1"
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
					Plugin:          "AWS",
					ExpectedResults: 1,
				},
				Status: ValidationResultStatus{
					State: ValidationSucceeded,
					Conditions: []clusterv1beta1.Condition{
						{
							Type:               SinkEmission,
							Reason:             string(SinkEmitSucceeded),
							Status:             corev1.ConditionTrue,
							LastTransitionTime: metav1.Now(),
						},
					},
					ValidationConditions: []ValidationCondition{
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
			expectedHash: "Jp+QKNlngLsgv2kTlKhR3w==",
		},
	}
	for _, c := range cs {
		hash := c.validationResult.Hash()
		if !reflect.DeepEqual(hash, c.expectedHash) {
			t.Errorf("expected (%s), got (%s)", c.expectedHash, hash)
		}
	}
}
