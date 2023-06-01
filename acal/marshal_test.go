package acal

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestPerformStandardValueExtraction(t *testing.T) {
	scenarios := []struct {
		desc string
		test func(t *testing.T)
	}{
		{
			desc: "nil Value",
			test: func(t *testing.T) {
				cache := NewMockIValueCache(t)

				valueOpsMock, cleanup := MockValueOps(t)
				defer cleanup()

				valueOpsMock.On("IsNilValue", nil).Return(true).Once()

				result := PerformStandardValueExtraction(nil, cache)

				assert.Equal(t, cache, result)
				cache.AssertNotCalled(t, "Take", mock.Anything)
			},
		},
		{
			desc: "already taken anchored Value",
			test: func(t *testing.T) {
				cacheMock := NewMockIValueCache(t)

				aValMock := newMockValueWithAllFeatures(t)
				aValMock.On("IsConditional").Return(false).Maybe()
				aValMock.On("HasFormula").Return(false).Maybe()

				valueOpsMock, cleanup := MockValueOps(t)
				defer cleanup()

				valueOpsMock.On("IsNilValue", aValMock).Return(false).Once()
				valueOpsMock.On("HasIdentity", aValMock).Return(true).Once()

				// Already taken Value
				cacheMock.On("Take", aValMock).Return(false).Once()

				result := PerformStandardValueExtraction(aValMock, cacheMock)

				assert.Equal(t, cacheMock, result)
				aValMock.AssertNotCalled(t, "IsConditional")
				aValMock.AssertNotCalled(t, "HasFormula")
			},
		},
		{
			desc: "unique anchored Value with condition",
			test: func(t *testing.T) {
				cacheMock := NewMockIValueCache(t)

				criteriaMock := NewMockTypedValue[bool](t)
				criteriaMock.On("ExtractValues", cacheMock).Return(cacheMock).Once()

				aValMock := newMockValueWithAllFeatures(t)
				aValMock.On("IsConditional").Return(true).Once()
				aValMock.On("GetCondition").Return(NewCondition(criteriaMock)).Once()
				aValMock.On("GetTags").Return(nil).Once()
				aValMock.On("HasFormula").Return(false).Once()
				aValMock.On("GetFormulaFn").Return(nil).Maybe()

				valueOpsMock, cleanup := MockValueOps(t)
				defer cleanup()

				valueOpsMock.On("IsNilValue", aValMock).Return(false).Once()
				valueOpsMock.On("HasIdentity", aValMock).Return(true).Once()

				// New Value
				cacheMock.On("Take", aValMock).Return(true).Once()

				result := PerformStandardValueExtraction(aValMock, cacheMock)

				assert.Equal(t, cacheMock, result)
				aValMock.AssertNotCalled(t, "GetFormulaFn")
			},
		},
		{
			desc: "unique anchored Value with formula",
			test: func(t *testing.T) {
				cacheMock := NewMockIValueCache(t)

				aValMock := newMockValueWithAllFeatures(t)
				aValMock.On("IsConditional").Return(false).Once()
				aValMock.On("GetCondition").Return(nil).Maybe()
				aValMock.On("GetTags").Return(nil).Once()
				aValMock.On("HasFormula").Return(true).Once()

				mockOperand1 := newMockValueWithAllFeatures(t)
				mockOperand1.On("ExtractValues", cacheMock).Return(cacheMock).Once()

				mockOperand2 := newMockValueWithAllFeatures(t)
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

				valueOpsMock, cleanup := MockValueOps(t)
				defer cleanup()

				valueOpsMock.On("IsNilValue", aValMock).Return(false).Once()
				valueOpsMock.On("HasIdentity", aValMock).Return(true).Once()

				// New Value
				cacheMock.On("Take", aValMock).Return(true).Once()

				result := PerformStandardValueExtraction(aValMock, cacheMock)

				assert.Equal(t, cacheMock, result)
				aValMock.AssertNotCalled(t, "GetCondition")
			},
		},
		{
			desc: "unique anchored Value with tags containing an Value",
			test: func(t *testing.T) {
				cacheMock := NewMockIValueCache(t)

				tag := Tag{Name: "TestName", Value: 5}

				aValMock := newMockValueWithAllFeatures(t)
				aValMock.On("IsConditional").Return(false).Once()
				aValMock.On("GetCondition").Return(nil).Maybe()
				aValMock.On("GetTags").Return(Tags{tag}).Once()
				aValMock.On("HasFormula").Return(false).Once()
				aValMock.On("GetFormulaFn").Return(nil).Maybe()

				valueOpsMock, cleanup := MockValueOps(t)
				defer cleanup()

				valueOpsMock.On("IsNilValue", aValMock).Return(false).Once()
				valueOpsMock.On("HasIdentity", aValMock).Return(true).Once()

				// New Value
				cacheMock.On("Take", aValMock).Return(true).Once()

				result := PerformStandardValueExtraction(aValMock, cacheMock)

				assert.Equal(t, cacheMock, result)
				aValMock.AssertNotCalled(t, "GetCondition")
			},
		},
	}

	for _, sc := range scenarios {
		t.Run(sc.desc, sc.test)
	}
}
