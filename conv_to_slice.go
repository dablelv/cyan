package util

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/spf13/cast"
)

// Map2Slice converts keys and values of map to slice
func Map2Slice(i interface{}) (k []interface{}, v []interface{}) {
	k, v, _ = Map2SliceE(i)
	return
}

// Map2SliceE converts keys and values of map to slice with error
func Map2SliceE(i interface{})(ks []interface{}, vs []interface{}, err error) {
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
func Map2StrSlice(i interface{})(k []string, v []string) {
	k, v, _ = Map2StrSliceE(i)
	return
}

// Map2StrSliceE converts keys and values of map to string slice with error
func Map2StrSliceE(i interface{})(k []string, v []string, err error) {
	slK, slV, err := Map2SliceE(i)
	if err != nil {
		return
	}
	k, err = cast.ToStringSliceE(slK)
	if err != nil {
		return
	}
	v, err = cast.ToStringSliceE(slV)
	return
}

// ToStrSlice casts an interface to a []string type.
func ToStrSlice(i interface{}) []string {
	v, _ := ToStrSliceE(i)
	return v
}

// ToStrSliceE casts an interface to a []string slice with error
func ToStrSliceE(i interface{}) ([]string, error) {
	var s []string
	// if i is a slice or array
	kind := reflect.TypeOf(i).Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		sl := reflect.ValueOf(i)
		for j := 0; j < sl.Len(); j++ {
			v, err := cast.ToStringE(sl.Index(j).Interface())
			if err != nil {
				return nil, err
			}
			s = append(s, v)
		}
		return s, nil
	}
	// if i is a single value
	switch v := i.(type) {
	case string:
		return strings.Fields(v), nil
	case interface{}:
		str, err := cast.ToStringE(v)
		if err != nil {
			return nil, err
		}
		return []string{str}, nil
	default:
		return nil, fmt.Errorf("unable to cast %#v of type %T to []string", i, i)
	}
}

// ToUint64SliceE casts an interface to []uint64 slice with error
func ToUint64SliceE(i interface{}) ([]uint64, error) {
	var u []uint64
	// if i is a slice or array
	kind := reflect.TypeOf(i).Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		sl := reflect.ValueOf(i)
		for j := 0; j < sl.Len(); j++ {
			v, err := cast.ToUint64E(sl.Index(j).Interface())
			if err != nil {
				return nil, err
			}
			u = append(u, v)
		}
		return u, nil
	}
	// if i is a single value
	u64, err := cast.ToUint64E(i)
	if err != nil {
		return nil, err
	}
	return []uint64{u64}, nil
}