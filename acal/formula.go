package acal

import (
	"encoding/json"
)

// FormulaProvider is the interface concrete Value needs to
// implement if they provide SyntaxNode.
type FormulaProvider interface {
	// HasFormula returns whether this FormulaProvider has a formula.
	HasFormula() bool
	// GetFormula returns the formula provided by this FormulaProvider.
	GetFormula() *SyntaxNode
}

// SyntaxNode is the representation of an operation performed on Value.
type SyntaxNode struct {
	category OpCategory
	op       Op
	opDesc   string
	operands []*SyntaxOperand
}

// NewSyntaxNode returns a new SyntaxNode with the provided fields.
func NewSyntaxNode(category OpCategory, op Op, opDesc string, operands []*SyntaxOperand) *SyntaxNode {
	return &SyntaxNode{
		category: category,
		op:       op,
		opDesc:   opDesc,
		operands: operands,
	}
}

// GetCategory returns the category of this SyntaxNode.
func (n *SyntaxNode) GetCategory() OpCategory {
	return n.category
}

// GetOp returns the operation performed of this SyntaxNode.
func (n *SyntaxNode) GetOp() Op {
	return n.op
}

// GetOpDesc returns the operation description of this SyntaxNode.
func (n *SyntaxNode) GetOpDesc() string {
	return n.opDesc
}

// GetOperands returns the operands of this SyntaxNode.
func (n *SyntaxNode) GetOperands() []*SyntaxOperand {
	return n.operands
}

// MarshalJSON returns the JSON encoding of this SyntaxNode.
func (n *SyntaxNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		&struct {
			Category  string
			Operation string           `json:",omitempty"`
			Operands  []*SyntaxOperand `json:",omitempty"`
		}{
			Category:  n.category.toString(),
			Operation: n.opDesc,
			Operands:  n.operands,
		},
	)
}

// SyntaxOperand represents how an operand of a SyntaxNode should be
// encoded in JSON format.
type SyntaxOperand struct {
	Name              string      `json:",omitempty"`
	StageIdx          int         `json:",omitempty"`
	StaticValue       string      `json:",omitempty"`
	Node              *SyntaxNode `json:",omitempty"`
	WrapInParentheses bool        `json:",omitempty"`
	value             Value
}

// NewSyntaxOperand returns a new SyntaxOperand for those Value
// that come without stageIdx.
func NewSyntaxOperand(v Value) *SyntaxOperand {
	return &SyntaxOperand{
		Name:  Identify(v),
		value: v,
	}
}

// NewSyntaxOperandWithStageIdx returns a new SyntaxOperand for those
// Value that come with stageIdx.
func NewSyntaxOperandWithStageIdx(v Value, stageIdx int) *SyntaxOperand {
	return &SyntaxOperand{
		Name:     Identify(v),
		StageIdx: stageIdx,
		value:    v,
	}
}

// NewSyntaxOperandWithStaticValue returns a new SyntaxOperand for static values.
func NewSyntaxOperandWithStaticValue(value string) *SyntaxOperand {
	return &SyntaxOperand{
		StaticValue: value,
	}
}

// NewSyntaxOperandWithFormula returns a new SyntaxOperand for those Value
// that are not anchored but have a formula.
func NewSyntaxOperandWithFormula(formula *SyntaxNode, wrapInParentheses bool) *SyntaxOperand {
	return &SyntaxOperand{
		Node:              formula,
		WrapInParentheses: wrapInParentheses,
	}
}
