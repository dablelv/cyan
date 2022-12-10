package conv

import (
	"strconv"

	"github.com/spf13/cast"

	huge "github.com/dablelv/go-huge-util"
)

//
// Desc: split a string to the specified type slice by the specified separator.
//

// SplitStrToSlice splits a string to a slice by the specified separator.
func SplitStrToSlice[T any](s, sep string) []T {
	v, _ := SplitStrToSliceE[T](s, sep)
	return v
}

// SplitStrToSliceE splits a string to a slice by the specified separator and returns an error if occurred.
func SplitStrToSliceE[T any](s, sep string) ([]T, error) {
	ss := huge.Split(s, sep)
	dst := make([]T, len(ss))
	var t T
	for i, v := range ss {
		switch any(t).(type) {
		case string:
			dst[i] = v
		case int:
			v, err := strconv.ParseInt(v, 0, 0)
			if err != nil {
				return nil, err
			}
			dst[i] = int(v)
		case int8:

		}
	}
	return dst, nil
}

// SplitStrToIntSlice splits a string to []int by the specified separator.
// For example, split "1,2,3" to []int{1,2,3} by comma.
func SplitStrToIntSlice(s, sep string) []int {
	ss := huge.Split(s, sep)
	dst := make([]int, len(ss))
	for i, v := range ss {
		dst[i] = cast.ToInt(v)
	}
	return dst
}

// SplitStrToInt8Slice splits a string to []int8 by the specified separator.
// For example, split "1,2,3" to []int8{1,2,3} by comma.
func SplitStrToInt8Slice(s, sep string) []int8 {
	arr := huge.Split(s, sep)
	dst := make([]int8, len(arr))
	for i, v := range arr {
		dst[i] = cast.ToInt8(v)
	}
	return dst
}

// SplitStrToInt16Slice splits a string to []int16 by the specified separator.
// For example, split "1,2,3" to []int16{1,2,3} by comma.
func SplitStrToInt16Slice(s, sep string) []int16 {
	arr := huge.Split(s, sep)
	dst := make([]int16, len(arr))
	for i, v := range arr {
		dst[i] = cast.ToInt16(v)
	}
	return dst
}

// SplitStrToInt32Slice splits a string to []int16 by the specified separator.
// For example, split "1,2,3" to []int32{1,2,3} by comma.
func SplitStrToInt32Slice(s, sep string) []int32 {
	arr := huge.Split(s, sep)
	dst := make([]int32, len(arr))
	for i, v := range arr {
		dst[i] = cast.ToInt32(v)
	}
	return dst
}

// SplitStrToInt64Slice splits a string to []int64 by the specified separator.
// For example, split "1,2,3" to []int64{1,2,3} by comma.
func SplitStrToInt64Slice(s, sep string) []int64 {
	arr := huge.Split(s, sep)
	dst := make([]int64, len(arr))
	for i, v := range arr {
		dst[i] = cast.ToInt64(v)
	}
	return dst
}

// SplitStrToUintSlice splits a string to []int64 by the specified separator.
// For example, split "1,2,3" to []uint{1,2,3} by comma.
func SplitStrToUintSlice(s, sep string) []uint {
	arr := huge.Split(s, sep)
	dst := make([]uint, len(arr))
	for i, v := range arr {
		dst[i] = cast.ToUint(v)
	}
	return dst
}

// SplitStrToUint8Slice splits a string to []uint8 by the specified separator.
// For example, split "1,2,3" to []uint8{1,2,3} by comma.
func SplitStrToUint8Slice(s, sep string) []uint8 {
	arr := huge.Split(s, sep)
	dst := make([]uint8, len(arr))
	for i, v := range arr {
		dst[i] = cast.ToUint8(v)
	}
	return dst
}

// SplitStrToUint16Slice splits a string to []uint16 by the specified separator.
// For example, split "1,2,3" to []uint16{1,2,3} by comma.
func SplitStrToUint16Slice(s, sep string) []uint16 {
	arr := huge.Split(s, sep)
	dst := make([]uint16, len(arr))
	for i, v := range arr {
		dst[i] = cast.ToUint16(v)
	}
	return dst
}

// SplitStrToUint32Slice splits a string to []uint32 by the specified separator.
// For example, split "1,2,3" to []uint32{1,2,3} by comma.
func SplitStrToUint32Slice(s, sep string) []uint32 {
	arr := huge.Split(s, sep)
	dst := make([]uint32, len(arr))
	for i, v := range arr {
		dst[i] = cast.ToUint32(v)
	}
	return dst
}

// SplitStrToUint64Slice splits a string to []uint64 by the specified separator.
// For example, split "1,2,3" to []uint64{1,2,3} by comma.
func SplitStrToUint64Slice(s, sep string) []uint64 {
	arr := huge.Split(s, sep)
	dst := make([]uint64, len(arr))
	for i, v := range arr {
		dst[i] = cast.ToUint64(v)
	}
	return dst
}

// SplitStrToFloat32Slice splits a string to []float32 by the specified separator.
// For example, split "1.1,2.2,3.3" to []float32{1.1,2.2,3.3} by comma.
func SplitStrToFloat32Slice(s string, sep string) []float32 {
	arr := huge.Split(s, sep)
	dst := make([]float32, len(arr))
	for i, v := range arr {
		dst[i] = cast.ToFloat32(v)
	}
	return dst
}

// SplitStrToFloat64Slice splits a string to []float32 by the specified separator.
// For example, split "1.1,2.2,3.3" to []float64{1.1,2.2,3.3} by comma.
func SplitStrToFloat64Slice(s string, sep string) []float64 {
	arr := huge.Split(s, sep)
	dst := make([]float64, len(arr))
	for i, v := range arr {
		dst[i] = cast.ToFloat64(v)
	}
	return dst
}
