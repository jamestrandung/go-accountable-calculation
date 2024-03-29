package boolean

import (
	"github.com/jamestrandung/go-accountable-calculation/acal"
)

// And applies AND operation on the value of this Simple and the given acal.TypedValue.
func (s Simple) And(s2 acal.TypedValue[bool]) Simple {
	return PerformBinaryLogicOp[bool](
		s, s2, opAnd, "AND", func(a, b bool) bool {
			return a && b
		},
	)
}

// Or applies OR operation on the value of this Simple and the given acal.TypedValue.
func (s Simple) Or(s2 acal.TypedValue[bool]) Simple {
	return PerformBinaryLogicOp[bool](
		s, s2, opOr, "OR", func(a, b bool) bool {
			return a || b
		},
	)
}

// Not returns the inverse value of this Simple.
func (s Simple) Not() Simple {
	return PerformUnaryLogicOp[bool](
		s, "NOT", func(b bool) bool {
			return !b
		},
	)
}

// Then does nothing and returns this Simple as-is. It's meant
// for separating code into more readable chunk.
func (s Simple) Then() Simple {
	return s
}
