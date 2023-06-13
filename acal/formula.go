package acal

import (
	"encoding/json"
	"fmt"
)

// formulaProvider is the interface concrete Value needs to
// implement if they provide SyntaxNode.
type formulaProvider interface {
	// HasFormula returns whether this formulaProvider has a formula.
	HasFormula() bool
	// GetFormulaFn returns the function to build a formula of this FormulaProvider.
	GetFormulaFn() func() *SyntaxNode
}

// syntaxOperandProvider is the interface concrete Value needs to
// implement if they can be converted into a SyntaxOperand.
type syntaxOperandProvider interface {
	// ToSyntaxOperand returns the SyntaxOperand representation of this Value.
	ToSyntaxOperand(nextOp Op) *SyntaxOperand
}

// SyntaxNode is the representation of an operation performed on Value.
type SyntaxNode struct {
	category OpCategory
	op       Op
	opDesc   string
	operands []any
}

// NewSyntaxNode returns a new SyntaxNode with the provided fields.
func NewSyntaxNode(category OpCategory, op Op, opDesc string, operands []any) *SyntaxNode {
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
func (n *SyntaxNode) GetOperands() []any {
	return n.operands
}

// MarshalJSON returns the JSON encoding of this SyntaxNode.
func (n *SyntaxNode) MarshalJSON() ([]byte, error) {
	operands := make([]*SyntaxOperand, 0, len(n.operands))
	for _, operand := range n.operands {
		if sop, ok := operand.(Value); ok {
			operands = append(operands, sop.ToSyntaxOperand(n.op))
			continue
		}

		operands = append(operands, NewSyntaxOperandWithStaticValue(fmt.Sprintf("%v", operand)))
	}

	return json.Marshal(
		&struct {
			Category  string
			Operation string           `json:",omitempty"`
			Operands  []*SyntaxOperand `json:",omitempty"`
		}{
			Category:  n.category.toString(),
			Operation: n.opDesc,
			Operands:  operands,
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
}

// NewSyntaxOperand returns a new SyntaxOperand for those Value
// that come without stageIdx.
func NewSyntaxOperand(v Value) *SyntaxOperand {
	return &SyntaxOperand{
		Name: Identify(v),
	}
}

// NewSyntaxOperandWithStageIdx returns a new SyntaxOperand for those
// Value that come with stageIdx.
func NewSyntaxOperandWithStageIdx(v Value, stageIdx int) *SyntaxOperand {
	return &SyntaxOperand{
		Name:     Identify(v),
		StageIdx: stageIdx,
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
