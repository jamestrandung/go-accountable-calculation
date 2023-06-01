package main

import (
	"fmt"
	"github.com/jamestrandung/go-accountable-calculation/acal"
	"github.com/jamestrandung/go-accountable-calculation/float"
	"math"
)

type RawInterface interface {
	GetValue() any
}

type TypedInterface[T any] interface {
	RawInterface
	GetTypedValue() T
}

type boolV struct {
}

func (boolV) GetValue() any {
	return false
}

func (boolV) GetTypedValue() bool {
	return false
}

func do(v TypedInterface[bool]) {
	fmt.Println(v)
}

type Implementation[T any] struct {
	value T
}

func (i Implementation[T]) GetTypedValue() T {
	return i.value
}

func (i Implementation[T]) Print() {
	fmt.Println(i.GetTypedValue())
}

type Bool struct {
	Implementation[bool]
}

func (Bool) GetTypedValue() bool {
	return true
}

func main() {
	var t1 acal.Value
	var t2 acal.TypedValue[bool]

	fmt.Println(acal.IsNilValue(t1))
	fmt.Println(acal.IsNilValue(t2))

	f := float.NewSimpleFromFloat("Something", math.MaxFloat64)

	fmt.Println(acal.ToString(f))

	f = float.NewSimpleFromFloat("Something", 2.123456789)

	fmt.Println(acal.ToString(f))

	f = float.NewSimpleFromFloat("Something", 12345678)

	fmt.Println(acal.ToString(f))
}
