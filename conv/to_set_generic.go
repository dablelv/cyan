package conv

import (
	"fmt"
	"reflect"
)

// ToSet converts a slice or array to map[T]struct{} and returns a nil if error occurred.
func ToSet[T comparable](i any) map[T]struct{} {
	m, _ := ToSetE[T](i)
	return m
}

// ToSetGE converts a slice or array to map[T]struct{} and returns an error if occurred.
// Note that the the element type of input must be equal to the map key type.
// Note that this function is implemented through 1.18 generics, so the element type needs to
// be specified when calling it, e.g. ToSetStrictE[int]([]int{1,2,3}).
func ToSetE[T comparable](i any) (map[T]struct{}, error) {
	// Check params.
	if i == nil {
		return nil, fmt.Errorf("the input i is nil")
	}
	t := reflect.TypeOf(i)
	kind := t.Kind()
	if kind != reflect.Slice && kind != reflect.Array {
		return nil, fmt.Errorf("the type %T of input %#v isn't a slice or array", i, i)
	}

	// Execute the conversion.
	v := reflect.ValueOf(i)
	mapT := reflect.MapOf(t.Elem(), reflect.TypeOf(struct{}{}))
	mapV := reflect.MakeMapWithSize(mapT, v.Len())
	for j := 0; j < v.Len(); j++ {
		mapV.SetMapIndex(v.Index(j), reflect.ValueOf(struct{}{}))
	}
	if v, ok := mapV.Interface().(map[T]struct{}); ok {
		return v, nil
	}
	var vt T
	return nil, fmt.Errorf("the input element type %v isn't %T", t.Elem().Name(), vt)
}
