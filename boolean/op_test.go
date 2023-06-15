package boolean

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIf(t *testing.T) {
	scenarios := []struct {
		desc string
		test func(t *testing.T)
	}{
		{
			desc: "nil value",
			test: func(t *testing.T) {
				mockBool := NewMockInterface(t)
				mockBool.On("IsNil").
					Return(true).
					Once()

				var dummy int
				doIfFn := func(b Interface) {
					assert.Equal(t, NilBool, b)
					dummy = 1
				}

				actual := If(mockBool, doIfFn)

				assert.False(t, actual)
				assert.Equal(t, 0, dummy)
			},
		},
		{
			desc: "false value",
			test: func(t *testing.T) {
				mockBool := NewMockInterface(t)
				mockBool.On("IsNil").
					Return(false).
					Once()
				mockBool.On("Bool").
					Return(false).
					Once()

				var dummy int
				doIfFn := func(b Interface) {
					assert.Equal(t, mockBool, b)
					dummy = 1
				}

				actual := If(mockBool, doIfFn)

				assert.False(t, actual)
				assert.Equal(t, 0, dummy)
			},
		},
		{
			desc: "true value",
			test: func(t *testing.T) {
				mockBool := NewMockInterface(t)
				mockBool.On("IsNil").
					Return(false).
					Once()
				mockBool.On("Bool").
					Return(true).
					Once()

				var dummy int
				doIfFn := func(b Interface) {
					assert.Equal(t, mockBool, b)
					dummy = 1
				}

				actual := If(mockBool, doIfFn)

				assert.True(t, actual)
				assert.Equal(t, 1, dummy)
			},
		},
	}

	for _, sc := range scenarios {
		t.Run(sc.desc, sc.test)
	}
}

func TestIfNot(t *testing.T) {
	scenarios := []struct {
		desc string
		test func(t *testing.T)
	}{
		{
			desc: "nil value",
			test: func(t *testing.T) {
				mockBool := NewMockInterface(t)
				mockBool.On("IsNil").
					Return(true).
					Once()

				var dummy int
				doIfFn := func(b Interface) {
					assert.True(t, b.Bool())
					dummy = 1
				}

				actual := IfNot(mockBool, doIfFn)

				assert.True(t, actual)
				assert.Equal(t, 1, dummy)
			},
		},
		{
			desc: "false value",
			test: func(t *testing.T) {
				dummyNot := MakeSimple("NotFalse", true)

				mockBool := NewMockInterface(t)
				mockBool.On("IsNil").
					Return(false).
					Once()
				mockBool.On("Bool").
					Return(false).
					Once()
				mockBool.On("Not").
					Return(dummyNot).
					Once()

				var dummy int
				doIfFn := func(b Interface) {
					assert.Equal(t, dummyNot, b)
					dummy = 1
				}

				actual := IfNot(mockBool, doIfFn)

				assert.True(t, actual)
				assert.Equal(t, 1, dummy)
			},
		},
		{
			desc: "true value",
			test: func(t *testing.T) {
				mockBool := NewMockInterface(t)
				mockBool.On("IsNil").
					Return(false).
					Once()
				mockBool.On("Bool").
					Return(true).
					Once()

				var dummy int
				doIfFn := func(b Interface) {
					assert.Equal(t, mockBool, b)
					dummy = 1
				}

				actual := IfNot(mockBool, doIfFn)

				assert.False(t, actual)
				assert.Equal(t, 0, dummy)
			},
		},
	}

	for _, sc := range scenarios {
		t.Run(sc.desc, sc.test)
	}
}

func TestIfElse(t *testing.T) {
	scenarios := []struct {
		desc string
		test func(t *testing.T)
	}{
		{
			desc: "nil value",
			test: func(t *testing.T) {
				mockBool := NewMockInterface(t)
				mockBool.On("IsNil").
					Return(true).
					Once()

				var dummy int
				doIfFn := func(b Interface) {
					dummy = 1
				}
				doElseFn := func(b Interface) {
					assert.True(t, b.Bool())
					dummy = 2
				}

				actual := IfElse(mockBool, doIfFn, doElseFn)

				assert.False(t, actual)
				assert.Equal(t, 2, dummy)
			},
		},
		{
			desc: "false value",
			test: func(t *testing.T) {
				dummyNot := MakeSimple("NotFalse", true)

				mockBool := NewMockInterface(t)
				mockBool.On("IsNil").
					Return(false).
					Once()
				mockBool.On("Bool").
					Return(false).
					Once()
				mockBool.On("Not").
					Return(dummyNot).
					Once()

				var dummy int
				doIfFn := func(b Interface) {
					dummy = 1
				}
				doElseFn := func(b Interface) {
					assert.Equal(t, dummyNot, b)
					dummy = 2
				}

				actual := IfElse(mockBool, doIfFn, doElseFn)

				assert.False(t, actual)
				assert.Equal(t, 2, dummy)
			},
		},
		{
			desc: "true value",
			test: func(t *testing.T) {
				mockBool := NewMockInterface(t)
				mockBool.On("IsNil").
					Return(false).
					Once()
				mockBool.On("Bool").
					Return(true).
					Once()

				var dummy int
				doIfFn := func(b Interface) {
					assert.Equal(t, mockBool, b)
					dummy = 1
				}
				doElseFn := func(b Interface) {
					dummy = 2
				}

				actual := IfElse(mockBool, doIfFn, doElseFn)

				assert.True(t, actual)
				assert.Equal(t, 1, dummy)
			},
		},
	}

	for _, sc := range scenarios {
		t.Run(sc.desc, sc.test)
	}
}
