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
