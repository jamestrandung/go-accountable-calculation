// Code generated by mockery v2.28.1. DO NOT EDIT.

package acal

import mock "github.com/stretchr/testify/mock"

// MockICondOps is an autogenerated mock type for the ICondOps type
type MockICondOps struct {
	mock.Mock
}

// ApplyProgressiveCondition provides a mock function with given fields: criteria, values
func (_m *MockICondOps) ApplyProgressiveCondition(criteria TypedValue[bool], values ...ProgressiveConditionalValue) CloseIfFunc {
	_va := make([]interface{}, len(values))
	for _i := range values {
		_va[_i] = values[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, criteria)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 CloseIfFunc
	if rf, ok := ret.Get(0).(func(TypedValue[bool], ...ProgressiveConditionalValue) CloseIfFunc); ok {
		r0 = rf(criteria, values...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(CloseIfFunc)
		}
	}

	return r0
}

// ApplyStaticCondition provides a mock function with given fields: criteria, values
func (_m *MockICondOps) ApplyStaticCondition(criteria TypedValue[bool], values ...StaticConditionalValue) {
	_va := make([]interface{}, len(values))
	for _i := range values {
		_va[_i] = values[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, criteria)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

type mockConstructorTestingTNewMockICondOps interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockICondOps creates a new instance of MockICondOps. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockICondOps(t mockConstructorTestingTNewMockICondOps) *MockICondOps {
	mock := &MockICondOps{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
