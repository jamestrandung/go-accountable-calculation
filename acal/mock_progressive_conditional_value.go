// Code generated by mockery v2.15.0. DO NOT EDIT.

package acal

import mock "github.com/stretchr/testify/mock"

// MockProgressiveConditionalValue is an autogenerated mock type for the ProgressiveConditionalValue type
type MockProgressiveConditionalValue struct {
	mock.Mock
}

// AddCondition provides a mock function with given fields: _a0
func (_m *MockProgressiveConditionalValue) AddCondition(_a0 TypedValue[bool]) CloseIfFunc {
	ret := _m.Called(_a0)

	var r0 CloseIfFunc
	if rf, ok := ret.Get(0).(func(TypedValue[bool]) CloseIfFunc); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(CloseIfFunc)
		}
	}

	return r0
}

type mockConstructorTestingTNewMockProgressiveConditionalValue interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockProgressiveConditionalValue creates a new instance of MockProgressiveConditionalValue. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockProgressiveConditionalValue(t mockConstructorTestingTNewMockProgressiveConditionalValue) *MockProgressiveConditionalValue {
	mock := &MockProgressiveConditionalValue{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
