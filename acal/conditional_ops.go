package acal

// ICondOps ...
//go:generate mockery --name=ICondOps --case underscore --inpackage
type ICondOps interface {
	// ApplyStaticCondition ...
	ApplyStaticCondition(criteria TypedValue[bool], values ...StaticConditionalValue)
	// ApplyProgressiveCondition ...
	ApplyProgressiveCondition(criteria TypedValue[bool], values ...ProgressiveConditionalValue) CloseIfFunc
}

type condOpsImpl struct{}

// ApplyStaticCondition ...
func (condOpsImpl) ApplyStaticCondition(criteria TypedValue[bool], values ...StaticConditionalValue) {
	c := NewCondition(criteria)

	for _, value := range values {
		value.AddCondition(c)
	}
}

// ApplyProgressiveCondition ...
func (condOpsImpl) ApplyProgressiveCondition(criteria TypedValue[bool], values ...ProgressiveConditionalValue) CloseIfFunc {
	closeIfFn := func() {}

	for _, value := range values {
		fn := value.AddCondition(criteria)

		prevFn := closeIfFn

		closeIfFn = func() {
			fn()
			prevFn()
		}
	}

	return closeIfFn
}

var condOps ICondOps = condOpsImpl{}

// ApplyStaticCondition creates a Condition and applies it on all given StaticConditionalValue.
func ApplyStaticCondition(criteria TypedValue[bool], values ...StaticConditionalValue) {
	condOps.ApplyStaticCondition(criteria, values...)
}

// ApplyProgressiveCondition applies the provided Condition on all given ProgressiveConditionalValue and returns a
// single CloseIfFunc to close all conditions at once.
func ApplyProgressiveCondition(criteria TypedValue[bool], values ...ProgressiveConditionalValue) CloseIfFunc {
	return condOps.ApplyProgressiveCondition(criteria, values...)
}
