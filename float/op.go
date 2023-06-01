package float

import (
    "github.com/jamestrandung/go-accountable-calculation/acal"
)

const (
    opPlus acal.Op = iota
    opMinus
    opMultiply
    opDivide
    opAbs
    opCeil
    opFloor
    opRound
    opPow

    maxOpName          = "Max"
    minOpName          = "Min"
    medianOpName       = "Median"
    boundBetweenOpName = "BoundBetween"
)

var (
    opDesc = map[acal.Op]string{
        opPlus:     "+",
        opMinus:    "-",
        opMultiply: "*",
        opDivide:   "/",
        opAbs:      "Abs",
        opCeil:     "Ceil",
        opFloor:    "Floor",
        opRound:    "Round",
        opPow:      "Pow",
    }

    opLevel = map[acal.Op]int{
        opPlus:     0,
        opMinus:    0,
        opMultiply: 1,
        opDivide:   1,
    }
)
