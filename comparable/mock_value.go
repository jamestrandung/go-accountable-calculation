// Code generated by mockery v2.22.1. DO NOT EDIT.

package comparable

import (
	acal "github.com/jamestrandung/go-accountable-calculation/acal"
	mock "github.com/stretchr/testify/mock"
)

// MockValue is an autogenerated mock type for the Value type
type MockValue struct {
	mock.Mock
}

// ExtractValues provides a mock function with given fields: cache
func (_m *MockValue) ExtractValues(cache acal.IValueCache) acal.IValueCache {
	ret := _m.Called(cache)

	var r0 acal.IValueCache
	if rf, ok := ret.Get(0).(func(acal.IValueCache) acal.IValueCache); ok {
		r0 = rf(cache)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(acal.IValueCache)
		}
	}

	return r0
}

// GetAlias provides a mock function with given fields:
func (_m *MockValue) GetAlias() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetName provides a mock function with given fields:
func (_m *MockValue) GetName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetValue provides a mock function with given fields:
func (_m *MockValue) GetValue() interface{} {
	ret := _m.Called()

	var r0 interface{}
	if rf, ok := ret.Get(0).(func() interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// IsNil provides a mock function with given fields:
func (_m *MockValue) IsNil() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// SelfReplaceIfNil provides a mock function with given fields:
func (_m *MockValue) SelfReplaceIfNil() acal.Value {
	ret := _m.Called()

	var r0 acal.Value
	if rf, ok := ret.Get(0).(func() acal.Value); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(acal.Value)
		}
	}

	return r0
}

// SetAlias provides a mock function with given fields: _a0
func (_m *MockValue) SetAlias(_a0 string) {
	_m.Called(_a0)
}

// String provides a mock function with given fields:
func (_m *MockValue) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ToSyntaxOperand provides a mock function with given fields: nextOp
func (_m *MockValue) ToSyntaxOperand(nextOp acal.Op) *acal.SyntaxOperand {
	ret := _m.Called(nextOp)

	var r0 *acal.SyntaxOperand
	if rf, ok := ret.Get(0).(func(acal.Op) *acal.SyntaxOperand); ok {
		r0 = rf(nextOp)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*acal.SyntaxOperand)
		}
	}

	return r0
}

type mockConstructorTestingTNewMockValue interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockValue creates a new instance of MockValue. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockValue(t mockConstructorTestingTNewMockValue) *MockValue {
	mock := &MockValue{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}