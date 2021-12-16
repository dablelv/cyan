package util

import (
	"errors"
	"reflect"
)

// DeleteSliceElms deletes the specified elements from the slice.
// Note that the original slice will not be modified.
func DeleteSliceElms(i interface{}, elms ...interface{}) interface{} {
	res, _ := DeleteSliceElmsE(i, elms...)
	return res
}

// DeleteSliceElmsE deletes the specified elements from the slice.
// Note that the original slice will not be modified.
func DeleteSliceElmsE(i interface{}, elms ...interface{})(interface{}, error) {
	// check params
	v := reflect.ValueOf(i)
	if v.Kind() != reflect.Slice {
		return nil, errors.New("the input isn't a slice")
	}
	if v.Len() == 0 || len(elms) == 0 {
		return i, nil
	}
	if reflect.TypeOf(i).Elem() != reflect.TypeOf(elms[0]) {
		return nil, errors.New("element type is ill")
	}
	// convert the elements to map set
	m := make(map[interface{}]struct{})
	for _, v := range elms {
		m[v] = struct{}{}
	}
	// filter out specified elements
	t := reflect.MakeSlice(reflect.TypeOf(i), 0, v.Len())
	for i := 0; i < v.Len(); i++ {
		if _, ok := m[v.Index(i).Interface()]; !ok {
			t = reflect.Append(t, v.Index(i))
		}
	}
	return t.Interface(), nil
}
