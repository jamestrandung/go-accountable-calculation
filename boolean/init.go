package boolean

import "github.com/jamestrandung/go-accountable-calculation/acal"

var NilBool *Simple

func init() {
	NilBool = Bool("NilBool", false)
	NilBool.From(acal.SourceHardcode)
}
