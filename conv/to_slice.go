package conv

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

//
// Desc: convert an any type value to the specified type slice.
//

// ToStrSlice converts an interface to a []string.
// For example, covert []int{1, 2, 3} to []string{"1", "2", "3"}.
func ToStrSlice(i any) []string {
	v, _ := ToStrSliceE(i)
	return v
}

// ToStrSliceE converts an interface to a []string with error.
func ToStrSliceE(i any) ([]string, error) {
	if i == nil {
		return nil, fmt.Errorf("unable to cast %#v of type %T to []string", i, i)
	}

	switch v := i.(type) {
	case []string:
		return v, nil
	}

	// If i is a slice or array.
	kind := reflect.TypeOf(i).Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		sl := reflect.ValueOf(i)
		s := make([]string, sl.Len())
		for j := 0; j < sl.Len(); j++ {
			v, err := ToStringE(sl.Index(j).Interface())
			if err != nil {
				return nil, err
			}
			s[j] = v
		}
		return s, nil
	}

	// If i is a single value.
	switch v := i.(type) {
	case string:
		return strings.Fields(v), nil
	default:
		s, err := ToStringE(v)
		if err != nil {
			return nil, err
		}
		return []string{s}, nil
	}
}

// ToIntSlice converts an interface to []int.
// E.g. covert []string{"1", "2", "3"} to []int{1, 2, 3}.
func ToIntSlice(i any) []int {
	v, _ := ToIntSliceE(i)
	return v
}

// ToIntSliceE converts any tpye to []int slice.
func ToIntSliceE(i interface{}) ([]int, error) {
	if i == nil {
		return []int{}, nil
	}

	switch v := i.(type) {
	case []int:
		return v, nil
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]int, s.Len())
		for j := 0; j < s.Len(); j++ {
			val, err := ToIntE(s.Index(j).Interface())
			if err != nil {
				return []int{}, fmt.Errorf("unable to cast %#v of type %T to []int", i, i)
			}
			a[j] = val
		}
		return a, nil
	default:
		return []int{}, fmt.Errorf("unable to cast %#v of type %T to []int", i, i)
	}
}

// ToInt8Slice converts an interface to []int8.
// For example, covert []string{"1", "2", "3"} to []int8{1, 2, 3}.
func ToInt8Slice(i any) []int8 {
	v, _ := ToInt8SliceE(i)
	return v
}

// ToInt8SliceE converts an interface to []int8 with error.
func ToInt8SliceE(i any) ([]int8, error) {
	if i == nil {
		return nil, fmt.Errorf("unable to cast %#v of type %T to []int8", i, i)
	}

	switch v := i.(type) {
	case []int8:
		return v, nil
	}

	// If i is a slice or array.
	kind := reflect.TypeOf(i).Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		sl := reflect.ValueOf(i)
		i8s := make([]int8, sl.Len())
		for j := 0; j < sl.Len(); j++ {
			v, err := ToInt8E(sl.Index(j).Interface())
			if err != nil {
				return nil, err
			}
			i8s[j] = v
		}
		return i8s, nil
	}

	// If i is a single value.
	v, err := ToInt8E(i)
	if err != nil {
		return nil, err
	}
	return []int8{v}, nil
}

// ToInt16Slice converts an interface to []int16.
// For example, covert []string{"1", "2", "3"} to []int16{1, 2, 3}.
func ToInt16Slice(i any) []int16 {
	v, _ := ToInt16SliceE(i)
	return v
}

// ToInt16SliceE converts an interface to []int16 with error.
func ToInt16SliceE(i any) ([]int16, error) {
	if i == nil {
		return nil, fmt.Errorf("unable to cast %#v of type %T to []int16", i, i)
	}

	switch v := i.(type) {
	case []int16:
		return v, nil
	}

	// If i is a slice or array.
	kind := reflect.TypeOf(i).Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		sl := reflect.ValueOf(i)
		i16s := make([]int16, sl.Len())
		for j := 0; j < sl.Len(); j++ {
			v, err := ToInt16E(sl.Index(j).Interface())
			if err != nil {
				return nil, err
			}
			i16s[j] = v
		}
		return i16s, nil
	}

	// If i is a single value.
	v, err := ToInt16E(i)
	if err != nil {
		return nil, err
	}
	return []int16{v}, nil
}

// ToInt32Slice converts an interface to []int32.
// For example, covert []string{"1", "2", "3"} to []int32{1, 2, 3}.
func ToInt32Slice(i any) []int32 {
	v, _ := ToInt32SliceE(i)
	return v
}

// ToInt32SliceE converts an interface to []int32 with error.
func ToInt32SliceE(i any) ([]int32, error) {
	if i == nil {
		return nil, fmt.Errorf("unable to cast %#v of type %T to []int32", i, i)
	}

	switch v := i.(type) {
	case []int32:
		return v, nil
	}

	// If i is a slice or array.
	kind := reflect.TypeOf(i).Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		sl := reflect.ValueOf(i)
		i32s := make([]int32, sl.Len())
		for j := 0; j < sl.Len(); j++ {
			v, err := ToInt32E(sl.Index(j).Interface())
			if err != nil {
				return nil, err
			}
			i32s[j] = v
		}
		return i32s, nil
	}

	// If i is a single value.
	v, err := ToInt32E(i)
	if err != nil {
		return nil, err
	}
	return []int32{v}, nil
}

// ToInt64Slice converts an interface to []int64 slice.
// For example, covert []string{"1", "2", "3"} to []int64{1, 2, 3}.
func ToInt64Slice(i any) []int64 {
	v, _ := ToInt64SliceE(i)
	return v
}

// ToInt64SliceE converts an interface to []int64 slice with error.
func ToInt64SliceE(i any) ([]int64, error) {
	if i == nil {
		return nil, fmt.Errorf("unable to cast %#v of type %T to []int64", i, i)
	}

	switch v := i.(type) {
	case []int64:
		return v, nil
	}

	// If i is a slice or array.
	kind := reflect.TypeOf(i).Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		sl := reflect.ValueOf(i)
		i64s := make([]int64, sl.Len())
		for j := 0; j < sl.Len(); j++ {
			v, err := ToInt64E(sl.Index(j).Interface())
			if err != nil {
				return nil, err
			}
			i64s[j] = v
		}
		return i64s, nil
	}

	// If i is a single value.
	v, err := ToInt64E(i)
	if err != nil {
		return nil, err
	}
	return []int64{v}, nil
}

// ToUintSlice converts an interface to []uint.
// For example, covert []string{"1", "2", "3"} to []uint{1, 2, 3}.
func ToUintSlice(i any) []uint {
	v, _ := ToUintSliceE(i)
	return v
}

// ToUintSliceE converts an interface to []uint with error.
func ToUintSliceE(i any) ([]uint, error) {
	if i == nil {
		return nil, fmt.Errorf("unable to cast %#v of type %T to []uint", i, i)
	}

	switch v := i.(type) {
	case []uint:
		return v, nil
	}

	// If i is a slice or array.
	kind := reflect.TypeOf(i).Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		sl := reflect.ValueOf(i)
		u := make([]uint, sl.Len())
		for j := 0; j < sl.Len(); j++ {
			v, err := ToUintE(sl.Index(j).Interface())
			if err != nil {
				return nil, err
			}
			u[j] = v
		}
		return u, nil
	}

	// If i is a single value.
	v, err := ToUintE(i)
	if err != nil {
		return nil, err
	}
	return []uint{v}, nil
}

// ToUint8Slice converts an interface to []uint8.
// E.g. covert []string{"1", "2", "3"} to []uint8{1, 2, 3}.
func ToUint8Slice(i any) []uint8 {
	v, _ := ToUint8SliceE(i)
	return v
}

// ToUint8SliceE converts an interface to []uint8 slice with error.
func ToUint8SliceE(i any) ([]uint8, error) {
	if i == nil {
		return nil, fmt.Errorf("unable to cast %#v of type %T to []uint8", i, i)
	}

	switch v := i.(type) {
	case []uint8:
		return v, nil
	}

	// if i is a slice or array.
	kind := reflect.TypeOf(i).Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		sl := reflect.ValueOf(i)
		u := make([]uint8, sl.Len())
		for j := 0; j < sl.Len(); j++ {
			v, err := ToUint8E(sl.Index(j).Interface())
			if err != nil {
				return nil, err
			}
			u[j] = v
		}
		return u, nil
	}

	// If i is a single value.
	u, err := ToUint8E(i)
	if err != nil {
		return nil, err
	}
	return []uint8{u}, nil
}

// ToUint16Slice converts an interface to []uint16.
// For example, covert []string{"1", "2", "3"} to []uint16{1, 2, 3}.
func ToUint16Slice(i any) []uint16 {
	v, _ := ToUint16SliceE(i)
	return v
}

// ToUint16SliceE converts an interface to []uint16 slice with error.
func ToUint16SliceE(i any) ([]uint16, error) {
	if i == nil {
		return nil, fmt.Errorf("unable to cast %#v of type %T to []uint16", i, i)
	}

	switch v := i.(type) {
	case []uint16:
		return v, nil
	}

	// If i is a slice or array.
	kind := reflect.TypeOf(i).Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		sl := reflect.ValueOf(i)
		u := make([]uint16, sl.Len())
		for j := 0; j < sl.Len(); j++ {
			v, err := ToUint16E(sl.Index(j).Interface())
			if err != nil {
				return nil, err
			}
			u[j] = v
		}
		return u, nil
	}

	// If i is a single value.
	u, err := ToUint16E(i)
	if err != nil {
		return nil, err
	}
	return []uint16{u}, nil
}

// ToUint32Slice converts an interface to []uint32.
// For example, covert []string{"1", "2", "3"} to []uint32{1, 2, 3}.
func ToUint32Slice(i any) []uint32 {
	v, _ := ToUint32SliceE(i)
	return v
}

// ToUint32SliceE converts an interface to []uint32 slice with error.
func ToUint32SliceE(i any) ([]uint32, error) {
	if i == nil {
		return nil, fmt.Errorf("unable to cast %#v of type %T to []uint32", i, i)
	}

	switch v := i.(type) {
	case []uint32:
		return v, nil
	}

	// If i is a slice or array.
	kind := reflect.TypeOf(i).Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		sl := reflect.ValueOf(i)
		u := make([]uint32, sl.Len())
		for j := 0; j < sl.Len(); j++ {
			v, err := ToUint32E(sl.Index(j).Interface())
			if err != nil {
				return nil, err
			}
			u[j] = v
		}
		return u, nil
	}

	// If i is a single value.
	u, err := ToUint32E(i)
	if err != nil {
		return nil, err
	}
	return []uint32{u}, nil
}

// ToUint64Slice converts an interface to []uint64.
// For example, covert []string{"1", "2", "3"} to []uint64{1, 2, 3}.
func ToUint64Slice(i any) []uint64 {
	v, _ := ToUint64SliceE(i)
	return v
}

// ToUint64SliceE converts an interface to []uint64 slice with error.
func ToUint64SliceE(i any) ([]uint64, error) {
	if i == nil {
		return nil, fmt.Errorf("unable to cast %#v of type %T to []uint64", i, i)
	}

	switch v := i.(type) {
	case []uint64:
		return v, nil
	}

	// If i is a slice or array.
	kind := reflect.TypeOf(i).Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		sl := reflect.ValueOf(i)
		u := make([]uint64, sl.Len())
		for j := 0; j < sl.Len(); j++ {
			v, err := ToUint64E(sl.Index(j).Interface())
			if err != nil {
				return nil, err
			}
			u[j] = v
		}
		return u, nil
	}

	// If i is a single value.
	u, err := ToUint64E(i)
	if err != nil {
		return nil, err
	}
	return []uint64{u}, nil
}

// ToByteSlice converts an interface to []byte.
// E.g. covert []string{"1", "2", "3"} to []byte{1, 2, 3}.
func ToByteSlice(i any) []byte {
	return ToUint8Slice(i)
}

// ToByteSliceE converts an interface to []byte slice with error
func ToByteSliceE(i any) ([]byte, error) {
	return ToUint8SliceE(i)
}

// ToBoolSlice converts an interface to []bool.
func ToBoolSlice(a any) []bool {
	v, _ := ToBoolSliceE(a)
	return v
}

// ToBoolSliceE converts an interface to []bool.
func ToBoolSliceE(i any) ([]bool, error) {
	if i == nil {
		return []bool{}, fmt.Errorf("unable to cast %#v of type %T to []bool", i, i)
	}

	switch v := i.(type) {
	case []bool:
		return v, nil
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]bool, s.Len())
		for j := 0; j < s.Len(); j++ {
			val, err := ToBoolE(s.Index(j).Interface())
			if err != nil {
				return []bool{}, fmt.Errorf("unable to cast %#v of type %T to []bool", i, i)
			}
			a[j] = val
		}
		return a, nil
	default:
		return []bool{}, fmt.Errorf("unable to cast %#v of type %T to []bool", i, i)
	}
}

// ToDurationSlice converts an interface to []time.Duration.
func ToDurationSlice(i any) []time.Duration {
	v, _ := ToDurationSliceE(i)
	return v
}

// ToDurationSliceE converts any type to []time.Duration with error.
func ToDurationSliceE(i interface{}) ([]time.Duration, error) {
	if i == nil {
		return []time.Duration{}, fmt.Errorf("unable to cast %#v of type %T to []time.Duration", i, i)
	}

	switch v := i.(type) {
	case []time.Duration:
		return v, nil
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]time.Duration, s.Len())
		for j := 0; j < s.Len(); j++ {
			val, err := ToDurationE(s.Index(j).Interface())
			if err != nil {
				return []time.Duration{}, fmt.Errorf("unable to cast %#v of type %T to []time.Duration", i, i)
			}
			a[j] = val
		}
		return a, nil
	default:
		return []time.Duration{}, fmt.Errorf("unable to cast %#v of type %T to []time.Duration", i, i)
	}
}
