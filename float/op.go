package float

import (
    "github.com/jamestrandung/go-accountable-calculation/acal"
    "github.com/shopspring/decimal"
)

var (
    opLevel = map[acal.Op]int{
        opAdd:      0,
        opSubtract: 0,
        opMultiply: 1,
        opDivide:   1,
    }
)

type opProvider struct {
    tv acal.TypedValue[decimal.Decimal]
}
