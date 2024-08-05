package validationresult

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"

	"github.com/validator-labs/validator/api/v1alpha1"
	"github.com/validator-labs/validator/pkg/constants"
	"github.com/validator-labs/validator/pkg/test"
	"github.com/validator-labs/validator/pkg/types"
	"github.com/validator-labs/validator/pkg/util"
)

var err = errors.New("error")

func res(s corev1.ConditionStatus, state v1alpha1.ValidationState) *types.ValidationRuleResult {
	return &types.ValidationRuleResult{
		Condition: &v1alpha1.ValidationCondition{
			Status:         s,
			ValidationRule: constants.ValidationRulePrefix,
		},
		State: util.Ptr(state),
	}
}

func vr(cs []corev1.ConditionStatus, state v1alpha1.ValidationState, err error) *v1alpha1.ValidationResult {
	vr := &v1alpha1.ValidationResult{
		Status: v1alpha1.ValidationResultStatus{
			ValidationConditions: make([]v1alpha1.ValidationCondition, 0),
			State:                state,
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
		vr.Status.ValidationConditions = append(vr.Status.ValidationConditions, condition)
	}
	return vr
}

func TestHandleExistingValidationResult(t *testing.T) {
	cs := []struct {
		name string
		vr   *v1alpha1.ValidationResult
	}{
		{
			name: "ValidationInProgress",
			vr:   vr(nil, v1alpha1.ValidationInProgress, nil),
		},
		{
			name: "ValidationFailed",
			vr:   vr([]corev1.ConditionStatus{corev1.ConditionFalse}, v1alpha1.ValidationFailed, nil),
		},
		{
			name: "ValidationSucceeded",
			vr:   vr(nil, v1alpha1.ValidationSucceeded, nil),
		},
	}
	for _, c := range cs {
		t.Log(c.name)
		HandleExisting(c.vr, logr.Logger{})
	}
}

func TestHandleNewValidationResult(t *testing.T) {
	cs := []struct {
		name     string
		client   test.ClientMock
		patcher  test.PatchHelperMock
		vr       *v1alpha1.ValidationResult
		expected error
	}{
		{
			name: "Pass",
			client: test.ClientMock{
				SubResourceMock: test.SubResourceMock{},
			},
			patcher:  test.PatchHelperMock{},
			vr:       vr(nil, v1alpha1.ValidationInProgress, nil),
			expected: nil,
		},
		{
			name: "Fail (create)",
			client: test.ClientMock{
				CreateErrors:    []error{errors.New("creation failed")},
				SubResourceMock: test.SubResourceMock{},
			},
			patcher:  test.PatchHelperMock{},
			vr:       vr([]corev1.ConditionStatus{corev1.ConditionFalse}, v1alpha1.ValidationFailed, nil),
			expected: errors.New("creation failed"),
		},
		{
			name: "Fail (get)",
			client: test.ClientMock{
				GetErrors:       []error{errors.New("get failed")},
				SubResourceMock: test.SubResourceMock{},
			},
			patcher:  test.PatchHelperMock{},
			vr:       vr([]corev1.ConditionStatus{corev1.ConditionFalse}, v1alpha1.ValidationFailed, nil),
			expected: errors.New("get failed"),
		},
		{
			name: "Fail (status update)",
			client: test.ClientMock{
				SubResourceMock: test.SubResourceMock{
					UpdateErrors: []error{errors.New("status update failed")},
				},
			},
			patcher:  test.PatchHelperMock{},
			vr:       vr(nil, v1alpha1.ValidationSucceeded, nil),
			expected: errors.New("status update failed"),
		},
		{
			name:   "Fail (patch)",
			client: test.ClientMock{},
			patcher: test.PatchHelperMock{
				PatchErrors: []error{errors.New("patch failed")},
			},
			vr:       vr(nil, v1alpha1.ValidationSucceeded, nil),
			expected: errors.New("patch failed"),
		},
	}
	for _, c := range cs {
		t.Log(c.name)
		err = HandleNew(context.Background(), c.client, c.patcher, c.vr, logr.Logger{})
		if err != nil && !reflect.DeepEqual(c.expected.Error(), err.Error()) {
			t.Errorf("expected (%v), got (%v)", c.expected, err)
		}
	}
}

func TestSafeUpdateValidationResult(t *testing.T) {
	cs := []struct {
		name    string
		client  test.ClientMock
		patcher test.PatchHelperMock
		vr      *v1alpha1.ValidationResult
		vrr     types.ValidationResponse
	}{
		{
			name:    "Pass",
			client:  test.ClientMock{},
			patcher: test.PatchHelperMock{},
			vr:      &v1alpha1.ValidationResult{},
			vrr: types.ValidationResponse{
				ValidationRuleResults: []*types.ValidationRuleResult{
					res(corev1.ConditionTrue, v1alpha1.ValidationSucceeded),
				},
				ValidationRuleErrors: []error{nil},
			},
		},
		{
			name: "Fail (get)",
			client: test.ClientMock{
				GetErrors: []error{errors.New("get failed")},
			},
			patcher: test.PatchHelperMock{},
			vr:      &v1alpha1.ValidationResult{},
			vrr: types.ValidationResponse{
				ValidationRuleResults: []*types.ValidationRuleResult{
					res(corev1.ConditionTrue, v1alpha1.ValidationSucceeded),
				},
				ValidationRuleErrors: []error{errors.New("get failed")},
			},
		},
		{
			name: "Fail (update)",
			client: test.ClientMock{
				SubResourceMock: test.SubResourceMock{
					UpdateErrors: []error{errors.New("status update failed")},
				},
			},
			patcher: test.PatchHelperMock{},
			vr:      &v1alpha1.ValidationResult{},
			vrr: types.ValidationResponse{
				ValidationRuleResults: []*types.ValidationRuleResult{
					res(corev1.ConditionTrue, v1alpha1.ValidationSucceeded),
				},
				ValidationRuleErrors: []error{errors.New("status update failed")},
			},
		},
		{
			name:   "Fail (patch)",
			client: test.ClientMock{},
			patcher: test.PatchHelperMock{
				PatchErrors: []error{errors.New("patch failed")},
			},
			vr: &v1alpha1.ValidationResult{},
			vrr: types.ValidationResponse{
				ValidationRuleResults: []*types.ValidationRuleResult{
					res(corev1.ConditionTrue, v1alpha1.ValidationSucceeded),
				},
				ValidationRuleErrors: []error{errors.New("patch failed")},
			},
		},
		{
			name: "Fail (nil)",
			client: test.ClientMock{
				SubResourceMock: test.SubResourceMock{
					UpdateErrors: []error{errors.New("status update failed")},
				},
			},
			patcher: test.PatchHelperMock{},
			vr:      &v1alpha1.ValidationResult{},
			vrr: types.ValidationResponse{
				ValidationRuleResults: []*types.ValidationRuleResult{nil},
				ValidationRuleErrors:  []error{errors.New("status update failed")},
			},
		},
	}
	for _, c := range cs {
		t.Log(c.name)
		SafeUpdate(context.Background(), c.patcher, c.vr, c.vrr, logr.Logger{})
	}
}

func TestUpdateValidationResultStatus(t *testing.T) {
	cs := []struct {
		name       string
		vrr        *types.ValidationRuleResult
		vrrErr     error
		vrCurr     *v1alpha1.ValidationResult
		vrExpected *v1alpha1.ValidationResult
	}{
		{
			name:       "nil -> Pass -> PASS",
			vrr:        res(corev1.ConditionTrue, v1alpha1.ValidationSucceeded),
			vrrErr:     nil,
			vrCurr:     vr(nil, v1alpha1.ValidationInProgress, nil),
			vrExpected: vr([]corev1.ConditionStatus{corev1.ConditionTrue}, v1alpha1.ValidationSucceeded, nil),
		},
		{
			name:       "nil -> Error -> FAIL",
			vrr:        res(corev1.ConditionFalse, v1alpha1.ValidationFailed),
			vrrErr:     err,
			vrCurr:     vr(nil, v1alpha1.ValidationInProgress, err),
			vrExpected: vr([]corev1.ConditionStatus{corev1.ConditionFalse}, v1alpha1.ValidationFailed, err),
		},
		{
			name:       "nil -> Fail -> FAIL",
			vrr:        res(corev1.ConditionFalse, v1alpha1.ValidationFailed),
			vrrErr:     nil,
			vrCurr:     vr(nil, v1alpha1.ValidationInProgress, nil),
			vrExpected: vr([]corev1.ConditionStatus{corev1.ConditionFalse}, v1alpha1.ValidationFailed, nil),
		},
		{
			name:       "Pass -> Pass -> PASS",
			vrr:        res(corev1.ConditionTrue, v1alpha1.ValidationSucceeded),
			vrrErr:     nil,
			vrCurr:     vr([]corev1.ConditionStatus{corev1.ConditionTrue}, v1alpha1.ValidationSucceeded, nil),
			vrExpected: vr([]corev1.ConditionStatus{corev1.ConditionTrue}, v1alpha1.ValidationSucceeded, nil),
		},
		{
			name:       "Pass -> Fail -> PASS",
			vrr:        res(corev1.ConditionFalse, v1alpha1.ValidationFailed),
			vrrErr:     nil,
			vrCurr:     vr([]corev1.ConditionStatus{corev1.ConditionTrue}, v1alpha1.ValidationSucceeded, nil),
			vrExpected: vr([]corev1.ConditionStatus{corev1.ConditionFalse}, v1alpha1.ValidationFailed, nil),
		},
		{
			name:       "Fail -> Pass -> PASS",
			vrr:        res(corev1.ConditionTrue, v1alpha1.ValidationSucceeded),
			vrrErr:     nil,
			vrCurr:     vr([]corev1.ConditionStatus{corev1.ConditionFalse}, v1alpha1.ValidationFailed, nil),
			vrExpected: vr([]corev1.ConditionStatus{corev1.ConditionTrue}, v1alpha1.ValidationSucceeded, nil),
		},
		{
			name:       "[Pass, Pass] -> Fail -> FAIL",
			vrr:        res(corev1.ConditionFalse, v1alpha1.ValidationFailed),
			vrrErr:     nil,
			vrCurr:     vr([]corev1.ConditionStatus{corev1.ConditionTrue, corev1.ConditionTrue}, v1alpha1.ValidationSucceeded, nil),
			vrExpected: vr([]corev1.ConditionStatus{corev1.ConditionFalse, corev1.ConditionTrue}, v1alpha1.ValidationFailed, nil),
		},
		{
			name:       "[Fail, Fail] -> Pass -> FAIL",
			vrr:        res(corev1.ConditionTrue, v1alpha1.ValidationSucceeded),
			vrrErr:     nil,
			vrCurr:     vr([]corev1.ConditionStatus{corev1.ConditionFalse, corev1.ConditionFalse}, v1alpha1.ValidationFailed, nil),
			vrExpected: vr([]corev1.ConditionStatus{corev1.ConditionTrue, corev1.ConditionFalse}, v1alpha1.ValidationFailed, nil),
		},
		{
			name:       "[Fail, Pass] -> Pass -> PASS",
			vrr:        res(corev1.ConditionTrue, v1alpha1.ValidationSucceeded),
			vrrErr:     nil,
			vrCurr:     vr([]corev1.ConditionStatus{corev1.ConditionFalse, corev1.ConditionTrue}, v1alpha1.ValidationFailed, nil),
			vrExpected: vr([]corev1.ConditionStatus{corev1.ConditionTrue, corev1.ConditionTrue}, v1alpha1.ValidationSucceeded, nil),
		},
	}
	for _, c := range cs {
		t.Log(c.name)
		updateStatus(c.vrCurr, c.vrr, c.vrrErr, logr.Logger{})
		if !reflect.DeepEqual(c.vrCurr.Hash(), c.vrExpected.Hash()) {
			t.Errorf("expected (%+v), got (%+v)", c.vrExpected, c.vrCurr)
		}
	}
}
