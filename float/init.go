package float

import (
    "github.com/jamestrandung/go-accountable-calculation/acal"
    "github.com/my-shopspring/decimal"
)

var (
    NilFloat Simple

    MinusOne   *acal.Constant[decimal.Decimal]
    Zero       *acal.Constant[decimal.Decimal]
    One        *acal.Constant[decimal.Decimal]
    Two        *acal.Constant[decimal.Decimal]
    Three      *acal.Constant[decimal.Decimal]
    Four       *acal.Constant[decimal.Decimal]
    Five       *acal.Constant[decimal.Decimal]
    Six        *acal.Constant[decimal.Decimal]
    Seven      *acal.Constant[decimal.Decimal]
    Eight      *acal.Constant[decimal.Decimal]
    Nine       *acal.Constant[decimal.Decimal]
    Ten        *acal.Constant[decimal.Decimal]
    TwentyFour *acal.Constant[decimal.Decimal]
    Sixty      *acal.Constant[decimal.Decimal]
    Hundred    *acal.Constant[decimal.Decimal]
    Thousand   *acal.Constant[decimal.Decimal]
    Million    *acal.Constant[decimal.Decimal]
)

func init() {
    NilFloat = MakeSimpleFromFloat("NilFloat", 0)
    NilFloat.From(acal.SourceHardcode)

    MinusOne = acal.NewConstantWithFormat[decimal.Decimal](decimal.NewFromInt(-1), FormatFn)
    Zero = acal.NewConstantWithFormat[decimal.Decimal](decimal.Zero, FormatFn)
    One = acal.NewConstantWithFormat[decimal.Decimal](decimal.NewFromInt(1), FormatFn)
    Two = acal.NewConstantWithFormat[decimal.Decimal](decimal.NewFromInt(2), FormatFn)
    Three = acal.NewConstantWithFormat[decimal.Decimal](decimal.NewFromInt(3), FormatFn)
    Four = acal.NewConstantWithFormat[decimal.Decimal](decimal.NewFromInt(4), FormatFn)
    Five = acal.NewConstantWithFormat[decimal.Decimal](decimal.NewFromInt(5), FormatFn)
    Six = acal.NewConstantWithFormat[decimal.Decimal](decimal.NewFromInt(6), FormatFn)
    Seven = acal.NewConstantWithFormat[decimal.Decimal](decimal.NewFromInt(7), FormatFn)
    Eight = acal.NewConstantWithFormat[decimal.Decimal](decimal.NewFromInt(8), FormatFn)
    Nine = acal.NewConstantWithFormat[decimal.Decimal](decimal.NewFromInt(9), FormatFn)
    Ten = acal.NewConstantWithFormat[decimal.Decimal](decimal.NewFromInt(10), FormatFn)
    TwentyFour = acal.NewConstantWithFormat[decimal.Decimal](decimal.NewFromInt(24), FormatFn)
    Sixty = acal.NewConstantWithFormat[decimal.Decimal](decimal.NewFromInt(60), FormatFn)
    Hundred = acal.NewConstantWithFormat[decimal.Decimal](decimal.NewFromInt(100), FormatFn)
    Thousand = acal.NewConstantWithFormat[decimal.Decimal](decimal.NewFromInt(1000), FormatFn)
    Million = acal.NewConstantWithFormat[decimal.Decimal](decimal.NewFromInt(1000000), FormatFn)
}
