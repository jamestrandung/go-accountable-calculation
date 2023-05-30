package main

import (
	"fmt"
	"github.com/jamestrandung/go-accountable-calculation/acal"
	"github.com/jamestrandung/go-accountable-calculation/boolean"
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

	b1 := boolean.NewSimpleFrom(acal.NewConstant(true))
	b1.Anchor("Bool")

	fmt.Print(acal.ToString(b1))

	t1 = acal.NewProgressive[int]("something")
	t2 = acal.NewProgressive[bool]("something")

	p, ok := t2.(*acal.Progressive[bool])
	fmt.Println(p, ok)
}
