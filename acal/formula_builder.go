package acal

import "testing"

// FormulaBuilder helps create formulas in different categories.
var FormulaBuilder IFormulaBuilder = formulaBuilderImpl{}

// IFormulaBuilder ...
//
//go:generate mockery --name=IFormulaBuilder --case underscore --inpackage
type IFormulaBuilder interface {
	// NewFormulaFunctionCall returns a new SyntaxNode representing a function call taking in the provided arguments.
	NewFormulaFunctionCall(fnName string, arguments ...any) *SyntaxNode
	// NewFormulaTwoValMiddleOp returns a new SyntaxNode representing a binary operation that has an operator in the middle of two operands.
	NewFormulaTwoValMiddleOp(v1 Value, v2 Value, op Op, opDesc string) *SyntaxNode
}

// MockFormulaBuilder can be used in tests to perform monkey-patching on FormulaBuilder
func MockFormulaBuilder(t *testing.T) (*MockIFormulaBuilder, func()) {
	old := FormulaBuilder
	mock := NewMockIFormulaBuilder(t)

	FormulaBuilder = mock
	return mock, func() {
		FormulaBuilder = old
	}
}

type formulaBuilderImpl struct{}

// NewFormulaFunctionCall ...
func (b formulaBuilderImpl) NewFormulaFunctionCall(fnName string, arguments ...any) *SyntaxNode {
	operands := make([]any, 0, len(arguments))

	for _, arg := range arguments {
		v, ok := arg.(Value)
		if !ok {
			operands = append(operands, arg)
			continue
		}

		v = replaceNilFromConcreteImplementation(v)
		operands = append(operands, v)
	}

	return NewSyntaxNode(
		OpCategoryFunctionCall,
		OpTransparent,
		fnName,
		operands,
	)
}

// NewFormulaTwoValMiddleOp ...
func (b formulaBuilderImpl) NewFormulaTwoValMiddleOp(v1 Value, v2 Value, op Op, opDesc string) *SyntaxNode {
	v1 = replaceNilFromConcreteImplementation(v1)
	v2 = replaceNilFromConcreteImplementation(v2)

	return NewSyntaxNode(OpCategoryTwoValMiddleOp, op, opDesc, []any{v1, v2})
}

func replaceNilFromConcreteImplementation(v Value) Value {
	if v != nil {
		return v.SelfReplaceIfNil()
	}

	return v
}
