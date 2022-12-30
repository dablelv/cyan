package slice

import (
	"reflect"
)

//
// Note that after Go 1.18, this file is deprecated.
// Please use the standard lib function https://pkg.go.dev/golang.org/x/exp/slices#Contains and ContainsFunc
// implemented by generics.
//

// Contains reports whether slice or array contains the target element.
// Note that if the target element is a numeric literal, please specify its type explicitly,
// otherwise it defaults to int.
// E.g. you might call like IsContains([]int32{1,2,3}, int32(1)).
func Contains(a, target any) bool {
	if a == nil {
		return false
	}
	t := reflect.TypeOf(a)
	if t.Kind() != reflect.Slice && t.Kind() != reflect.Array {
		return false
	}
	v := reflect.ValueOf(a)
	for i := 0; i < v.Len(); i++ {
		if target == v.Index(i).Interface() {
			return true
		}
	}
	return false
}
