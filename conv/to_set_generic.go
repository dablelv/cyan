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

// ToSetE converts a slice or array to map[T]struct{} and returns an error if occurred.
// Note that the the element type of input don't need to be equal to the map key type.
// For example, []uint64{1, 2, 3} can be converted to map[uint64]struct{}{1:{},2:{},3:{}}
// and also can be converted to map[string]struct{}{"1":{},"2":{},"3":{}} if you want.
// Note that this function is implemented through 1.18 generics, so the element type needs to
// be specified when calling it, e.g. ToSetE[int]([]int{1,2,3}).
func ToSetE[T comparable](i any) (map[T]struct{}, error) {
	// Check param.
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
	// Convert the element to the type T.
	set := make(map[T]struct{}, v.Len())
	for _, k := range mapV.MapKeys() {
		v, err := ToAnyE[T](k.Interface())
		if err != nil {
			return nil, err
		}
		set[v] = struct{}{}
	}
	return set, nil
}
