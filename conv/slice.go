package conv

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// ToSlice converts any type slice or array to the specified type slice.
func ToSlice[T any](a any) []T {
	r, _ := ToSliceE[T](a)
	return r
}

// ToSliceE converts any type slice or array to the specified type slice.
// An error will be returned if an error occurred.
func ToSliceE[T any](a any) ([]T, error) {
	if a == nil {
		return nil, nil
	}
	switch v := a.(type) {
	case []T:
		return v, nil
	case string:
		return ToSliceE[T](strings.Fields(v))
	}

	kind := reflect.TypeOf(a).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		// If input is a slice or array.
		v := reflect.ValueOf(a)
		if kind == reflect.Slice && v.IsNil() {
			return nil, nil
		}
		s := make([]T, v.Len())
		for i := 0; i < v.Len(); i++ {
			val, err := ToAnyE[T](v.Index(i).Interface())
			if err != nil {
				return nil, err
			}
			s[i] = val
		}
		return s, nil
	default:
		// If input is a single value.
		v, err := ToAnyE[T](a)
		if err != nil {
			return nil, err
		}
		return []T{v}, nil
	}
}

// JsonToSlice converts the JSON-encoded data to any type slice with no error returned.
func JsonToSlice[S ~[]E, E any](data []byte) S {
	s, _ := JsonToSliceE[S](data)
	return s
}

// JsonToSliceE converts the JSON-encoded data to any type slice.
// E.g. a JSON value ["foo", "bar", "baz"] can be converted to []string{"foo", "bar", "baz"}
// when calling JsonToSliceE[[]string](`["foo", "bar", "baz"]`).
func JsonToSliceE[S ~[]E, E any](data []byte) (s S, err error) {
	err = json.Unmarshal(data, &s)
	return
}

//
// Convert map keys and values to slice in indeterminate order.
// E.g. covert map[string]int{"a":1,"b":2, "c":3} to []string{"a", "c", "b"} and []int{1, 3, 2}.
//

// MapKeys returns a slice of all the keys in m.
// The keys returned are in indeterminate order.
// Deprecated: As of Go 1.18, please use standard library golang.org/x/exp/maps#Keys.
func MapKeys[K comparable, V any, M ~map[K]V](m M) []K {
	s := make([]K, 0, len(m))
	for k := range m {
		s = append(s, k)
	}
	return s
}

// MapVals returns a slice of all the values in m.
// The values returned are in indeterminate order.
// Deprecated: As of Go 1.18, please use standard library golang.org/x/exp/maps#Values.
func MapVals[K comparable, V any, M ~map[K]V](m M) []V {
	s := make([]V, 0, len(m))
	for _, v := range m {
		s = append(s, v)
	}
	return s
}

// MapKeyVals returns two slice of all the keys and values in m.
// The keys and values are returned in an indeterminate order.
func MapKeyVals[K comparable, V any, M ~map[K]V](m M) ([]K, []V) {
	ks, vs := make([]K, 0, len(m)), make([]V, 0, len(m))
	for k, v := range m {
		ks = append(ks, k)
		vs = append(vs, v)
	}
	return ks, vs
}

// MapToSlice converts map keys and values to slice in indeterminate order.
// Deprecated: As of Go 1.18, please use standard library golang.org/x/exp/maps#Keys and Values.
func MapToSlice(a any) (ks, vs any) {
	ks, vs, _ = MapToSliceE(a)
	return
}

// MapToSliceE converts keys and values of map to slice in indeterminate order with error.
// Deprecated: As of Go 1.18, please use standard library golang.org/x/exp/maps#Keys and Values.
func MapToSliceE(a any) (ks, vs any, err error) {
	t := reflect.TypeOf(a)
	if t.Kind() != reflect.Map {
		err = fmt.Errorf("the input %#v of type %T isn't a map", a, a)
		return
	}

	// Convert.
	m := reflect.ValueOf(a)
	keys := m.MapKeys()
	ksT, vsT := reflect.SliceOf(t.Key()), reflect.SliceOf(t.Elem())
	ksV, vsV := reflect.MakeSlice(ksT, 0, m.Len()), reflect.MakeSlice(vsT, 0, m.Len())
	for _, k := range keys {
		ksV = reflect.Append(ksV, k)
		vsV = reflect.Append(vsV, m.MapIndex(k))
	}
	return ksV.Interface(), vsV.Interface(), nil
}

// SplitStrToSlice splits a string to a slice by the specified separator.
func SplitStrToSlice[T any](s, sep string) []T {
	v, _ := SplitStrToSliceE[T](s, sep)
	return v
}

// SplitStrToSliceE splits a string to a slice by the specified separator and returns an error if occurred.
// Note that this function is implemented through 1.18 generics, so the element type needs to
// be specified when calling it, e.g. SplitStrToSliceE[int]("1,2,3", ",").
func SplitStrToSliceE[T any](s, sep string) ([]T, error) {
	ss := strings.Split(s, sep)
	r := make([]T, len(ss))
	for i := range ss {
		v, err := ToAnyE[T](ss[i])
		if err != nil {
			return nil, err
		}
		r[i] = v
	}
	return r, nil
}
