package slice

import (
	"fmt"
	"reflect"

	"github.com/dablelv/go-huge-util/math"
)

func GetElemIndexesSlice(slice interface{}, value interface{}) []int {
	indexes, _ := GetElemIndexesSliceE(slice, value)
	return indexes
}

// GetEleIndexesSliceE finds all indexes of the specified element in a slice.
func GetElemIndexesSliceE(slice interface{}, value interface{}) ([]int, error) {
	// check params
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return nil, fmt.Errorf("the input %#v of type %T isn't a slice", slice, slice)
	}
	// get indexes
	var indexes []int
	for i := 0; i < v.Len(); i++ {
		if v.Index(i).Interface() == value {
			indexes = append(indexes, i)
		}
	}
	return indexes, nil
}

// GetRandomSliceElem get a random element from a slice or array.
// If the length of slice or array is zero it will panic.
func GetRandomSliceElem(i interface{}) interface{} {
	v := reflect.ValueOf(i)
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		return i
	}
	idx := math.RandIntn(v.Len())
	return v.Index(idx).Interface()
}
