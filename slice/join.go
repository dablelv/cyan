package slice

import (
	"fmt"
	"reflect"

	"github.com/dablelv/go-huge-util/conv"
)

// Join joins all elements of slice or array with a separator.
// E.g. input []int32{1, 2, 3} and separator "#", output is a string "1#2#3".
func Join(a any, sep string) string {
	s, _ := JoinE(a, sep)
	return s
}

// JoinE joins all elements in slice or array with separator with returned error.
func JoinE(a any, sep string) (string, error) {
	// Check arguments.
	v := reflect.ValueOf(a)
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		return "", fmt.Errorf("the input %#v of type %T isn't a slice or array", a, a)
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
