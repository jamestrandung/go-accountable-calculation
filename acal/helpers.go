package acal

import (
	"reflect"
)

// isNil returns whether the given value is nil.
func isNil(i any) bool {
	if i == nil {
		return true
	}

	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Func, reflect.Chan, reflect.Slice, reflect.Interface:
		return reflect.ValueOf(i).IsNil()
	default:
		return false
	}
}
