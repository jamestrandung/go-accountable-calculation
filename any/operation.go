package any

import (
	"github.com/jamestrandung/go-accountable-calculation/acal"
	"github.com/jamestrandung/go-accountable-calculation/boolean"
)

// performBinaryLogicOp returns a boolean.Simple to represent the result of
// performing binaryOpFn on the values of the provided Simple.
func performBinaryLogicOp[T any](a1 Value, a2 Value, op acal.Op, binaryOpFn func(a, b T) bool) *boolean.Simple {
	if acal.IsNilValue(a1) && acal.IsNilValue(a2) {
		return nil
	}

	return boolean.NewSimpleWithFormula(
		binaryOpFn(a1.GetValue(), a2.GetValue()), func() *acal.SyntaxNode {
			return acal.FormulaBuilder.NewFormulaTwoValMiddleOp(a1, a2, op, anyLogicOpDesc[op])
		},
	)
}
