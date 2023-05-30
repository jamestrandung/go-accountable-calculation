package acal

// valueWithFormula ...
//
//go:generate mockery --name=valueWithFormula --case underscore --inpackage
type valueWithFormula interface {
	Value
	FormulaProvider
}

//go:generate mockery --name=valueWithAllFeatures --case underscore --inpackage
type valueWithAllFeatures interface {
	Value
	FormulaProvider
	ITagger
	StaticConditionalValue
}
