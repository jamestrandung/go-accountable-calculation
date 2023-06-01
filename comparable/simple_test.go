package comparable

import (
	"github.com/jamestrandung/go-accountable-calculation/acal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimple_IsNil(t *testing.T) {
	var nilSimple *Simple[int]
	assert.True(t, nilSimple.IsNil())

	fSimple := NewSimple[bool]("false", false)
	assert.False(t, fSimple.IsNil())
}

func TestSimple_SelfReplaceIfNil(t *testing.T) {
	var nilSimple *Simple[int]
	assert.Equal(t, acal.ZeroSimple[int]("NilComparable"), nilSimple.SelfReplaceIfNil())

	simple := NewSimple("Simple", 1)
	assert.Equal(t, simple, simple.SelfReplaceIfNil())
}

func TestSimple_Anchor(t *testing.T) {
	scenarios := []struct {
		desc string
		test func(t *testing.T)
	}{
		{
			desc: "nil Simple",
			test: func(t *testing.T) {
				var nilSimple *Simple[int]

				actual := nilSimple.Anchor("Something")

				assert.Equal(t, "Something", actual.GetName())
				assert.Equal(t, "0", actual.Stringify())
			},
		},
		{
			desc: "same Simple",
			test: func(t *testing.T) {
				simple := NewSimpleFrom[bool](acal.NewConstant(true))

				actual := simple.Anchor("Something")

				assert.Equal(t, "Something", actual.GetName())
				assert.Equal(t, "true", actual.Stringify())
				assert.Equal(t, simple, actual)
			},
		},
		{
			desc: "new Simple",
			test: func(t *testing.T) {
				simple := NewSimple("AlreadyAnchored", true)

				actual := simple.Anchor("Something")

				assert.Equal(t, "Something", actual.GetName())
				assert.Equal(t, "true", actual.Stringify())
				assert.NotEqual(t, simple, actual)
			},
		},
	}

	for _, sc := range scenarios {
		t.Run(sc.desc, sc.test)
	}
}
