package boolean

import "github.com/jamestrandung/go-accountable-calculation/acal"

var NilBool Simple

func init() {
	NilBool = MakeSimple("NilBool", false)
	NilBool.From(acal.SourceHardcode)
}
