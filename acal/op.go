package acal

// OpCategory represents the kind of operation being performed,
// which can decide how such operation should be displayed in
// a formula.
type OpCategory int

const (
	// OpCategoryFunctionCall represents an operation that looks like a function call.
	OpCategoryFunctionCall OpCategory = iota
	// OpCategoryTwoValMiddleOp represents an operation with the operator lies in the middle of 2 operands.
	OpCategoryTwoValMiddleOp
	// OpCategoryAssignVariable represents an assign operation where a variable is assigned as the value of another variable.
	OpCategoryAssignVariable
	// OpCategoryAssignStatic represents an assign operation where a static value is assigned as the value of a variable.
	OpCategoryAssignStatic
)

func (c OpCategory) toString() string {
	switch c {
	case OpCategoryFunctionCall:
		return "FunctionCall"
	case OpCategoryTwoValMiddleOp:
		return "TwoValMiddleOp"
	case OpCategoryAssignVariable:
		return "AssignVariable"
	case OpCategoryAssignStatic:
		return "AssignStatic"
	}

	return "Unknown"
}

// Op represents an operation that can be performed on Value.
type Op int

// IsTransparent returns whether this Op is the special OpTransparent.
func (o Op) IsTransparent() bool {
	return o == OpTransparent
}

const (
	// OpTransparent is a transparent op that does not show up in formulas.
	OpTransparent Op = -1
)
