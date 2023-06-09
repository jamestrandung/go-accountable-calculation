package acal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLocal_IsNil(t *testing.T) {
	var nilLocal *Local[int]

	assert.True(t, nilLocal.IsNil())

	local := NewLocal[int]("Local", NewSimple("Something", 2))

	assert.False(t, local.IsNil())
}

func TestLocal_ToSyntaxOperand(t *testing.T) {
	local := NewLocal[int]("Local", NewSimple("Something", 2))

	actual := local.ToSyntaxOperand(OpTransparent)

	assert.Equal(
		t, &SyntaxOperand{
			Name:  "Local",
			value: local,
		}, actual,
	)
}

func TestLocal_ExtractValues(t *testing.T) {
	defer func(original func(v Value, cache IValueCache) IValueCache) {
		PerformStandardValueExtraction = original
	}(PerformStandardValueExtraction)

	local := NewLocal[int]("Local", NewSimple("Something", 2))
	mockCache := NewMockIValueCache(t)

	PerformStandardValueExtraction = func(v Value, cache IValueCache) IValueCache {
		assert.Equal(t, local, v)
		assert.Equal(t, mockCache, cache)

		return cache
	}

	actual := local.ExtractValues(mockCache)

	assert.Equal(t, mockCache, actual)
}

func TestLocal_SelfReplaceIfNil(t *testing.T) {
	var nilLocal *Local[int]
	assert.Equal(t, ZeroSimple[int]("NilLocal"), nilLocal.SelfReplaceIfNil())

	local := NewLocal[int]("Local", NewSimple("Something", 2))
	assert.Equal(t, local, local.SelfReplaceIfNil())
}

func TestLocal_DoAnchor(t *testing.T) {
	local := NewLocal[int]("Local", NewSimple("Simple", 2))

	simple := local.DoAnchor("Something")

	assert.Equal(t, "Something", simple.GetName())
	assert.Equal(t, 2, simple.GetTypedValue())
}

func TestLocal_MarshallJSON(t *testing.T) {
	scenarios := []struct {
		desc string
		test func(t *testing.T)
	}{
		{
			desc: "nil Local",
			test: func(t *testing.T) {
				var nilLocal *Local[int]

				actual, err := nilLocal.MarshalJSON()
				assert.Nil(t, actual)
				assert.Nil(t, err)
			},
		},
		{
			desc: "non-nil Local",
			test: func(t *testing.T) {
				dummyInt := NewSimple("DummyInt", 1)

				toMarshall := NewLocal[int]("Local", NewSimple("Something", 2))
				toMarshall.Tag(NewTagFrom(dummyInt))
				toMarshall.AddCondition(NewCondition(NewSimple("Criteria", true)))

				wanted := `{"Value":"2","Source":"dependent_calculation","DependentField":"Something","Calculation":{"Something":{"Value":"2","Source":"unknown"}},"Tags":{"DummyInt":{"Value":"1","IsValue":true}},"Condition":{"Formula":{"Category":"AssignVariable","Operands":[{"Name":"Criteria"}]}}}`

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
