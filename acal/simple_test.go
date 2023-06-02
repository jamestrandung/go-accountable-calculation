package acal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimple_IsNil(t *testing.T) {
	var nilSimple *Simple[int]

	assert.True(t, nilSimple.IsNil())

	simple := NewSimple("Simple", 2)

	assert.False(t, simple.IsNil())
}

func TestSimple_ToSyntaxOperand(t *testing.T) {
	scenarios := []struct {
		desc string
		test func(t *testing.T)
	}{
		{
			desc: "anchored value",
			test: func(t *testing.T) {
				simple := NewSimple("Simple", 1)

				operand := simple.ToSyntaxOperand(OpTransparent)

				assert.Equal(t, NewSyntaxOperand(simple), operand)
			},
		},
		{
			desc: "nil value",
			test: func(t *testing.T) {
				var simple *Simple[int]

				operand := simple.ToSyntaxOperand(OpTransparent)

				assert.Equal(t, NewSyntaxOperandWithStaticValue(Describe(simple)), operand)
			},
		},
		{
			desc: "value with no formula",
			test: func(t *testing.T) {
				simple := &Simple[int]{
					value: 1,
				}

				operand := simple.ToSyntaxOperand(OpTransparent)

				assert.Equal(t, NewSyntaxOperandWithStaticValue(Describe(simple)), operand)
			},
		},
		{
			desc: "value with formula",
			test: func(t *testing.T) {
				simple := NewSimpleWithFormula(1, &SyntaxNode{})

				operand := simple.ToSyntaxOperand(OpTransparent)

				assert.Nil(t, operand)
			},
		},
	}

	for _, sc := range scenarios {
		t.Run(sc.desc, sc.test)
	}
}

func TestSimple_ExtractValues(t *testing.T) {
	defer func(original func(v Value, cache IValueCache) IValueCache) {
		PerformStandardValueExtraction = original
	}(PerformStandardValueExtraction)

	simple := NewSimple("Simple", 1)
	mockCache := NewMockIValueCache(t)

	PerformStandardValueExtraction = func(v Value, cache IValueCache) IValueCache {
		assert.Equal(t, simple, v)
		assert.Equal(t, mockCache, cache)

		return cache
	}

	actual := simple.ExtractValues(mockCache)

	assert.Equal(t, mockCache, actual)
}

func TestSimple_SelfReplaceIfNil(t *testing.T) {
	var nilSimple *Simple[int]
	assert.Equal(t, ZeroSimple[int]("NilSimple"), nilSimple.SelfReplaceIfNil())

	simple := NewSimple("Simple", 1)
	assert.Equal(t, simple, simple.SelfReplaceIfNil())
}

func TestSimple_DoAnchor(t *testing.T) {
	scenarios := []struct {
		desc string
		test func(t *testing.T)
	}{
		{
			desc: "anchored value",
			test: func(t *testing.T) {
				simple := NewSimple("Simple", 1)

				actual, isNew := simple.DoAnchor("Something")
				assert.Equal(t, "Something", actual.GetName())
				assert.True(t, isNew)
				assert.NotEqual(t, simple, actual)
			},
		},
		{
			desc: "non-anchored value",
			test: func(t *testing.T) {
				simple := NewSimpleFrom[int](NewConstant(1))

				actual, isNew := simple.DoAnchor("Something")
				assert.Equal(t, "Something", actual.GetName())
				assert.False(t, isNew)
				assert.Equal(t, simple, actual)
			},
		},
	}

	for _, sc := range scenarios {
		t.Run(sc.desc, sc.test)
	}
}

func TestSimple_MarshallJSON(t *testing.T) {
	scenarios := []struct {
		desc string
		test func(t *testing.T)
	}{
		{
			desc: "nil Simple",
			test: func(t *testing.T) {
				var nilSimple *Simple[int]

				actual, err := nilSimple.MarshalJSON()
				assert.Nil(t, actual)
				assert.Nil(t, err)
			},
		},
		{
			desc: "Simple with formula",
			test: func(t *testing.T) {
				dummyInt := NewSimple("DummyInt", 1)
				dummyBool := NewSimple("DummyBool", true)

				toMarshall := NewSimpleFrom[bool](dummyBool)
				toMarshall.Tag(NewTagFrom(dummyInt))
				toMarshall.From(SourceHardcode) // Should be ignored

				wanted := `{"Value":"true","Source":"static_calculation","Tags":{"DummyInt":{"Value":"1","IsValue":true}},"Formula":{"Category":"AssignVariable","Operands":[{"Name":"DummyBool"}]}}`

				actual, err := toMarshall.MarshalJSON()
				assert.Equal(t, wanted, string(actual))
				assert.Nil(t, err)
			},
		},
		{
			desc: "Simple with NO formula",
			test: func(t *testing.T) {
				dummyInt := NewSimple("DummyInt", 1)

				toMarshall := NewSimple("DummyBool", true)
				toMarshall.Tag(NewTagFrom(dummyInt))
				toMarshall.From(SourceHardcode)
				toMarshall.AddCondition(NewCondition(NewSimple("Criteria", true)))

				wanted := `{"Value":"true","Source":"hardcode","Tags":{"DummyInt":{"Value":"1","IsValue":true}},"Condition":{"Formula":{"Category":"AssignVariable","Operands":[{"Name":"Criteria"}]}}}`

				actual, err := toMarshall.MarshalJSON()
				assert.Equal(t, wanted, string(actual))
				assert.Nil(t, err)
			},
		},
	}

	for _, sc := range scenarios {
		t.Run(sc.desc, sc.test)
	}
}
