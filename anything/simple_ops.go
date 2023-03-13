package anything

import (
	"github.com/jamestrandung/go-accountable-calculation/acal"
	"github.com/jamestrandung/go-accountable-calculation/boolean"
	"github.com/jamestrandung/go-accountable-calculation/op"
)

// EqualsRaw returns whether the value of this Simple equals to the raw input value.
func (s *Simple[T]) EqualsRaw(v T) *boolean.Simple {
	return op.PerformBinaryLogicOp[T](
		s, acal.NewConstant[T](v), anyOpEquals, "==", func(a, b T) bool {
			return a == b
		},
	)
}

// Equals returns whether the value of this Simple equals to the input value.
func (s *Simple[T]) Equals(v acal.TypedValue[T]) *boolean.Simple {
	return op.PerformBinaryLogicOp[T](
		s, v, anyOpEquals, "==", func(a, b T) bool {
			return a == b
		},
	)
}

// NotEqualsRaw returns whether the value of this Simple does not equal to the raw input value.
func (s *Simple[T]) NotEqualsRaw(v T) *boolean.Simple {
	return op.PerformBinaryLogicOp[T](
		s, acal.NewConstant[T](v), anyOpNotEquals, "!=", func(a, b T) bool {
			return a != b
		},
	)
}

// NotEquals returns whether the value of this Simple does not equal to the input value.
func (s *Simple[T]) NotEquals(v acal.TypedValue[T]) *boolean.Simple {
	return op.PerformBinaryLogicOp[T](
		s, v, anyOpNotEquals, "!=", func(a, b T) bool {
			return a != b
		},
	)
}
