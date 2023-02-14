package float

import "math"

// roundToDecimalPlace rounds the given float64 to the provided decimal place
func roundToDecimalPlace(f float64, decimalPlace int) float64 {
	roundingFactor := math.Pow(10, float64(decimalPlace))
	return math.Round(f*roundingFactor) / roundingFactor
}
