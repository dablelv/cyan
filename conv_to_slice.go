package util

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/spf13/cast"
)

// Map2Slice converts keys and values of map to slice
func Map2Slice(i interface{}) (k []interface{}, v []interface{}) {
	k, v, _ = Map2SliceE(i)
	return
}

// Map2SliceE converts keys and values of map to slice with error
func Map2SliceE(i interface{}) (ks []interface{}, vs []interface{}, err error) {
	kind := reflect.TypeOf(i).Kind()
	if kind != reflect.Map {
		err = fmt.Errorf("the input %#v of type %T isn't a map", i, i)
		return
	}
	m := reflect.ValueOf(i)
	keys := m.MapKeys()
	ks, vs = make([]interface{}, 0, len(keys)), make([]interface{}, 0, len(keys))
	for _, k := range keys {
		ks = append(ks, k.Interface())
		v := m.MapIndex(k)
		vs = append(vs, v.Interface())
	}
	return
}

// Map2StrSlice converts keys and values of map to string slice
func Map2StrSlice(i interface{}) (k []string, v []string) {
	k, v, _ = Map2StrSliceE(i)
	return
}

// Map2StrSliceE converts keys and values of map to string slice with error
func Map2StrSliceE(i interface{}) (k []string, v []string, err error) {
	slK, slV, err := Map2SliceE(i)
	if err != nil {
		return
	}
	k, err = ToStrSliceE(slK)
	if err != nil {
		return
	}
	v, err = ToStrSliceE(slV)
	return
}

// ToStrSlice converts an interface to a []string type.
func ToStrSlice(i interface{}) []string {
	v, _ := ToStrSliceE(i)
	return v
}

// ToStrSliceE converts an interface to a []string slice with error.
func ToStrSliceE(i interface{}) ([]string, error) {
	if i == nil {
		return nil, fmt.Errorf("unable to cast %#v of type %T to []string", i, i)
	}

	switch v := i.(type) {
	case []string:
		return v, nil
	}

	// if i is a slice or array
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

	// if i is a single value
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

// ToUintSlice converts an interface to []uint slice
func ToUintSlice(i interface{}) []uint {
	v, _ := ToUintSliceE(i)
	return v
}

// ToUintSliceE converts an interface to []uint slice with error
func ToUintSliceE(i interface{}) ([]uint, error) {
	if i == nil {
		return nil, fmt.Errorf("unable to cast %#v of type %T to []uint", i, i)
	}

	switch v := i.(type) {
	case []uint:
		return v, nil
	}

	// if i is a slice or array
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

	// if i is a single value
	v, err := cast.ToUintE(i)
	if err != nil {
		return nil, err
	}
	return []uint{v}, nil
}

// ToInt8Slice converts an interface to []int8 slice
func ToInt8Slice(i interface{}) []int8 {
	v, _ := ToInt8SliceE(i)
	return v
}

// ToInt8SliceE converts an interface to []int8 slice with error
func ToInt8SliceE(i interface{}) ([]int8, error) {
	if i == nil {
		return nil, fmt.Errorf("unable to cast %#v of type %T to []int8", i, i)
	}

	switch v := i.(type) {
	case []int8:
		return v, nil
	}

	// if i is a slice or array
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

	// if i is a single value
	v, err := cast.ToInt8E(i)
	if err != nil {
		return nil, err
	}
	return []int8{v}, nil
}

// ToInt16Slice converts an interface to []int16 slice
func ToInt16Slice(i interface{}) []int16 {
	v, _ := ToInt16SliceE(i)
	return v
}

// ToInt16SliceE converts an interface to []int16 slice with error
func ToInt16SliceE(i interface{}) ([]int16, error) {
	if i == nil {
		return nil, fmt.Errorf("unable to cast %#v of type %T to []int16", i, i)
	}

	switch v := i.(type) {
	case []int16:
		return v, nil
	}

	// if i is a slice or array
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

	// if i is a single value
	v, err := cast.ToInt16E(i)
	if err != nil {
		return nil, err
	}
	return []int16{v}, nil
}

// ToInt32Slice converts an interface to []int32 slice
func ToInt32Slice(i interface{}) []int32 {
	v, _ := ToInt32SliceE(i)
	return v
}

// ToInt32SliceE converts an interface to []int32 slice with error
func ToInt32SliceE(i interface{}) ([]int32, error) {
	if i == nil {
		return nil, fmt.Errorf("unable to cast %#v of type %T to []int32", i, i)
	}

	switch v := i.(type) {
	case []int32:
		return v, nil
	}

	// if i is a slice or array
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

	// if i is a single value
	v, err := cast.ToInt32E(i)
	if err != nil {
		return nil, err
	}
	return []int32{v}, nil
}

// ToInt64Slice converts an interface to []int64 slice
func ToInt64Slice(i interface{}) []int64 {
	v, _ := ToInt64SliceE(i)
	return v
}

// ToInt64SliceE converts an interface to []int64 slice with error
func ToInt64SliceE(i interface{}) ([]int64, error) {
	if i == nil {
		return nil, fmt.Errorf("unable to cast %#v of type %T to []int64", i, i)
	}

	switch v := i.(type) {
	case []int64:
		return v, nil
	}

	// if i is a slice or array
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

	// if i is a single value
	v, err := cast.ToInt64E(i)
	if err != nil {
		return nil, err
	}
	return []int64{v}, nil
}

// ToUint8Slice converts an interface to []uint8 slice
func ToUint8Slice(i interface{}) []uint8 {
	v, _ := ToUint8SliceE(i)
	return v
}

// ToUint8SliceE converts an interface to []uint8 slice with error
func ToUint8SliceE(i interface{}) ([]uint8, error) {
	if i == nil {
		return nil, fmt.Errorf("unable to cast %#v of type %T to []uint8", i, i)
	}

	switch v := i.(type) {
	case []uint8:
		return v, nil
	}

	// if i is a slice or array
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

	// if i is a single value
	u, err := cast.ToUint8E(i)
	if err != nil {
		return nil, err
	}
	return []uint8{u}, nil
}

// ToUint16Slice converts an interface to []uint16 slice
func ToUint16Slice(i interface{}) []uint16 {
	v, _ := ToUint16SliceE(i)
	return v
}

// ToUint16SliceE converts an interface to []uint16 slice with error
func ToUint16SliceE(i interface{}) ([]uint16, error) {
	if i == nil {
		return nil, fmt.Errorf("unable to cast %#v of type %T to []uint16", i, i)
	}

	switch v := i.(type) {
	case []uint16:
		return v, nil
	}

	// if i is a slice or array
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

	// if i is a single value
	u, err := cast.ToUint16E(i)
	if err != nil {
		return nil, err
	}
	return []uint16{u}, nil
}

// ToUint32Slice converts an interface to []uint32 slice
func ToUint32Slice(i interface{}) []uint32 {
	v, _ := ToUint32SliceE(i)
	return v
}

// ToUint32SliceE converts an interface to []uint32 slice with error
func ToUint32SliceE(i interface{}) ([]uint32, error) {
	if i == nil {
		return nil, fmt.Errorf("unable to cast %#v of type %T to []uint32", i, i)
	}

	switch v := i.(type) {
	case []uint32:
		return v, nil
	}

	// if i is a slice or array
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

	// if i is a single value
	u, err := cast.ToUint32E(i)
	if err != nil {
		return nil, err
	}
	return []uint32{u}, nil
}

// ToUint64Slice converts an interface to []uint64 slice
func ToUint64Slice(i interface{}) []uint64 {
	v, _ := ToUint64SliceE(i)
	return v
}

// ToUint64SliceE converts an interface to []uint64 slice with error
func ToUint64SliceE(i interface{}) ([]uint64, error) {
	if i == nil {
		return nil, fmt.Errorf("unable to cast %#v of type %T to []uint64", i, i)
	}

	switch v := i.(type) {
	case []uint64:
		return v, nil
	}

	// if i is a slice or array
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

	// if i is a single value
	u, err := cast.ToUint64E(i)
	if err != nil {
		return nil, err
	}
	return []uint64{u}, nil
}

// ToByteSlice converts an interface to []byte slice
func ToByteSlice(i interface{}) []byte {
	return ToUint8Slice(i)
}

// ToByteSliceE converts an interface to []byte slice with error
func ToByteSliceE(i interface{}) ([]byte, error) {
	return ToUint8SliceE(i)
}

// ToIntSlice converts an interface to []int slice.
// Please note that ToIntSlice is alias to cast.ToIntSlice.
func ToIntSlice(i interface{}) []int {
	v, _ := cast.ToIntSliceE(i)
	return v
}

// ToBoolSlice converts an interface to []bool slice.
// Please note that ToBoolSlice is alias to cast.ToBoolSlice.
func ToBoolSlice(i interface{}) []bool {
	v, _ := cast.ToBoolSliceE(i)
	return v
}

// ToDurationSlice converts an interface to []time.Duration slice.
// Please note that ToDurationSlice is alias to cast.ToDurationSlice.
func ToDurationSlice(i interface{}) []time.Duration {
	v, _ := cast.ToDurationSliceE(i)
	return v
}
