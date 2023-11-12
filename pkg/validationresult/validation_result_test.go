package validationresult

import (
	"errors"
	"reflect"
	"testing"

	corev1 "k8s.io/api/core/v1"

	"github.com/spectrocloud-labs/validator/api/v1alpha1"
	"github.com/spectrocloud-labs/validator/pkg/types"
	"github.com/spectrocloud-labs/validator/pkg/util/ptr"
)

var err = errors.New("error")

func res(s corev1.ConditionStatus, state v1alpha1.ValidationState) *types.ValidationResult {
	return &types.ValidationResult{
		Condition: &v1alpha1.ValidationCondition{
			Status: s,
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
			Status: c,
		}
		if err != nil {
			condition.Message = validationErrorMsg
			condition.Failures = append(condition.Failures, err.Error())
		}
		vr.Status.Conditions = append(vr.Status.Conditions, condition)
	}
	return vr
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
