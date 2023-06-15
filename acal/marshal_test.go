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

                mockValueOps, cleanup := MockValueOps(t)
                defer cleanup()

                mockValueOps.On("IsNilValue", nil).Return(true).Once()

                result := PerformStandardValueExtraction(nil, cache)

                assert.Equal(t, cache, result)
                cache.AssertNotCalled(t, "Take", mock.Anything)
            },
        },
        {
            desc: "already taken anchored Value",
            test: func(t *testing.T) {
                mockCache := NewMockIValueCache(t)

                mockValue := newMockValueWithAllFeatures(t)
                mockValue.On("IsConditional").Return(false).Maybe()
                mockValue.On("HasFormula").Return(false).Maybe()

                mockValueOps, cleanup := MockValueOps(t)
                defer cleanup()

                mockValueOps.On("IsNilValue", mockValue).Return(false).Once()
                mockValueOps.On("HasIdentity", mockValue).Return(true).Once()

                // Already taken Value
                mockCache.On("Take", mockValue).Return(false).Once()

                result := PerformStandardValueExtraction(mockValue, mockCache)

                assert.Equal(t, mockCache, result)
                mockValue.AssertNotCalled(t, "IsConditional")
                mockValue.AssertNotCalled(t, "HasFormula")
            },
        },
        {
            desc: "unique anchored Value with condition",
            test: func(t *testing.T) {
                mockCache := NewMockIValueCache(t)

                mockCriteria := NewMockTypedValue[bool](t)
                mockCriteria.On("ExtractValues", mockCache).Return(mockCache).Once()

                mockValue := newMockValueWithAllFeatures(t)
                mockValue.On("IsConditional").Return(true).Once()
                mockValue.On("GetCondition").Return(NewCondition(mockCriteria)).Once()
                mockValue.On("GetTags").Return(nil).Once()
                mockValue.On("HasFormula").Return(false).Once()
                mockValue.On("GetFormula").Return(nil).Maybe()

                mockValueOps, cleanup := MockValueOps(t)
                defer cleanup()

                mockValueOps.On("IsNilValue", mockValue).Return(false).Once()
                mockValueOps.On("HasIdentity", mockValue).Return(true).Once()

                // New Value
                mockCache.On("Take", mockValue).Return(true).Once()

                result := PerformStandardValueExtraction(mockValue, mockCache)

                assert.Equal(t, mockCache, result)
                mockValue.AssertNotCalled(t, "GetFormula")
            },
        },
        {
            desc: "unique anchored Value with formula",
            test: func(t *testing.T) {
                mockCache := NewMockIValueCache(t)

                mockValue := newMockValueWithAllFeatures(t)
                mockValue.On("IsConditional").Return(false).Once()
                mockValue.On("GetCondition").Return(nil).Maybe()
                mockValue.On("GetTags").Return(nil).Once()
                mockValue.On("HasFormula").Return(true).Once()

                mockOperand1 := newMockValueWithAllFeatures(t)
                mockOperand1.On("ExtractValues", mockCache).Return(mockCache).Once()

                mockOperand2 := newMockValueWithAllFeatures(t)
                mockOperand2.On("ExtractValues", mockCache).Return(mockCache).Once()

                var dummyOpCategory OpCategory = 99

                mockValue.On("GetFormulaFn").Return(
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

                mockValueOps, cleanup := MockValueOps(t)
                defer cleanup()

                mockValueOps.On("IsNilValue", mockValue).Return(false).Once()
                mockValueOps.On("HasIdentity", mockValue).Return(true).Once()

                // New Value
                mockCache.On("Take", mockValue).Return(true).Once()

                result := PerformStandardValueExtraction(mockValue, mockCache)

                assert.Equal(t, mockCache, result)
                mockValue.AssertNotCalled(t, "GetCondition")
            },
        },
        {
            desc: "unique anchored Value with tags containing an Value",
            test: func(t *testing.T) {
                mockCache := NewMockIValueCache(t)

                tag := Tag{Name: "TestName", Value: 5}

                mockValue := newMockValueWithAllFeatures(t)
                mockValue.On("IsConditional").Return(false).Once()
                mockValue.On("GetCondition").Return(nil).Maybe()
                mockValue.On("GetTags").Return(Tags{tag}).Once()
                mockValue.On("HasFormula").Return(false).Once()
                mockValue.On("GetFormulaFn").Return(nil).Maybe()

                mockValueOps, cleanup := MockValueOps(t)
                defer cleanup()

                mockValueOps.On("IsNilValue", mockValue).Return(false).Once()
                mockValueOps.On("HasIdentity", mockValue).Return(true).Once()

                // New Value
                mockCache.On("Take", mockValue).Return(true).Once()

                result := PerformStandardValueExtraction(mockValue, mockCache)

                assert.Equal(t, mockCache, result)
                mockValue.AssertNotCalled(t, "GetCondition")
            },
        },
    }

    for _, sc := range scenarios {
        t.Run(sc.desc, sc.test)
    }
}
