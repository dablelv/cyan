package slice

import (
	"reflect"

	"github.com/dablelv/go-huge-util/math"
)

// RandomElem returns a random element from a slice or array.
// If the length of slice or array is zero it will panic.
func RandomElem(a any) any {
	v := reflect.ValueOf(a)
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		return a
	}
	i := math.RandIntn(v.Len())
	return v.Index(i).Interface()
}
