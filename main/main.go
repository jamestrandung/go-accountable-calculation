package main

import (
	"fmt"
	"github.com/jamestrandung/go-accountable-calculation/acal"
	"github.com/jamestrandung/go-accountable-calculation/float"
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

	f1 := float.MakeSimpleFromFloat("f1", 1)
	//f2 := float.MakeSimpleFromFloat("f2", 2)
	//
	//f1.LargerThan(f2).Anchor("b1")

	p1 := float.MakeProgressive("p1")

	p1.Update(float.Zero)

	b2 := f1.LargerThan(p1).Anchor("b2")

	p1.Update(float.One)

	fmt.Println(acal.ToString(b2))

	var f2 float.Simple
	fmt.Println(f2.Float())
	fmt.Println(f2.Decimal())

	f3 := f1.Plus(f2).Anchor("f3")
	fmt.Println(acal.ToString(f3))

}
