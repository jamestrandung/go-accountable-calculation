package acal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// MockValueOps can be used in tests to mock IValueOps.
func MockValueOps(t *testing.T) (*MockIValueOps, func()) {
	old := valueOps
	mock := NewMockIValueOps(t)

	valueOps = mock
	return mock, func() {
		valueOps = old
	}
}

func TestValueOpsImpl_IsNilValue(t *testing.T) {
	mockValue := newMockValueWithFormula(t)

	scenarios := []struct {
		desc  string
		aVal  Value
		setup func()
		want  bool
	}{
		{
			desc: "nil Value",
			aVal: nil,
			want: true,
		},
		{
			desc: "Value with nil value",
			aVal: mockValue,
			setup: func() {
				mockValue.On("IsNil").Return(true).Once()
			},
			want: true,
		},
		{
			desc: "Value with non-nil value",
			aVal: mockValue,
			setup: func() {
				mockValue.On("IsNil").Return(false).Once()
			},
			want: false,
		},
	}

	for _, scenario := range scenarios {
		sc := scenario
		t.Run(
			sc.desc, func(t *testing.T) {
				if sc.setup != nil {
					sc.setup()
				}

				ops := valueOpsImpl{}

				actual := ops.IsNilValue(sc.aVal)
				assert.Equal(t, sc.want, actual)
			},
		)
	}
}

func TestValueOpsImpl_HasIdentity(t *testing.T) {
	mockValue := newMockValueWithFormula(t)

	mockValueOps, cleanup := MockValueOps(t)
	defer cleanup()

	scenarios := []struct {
		desc  string
		setup func()
		want  bool
	}{
		{
			desc: "nil Value",
			setup: func() {
				mockValueOps.On("IsNilValue", mockValue).
					Return(true).
					Once()
			},
			want: false,
		},
		{
			desc: "Value with no identity",
			setup: func() {
				mockValueOps.On("IsNilValue", mockValue).
					Return(false).
					Once()

				mockValue.On("HasIdentity").Return(false).Once()
			},
			want: false,
		},
		{
			desc: "Value with identity",
			setup: func() {
				mockValueOps.On("IsNilValue", mockValue).
					Return(false).
					Once()

				mockValue.On("HasIdentity").Return(true).Once()
			},
			want: true,
		},
	}

	for _, scenario := range scenarios {
		sc := scenario
		t.Run(
			sc.desc, func(t *testing.T) {
				if sc.setup != nil {
					sc.setup()
				}

				ops := valueOpsImpl{}

				actual := ops.HasIdentity(mockValue)
				assert.Equal(t, sc.want, actual)
			},
		)
	}
}

func TestValueOpsImpl_Identify(t *testing.T) {
	mockValue := newMockValueWithFormula(t)

	mockValueOps, cleanup := MockValueOps(t)
	defer cleanup()

	scenarios := []struct {
		desc  string
		setup func()
		want  string
	}{
		{
			desc: "nil Value",
			setup: func() {
				mockValueOps.On("IsNilValue", mockValue).
					Return(true).
					Once()
			},
			want: UnknownValueName,
		},
		{
			desc: "Value with no identity",
			setup: func() {
				mockValueOps.On("IsNilValue", mockValue).
					Return(false).
					Once()

				mockValue.On("Identify").Return("").Once()
			},
			want: "",
		},
		{
			desc: "Value with identity",
			setup: func() {
				mockValueOps.On("IsNilValue", mockValue).
					Return(false).
					Once()

				mockValue.On("Identify").Return("identity").Once()
			},
			want: "identity",
		},
	}

	for _, scenario := range scenarios {
		sc := scenario
		t.Run(
			sc.desc, func(t *testing.T) {
				if sc.setup != nil {
					sc.setup()
				}

				ops := valueOpsImpl{}

				actual := ops.Identify(mockValue)
				assert.Equal(t, sc.want, actual)
			},
		)
	}
}

func TestValueOpsImpl_Stringify(t *testing.T) {
	scenarios := []struct {
		desc string
		test func(t *testing.T)
	}{
		{
			desc: "nil Value",
			test: func(t *testing.T) {
				mockValue := newMockValueWithFormula(t)

				mockValueOps, cleanup := MockValueOps(t)
				defer cleanup()

				mockValueOps.On("IsNilValue", mockValue).
					Return(true).
					Once()

				ops := valueOpsImpl{}

				actual := ops.Stringify(mockValue)
				assert.Equal(t, "", actual)
			},
		},
		{
			desc: "non-nil Value",
			test: func(t *testing.T) {
				mockValue := newMockValueWithFormula(t)
				mockValue.On("Stringify").Return("5").Once()

				mockValueOps, cleanup := MockValueOps(t)
				defer cleanup()

				mockValueOps.On("IsNilValue", mockValue).
					Return(false).
					Once()

				ops := valueOpsImpl{}

				actual := ops.Stringify(mockValue)
				assert.Equal(t, "5", actual)
			},
		},
	}

	for _, sc := range scenarios {
		t.Run(sc.desc, sc.test)
	}
}

func TestValueOpsImpl_Describe(t *testing.T) {
	scenarios := []struct {
		desc string
		test func(t *testing.T)
	}{
		{
			desc: "nil Value",
			test: func(t *testing.T) {
				mockValue := newMockValueWithFormula(t)

				mockValueOps, cleanup := MockValueOps(t)
				defer cleanup()

				mockValueOps.On("IsNilValue", mockValue).
					Return(true).
					Once()

				ops := valueOpsImpl{}

				actual := ops.Describe(mockValue)
				assert.Equal(t, "?[?]", actual)
			},
		},
		{
			desc: "Value with empty identity",
			test: func(t *testing.T) {
				mockValue := newMockValueWithFormula(t)
				mockValue.On("Stringify").Return("5").Once()

				mockValueOps, cleanup := MockValueOps(t)
				defer cleanup()

				mockValueOps.On("IsNilValue", mockValue).
					Return(false).
					Once()
				mockValueOps.On("Identify", mockValue).
					Return("").
					Once()

				ops := valueOpsImpl{}

				actual := ops.Describe(mockValue)
				assert.Equal(t, "?[5]", actual)
			},
		},
		{
			desc: "Value with non-empty identity",
			test: func(t *testing.T) {
				mockValue := newMockValueWithFormula(t)
				mockValue.On("Stringify").Return("true").Once()

				mockValueOps, cleanup := MockValueOps(t)
				defer cleanup()

				mockValueOps.On("IsNilValue", mockValue).
					Return(false).
					Once()
				mockValueOps.On("Identify", mockValue).
					Return("TestIdentity").
					Once()

				ops := valueOpsImpl{}

				actual := ops.Describe(mockValue)
				assert.Equal(t, "TestIdentity[true]", actual)
			},
		},
	}

	for _, sc := range scenarios {
		t.Run(sc.desc, sc.test)
	}
}

func TestValueOpsImpl_DescribeValueAsFormula(t *testing.T) {
	mockValue := newMockValueWithFormula(t)
	dummyFormula := &SyntaxNode{}

	mockValueOps, cleanup := MockValueOps(t)
	defer cleanup()

	scenarios := []struct {
		desc  string
		setup func()
		want  *SyntaxNode
	}{
		{
			desc: "anchored Value",
			setup: func() {
				mockValueOps.On("HasIdentity", mockValue).
					Return(true).
					Once()
			},
			want: &SyntaxNode{
				category: OpCategoryAssignVariable,
				op:       OpTransparent,
				operands: []any{
					mockValue,
				},
			},
		},
		{
			desc: "un-anchored Value with no formula",
			setup: func() {
				mockValueOps.On("HasIdentity", mockValue).
					Return(false).
					Once()

				mockValueOps.On("Describe", mockValue).
					Return("TestDescription").
					Once()

				mockValue.On("HasFormula").Return(false).Once()
			},
			want: &SyntaxNode{
				category: OpCategoryAssignStatic,
				op:       OpTransparent,
				opDesc:   "TestDescription",
			},
		},
		{
			desc: "un-anchored Value with formula",
			setup: func() {
				mockValueOps.On("HasIdentity", mockValue).
					Return(false).
					Once()

				mockValue.On("HasFormula").Return(true).Once()
				mockValue.On("GetFormulaFn").Return(
					func() *SyntaxNode {
						return dummyFormula
					},
				).Once()
			},
			want: dummyFormula,
		},
	}

	for _, scenario := range scenarios {
		sc := scenario
		t.Run(
			sc.desc, func(t *testing.T) {
				if sc.setup != nil {
					sc.setup()
				}

				ops := valueOpsImpl{}

				actual := ops.DescribeValueAsFormula(mockValue)()
				assert.Equal(t, sc.want, actual)
			},
		)
	}
}

func TestExtractTypedValue(t *testing.T) {
	scenarios := []struct {
		desc string
		test func(t *testing.T)
	}{
		{
			desc: "nil TypedValue",
			test: func(t *testing.T) {
				mockValue := NewMockTypedValue[int](t)

				mockValueOps, cleanup := MockValueOps(t)
				defer cleanup()

				mockValueOps.On("IsNilValue", mockValue).
					Return(true).
					Once()

				actual := ExtractTypedValue[int](mockValue)

				assert.Equal(t, 0, actual)
			},
		},
		{
			desc: "non-nil TypedValue",
			test: func(t *testing.T) {
				mockValue := NewMockTypedValue[int](t)
				mockValue.On("GetTypedValue").
					Return(1).
					Once()

				mockValueOps, cleanup := MockValueOps(t)
				defer cleanup()

				mockValueOps.On("IsNilValue", mockValue).
					Return(false).
					Once()

				actual := ExtractTypedValue[int](mockValue)

				assert.Equal(t, 1, actual)
			},
		},
	}

	for _, sc := range scenarios {
		t.Run(sc.desc, sc.test)
	}
}
