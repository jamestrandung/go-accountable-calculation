// Code generated by mockery v2.15.0. DO NOT EDIT.

package acal

import mock "github.com/stretchr/testify/mock"

// MockIValueOps is an autogenerated mock type for the IValueOps type
type MockIValueOps struct {
	mock.Mock
}

// Describe provides a mock function with given fields: v
func (_m *MockIValueOps) Describe(v Value) string {
	ret := _m.Called(v)

	var r0 string
	if rf, ok := ret.Get(0).(func(Value) string); ok {
		r0 = rf(v)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// DescribeValueAsFormula provides a mock function with given fields: v
func (_m *MockIValueOps) DescribeValueAsFormula(v Value) func() *SyntaxNode {
	ret := _m.Called(v)

	var r0 func() *SyntaxNode
	if rf, ok := ret.Get(0).(func(Value) func() *SyntaxNode); ok {
		r0 = rf(v)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(func() *SyntaxNode)
		}
	}

	return r0
}

// Identify provides a mock function with given fields: v
func (_m *MockIValueOps) Identify(v Value) string {
	ret := _m.Called(v)

	var r0 string
	if rf, ok := ret.Get(0).(func(Value) string); ok {
		r0 = rf(v)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// IsAnchored provides a mock function with given fields: v
func (_m *MockIValueOps) IsAnchored(v Value) bool {
	ret := _m.Called(v)

	var r0 bool
	if rf, ok := ret.Get(0).(func(Value) bool); ok {
		r0 = rf(v)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsNilValue provides a mock function with given fields: v
func (_m *MockIValueOps) IsNilValue(v Value) bool {
	ret := _m.Called(v)

	var r0 bool
	if rf, ok := ret.Get(0).(func(Value) bool); ok {
		r0 = rf(v)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

type mockConstructorTestingTNewMockIValueOps interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockIValueOps creates a new instance of MockIValueOps. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockIValueOps(t mockConstructorTestingTNewMockIValueOps) *MockIValueOps {
	mock := &MockIValueOps{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
