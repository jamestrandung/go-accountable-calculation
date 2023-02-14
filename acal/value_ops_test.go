package acal

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

// MockValueOps can be used in tests to mock IValueOps.
func MockValueOps() (*MockIValueOps, func()) {
	old := valueOps
	mock := &MockIValueOps{}

	valueOps = mock
	return mock, func() {
		valueOps = old
	}
}

func TestValueOpsImpl_IsNilValue(t *testing.T) {
	aValMock := &mockValueWithFormula{}

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
			aVal: aValMock,
			setup: func() {
				aValMock.On("IsNil").Return(true).Once()
			},
			want: true,
		},
		{
			desc: "Value with non-nil value",
			aVal: aValMock,
			setup: func() {
				aValMock.On("IsNil").Return(false).Once()
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
				mock.AssertExpectationsForObjects(t, aValMock)
			},
		)
	}
}

func TestValueOpsImpl_IsAnchored(t *testing.T) {
	aValMock := &mockValueWithFormula{}

	valueOpsMock, cleanup := MockValueOps()
	defer cleanup()

	scenarios := []struct {
		desc  string
		setup func()
		want  bool
	}{
		{
			desc: "nil Value",
			setup: func() {
				valueOpsMock.On("IsNilValue", aValMock).
					Return(true).
					Once()
			},
			want: false,
		},
		{
			desc: "Value with empty name & empty alias",
			setup: func() {
				valueOpsMock.On("IsNilValue", aValMock).
					Return(false).
					Once()

				aValMock.On("GetName").Return("").Once()
				aValMock.On("GetAlias").Return("").Once()
			},
			want: false,
		},
		{
			desc: "Value with non-empty name",
			setup: func() {
				valueOpsMock.On("IsNilValue", aValMock).
					Return(false).
					Once()

				aValMock.On("GetName").Return("TestName").Once()
			},
			want: true,
		},
		{
			desc: "Value with empty name & non-empty alias",
			setup: func() {
				valueOpsMock.On("IsNilValue", aValMock).
					Return(false).
					Once()

				aValMock.On("GetName").Return("").Once()
				aValMock.On("GetAlias").Return("TestAlias").Once()
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

				actual := ops.IsAnchored(aValMock)
				assert.Equal(t, sc.want, actual)
				mock.AssertExpectationsForObjects(t, aValMock, valueOpsMock)
			},
		)
	}
}

func TestValueOpsImpl_Identify(t *testing.T) {
	aValMock := &mockValueWithFormula{}

	valueOpsMock, cleanup := MockValueOps()
	defer cleanup()

	scenarios := []struct {
		desc  string
		setup func()
		want  string
	}{
		{
			desc: "nil Value",
			setup: func() {
				valueOpsMock.On("IsNilValue", aValMock).
					Return(true).
					Once()
			},
			want: UnknownValueName,
		},
		{
			desc: "Value with empty name & empty alias",
			setup: func() {
				valueOpsMock.On("IsNilValue", aValMock).
					Return(false).
					Once()

				aValMock.On("GetName").Return("").Once()
				aValMock.On("GetAlias").Return("").Once()
			},
			want: "",
		},
		{
			desc: "Value with non-empty name & empty alias",
			setup: func() {
				valueOpsMock.On("IsNilValue", aValMock).
					Return(false).
					Once()

				aValMock.On("GetName").Return("TestName").Once()
				aValMock.On("GetAlias").Return("").Once()
			},
			want: "TestName",
		},
		{
			desc: "Value with empty name & non-empty alias",
			setup: func() {
				valueOpsMock.On("IsNilValue", aValMock).
					Return(false).
					Once()

				aValMock.On("GetAlias").Return("TestAlias").Twice()
			},
			want: "TestAlias",
		},
		{
			desc: "Value with non-empty name & non-empty alias",
			setup: func() {
				valueOpsMock.On("IsNilValue", aValMock).
					Return(false).
					Once()

				aValMock.On("GetAlias").Return("TestAlias").Twice()
			},
			want: "TestAlias",
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

				actual := ops.Identify(aValMock)
				assert.Equal(t, sc.want, actual)
				mock.AssertExpectationsForObjects(t, aValMock, valueOpsMock)
			},
		)
	}
}

func TestValueOpsImpl_Describe(t *testing.T) {
	aValMock := &mockValueWithFormula{}

	valueOpsMock, cleanup := MockValueOps()
	defer cleanup()

	scenarios := []struct {
		desc  string
		setup func()
		want  string
	}{
		{
			desc: "nil Value",
			setup: func() {
				valueOpsMock.On("IsNilValue", aValMock).
					Return(true).
					Once()
			},
			want: "?[?]",
		},
		{
			desc: "Value with empty identity",
			setup: func() {
				valueOpsMock.On("IsNilValue", aValMock).
					Return(false).
					Once()

				valueOpsMock.On("Identify", aValMock).
					Return("").
					Once()

				aValMock.On("GetValue").Return(5).Once()
			},
			want: "?[5]",
		},
		{
			desc: "Value with non-empty identity",
			setup: func() {
				valueOpsMock.On("IsNilValue", aValMock).
					Return(false).
					Once()

				valueOpsMock.On("Identify", aValMock).
					Return("TestIdentity").
					Once()

				aValMock.On("GetValue").Return(true).Once()
			},
			want: "TestIdentity[true]",
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

				actual := ops.Describe(aValMock)
				assert.Equal(t, sc.want, actual)
				mock.AssertExpectationsForObjects(t, aValMock, valueOpsMock)
			},
		)
	}
}

func TestValueOpsImpl_DescribeValueAsFormula(t *testing.T) {
	aValMock := &mockValueWithFormula{}
	dummyFormula := &SyntaxNode{}

	valueOpsMock, cleanup := MockValueOps()
	defer cleanup()

	scenarios := []struct {
		desc  string
		setup func()
		want  *SyntaxNode
	}{
		{
			desc: "anchored Value",
			setup: func() {
				valueOpsMock.On("IsAnchored", aValMock).
					Return(true).
					Once()
			},
			want: &SyntaxNode{
				category: OpCategoryAssignVariable,
				op:       OpTransparent,
				operands: []any{
					aValMock,
				},
			},
		},
		{
			desc: "un-anchored Value with no formula",
			setup: func() {
				valueOpsMock.On("IsAnchored", aValMock).
					Return(false).
					Once()

				valueOpsMock.On("Describe", aValMock).
					Return("TestDescription").
					Once()

				aValMock.On("HasFormula").Return(false).Once()
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
				valueOpsMock.On("IsAnchored", aValMock).
					Return(false).
					Once()

				aValMock.On("HasFormula").Return(true).Once()
				aValMock.On("GetFormulaFn").Return(
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

				actual := ops.DescribeValueAsFormula(aValMock)
				assert.Equal(t, sc.want, actual())
				mock.AssertExpectationsForObjects(t, aValMock, valueOpsMock)
			},
		)
	}
}
