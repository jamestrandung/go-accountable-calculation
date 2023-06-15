// Code generated by mockery v2.28.1. DO NOT EDIT.

package acal

import mock "github.com/stretchr/testify/mock"

// MockStaticConditionalValue is an autogenerated mock type for the StaticConditionalValue type
type MockStaticConditionalValue struct {
	mock.Mock
}

// AddCondition provides a mock function with given fields: _a0
func (_m *MockStaticConditionalValue) AddCondition(_a0 *Condition) {
	_m.Called(_a0)
}

// GetCondition provides a mock function with given fields:
func (_m *MockStaticConditionalValue) GetCondition() *Condition {
	ret := _m.Called()

	var r0 *Condition
	if rf, ok := ret.Get(0).(func() *Condition); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Condition)
		}
	}

	return r0
}

// IsConditional provides a mock function with given fields:
func (_m *MockStaticConditionalValue) IsConditional() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

type mockConstructorTestingTNewMockStaticConditionalValue interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockStaticConditionalValue creates a new instance of MockStaticConditionalValue. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockStaticConditionalValue(t mockConstructorTestingTNewMockStaticConditionalValue) *MockStaticConditionalValue {
	mock := &MockStaticConditionalValue{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
