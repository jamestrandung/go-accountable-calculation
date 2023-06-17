// Code generated by mockery v2.28.1. DO NOT EDIT.

package float

import (
	acal "github.com/jamestrandung/go-accountable-calculation/acal"
	boolean "github.com/jamestrandung/go-accountable-calculation/boolean"

	decimal "github.com/my-shopspring/decimal"

	mock "github.com/stretchr/testify/mock"
)

// MockInterface is an autogenerated mock type for the Interface type
type MockInterface struct {
	mock.Mock
}

// Abs provides a mock function with given fields:
func (_m *MockInterface) Abs() {
	_m.Called()
}

// Add provides a mock function with given fields: _a0
func (_m *MockInterface) Add(_a0 acal.TypedValue[decimal.Decimal]) Simple {
	ret := _m.Called(_a0)

	var r0 Simple
	if rf, ok := ret.Get(0).(func(acal.TypedValue[decimal.Decimal]) Simple); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(Simple)
	}

	return r0
}

// Ceil provides a mock function with given fields:
func (_m *MockInterface) Ceil() {
	_m.Called()
}

// Decimal provides a mock function with given fields:
func (_m *MockInterface) Decimal() decimal.Decimal {
	ret := _m.Called()

	var r0 decimal.Decimal
	if rf, ok := ret.Get(0).(func() decimal.Decimal); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(decimal.Decimal)
	}

	return r0
}

// Div provides a mock function with given fields: _a0
func (_m *MockInterface) Div(_a0 acal.TypedValue[decimal.Decimal]) Simple {
	ret := _m.Called(_a0)

	var r0 Simple
	if rf, ok := ret.Get(0).(func(acal.TypedValue[decimal.Decimal]) Simple); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(Simple)
	}

	return r0
}

// Equals provides a mock function with given fields: _a0
func (_m *MockInterface) Equals(_a0 acal.TypedValue[decimal.Decimal]) boolean.Simple {
	ret := _m.Called(_a0)

	var r0 boolean.Simple
	if rf, ok := ret.Get(0).(func(acal.TypedValue[decimal.Decimal]) boolean.Simple); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(boolean.Simple)
	}

	return r0
}

// ExtractValues provides a mock function with given fields: cache
func (_m *MockInterface) ExtractValues(cache acal.IValueCache) acal.IValueCache {
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

// Float provides a mock function with given fields:
func (_m *MockInterface) Float() float64 {
	ret := _m.Called()

	var r0 float64
	if rf, ok := ret.Get(0).(func() float64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(float64)
	}

	return r0
}

// Floor provides a mock function with given fields:
func (_m *MockInterface) Floor() Simple {
	ret := _m.Called()

	var r0 Simple
	if rf, ok := ret.Get(0).(func() Simple); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(Simple)
	}

	return r0
}

// GetAlias provides a mock function with given fields:
func (_m *MockInterface) GetAlias() string {
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
func (_m *MockInterface) GetName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetTypedValue provides a mock function with given fields:
func (_m *MockInterface) GetTypedValue() decimal.Decimal {
	ret := _m.Called()

	var r0 decimal.Decimal
	if rf, ok := ret.Get(0).(func() decimal.Decimal); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(decimal.Decimal)
	}

	return r0
}

// GetValue provides a mock function with given fields:
func (_m *MockInterface) GetValue() interface{} {
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

// HasIdentity provides a mock function with given fields:
func (_m *MockInterface) HasIdentity() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Identify provides a mock function with given fields:
func (_m *MockInterface) Identify() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Inv provides a mock function with given fields:
func (_m *MockInterface) Inv() Simple {
	ret := _m.Called()

	var r0 Simple
	if rf, ok := ret.Get(0).(func() Simple); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(Simple)
	}

	return r0
}

// IsNegative provides a mock function with given fields:
func (_m *MockInterface) IsNegative() boolean.Simple {
	ret := _m.Called()

	var r0 boolean.Simple
	if rf, ok := ret.Get(0).(func() boolean.Simple); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(boolean.Simple)
	}

	return r0
}

// IsNil provides a mock function with given fields:
func (_m *MockInterface) IsNil() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsPositive provides a mock function with given fields:
func (_m *MockInterface) IsPositive() boolean.Simple {
	ret := _m.Called()

	var r0 boolean.Simple
	if rf, ok := ret.Get(0).(func() boolean.Simple); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(boolean.Simple)
	}

	return r0
}

// IsZero provides a mock function with given fields:
func (_m *MockInterface) IsZero() boolean.Simple {
	ret := _m.Called()

	var r0 boolean.Simple
	if rf, ok := ret.Get(0).(func() boolean.Simple); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(boolean.Simple)
	}

	return r0
}

// LargerThan provides a mock function with given fields: _a0
func (_m *MockInterface) LargerThan(_a0 acal.TypedValue[decimal.Decimal]) boolean.Simple {
	ret := _m.Called(_a0)

	var r0 boolean.Simple
	if rf, ok := ret.Get(0).(func(acal.TypedValue[decimal.Decimal]) boolean.Simple); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(boolean.Simple)
	}

	return r0
}

// LargerThanEquals provides a mock function with given fields: _a0
func (_m *MockInterface) LargerThanEquals(_a0 acal.TypedValue[decimal.Decimal]) boolean.Simple {
	ret := _m.Called(_a0)

	var r0 boolean.Simple
	if rf, ok := ret.Get(0).(func(acal.TypedValue[decimal.Decimal]) boolean.Simple); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(boolean.Simple)
	}

	return r0
}

// Mul provides a mock function with given fields: _a0
func (_m *MockInterface) Mul(_a0 acal.TypedValue[decimal.Decimal]) Simple {
	ret := _m.Called(_a0)

	var r0 Simple
	if rf, ok := ret.Get(0).(func(acal.TypedValue[decimal.Decimal]) Simple); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(Simple)
	}

	return r0
}

// Neg provides a mock function with given fields:
func (_m *MockInterface) Neg() Simple {
	ret := _m.Called()

	var r0 Simple
	if rf, ok := ret.Get(0).(func() Simple); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(Simple)
	}

	return r0
}

// NotEquals provides a mock function with given fields: _a0
func (_m *MockInterface) NotEquals(_a0 acal.TypedValue[decimal.Decimal]) boolean.Simple {
	ret := _m.Called(_a0)

	var r0 boolean.Simple
	if rf, ok := ret.Get(0).(func(acal.TypedValue[decimal.Decimal]) boolean.Simple); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(boolean.Simple)
	}

	return r0
}

// Round provides a mock function with given fields: decimalPlace
func (_m *MockInterface) Round(decimalPlace acal.TypedValue[decimal.Decimal]) Simple {
	ret := _m.Called(decimalPlace)

	var r0 Simple
	if rf, ok := ret.Get(0).(func(acal.TypedValue[decimal.Decimal]) Simple); ok {
		r0 = rf(decimalPlace)
	} else {
		r0 = ret.Get(0).(Simple)
	}

	return r0
}

// SetAlias provides a mock function with given fields: _a0
func (_m *MockInterface) SetAlias(_a0 string) {
	_m.Called(_a0)
}

// SmallerThan provides a mock function with given fields: _a0
func (_m *MockInterface) SmallerThan(_a0 acal.TypedValue[decimal.Decimal]) boolean.Simple {
	ret := _m.Called(_a0)

	var r0 boolean.Simple
	if rf, ok := ret.Get(0).(func(acal.TypedValue[decimal.Decimal]) boolean.Simple); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(boolean.Simple)
	}

	return r0
}

// SmallerThanEquals provides a mock function with given fields: _a0
func (_m *MockInterface) SmallerThanEquals(_a0 acal.TypedValue[decimal.Decimal]) boolean.Simple {
	ret := _m.Called(_a0)

	var r0 boolean.Simple
	if rf, ok := ret.Get(0).(func(acal.TypedValue[decimal.Decimal]) boolean.Simple); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(boolean.Simple)
	}

	return r0
}

// Stringify provides a mock function with given fields:
func (_m *MockInterface) Stringify() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Sub provides a mock function with given fields: _a0
func (_m *MockInterface) Sub(_a0 acal.TypedValue[decimal.Decimal]) Simple {
	ret := _m.Called(_a0)

	var r0 Simple
	if rf, ok := ret.Get(0).(func(acal.TypedValue[decimal.Decimal]) Simple); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(Simple)
	}

	return r0
}

// ToSyntaxOperand provides a mock function with given fields: nextOp
func (_m *MockInterface) ToSyntaxOperand(nextOp acal.Op) *acal.SyntaxOperand {
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
func NewMockInterface(t mockConstructorTestingTNewMockInterface) *MockInterface {
	mock := &MockInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
