package conv

import (
	"fmt"
	"strconv"

	huge "github.com/dablelv/go-huge-util"
)

// SplitStrToSlice splits a string to a slice by the specified separator.
func SplitStrToSlice[T any](s, sep string) []T {
	v, _ := SplitStrToSliceE[T](s, sep)
	return v
}

// SplitStrToSliceE splits a string to a slice by the specified separator and returns an error if occurred.
// Note that this function is implemented through 1.18 generics, so the element type needs to
// be specified when calling it, e.g. SplitStrToSliceE[int]("1,2,3", ",").
func SplitStrToSliceE[T any](s, sep string) ([]T, error) {
	ss := huge.Split(s, sep)
	dst := make([]T, len(ss))
	var t T
	for i, v := range ss {
		switch any(t).(type) {
		case string:
			dst[i] = any(v).(T)
		case int:
			v, err := strconv.ParseInt(v, 0, 0)
			if err != nil {
				return nil, err
			}
			dst[i] = any(int(v)).(T)
		case int8:
			v, err := strconv.ParseInt(v, 0, 8)
			if err != nil {
				return nil, err
			}
			dst[i] = any(int8(v)).(T)
		case int16:
			v, err := strconv.ParseInt(v, 0, 16)
			if err != nil {
				return nil, err
			}
			dst[i] = any(int16(v)).(T)
		case int32:
			v, err := strconv.ParseInt(v, 0, 32)
			if err != nil {
				return nil, err
			}
			dst[i] = any(int32(v)).(T)
		case int64:
			v, err := strconv.ParseInt(v, 0, 32)
			if err != nil {
				return nil, err
			}
			dst[i] = any(v).(T)
		case uint:
			v, err := strconv.ParseUint(v, 0, 0)
			if err != nil {
				return nil, err
			}
			dst[i] = any(uint(v)).(T)
		case uint8:
			v, err := strconv.ParseUint(v, 0, 8)
			if err != nil {
				return nil, err
			}
			dst[i] = any(uint8(v)).(T)
		case uint16:
			v, err := strconv.ParseUint(v, 0, 16)
			if err != nil {
				return nil, err
			}
			dst[i] = any(uint16(v)).(T)
		case uint32:
			v, err := strconv.ParseUint(v, 0, 32)
			if err != nil {
				return nil, err
			}
			dst[i] = any(uint32(v)).(T)
		case uint64:
			v, err := strconv.ParseUint(v, 0, 64)
			if err != nil {
				return nil, err
			}
			dst[i] = any(v).(T)
		case float32:
			v, err := strconv.ParseFloat(v, 32)
			if err != nil {
				return nil, err
			}
			dst[i] = any(float32(v)).(T)
		case float64:
			v, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return nil, err
			}
			dst[i] = any(v).(T)
		case bool:
			v, err := strconv.ParseBool(v)
			if err != nil {
				return nil, err
			}
			dst[i] = any(v).(T)
		default:
			return nil, fmt.Errorf("the type %T is not supported", t)
		}
	}
	return dst, nil
}
