package acal

import (
	"encoding/json"
)

// CloseIfFunc is to be called when an if clause ends so that
// progressive variables can record the close if stage index.
type CloseIfFunc func()

// DoNothingCloseIfFunc is a CloseIfFunc that does nothing.
var DoNothingCloseIfFunc CloseIfFunc = func() {}

// StaticConditionalValue needs to be implemented by static
// Value that holds a Condition.
//go:generate mockery --name=StaticConditionalValue --case underscore --inpackage
type StaticConditionalValue interface {
	// IsConditional returns whether a Condition is attached to this StaticConditionalValue.
	IsConditional() bool
	// GetCondition returns the Condition attached to this StaticConditionalValue.
	GetCondition() *Condition
	// AddCondition attaches the given Condition to this StaticConditionalValue.
	AddCondition(*Condition)
}

// ProgressiveConditionalValue needs to be implemented by progressive
// Value that might involve conditional calculation.
//go:generate mockery --name=ProgressiveConditionalValue --case underscore --inpackage
type ProgressiveConditionalValue interface {
	// AddCondition attaches the given Condition to this ProgressiveConditionalValue
	// and returns a CloseIfFunc to close the Condition.
	AddCondition(TypedValue[bool]) CloseIfFunc
}

// Condition represents an if clause in a calculation algorithm.
type Condition struct {
	criteria      TypedValue[bool]
	prevCondition *Condition

	openIfStageIdx  int
	closeIfStageIdx int
}

// NewCondition returns a new Condition with the provided fields.
func NewCondition(criteria TypedValue[bool]) *Condition {
	return &Condition{
		criteria: criteria,
	}
}

// NewProgressiveCondition returns a new Condition with the provided fields.
func NewProgressiveCondition(criteria TypedValue[bool], prevCondition *Condition, openIfStageIdx int) *Condition {
	return &Condition{
		criteria:       criteria,
		prevCondition:  prevCondition,
		openIfStageIdx: openIfStageIdx,
	}
}

// isValidProgressiveCondition returns whether this Condition contains at
// least 1 progressive stage between open and close index.
func (c *Condition) isValidProgressiveCondition() bool {
	return c.closeIfStageIdx >= c.openIfStageIdx
}

// setCloseIfStageIdx updates the closeIfStageIdx of this Condition to the provided value.
func (c *Condition) setCloseIfStageIdx(closeIfStageIdx int) {
	c.closeIfStageIdx = closeIfStageIdx
}

// MarshalJSON returns the JSON encoding of this Condition.
func (c *Condition) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		&struct {
			Formula         *SyntaxNode
			CloseIfStageIdx int `json:",omitempty"`
		}{
			Formula:         DescribeValueAsFormula(c.criteria)(),
			CloseIfStageIdx: c.closeIfStageIdx,
		},
	)
}
