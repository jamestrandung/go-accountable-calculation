package float

import (
	"github.com/shopspring/decimal"
	"math"
	"strconv"
)

const formattingDecimalPlace = 6

var formattingDigitThreshold = decimal.NewFromFloat(math.Pow10(8))

var floatFormatFn = func(d decimal.Decimal) string {
	d = d.Round(formattingDecimalPlace)
	if d.LessThan(formattingDigitThreshold) {
		return d.String()
	}

	return strconv.FormatFloat(d.InexactFloat64(), 'g', -1, 64)
}
