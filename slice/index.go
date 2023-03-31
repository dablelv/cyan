package slice

import (
	"fmt"
	"reflect"

	"github.com/dablelv/go-huge-util/math"
)

//
// Note that after Go 1.18, this file is deprecated.
// Please use the standard lib function https://pkg.go.dev/golang.org/x/exp/slices#Index and IndexFunc
// implemented by generics.
//

// GetElemIndexesSlice finds all indexes for the specified element in a slice.
func GetElemIndexesSlice(slice any, value any) []int {
	indexes, _ := GetElemIndexesSliceE(slice, value)
	return indexes
}

// GetEleIndexesSliceE finds all indexes for the specified element in a slice
// and returns an error if error occurred.
func GetElemIndexesSliceE(slice any, value any) ([]int, error) {
	// Check param.
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return nil, fmt.Errorf("the input %#v of type %T isn't a slice", slice, slice)
	}
	// Get indexes.
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
func GetRandomSliceElem(i any) any {
	v := reflect.ValueOf(i)
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		return i
	}
	idx := math.RandIntn(v.Len())
	return v.Index(idx).Interface()
}
