package float

import (
	"github.com/jamestrandung/go-accountable-calculation/acal"
	"github.com/jamestrandung/go-accountable-calculation/base"
)

var (
	NilFloat *base.Any[float64]
)

func init() {
	NilFloat = base.NewAny[float64]("NilFloat", 0)
	NilFloat.From(acal.SourceHardcode)
}
