package acal

import (
	"fmt"
)

// PerformStandardValueExtraction extracts all Value that were used to calculate
// the given value and put them in the provided cache.
var PerformStandardValueExtraction = func(v Value, cache IValueCache) IValueCache {
	if IsNilValue(v) {
		return cache
	}

	if HasIdentity(v) {
		if !cache.Take(v) {
			return cache
		}
	}

	if c, ok := v.(StaticConditionalValue); ok && c.IsConditional() {
		c.GetCondition().criteria.ExtractValues(cache)
	}

	if t, ok := v.(ITagger); ok {
		for _, tag := range t.GetTags() {
			if tag.aVal != nil {
				tag.aVal.ExtractValues(cache)
			}
		}
	}

	fp, ok := v.(formulaProvider)
	if !ok || !fp.HasFormula() {
		return cache
	}

	formula := fp.GetFormulaFn()()

	for _, operand := range formula.GetOperands() {
		if ev, ok := operand.(extractable); ok {
			ev.ExtractValues(cache)
		}
	}

	return cache
}

// ToString returns the JSON structure representing the given Value as a string
func ToString(aVal ...Value) string {
	json, err := MarshalJSON(aVal...)
	if err != nil {
		return fmt.Sprintf("<failed to marshal Value: %s>", err.Error())
	}

	return string(json)
}
