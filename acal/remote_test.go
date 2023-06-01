package acal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemote_IsNil(t *testing.T) {
	var nilRemote *Remote[int]

	assert.True(t, nilRemote.IsNil())

	remote := NewRemote("Remote", 2, "RemoteName", "LogKey")

	assert.False(t, remote.IsNil())
}

func TestRemote_GetTypedValue(t *testing.T) {
	var nilRemote *Remote[int]

	assert.Equal(t, 0, nilRemote.GetTypedValue())

	remote := NewRemote("Remote", 2, "RemoteName", "LogKey")

	assert.Equal(t, 2, remote.GetTypedValue())
}

func TestRemote_ToSyntaxOperand(t *testing.T) {
	remote := NewRemote("Remote", 2, "RemoteName", "LogKey")

	actual := remote.ToSyntaxOperand(OpTransparent)

	assert.Equal(
		t, &SyntaxOperand{
			Name: "Remote",
		}, actual,
	)
}

func TestRemote_ExtractValues(t *testing.T) {
	defer func(original func(v Value, cache IValueCache) IValueCache) {
		PerformStandardValueExtraction = original
	}(PerformStandardValueExtraction)

	remote := NewRemote("Remote", 2, "RemoteName", "LogKey")
	mockCache := NewMockIValueCache(t)

	PerformStandardValueExtraction = func(v Value, cache IValueCache) IValueCache {
		assert.Equal(t, remote, v)
		assert.Equal(t, mockCache, cache)

		return cache
	}

	actual := remote.ExtractValues(mockCache)

	assert.Equal(t, mockCache, actual)
}

func TestRemote_SelfReplaceIfNil(t *testing.T) {
	var nilRemote *Remote[int]
	assert.Equal(t, ZeroSimple[int]("NilRemote"), nilRemote.SelfReplaceIfNil())

	remote := NewRemote("Remote", 2, "RemoteName", "LogKey")
	assert.Equal(t, remote, remote.SelfReplaceIfNil())
}

func TestRemote_DoAnchor(t *testing.T) {
	remote := NewRemote("Remote", 2, "RemoteName", "LogKey")

	simple := remote.DoAnchor("Something")

	assert.Equal(t, "Something", simple.GetName())
	assert.Equal(t, 2, simple.GetTypedValue())
}

func TestRemote_MarshallJSON(t *testing.T) {
	scenarios := []struct {
		desc string
		test func(t *testing.T)
	}{
		{
			desc: "nil Remote",
			test: func(t *testing.T) {
				var nilRemote *Remote[int]

				actual, err := nilRemote.MarshalJSON()
				assert.Nil(t, actual)
				assert.Nil(t, err)
			},
		},
		{
			desc: "non-nil Remote",
			test: func(t *testing.T) {
				dummyInt := NewSimple("DummyInt", 1)

				toMarshall := NewRemote("Remote", 2, "RemoteName", "LogKey")
				toMarshall.Tag(NewTagFrom(dummyInt))
				toMarshall.AddCondition(NewCondition(NewSimple("Criteria", true)))

				wanted := `{"Value":"2","Source":"remote_calculation","DependentField":"RemoteName","LogKey":"LogKey","Tags":{"DummyInt":{"Value":"1","IsValue":true}},"Condition":{"Formula":{"Category":"AssignVariable","Operands":[{"Name":"Criteria"}]}}}`

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
