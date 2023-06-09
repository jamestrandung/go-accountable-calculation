package float

import (
	"github.com/jamestrandung/go-accountable-calculation/acal"
	"github.com/shopspring/decimal"
)

var (
	NilFloat Simple

	Zero *acal.Constant[decimal.Decimal]
	One  *acal.Constant[decimal.Decimal]
)

func init() {
	NilFloat = MakeSimpleFromFloat("NilFloat", 0)
	NilFloat.From(acal.SourceHardcode)

	Zero = acal.NewConstantWithFormat[decimal.Decimal](decimal.Zero, FormatFn)
	One = acal.NewConstantWithFormat[decimal.Decimal](decimal.NewFromInt(1), FormatFn)
}
