package acal

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

// MockCondOps can be used in tests to mock ICondOps.
func MockCondOps() (*MockICondOps, func()) {
	old := condOps
	mock := &MockICondOps{}

	condOps = mock
	return mock, func() {
		condOps = old
	}
}

func TestCondOpsImpl_ApplyStaticCondition(t *testing.T) {
	criteria := &MockTypedValue[bool]{}

	expectedCond := NewCondition(criteria)

	mockStaticConditionalValue1 := &MockStaticConditionalValue{}
	mockStaticConditionalValue1.On("AddCondition", expectedCond).Once()

	mockStaticConditionalValue2 := &MockStaticConditionalValue{}
	mockStaticConditionalValue2.On("AddCondition", expectedCond).Once()

	ops := condOpsImpl{}

	ops.ApplyStaticCondition(criteria, mockStaticConditionalValue1, mockStaticConditionalValue2)

	mock.AssertExpectationsForObjects(t, mockStaticConditionalValue1, mockStaticConditionalValue2)
}

func TestCondOpsImpl_ApplyProgressiveCondition(t *testing.T) {
	criteria := &MockTypedValue[bool]{}

	closeIf1 := false
	mockProgressiveConditionalValue1 := &MockProgressiveConditionalValue{}
	mockProgressiveConditionalValue1.On("AddCondition", criteria).
		Return(
			CloseIfFunc(
				func() {
					closeIf1 = true
				},
			),
		).
		Once()

	closeIf2 := false
	mockProgressiveConditionalValue2 := &MockProgressiveConditionalValue{}
	mockProgressiveConditionalValue2.On("AddCondition", criteria).
		Return(
			CloseIfFunc(
				func() {
					closeIf2 = true
				},
			),
		).
		Once()

	ops := condOpsImpl{}

	closeIfFn := ops.ApplyProgressiveCondition(criteria, mockProgressiveConditionalValue1, mockProgressiveConditionalValue2)

	mock.AssertExpectationsForObjects(t, mockProgressiveConditionalValue1, mockProgressiveConditionalValue2)

	closeIfFn()
	assert.Equal(t, true, closeIf1)
	assert.Equal(t, true, closeIf2)
}
