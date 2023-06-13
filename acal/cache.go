package acal

import "strconv"

type extractableValue interface {
	// ExtractValues extracts this Value and all Value that were used to calculate it.
	ExtractValues(cache IValueCache) IValueCache
}

// IValueCache keeps extracted values during marshalling so that we can
// build a complete JSON structure containing all calculated variables
// from a small set of input variables.
//
//go:generate mockery --name=IValueCache --case underscore --inpackage
type IValueCache interface {
	// Take accepts a Value to put into ValueCache. It will automatically de-duplicate
	// variables under the same name by giving duplicates an alias if it doesn't have
	// one already. The return value indicates whether the given value was accepted
	// into the cache.
	//
	// Format for alias of de-duplicated variables: <name_of_variable><count>.
	// - count starts from 2
	// - example: duplicate2, duplicate3
	Take(v Value) bool
	// Flatten returns a map of calculated variables that is ready for marshalling into
	// JSON. Variables with colliding alias will be de-duplicated by appending a count
	// suffix to the alias of all duplicates.
	//
	// Format for alias of de-duplicated variables: <name_of_variable>_<count>.
	// - count starts from 2
	// - example: duplicate_2, duplicate3_2
	Flatten() map[string]Value
}

type valueCache struct {
	values map[string][]Value
}

// NewValueCache ...
func NewValueCache() IValueCache {
	return valueCache{
		values: make(map[string][]Value),
	}
}

// Take ...
func (c valueCache) Take(v Value) bool {
	existings, ok := c.values[v.GetName()]
	if !ok {
		c.values[v.GetName()] = []Value{
			v,
		}

		return true
	}

	for _, existing := range existings {
		// Already took this v before
		if existing == v {
			return false
		}
	}

	if v.GetAlias() == "" {
		alias := v.GetName() + strconv.Itoa(len(existings)+1)
		v.SetAlias(alias)
	}

	c.values[v.GetName()] = append(existings, v)

	return true
}

// Flatten ...
func (c valueCache) Flatten() map[string]Value {
	flattens := make(map[string]Value)

	for _, values := range c.values {
		for _, value := range values {
			identity := Identify(value)
			if flattens[identity] == nil {
				flattens[identity] = value
				continue
			}

			// When name and alias of 2 different values are the same,
			// it means the alias of these values was set manually by
			// clients to mark that they are the same thing.
			if flattens[identity].GetName() == value.GetName() {
				continue
			}

			count := 2
			key := identity + "_" + strconv.Itoa(count)
			for flattens[key] != nil {
				count++
				key = identity + "_" + strconv.Itoa(count)
			}

			value.SetAlias(key)
			flattens[key] = value
		}
	}

	return flattens
}
