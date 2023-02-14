package main

import (
	"fmt"
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
	b := Bool{}
	fmt.Println(b.GetTypedValue())
	b.Print()

	b2 := Implementation[bool]{}
	fmt.Println(b2.GetTypedValue())
	b2.Print()

	var b3 Implementation[bool] = b
}
