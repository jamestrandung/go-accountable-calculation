package acal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// MockCondOps can be used in tests to mock ICondOps.
func MockCondOps(t *testing.T) (*MockICondOps, func()) {
	old := condOps
	mock := NewMockICondOps(t)

	condOps = mock
	return mock, func() {
		condOps = old
	}
}

func TestCondOpsImpl_ApplyStaticCondition(t *testing.T) {
	criteria := NewMockTypedValue[bool](t)

	expectedCond := NewCondition(criteria)

	mockStaticConditionalValue1 := NewMockStaticConditionalValue(t)
	mockStaticConditionalValue1.On("AddCondition", expectedCond).Once()

	mockStaticConditionalValue2 := NewMockStaticConditionalValue(t)
	mockStaticConditionalValue2.On("AddCondition", expectedCond).Once()

	ops := condOpsImpl{}

	ops.ApplyStaticCondition(criteria, mockStaticConditionalValue1, mockStaticConditionalValue2)
}

func TestCondOpsImpl_ApplyProgressiveCondition(t *testing.T) {
	criteria := NewMockTypedValue[bool](t)

	closeIf1 := false
	mockProgressiveConditionalValue1 := NewMockProgressiveConditionalValue(t)
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
	mockProgressiveConditionalValue2 := NewMockProgressiveConditionalValue(t)
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

	closeIfFn()
	assert.Equal(t, true, closeIf1)
	assert.Equal(t, true, closeIf2)
}
