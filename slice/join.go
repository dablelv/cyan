package slice

import (
	"fmt"
	"reflect"

	"github.com/dablelv/go-huge-util/conv"
)

// JoinSliceWithSep joins all elements in slice with a separator.
// E.g. input []int32{1, 2, 3} and separator "#", output is a string "1#2#3".
func JoinSliceWithSep(slice any, sep string) string {
	s, _ := JoinSliceWithSepE(slice, sep)
	return s
}

// JoinSliceWithSepE joins all elements in slice or array with separator and return an error if occurred.
func JoinSliceWithSepE(slice any, sep string) (string, error) {
	// Check param.
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		return "", fmt.Errorf("the input %#v of type %T isn't a slice or array", slice, slice)
	}
	// Join slice or array.
	var s string
	for i := 0; i < v.Len(); i++ {
		if len(s) > 0 {
			s += sep
		}
		str, err := conv.ToStringE(v.Index(i).Interface())
		if err != nil {
			return "", err
		}
		s += str
	}
	return s, nil
}
