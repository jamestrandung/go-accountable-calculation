package acal

import (
	"fmt"
	"strconv"
)

// FormatFloatForMarshalling formats the given float so that it won't cause
// errors such as "json: unsupported value: +Inf" during marshalling.
func FormatFloatForMarshalling(f float64) string {
	return strconv.FormatFloat(f, 'g', -1, 64)
}

// PerformStandardValueExtraction extracts all Value that were used to calculate
// the given value and put them in the provided cache.
var PerformStandardValueExtraction = func(v Value, cache IValueCache) IValueCache {
	if IsNilValue(v) {
		return cache
	}

	if IsAnchored(v) {
		if !cache.Take(v) {
			return cache
		}
	}

	if c, ok := v.(StaticConditionalValue); ok && c.IsConditional() {
		c.GetCondition().criteria.ExtractValues(cache)
	}

	if t, ok := v.(Tagger); ok {
		for _, tag := range t.GetTags() {
			if tag.aVal != nil {
				tag.aVal.ExtractValues(cache)
			}
		}
	}

	fp, ok := v.(FormulaProvider)
	if !ok || !fp.HasFormula() {
		return cache
	}

	formula := fp.GetFormulaFn()()

	for _, operand := range formula.GetOperands() {
		if aVal, ok := operand.(Value); ok {
			aVal.ExtractValues(cache)
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
