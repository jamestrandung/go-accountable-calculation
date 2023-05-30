package acal

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestFormatFloatForMarshalling(t *testing.T) {
	scenarios := []struct {
		desc string
		f    float64
		want string
	}{
		{
			desc: "small float",
			f:    24.2,
			want: "24.2",
		},
		{
			desc: "huge float",
			f:    17976931348623157000000000000000000000000000000000000000000000000000000000,
			want: "1.7976931348623158e+73",
		},
	}

	for _, scenario := range scenarios {
		sc := scenario
		t.Run(
			sc.desc, func(t *testing.T) {
				actual := FormatFloatForMarshalling(sc.f)

				assert.Equal(t, sc.want, actual)
			},
		)
	}
}

func TestPerformStandardValueExtraction(t *testing.T) {
	valueOpsMock, cleanup := MockValueOps()
	defer cleanup()

	scenarios := []struct {
		desc string
		test func(t *testing.T)
	}{
		{
			desc: "nil Value",
			test: func(t *testing.T) {
				cache := &MockIValueCache{}

				valueOpsMock.On("IsNilValue", nil).Return(true).Once()

				result := PerformStandardValueExtraction(nil, cache)

				assert.Equal(t, cache, result)
				cache.AssertNotCalled(t, "Take", mock.Anything)
				mock.AssertExpectationsForObjects(t, valueOpsMock)
			},
		},
		{
			desc: "already taken anchored Value",
			test: func(t *testing.T) {
				cacheMock := &MockIValueCache{}

				aValMock := &mockValueWithAllFeatures{}
				aValMock.On("IsConditional").Return(false).Maybe()
				aValMock.On("HasFormula").Return(false).Maybe()

				valueOpsMock.On("IsNilValue", aValMock).Return(false).Once()
				valueOpsMock.On("IsAnchored", aValMock).Return(true).Once()

				// Already taken Value
				cacheMock.On("Take", aValMock).Return(false).Once()

				result := PerformStandardValueExtraction(aValMock, cacheMock)

				assert.Equal(t, cacheMock, result)
				aValMock.AssertNotCalled(t, "IsConditional")
				aValMock.AssertNotCalled(t, "HasFormula")
				mock.AssertExpectationsForObjects(t, aValMock, valueOpsMock, cacheMock)
			},
		},
		{
			desc: "unique anchored Value with condition",
			test: func(t *testing.T) {
				cacheMock := &MockIValueCache{}

				criteriaMock := &MockTypedValue[bool]{}
				criteriaMock.On("ExtractValues", cacheMock).Return(cacheMock).Once()

				aValMock := &mockValueWithAllFeatures{}
				aValMock.On("IsConditional").Return(true).Once()
				aValMock.On("GetCondition").Return(NewCondition(criteriaMock)).Once()
				aValMock.On("GetTags").Return(nil).Once()
				aValMock.On("HasFormula").Return(false).Once()
				aValMock.On("GetFormulaFn").Return(nil).Maybe()

				valueOpsMock.On("IsNilValue", aValMock).Return(false).Once()
				valueOpsMock.On("IsAnchored", aValMock).Return(true).Once()

				// New Value
				cacheMock.On("Take", aValMock).Return(true).Once()

				result := PerformStandardValueExtraction(aValMock, cacheMock)

				assert.Equal(t, cacheMock, result)
				aValMock.AssertNotCalled(t, "GetFormulaFn")
				mock.AssertExpectationsForObjects(t, aValMock, valueOpsMock, cacheMock, criteriaMock)
			},
		},
		{
			desc: "unique anchored Value with formula",
			test: func(t *testing.T) {
				cacheMock := &MockIValueCache{}

				aValMock := &mockValueWithAllFeatures{}
				aValMock.On("IsConditional").Return(false).Once()
				aValMock.On("GetCondition").Return(nil).Maybe()
				aValMock.On("GetTags").Return(nil).Once()
				aValMock.On("HasFormula").Return(true).Once()

				mockOperand1 := &mockValueWithAllFeatures{}
				mockOperand1.On("ExtractValues", cacheMock).Return(cacheMock).Once()

				mockOperand2 := &mockValueWithAllFeatures{}
				mockOperand2.On("ExtractValues", cacheMock).Return(cacheMock).Once()

				var dummyOpCategory OpCategory = 99

				aValMock.On("GetFormulaFn").Return(
					func() *SyntaxNode {
						return NewSyntaxNode(
							dummyOpCategory, OpTransparent, "TestOpDesc", []any{
								mockOperand1,
								mockOperand2,
								"staticValue",
							},
						)
					},
				).Maybe()

				valueOpsMock.On("IsNilValue", aValMock).Return(false).Once()
				valueOpsMock.On("IsAnchored", aValMock).Return(true).Once()

				// New Value
				cacheMock.On("Take", aValMock).Return(true).Once()

				result := PerformStandardValueExtraction(aValMock, cacheMock)

				assert.Equal(t, cacheMock, result)
				aValMock.AssertNotCalled(t, "GetCondition")
				mock.AssertExpectationsForObjects(t, aValMock, valueOpsMock, cacheMock, mockOperand1, mockOperand2)
			},
		},
		{
			desc: "unique anchored Value with tags containing an Value",
			test: func(t *testing.T) {
				cacheMock := &MockIValueCache{}

				tag := Tag{Name: "TestName", Value: 5}

				aValMock := &mockValueWithAllFeatures{}
				aValMock.On("IsConditional").Return(false).Once()
				aValMock.On("GetCondition").Return(nil).Maybe()
				aValMock.On("GetTags").Return(Tags{tag}).Once()
				aValMock.On("HasFormula").Return(false).Once()
				aValMock.On("GetFormulaFn").Return(nil).Maybe()

				valueOpsMock.On("IsNilValue", aValMock).Return(false).Once()
				valueOpsMock.On("IsAnchored", aValMock).Return(true).Once()

				// New Value
				cacheMock.On("Take", aValMock).Return(true).Once()

				result := PerformStandardValueExtraction(aValMock, cacheMock)

				assert.Equal(t, cacheMock, result)
				aValMock.AssertNotCalled(t, "GetCondition")
				mock.AssertExpectationsForObjects(t, aValMock, valueOpsMock, cacheMock)
			},
		},
	}

	for _, scenario := range scenarios {
		sc := scenario
		t.Run(
			sc.desc, func(t *testing.T) {
				sc.test(t)
			},
		)
	}
}
