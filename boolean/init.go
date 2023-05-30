package boolean

import "github.com/jamestrandung/go-accountable-calculation/acal"

var NilBool *Simple

func init() {
	NilBool = NewSimple("NilBool", false)
	NilBool.From(acal.SourceHardcode)
}
