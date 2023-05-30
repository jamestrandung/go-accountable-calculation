// Code generated by mockery v2.22.1. DO NOT EDIT.

package comparable

import (
	acal "github.com/jamestrandung/go-accountable-calculation/acal"
	boolean "github.com/jamestrandung/go-accountable-calculation/boolean"

	mock "github.com/stretchr/testify/mock"
)

// MockInterface is an autogenerated mock type for the Interface type
type MockInterface[T interface{}] struct {
	mock.Mock
}

// Equals provides a mock function with given fields: _a0
func (_m *MockInterface[T]) Equals(_a0 T) *boolean.Simple {
	ret := _m.Called(_a0)

	var r0 *boolean.Simple
	if rf, ok := ret.Get(0).(func(T) *boolean.Simple); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*boolean.Simple)
		}
	}

	return r0
}

// EqualsRaw provides a mock function with given fields: v
func (_m *MockInterface[T]) EqualsRaw(v T) *boolean.Simple {
	ret := _m.Called(v)

	var r0 *boolean.Simple
	if rf, ok := ret.Get(0).(func(T) *boolean.Simple); ok {
		r0 = rf(v)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*boolean.Simple)
		}
	}

	return r0
}

// ExtractValues provides a mock function with given fields: cache
func (_m *MockInterface[T]) ExtractValues(cache acal.IValueCache) acal.IValueCache {
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
func (_m *MockInterface[T]) GetAlias() string {
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
func (_m *MockInterface[T]) GetName() string {
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
func (_m *MockInterface[T]) GetValue() interface{} {
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
func (_m *MockInterface[T]) IsNil() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// NotEquals provides a mock function with given fields: _a0
func (_m *MockInterface[T]) NotEquals(_a0 T) *boolean.Simple {
	ret := _m.Called(_a0)

	var r0 *boolean.Simple
	if rf, ok := ret.Get(0).(func(T) *boolean.Simple); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*boolean.Simple)
		}
	}

	return r0
}

// NotEqualsRaw provides a mock function with given fields: v
func (_m *MockInterface[T]) NotEqualsRaw(v T) *boolean.Simple {
	ret := _m.Called(v)

	var r0 *boolean.Simple
	if rf, ok := ret.Get(0).(func(T) *boolean.Simple); ok {
		r0 = rf(v)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*boolean.Simple)
		}
	}

	return r0
}

// SelfReplaceIfNil provides a mock function with given fields:
func (_m *MockInterface[T]) SelfReplaceIfNil() acal.Value {
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
func (_m *MockInterface[T]) SetAlias(_a0 string) {
	_m.Called(_a0)
}

// String provides a mock function with given fields:
func (_m *MockInterface[T]) String() string {
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
func (_m *MockInterface[T]) ToSyntaxOperand(nextOp acal.Op) *acal.SyntaxOperand {
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

type mockConstructorTestingTNewMockInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockInterface creates a new instance of MockInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockInterface[T interface{}](t mockConstructorTestingTNewMockInterface) *MockInterface[T] {
	mock := &MockInterface[T]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}