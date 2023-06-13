package acal

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

// IMarshalOps ...
//
//go:generate mockery --name=IMarshalOps --case underscore --inpackage
type IMarshalOps interface {
	// MarshalJSON ...
	MarshalJSON(values ...Value) ([]byte, error)
	// MarshalJSONByFields ...
	MarshalJSONByFields(data any) ([]byte, error)
}

type marshalOpsImpl struct{}

var jsonMarshalFn = json.Marshal

// MarshalJSON ...
func (marshalOpsImpl) MarshalJSON(values ...Value) ([]byte, error) {
	cache := NewValueCache()

	for _, value := range values {
		if IsNilValue(value) {
			continue
		}

		if !HasIdentity(value) {
			value.SetAlias(UnknownValueName)
		}

		cache = value.ExtractValues(cache)
	}

	return jsonMarshalFn(cache.Flatten())
}

// MarshalJSONByFields ...
func (marshalOpsImpl) MarshalJSONByFields(data any) ([]byte, error) {
	if isNil(data) {
		return nil, fmt.Errorf("cannot marshal nil data")
	}

	if v, ok := data.(Value); ok {
		return MarshalJSON(v)
	}

	if reflect.TypeOf(data).Kind() != reflect.Struct {
		return nil, fmt.Errorf("cannot marshal non-struct data")
	}

	toMarshal := make(map[string]any)

	cache := NewValueCache()

	rv := reflect.ValueOf(data)
	rt := reflect.TypeOf(data)
	for i := 0; i < rt.NumField(); i++ {
		rvf := rv.Field(i)

		// Un-exported fields
		if !rvf.CanInterface() {
			continue
		}

		rtf := rt.Field(i)

		actualValue := rvf.Interface()

		v, ok := actualValue.(Value)
		if !ok {
			toMarshal[rtf.Name] = actualValue
			continue
		}

		if !IsNilValue(v) {
			if !HasIdentity(v) {
				v.SetAlias(UnknownValueName)
			}

			cache = v.ExtractValues(cache)
		}
	}

	for key, extractedValue := range cache.Flatten() {
		if toMarshal[key] == nil {
			toMarshal[key] = extractedValue
			continue
		}

		count := 2
		newKey := key + "." + strconv.Itoa(count)
		for toMarshal[newKey] != nil {
			count++
			newKey = key + "." + strconv.Itoa(count)
		}

		extractedValue.SetAlias(newKey)
		toMarshal[newKey] = extractedValue
	}

	return jsonMarshalFn(toMarshal)
}

var marshalOps IMarshalOps = marshalOpsImpl{}

// MarshalJSON puts the provided Value and all other Value that were used to calculate
// the given Value into a map and returns a JSON encoding of this map.
func MarshalJSON(values ...Value) ([]byte, error) {
	return marshalOps.MarshalJSON(values...)
}

// MarshalJSONByFields will execute MarshalJSON on the provided data if it's a Value. Otherwise,
// each exported field of the provided data will be put into a map. If such field is a Value, all
// other Value that were used to calculate this field will also be added to this map. Finally, a
// JSON encoding of this map will be returned.
//
// If the alias of any Value happens to collide with the name of non-Value fields, such Value will
// be de-duplicated in the JSON map by appending a suffix to the alias of all duplicates. Format for
// alias of de-duplicated variables: <name_of_variable>.<count>.
// - count starts from 2
// - example: duplicate.2, duplicate3.2
func MarshalJSONByFields(data any) ([]byte, error) {
	return marshalOps.MarshalJSONByFields(data)
}
