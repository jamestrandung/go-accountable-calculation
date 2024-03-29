package acal

// IValueOps ...
//
//go:generate mockery --name=IValueOps --case underscore --inpackage
type IValueOps interface {
	// IsNilValue ...
	IsNilValue(v Value) bool
	// HasIdentity ...
	HasIdentity(v Value) bool
	// Identify ...
	Identify(v Value) string
	// Stringify ...
	Stringify(v Value) string
	// Describe ...
	Describe(v Value) string
	// DescribeValueAsFormula ...
	DescribeValueAsFormula(v Value) func() *SyntaxNode
}

type valueOpsImpl struct{}

// IsNilValue ...
func (valueOpsImpl) IsNilValue(v Value) bool {
	return v == nil || v.IsNil()
}

// HasIdentity ...
func (valueOpsImpl) HasIdentity(v Value) bool {
	if IsNilValue(v) {
		return false
	}

	return v.HasIdentity()
}

// Identify ...
func (valueOpsImpl) Identify(v Value) string {
	if IsNilValue(v) {
		return UnknownValueName
	}

	return v.Identify()
}

// Stringify ...
func (valueOpsImpl) Stringify(v Value) string {
	if IsNilValue(v) {
		return ""
	}

	return v.Stringify()
}

// Describe ...
func (valueOpsImpl) Describe(v Value) string {
	if IsNilValue(v) {
		return "?[?]"
	}

	identity := Identify(v)
	if identity == "" {
		return "?[" + v.Stringify() + "]"
	}

	return identity + "[" + v.Stringify() + "]"
}

// DescribeValueAsFormula ...
func (valueOpsImpl) DescribeValueAsFormula(v Value) func() *SyntaxNode {
	if HasIdentity(v) {
		return func() *SyntaxNode {
			return &SyntaxNode{
				category: OpCategoryAssignVariable,
				op:       OpTransparent,
				operands: []any{
					v,
				},
			}
		}
	}

	fp, ok := v.(formulaProvider)
	if ok && fp.HasFormula() {
		return fp.GetFormulaFn()
	}

	return func() *SyntaxNode {
		return &SyntaxNode{
			category: OpCategoryAssignStatic,
			op:       OpTransparent,
			opDesc:   Describe(v),
		}
	}
}

var valueOps IValueOps = valueOpsImpl{}

// IsNilValue returns whether the given Value is nil.
func IsNilValue(v Value) bool {
	return valueOps.IsNilValue(v)
}

// HasIdentity returns whether the given Value is anchored.
func HasIdentity(v Value) bool {
	return valueOps.HasIdentity(v)
}

// Identify returns the identity of the given Value, which is its
// alias if not nil. Otherwise, the result should be its name.
func Identify(v Value) string {
	return valueOps.Identify(v)
}

// Stringify returns the value this Value contains as a string.
func Stringify(v Value) string {
	return valueOps.Stringify(v)
}

// Describe returns a full description of the given Value, including
// its identity and current value in string format.
func Describe(v Value) string {
	return valueOps.Describe(v)
}

// DescribeValueAsFormula returns a full description of the given Value,
// in the form of a formula.
func DescribeValueAsFormula[T any](v TypedValue[T]) func() *SyntaxNode {
	return valueOps.DescribeValueAsFormula(PreProcessOperand(v))
}

// ExtractTypedValue returns the typed value the given TypedValue contains.
func ExtractTypedValue[T any](tv TypedValue[T]) T {
	if IsNilValue(tv) {
		var temp T
		return temp
	}

	return tv.GetTypedValue()
}
