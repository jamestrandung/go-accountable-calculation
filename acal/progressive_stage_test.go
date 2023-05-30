package acal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStage_IsNil(t *testing.T) {
	var nilStage *Stage[int]

	assert.True(t, nilStage.IsNil())

	stage := &Stage[int]{}

	assert.False(t, stage.IsNil())
}

func TestStage_GetTypedValue(t *testing.T) {
	var nilStage *Stage[int]

	assert.Equal(t, 0, nilStage.GetTypedValue())

	stage := &Stage[int]{
		value: 2,
	}

	assert.Equal(t, 2, stage.GetTypedValue())
}

func TestStage_ToSyntaxOperand(t *testing.T) {
	stage := &Stage[int]{
		self:  NewProgressive[int]("Something"),
		value: 2,
		idx:   3,
	}

	actual := stage.ToSyntaxOperand(OpTransparent)

	assert.Equal(
		t, &SyntaxOperand{
			Name:     "Something",
			StageIdx: 3,
		}, actual,
	)
}

func TestStage_SelfReplaceIfNil(t *testing.T) {
	var nilStage *Stage[int]
	assert.Equal(t, ZeroSimple[int]("NilStage"), nilStage.SelfReplaceIfNil())

	stage := &Stage[int]{}
	assert.Equal(t, stage, stage.SelfReplaceIfNil())
}
