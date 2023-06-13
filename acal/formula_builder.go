package acal

import "testing"

// FormulaBuilder helps create formulas in different categories.
var FormulaBuilder IFormulaBuilder = formulaBuilderImpl{}

// IFormulaBuilder ...
//
//go:generate mockery --name=IFormulaBuilder --case underscore --inpackage
type IFormulaBuilder interface {
	// NewFormulaFunctionCall returns a new SyntaxNode representing a function call taking in
	// the provided arguments. Clients must make sure to call PreProcessOperand on all args of
	// Value type before sending them into this method.
	NewFormulaFunctionCall(fnName string, arguments ...any) *SyntaxNode
	// NewFormulaTwoValMiddleOp returns a new SyntaxNode representing a binary operation that
	// has an operator lied in the middle of two operands. Clients must make sure that to call
	// PreProcessOperand on both v1 and v2 before sending them into this method.
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
		if v, ok := arg.(Value); ok {
			operands = append(operands, v)
			continue
		}

		operands = append(operands, arg)
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
	return NewSyntaxNode(OpCategoryTwoValMiddleOp, op, opDesc, []any{v1, v2})
}

// PreProcessOperand returns a replacement for the input value if it's nil.
func PreProcessOperand[T any](tv TypedValue[T]) TypedValue[T] {
	if IsNilValue(tv) {
		return ZeroSimple[T]("NilReplacement")
	}

	if ss, ok := tv.(snapshooter[T]); ok {
		return ss.getSnapshot()
	}

	return tv
}
