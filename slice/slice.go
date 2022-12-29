package slice

import (
	"fmt"
	"reflect"

	"github.com/dablelv/go-huge-util/conv"
)

// SumSlice calculates the sum of slice elements.
// E.g. input []int32{1, 2, 3} and output is 6.
func SumSlice(slice interface{}) float64 {
	v, _ := SumSliceE(slice)
	return v
}

// SumSliceE returns the sum of slice elements and an error if occurred.
func SumSliceE(slice interface{}) (float64, error) {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return 0.0, fmt.Errorf("the input %#v of type %T isn't a slice", slice, slice)
	}

	var sum float64
	for i := 0; i < v.Len(); i++ {
		switch v := v.Index(i).Interface().(type) {
		case int:
			sum += float64(v)
		case int8:
			sum += float64(v)
		case int16:
			sum += float64(v)
		case int32:
			sum += float64(v)
		case int64:
			sum += float64(v)
		case uint:
			sum += float64(v)
		case uint8:
			sum += float64(v)
		case uint16:
			sum += float64(v)
		case uint32:
			sum += float64(v)
		case uint64:
			sum += float64(v)
		case float32:
			sum += float64(v)
		case float64:
			sum += v
		default:
			return 0.0, fmt.Errorf("the element %#v of slice type %T isn't numerical type", v, v)
		}
	}
	return sum, nil
}

// IsContains checks whether slice or array contains the target element.
// Note that if the target element is a numeric literal, please specify its type explicitly, otherwise it defaults to int.
// For example you might call like IsContains([]int32{1,2,3}, int32(1)).
func IsContains(i interface{}, target interface{}) bool {
	if i == nil {
		return false
	}
	t := reflect.TypeOf(i)
	if t.Kind() != reflect.Slice && t.Kind() != reflect.Array {
		return false
	}
	v := reflect.ValueOf(i)
	for i := 0; i < v.Len(); i++ {
		if target == v.Index(i).Interface() {
			return true
		}
	}
	return false
}

// JoinSliceWithSep joins all elements in slice with a separator.
// E.g. input []int32{1, 2, 3} and separator "#", output is a string "1#2#3".
func JoinSliceWithSep(slice interface{}, sep string) string {
	s, _ := JoinSliceWithSepE(slice, sep)
	return s
}

// JoinSliceWithSepE joins all elements in slice or array with separator and return an error if occurred.
func JoinSliceWithSepE(slice interface{}, sep string) (string, error) {
	// Check param.
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		return "", fmt.Errorf("the input %#v of type %T isn't a slice or array", slice, slice)
	}
	// Join the slice or array.
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
