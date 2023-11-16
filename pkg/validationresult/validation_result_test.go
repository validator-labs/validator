package validationresult

import (
	"errors"
	"reflect"
	"testing"

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	ktypes "k8s.io/apimachinery/pkg/types"

	"github.com/spectrocloud-labs/validator/api/v1alpha1"
	"github.com/spectrocloud-labs/validator/internal/test"
	"github.com/spectrocloud-labs/validator/pkg/constants"
	"github.com/spectrocloud-labs/validator/pkg/types"
	"github.com/spectrocloud-labs/validator/pkg/util/ptr"
)

var err = errors.New("error")

func res(s corev1.ConditionStatus, state v1alpha1.ValidationState) *types.ValidationResult {
	return &types.ValidationResult{
		Condition: &v1alpha1.ValidationCondition{
			Status:         s,
			ValidationRule: constants.ValidationRulePrefix,
		},
		State: ptr.Ptr(state),
	}
}

func vr(cs []corev1.ConditionStatus, state v1alpha1.ValidationState, err error) *v1alpha1.ValidationResult {
	vr := &v1alpha1.ValidationResult{
		Status: v1alpha1.ValidationResultStatus{
			Conditions: make([]v1alpha1.ValidationCondition, 0),
			State:      state,
		},
	}
	for _, c := range cs {
		condition := v1alpha1.ValidationCondition{
			Status:         c,
			ValidationRule: constants.ValidationRulePrefix,
		}
		if err != nil {
			condition.Message = validationErrorMsg
			condition.Failures = append(condition.Failures, err.Error())
		}
		vr.Status.Conditions = append(vr.Status.Conditions, condition)
	}
	return vr
}

func TestHandleExistingValidationResult(t *testing.T) {
	cs := []struct {
		name string
		nn   ktypes.NamespacedName
		res  *v1alpha1.ValidationResult
	}{
		{
			name: "ValidationInProgress",
			nn:   ktypes.NamespacedName{},
			res:  vr(nil, v1alpha1.ValidationInProgress, nil),
		},
		{
			name: "ValidationFailed",
			nn:   ktypes.NamespacedName{},
			res:  vr([]corev1.ConditionStatus{corev1.ConditionFalse}, v1alpha1.ValidationFailed, nil),
		},
		{
			name: "ValidationSucceeded",
			nn:   ktypes.NamespacedName{},
			res:  vr(nil, v1alpha1.ValidationSucceeded, nil),
		},
	}
	for _, c := range cs {
		t.Log(c.name)
		HandleExistingValidationResult(c.nn, c.res, logr.Logger{})
	}
}

func TestHandleNewValidationResult(t *testing.T) {
	cs := []struct {
		name     string
		client   test.ClientMock
		res      *v1alpha1.ValidationResult
		expected error
	}{
		{
			name: "Pass",
			client: test.ClientMock{
				SubResourceMock: test.SubResourceMock{},
			},
			res:      vr(nil, v1alpha1.ValidationInProgress, nil),
			expected: nil,
		},
		{
			name: "Fail (create)",
			client: test.ClientMock{
				CreateErrors:    []error{errors.New("creation failed")},
				SubResourceMock: test.SubResourceMock{},
			},
			res:      vr([]corev1.ConditionStatus{corev1.ConditionFalse}, v1alpha1.ValidationFailed, nil),
			expected: errors.New("creation failed"),
		},
		{
			name: "Fail (status update)",
			client: test.ClientMock{
				SubResourceMock: test.SubResourceMock{
					UpdateErrors: []error{errors.New("update failed")},
				},
			},
			res:      vr(nil, v1alpha1.ValidationSucceeded, nil),
			expected: errors.New("update failed"),
		},
	}
	for _, c := range cs {
		t.Log(c.name)
		err := HandleNewValidationResult(c.client, c.res, logr.Logger{})
		if err != nil && !reflect.DeepEqual(c.expected.Error(), err.Error()) {
			t.Errorf("expected (%v), got (%v)", c.expected, err)
		}
	}
}

func TestUpdateValidationResult(t *testing.T) {
	cs := []struct {
		name       string
		res        *types.ValidationResult
		resErr     error
		vrCurr     *v1alpha1.ValidationResult
		vrExpected *v1alpha1.ValidationResult
	}{
		{
			name:       "nil -> Pass -> PASS",
			res:        res(corev1.ConditionTrue, v1alpha1.ValidationSucceeded),
			resErr:     nil,
			vrCurr:     vr(nil, v1alpha1.ValidationInProgress, nil),
			vrExpected: vr([]corev1.ConditionStatus{corev1.ConditionTrue}, v1alpha1.ValidationSucceeded, nil),
		},
		{
			name:       "nil -> Error -> FAIL",
			res:        res(corev1.ConditionFalse, v1alpha1.ValidationFailed),
			resErr:     err,
			vrCurr:     vr(nil, v1alpha1.ValidationInProgress, err),
			vrExpected: vr([]corev1.ConditionStatus{corev1.ConditionFalse}, v1alpha1.ValidationFailed, err),
		},
		{
			name:       "nil -> Fail -> FAIL",
			res:        res(corev1.ConditionFalse, v1alpha1.ValidationFailed),
			resErr:     nil,
			vrCurr:     vr(nil, v1alpha1.ValidationInProgress, nil),
			vrExpected: vr([]corev1.ConditionStatus{corev1.ConditionFalse}, v1alpha1.ValidationFailed, nil),
		},
		{
			name:       "Pass -> Pass -> PASS",
			res:        res(corev1.ConditionTrue, v1alpha1.ValidationSucceeded),
			resErr:     nil,
			vrCurr:     vr([]corev1.ConditionStatus{corev1.ConditionTrue}, v1alpha1.ValidationSucceeded, nil),
			vrExpected: vr([]corev1.ConditionStatus{corev1.ConditionTrue}, v1alpha1.ValidationSucceeded, nil),
		},
		{
			name:       "Pass -> Fail -> PASS",
			res:        res(corev1.ConditionFalse, v1alpha1.ValidationFailed),
			resErr:     nil,
			vrCurr:     vr([]corev1.ConditionStatus{corev1.ConditionTrue}, v1alpha1.ValidationSucceeded, nil),
			vrExpected: vr([]corev1.ConditionStatus{corev1.ConditionFalse}, v1alpha1.ValidationFailed, nil),
		},
		{
			name:       "Fail -> Pass -> PASS",
			res:        res(corev1.ConditionTrue, v1alpha1.ValidationSucceeded),
			resErr:     nil,
			vrCurr:     vr([]corev1.ConditionStatus{corev1.ConditionFalse}, v1alpha1.ValidationFailed, nil),
			vrExpected: vr([]corev1.ConditionStatus{corev1.ConditionTrue}, v1alpha1.ValidationSucceeded, nil),
		},
		{
			name:       "[Pass, Pass] -> Fail -> FAIL",
			res:        res(corev1.ConditionFalse, v1alpha1.ValidationFailed),
			resErr:     nil,
			vrCurr:     vr([]corev1.ConditionStatus{corev1.ConditionTrue, corev1.ConditionTrue}, v1alpha1.ValidationSucceeded, nil),
			vrExpected: vr([]corev1.ConditionStatus{corev1.ConditionFalse, corev1.ConditionTrue}, v1alpha1.ValidationFailed, nil),
		},
		{
			name:       "[Fail, Fail] -> Pass -> FAIL",
			res:        res(corev1.ConditionTrue, v1alpha1.ValidationSucceeded),
			resErr:     nil,
			vrCurr:     vr([]corev1.ConditionStatus{corev1.ConditionFalse, corev1.ConditionFalse}, v1alpha1.ValidationFailed, nil),
			vrExpected: vr([]corev1.ConditionStatus{corev1.ConditionTrue, corev1.ConditionFalse}, v1alpha1.ValidationFailed, nil),
		},
		{
			name:       "[Fail, Pass] -> Pass -> PASS",
			res:        res(corev1.ConditionTrue, v1alpha1.ValidationSucceeded),
			resErr:     nil,
			vrCurr:     vr([]corev1.ConditionStatus{corev1.ConditionFalse, corev1.ConditionTrue}, v1alpha1.ValidationFailed, nil),
			vrExpected: vr([]corev1.ConditionStatus{corev1.ConditionTrue, corev1.ConditionTrue}, v1alpha1.ValidationSucceeded, nil),
		},
	}
	for _, c := range cs {
		t.Log(c.name)
		updateValidationResult(c.vrCurr, c.res, c.resErr)
		if !reflect.DeepEqual(c.vrCurr.Hash(), c.vrExpected.Hash()) {
			t.Errorf("expected (%+v), got (%+v)", c.vrExpected, c.vrCurr)
		}
	}
}
