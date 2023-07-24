package slice

import (
	"reflect"
)

//
// Reverse a slice, e.g. input []int32{1, 2, 3} and output is []int32{3, 2, 1}.
//

// Reverse reverses the specified slice without modifying the original slice.
// Reverse implemented by generics is recommended to be used.
func Reverse[E any, S ~[]E](s S) S {
	if s == nil {
		return s
	}
	r := make([]E, 0, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		r = append(r, s[i])
	}
	return r
}

// ReverseRef implemented by reflect reverses the specified slice without modifying the original slice.
// If a isn't a slice ReverseRef will panic.
func ReverseRef(a any) any {
	v := reflect.ValueOf(a)
	if v.IsNil() {
		return a
	}
	r := reflect.MakeSlice(reflect.TypeOf(a), 0, v.Len())
	for i := v.Len() - 1; i >= 0; i-- {
		r = reflect.Append(r, v.Index(i))
	}
	return r.Interface()
}
