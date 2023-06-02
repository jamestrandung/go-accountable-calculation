package float

import "github.com/jamestrandung/go-accountable-calculation/acal"

var NilFloat Simple

func init() {
	NilFloat = MakeSimpleFromFloat("NilFloat", 0)
	NilFloat.From(acal.SourceHardcode)
}
