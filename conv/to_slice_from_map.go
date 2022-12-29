package conv

import (
	"fmt"
	"reflect"
)

//
// Convert map keys and values to slice in random order.
// E.g. covert map[string]int{"a":1,"b":2, "c":3} to []string{"a", "c", "b"} and []int{1, 3, 2}.
//

// MapKeys returns a slice of all the keys in m.
// The keys are returned in random order.
func MapKeys[Key comparable, Val any](m map[Key]Val) []Key {
	s := make([]Key, 0, len(m))
	for k := range m {
		s = append(s, k)
	}
	return s
}

// MapVals returns a slice of all the values in m.
// The values are returned in random order.
func MapVals[Key comparable, Val any](m map[Key]Val) []Val {
	s := make([]Val, 0, len(m))
	for _, v := range m {
		s = append(s, v)
	}
	return s
}

// MapKeyVals returns two slice of all the keys and values in m.
// The keys and values are returned in random order.
func MapKeysVals[Key comparable, Val any](m map[Key]Val) ([]Key, []Val) {
	ks, vs := make([]Key, 0, len(m)), make([]Val, 0, len(m))
	for k, v := range m {
		ks = append(ks, k)
		vs = append(vs, v)
	}
	return ks, vs
}

// Map2Slice converts map keys and values to slice in random order.
func Map2Slice(a any) (ks any, vs any) {
	ks, vs, _ = Map2SliceE(a)
	return
}

// Map2SliceE converts keys and values of map to slice in random order with error.
func Map2SliceE(a any) (ks any, vs any, err error) {
	// Check param.
	if a == nil {
		return nil, nil, fmt.Errorf("unable to converts %#v of type %T to map[any]struct{}", a, a)
	}
	t := reflect.TypeOf(a)
	if t.Kind() != reflect.Map {
		err = fmt.Errorf("the input %#v of type %T isn't a map", a, a)
		return
	}

	// Execute the conversion.
	m := reflect.ValueOf(a)
	l := m.Len()
	keys := m.MapKeys()
	ksT, vsT := reflect.SliceOf(t.Key()), reflect.SliceOf(t.Elem())
	ksV, vsV := reflect.MakeSlice(ksT, 0, l), reflect.MakeSlice(vsT, 0, l)
	for _, k := range keys {
		ksV = reflect.Append(ksV, k)
		vsV = reflect.Append(vsV, m.MapIndex(k))
	}
	return ksV.Interface(), vsV.Interface(), nil
}
