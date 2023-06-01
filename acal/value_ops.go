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
	// Describe ...
	Describe(v Value) string
	// DescribeValueAsFormula ...
	DescribeValueAsFormula(v Value) *SyntaxNode
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
func (valueOpsImpl) DescribeValueAsFormula(v Value) *SyntaxNode {
	if HasIdentity(v) {
		return &SyntaxNode{
			category: OpCategoryAssignVariable,
			op:       OpTransparent,
			operands: []any{
				v,
			},
		}
	}

	fp, ok := v.(FormulaProvider)
	if ok && fp.HasFormula() {
		return fp.GetFormula()
	}

	return &SyntaxNode{
		category: OpCategoryAssignStatic,
		op:       OpTransparent,
		opDesc:   Describe(v),
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

// Describe returns a full description of the given Value, including
// its identity and current value in string format.
func Describe(v Value) string {
	return valueOps.Describe(v)
}

// DescribeValueAsFormula returns a full description of the given Value,
// in the form of a formula.
func DescribeValueAsFormula(v Value) *SyntaxNode {
	return valueOps.DescribeValueAsFormula(v)
}
