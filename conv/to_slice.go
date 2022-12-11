package conv

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/spf13/cast"
)

//
// Part 1: convert map keys and values to slice in random order.
// For example, covert map[string]int{"a":1,"b":2, "c":3} to []string{"a", "c", "b"} and []int{1, 3, 2}.
//

// Map2Slice converts map keys and values to slice in random order.
func Map2Slice(i any) (ks any, vs any) {
	ks, vs, _ = Map2SliceE(i)
	return
}

// Map2SliceE converts keys and values of map to slice in random order with error.
func Map2SliceE(i any) (ks any, vs any, err error) {
	// Check params.
	if i == nil {
		return nil, nil, fmt.Errorf("unable to converts %#v of type %T to map[any]struct{}", i, i)
	}
	t := reflect.TypeOf(i)
	if t.Kind() != reflect.Map {
		err = fmt.Errorf("the input %#v of type %T isn't a map", i, i)
		return
	}

	// Execute the conversion.
	m := reflect.ValueOf(i)
	l := m.Len()
	keys := m.MapKeys()
	ksT, vsT := reflect.SliceOf(t.Key()), reflect.SliceOf(t.Elem())
	ksV, vsV := reflect.MakeSlice(ksT, 0, l), reflect.MakeSlice(vsT, 0, l)
	for _, k := range keys {
		ksV = reflect.Append(ksV, k)
		vsV = reflect.Append(vsV, m.MapIndex(k))
	}
	return ksV.Interface(), vsV.Interface(), nil
}

//
// Part 2: convert an any type value to the specified type slice.
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

	// if i is a slice or array.
	kind := reflect.TypeOf(i).Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		sl := reflect.ValueOf(i)
		s := make([]string, sl.Len())
		for j := 0; j < sl.Len(); j++ {
			v, err := cast.ToStringE(sl.Index(j).Interface())
			if err != nil {
				return nil, err
			}
			s[j] = v
		}
		return s, nil
	}

	// if i is a single value.
	switch v := i.(type) {
	case string:
		return strings.Fields(v), nil
	default:
		s, err := cast.ToStringE(v)
		if err != nil {
			return nil, err
		}
		return []string{s}, nil
	}
}

// ToIntSlice converts an interface to []int.
// For example, covert []string{"1", "2", "3"} to []int{1, 2, 3}.
func ToIntSlice(i any) []int {
	v, _ := ToIntSliceE(i)
	return v
}

// ToIntSliceE converts an interface to []int slice.
// Note that ToIntSliceE is alias to cast.ToIntSliceE.
func ToIntSliceE(i any) ([]int, error) {
	return cast.ToIntSliceE(i)
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

	// if i is a slice or array.
	kind := reflect.TypeOf(i).Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		sl := reflect.ValueOf(i)
		i8s := make([]int8, sl.Len())
		for j := 0; j < sl.Len(); j++ {
			v, err := cast.ToInt8E(sl.Index(j).Interface())
			if err != nil {
				return nil, err
			}
			i8s[j] = v
		}
		return i8s, nil
	}

	// if i is a single value.
	v, err := cast.ToInt8E(i)
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

	// if i is a slice or array.
	kind := reflect.TypeOf(i).Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		sl := reflect.ValueOf(i)
		i16s := make([]int16, sl.Len())
		for j := 0; j < sl.Len(); j++ {
			v, err := cast.ToInt16E(sl.Index(j).Interface())
			if err != nil {
				return nil, err
			}
			i16s[j] = v
		}
		return i16s, nil
	}

	// if i is a single value.
	v, err := cast.ToInt16E(i)
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

	// if i is a slice or array.
	kind := reflect.TypeOf(i).Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		sl := reflect.ValueOf(i)
		i32s := make([]int32, sl.Len())
		for j := 0; j < sl.Len(); j++ {
			v, err := cast.ToInt32E(sl.Index(j).Interface())
			if err != nil {
				return nil, err
			}
			i32s[j] = v
		}
		return i32s, nil
	}

	// if i is a single value.
	v, err := cast.ToInt32E(i)
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

	// if i is a slice or array.
	kind := reflect.TypeOf(i).Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		sl := reflect.ValueOf(i)
		i64s := make([]int64, sl.Len())
		for j := 0; j < sl.Len(); j++ {
			v, err := cast.ToInt64E(sl.Index(j).Interface())
			if err != nil {
				return nil, err
			}
			i64s[j] = v
		}
		return i64s, nil
	}

	// if i is a single value.
	v, err := cast.ToInt64E(i)
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

	// if i is a slice or array.
	kind := reflect.TypeOf(i).Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		sl := reflect.ValueOf(i)
		u := make([]uint, sl.Len())
		for j := 0; j < sl.Len(); j++ {
			v, err := cast.ToUintE(sl.Index(j).Interface())
			if err != nil {
				return nil, err
			}
			u[j] = v
		}
		return u, nil
	}

	// if i is a single value.
	v, err := cast.ToUintE(i)
	if err != nil {
		return nil, err
	}
	return []uint{v}, nil
}

// ToUint8Slice converts an interface to []uint8.
// For example, covert []string{"1", "2", "3"} to []uint8{1, 2, 3}.
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
			v, err := cast.ToUint8E(sl.Index(j).Interface())
			if err != nil {
				return nil, err
			}
			u[j] = v
		}
		return u, nil
	}

	// if i is a single value.
	u, err := cast.ToUint8E(i)
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

	// if i is a slice or array.
	kind := reflect.TypeOf(i).Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		sl := reflect.ValueOf(i)
		u := make([]uint16, sl.Len())
		for j := 0; j < sl.Len(); j++ {
			v, err := cast.ToUint16E(sl.Index(j).Interface())
			if err != nil {
				return nil, err
			}
			u[j] = v
		}
		return u, nil
	}

	// if i is a single value.
	u, err := cast.ToUint16E(i)
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

	// if i is a slice or array.
	kind := reflect.TypeOf(i).Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		sl := reflect.ValueOf(i)
		u := make([]uint32, sl.Len())
		for j := 0; j < sl.Len(); j++ {
			v, err := cast.ToUint32E(sl.Index(j).Interface())
			if err != nil {
				return nil, err
			}
			u[j] = v
		}
		return u, nil
	}

	// if i is a single value.
	u, err := cast.ToUint32E(i)
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

	// if i is a slice or array.
	kind := reflect.TypeOf(i).Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		sl := reflect.ValueOf(i)
		u := make([]uint64, sl.Len())
		for j := 0; j < sl.Len(); j++ {
			v, err := cast.ToUint64E(sl.Index(j).Interface())
			if err != nil {
				return nil, err
			}
			u[j] = v
		}
		return u, nil
	}

	// if i is a single value.
	u, err := cast.ToUint64E(i)
	if err != nil {
		return nil, err
	}
	return []uint64{u}, nil
}

// ToByteSlice converts an interface to []byte.
// For example, covert []string{"1", "2", "3"} to []byte{1, 2, 3}.
func ToByteSlice(i any) []byte {
	return ToUint8Slice(i)
}

// ToByteSliceE converts an interface to []byte slice with error
func ToByteSliceE(i any) ([]byte, error) {
	return ToUint8SliceE(i)
}

// ToBoolSlice converts an interface to []bool.
// Note that ToBoolSlice is alias to cast.ToBoolSlice.
func ToBoolSlice(i any) []bool {
	return cast.ToBoolSlice(i)
}

// ToBoolSliceE converts an interface to []bool.
// Note that ToBoolSliceE is alias to cast.ToBoolSliceE.
func ToBoolSliceE(i any) ([]bool, error) {
	return cast.ToBoolSliceE(i)
}

// ToDurationSlice converts an interface to []time.Duration.
// Note that ToDurationSlice is alias to cast.ToDurationSlice.
func ToDurationSlice(i any) []time.Duration {
	return cast.ToDurationSlice(i)
}

// ToDurationSliceE converts an interface to []time.Duration with error.
// Note that ToDurationSliceE is alias to cast.ToDurationSliceE.
func ToDurationSliceE(i any) ([]time.Duration, error) {
	return cast.ToDurationSliceE(i)
}
