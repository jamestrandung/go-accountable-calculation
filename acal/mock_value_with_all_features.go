// Code generated by mockery v2.28.1. DO NOT EDIT.

package acal

import mock "github.com/stretchr/testify/mock"

// mockValueWithAllFeatures is an autogenerated mock type for the valueWithAllFeatures type
type mockValueWithAllFeatures struct {
	mock.Mock
}

// AddCondition provides a mock function with given fields: _a0
func (_m *mockValueWithAllFeatures) AddCondition(_a0 *Condition) {
	_m.Called(_a0)
}

// ExtractValues provides a mock function with given fields: cache
func (_m *mockValueWithAllFeatures) ExtractValues(cache IValueCache) IValueCache {
	ret := _m.Called(cache)

	var r0 IValueCache
	if rf, ok := ret.Get(0).(func(IValueCache) IValueCache); ok {
		r0 = rf(cache)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(IValueCache)
		}
	}

	return r0
}

// GetAlias provides a mock function with given fields:
func (_m *mockValueWithAllFeatures) GetAlias() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetCondition provides a mock function with given fields:
func (_m *mockValueWithAllFeatures) GetCondition() *Condition {
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

// GetFormula provides a mock function with given fields:
func (_m *mockValueWithAllFeatures) GetFormula() *SyntaxNode {
	ret := _m.Called()

	var r0 *SyntaxNode
	if rf, ok := ret.Get(0).(func() *SyntaxNode); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*SyntaxNode)
		}
	}

	return r0
}

// GetName provides a mock function with given fields:
func (_m *mockValueWithAllFeatures) GetName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetTags provides a mock function with given fields:
func (_m *mockValueWithAllFeatures) GetTags() Tags {
	ret := _m.Called()

	var r0 Tags
	if rf, ok := ret.Get(0).(func() Tags); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Tags)
		}
	}

	return r0
}

// GetValue provides a mock function with given fields:
func (_m *mockValueWithAllFeatures) GetValue() interface{} {
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

// HasFormula provides a mock function with given fields:
func (_m *mockValueWithAllFeatures) HasFormula() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// HasIdentity provides a mock function with given fields:
func (_m *mockValueWithAllFeatures) HasIdentity() bool {
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
func (_m *mockValueWithAllFeatures) Identify() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// IsConditional provides a mock function with given fields:
func (_m *mockValueWithAllFeatures) IsConditional() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsNil provides a mock function with given fields:
func (_m *mockValueWithAllFeatures) IsNil() bool {
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
func (_m *mockValueWithAllFeatures) SelfReplaceIfNil() Value {
	ret := _m.Called()

	var r0 Value
	if rf, ok := ret.Get(0).(func() Value); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Value)
		}
	}

	return r0
}

// SetAlias provides a mock function with given fields: _a0
func (_m *mockValueWithAllFeatures) SetAlias(_a0 string) {
	_m.Called(_a0)
}

// Stringify provides a mock function with given fields:
func (_m *mockValueWithAllFeatures) Stringify() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Tag provides a mock function with given fields: _a0
func (_m *mockValueWithAllFeatures) Tag(_a0 ...Tag) {
	_va := make([]interface{}, len(_a0))
	for _i := range _a0 {
		_va[_i] = _a0[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// ToSyntaxOperand provides a mock function with given fields: nextOp
func (_m *mockValueWithAllFeatures) ToSyntaxOperand(nextOp Op) *SyntaxOperand {
	ret := _m.Called(nextOp)

	var r0 *SyntaxOperand
	if rf, ok := ret.Get(0).(func(Op) *SyntaxOperand); ok {
		r0 = rf(nextOp)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*SyntaxOperand)
		}
	}

	return r0
}

type mockConstructorTestingTnewMockValueWithAllFeatures interface {
	mock.TestingT
	Cleanup(func())
}

// newMockValueWithAllFeatures creates a new instance of mockValueWithAllFeatures. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newMockValueWithAllFeatures(t mockConstructorTestingTnewMockValueWithAllFeatures) *mockValueWithAllFeatures {
	mock := &mockValueWithAllFeatures{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
