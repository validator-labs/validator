package plugins

import (
	"testing"
	"time"

	ctrl "sigs.k8s.io/controller-runtime"
)

func TestFrequencyFromAnnotations(t *testing.T) {

	logger := ctrl.Log.WithName("TestFrequencyAnnotation")

	testCases := []struct {
		name        string
		annotations map[string]string
		result      ctrl.Result
	}{
		{
			name: "CorrectAnnotationValue",
			annotations: map[string]string{
				"validation.validator.labs/reconciliation-frequency": "20",
			},
			result: ctrl.Result{
				RequeueAfter: time.Duration(time.Second * 20),
			},
		},
		{
			name: "IncorrectAnnotationValue",
			annotations: map[string]string{
				"validation.validator.labs/reconciliation-frequency": "nonint",
			},
			result: ctrl.Result{
				RequeueAfter: time.Duration(time.Second * 120),
			},
		},
		{
			name:        "NoAnnotation",
			annotations: map[string]string{},
			result: ctrl.Result{
				RequeueAfter: time.Duration(time.Second * 120),
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := FrequencyFromAnnotations(logger, testCase.annotations)
			if actual != testCase.result {
				t.Errorf("expected %v, got %v", testCase.result, actual)
			}
		})
	}

}
