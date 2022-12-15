package conv

import (
	"github.com/spf13/cast"

	"github.com/dablelv/go-huge-util/str"
)

//
// Desc: split a string to the specified type slice by the specified separator.
//

// SplitStrToIntSlice splits a string to []int by the specified separator.
// For example, split "1,2,3" to []int{1,2,3} by comma.
func SplitStrToIntSlice(s, sep string) []int {
	ss := str.Split(s, sep)
	dst := make([]int, len(ss))
	for i, v := range ss {
		dst[i] = cast.ToInt(v)
	}
	return dst
}

// SplitStrToInt8Slice splits a string to []int8 by the specified separator.
// For example, split "1,2,3" to []int8{1,2,3} by comma.
func SplitStrToInt8Slice(s, sep string) []int8 {
	arr := str.Split(s, sep)
	dst := make([]int8, len(arr))
	for i, v := range arr {
		dst[i] = cast.ToInt8(v)
	}
	return dst
}

// SplitStrToInt16Slice splits a string to []int16 by the specified separator.
// For example, split "1,2,3" to []int16{1,2,3} by comma.
func SplitStrToInt16Slice(s, sep string) []int16 {
	arr := str.Split(s, sep)
	dst := make([]int16, len(arr))
	for i, v := range arr {
		dst[i] = cast.ToInt16(v)
	}
	return dst
}

// SplitStrToInt32Slice splits a string to []int16 by the specified separator.
// For example, split "1,2,3" to []int32{1,2,3} by comma.
func SplitStrToInt32Slice(s, sep string) []int32 {
	arr := str.Split(s, sep)
	dst := make([]int32, len(arr))
	for i, v := range arr {
		dst[i] = cast.ToInt32(v)
	}
	return dst
}

// SplitStrToInt64Slice splits a string to []int64 by the specified separator.
// For example, split "1,2,3" to []int64{1,2,3} by comma.
func SplitStrToInt64Slice(s, sep string) []int64 {
	arr := str.Split(s, sep)
	dst := make([]int64, len(arr))
	for i, v := range arr {
		dst[i] = cast.ToInt64(v)
	}
	return dst
}

// SplitStrToUintSlice splits a string to []int64 by the specified separator.
// For example, split "1,2,3" to []uint{1,2,3} by comma.
func SplitStrToUintSlice(s, sep string) []uint {
	arr := str.Split(s, sep)
	dst := make([]uint, len(arr))
	for i, v := range arr {
		dst[i] = cast.ToUint(v)
	}
	return dst
}

// SplitStrToUint8Slice splits a string to []uint8 by the specified separator.
// For example, split "1,2,3" to []uint8{1,2,3} by comma.
func SplitStrToUint8Slice(s, sep string) []uint8 {
	arr := str.Split(s, sep)
	dst := make([]uint8, len(arr))
	for i, v := range arr {
		dst[i] = cast.ToUint8(v)
	}
	return dst
}

// SplitStrToUint16Slice splits a string to []uint16 by the specified separator.
// For example, split "1,2,3" to []uint16{1,2,3} by comma.
func SplitStrToUint16Slice(s, sep string) []uint16 {
	arr := str.Split(s, sep)
	dst := make([]uint16, len(arr))
	for i, v := range arr {
		dst[i] = cast.ToUint16(v)
	}
	return dst
}

// SplitStrToUint32Slice splits a string to []uint32 by the specified separator.
// For example, split "1,2,3" to []uint32{1,2,3} by comma.
func SplitStrToUint32Slice(s, sep string) []uint32 {
	arr := str.Split(s, sep)
	dst := make([]uint32, len(arr))
	for i, v := range arr {
		dst[i] = cast.ToUint32(v)
	}
	return dst
}

// SplitStrToUint64Slice splits a string to []uint64 by the specified separator.
// For example, split "1,2,3" to []uint64{1,2,3} by comma.
func SplitStrToUint64Slice(s, sep string) []uint64 {
	arr := str.Split(s, sep)
	dst := make([]uint64, len(arr))
	for i, v := range arr {
		dst[i] = cast.ToUint64(v)
	}
	return dst
}

// SplitStrToFloat32Slice splits a string to []float32 by the specified separator.
// For example, split "1.1,2.2,3.3" to []float32{1.1,2.2,3.3} by comma.
func SplitStrToFloat32Slice(s string, sep string) []float32 {
	arr := str.Split(s, sep)
	dst := make([]float32, len(arr))
	for i, v := range arr {
		dst[i] = cast.ToFloat32(v)
	}
	return dst
}

// SplitStrToFloat64Slice splits a string to []float64 by the specified separator.
// For example, split "1.1,2.2,3.3" to []float64{1.1,2.2,3.3} by comma.
func SplitStrToFloat64Slice(s string, sep string) []float64 {
	arr := str.Split(s, sep)
	dst := make([]float64, len(arr))
	for i, v := range arr {
		dst[i] = cast.ToFloat64(v)
	}
	return dst
}

// SplitStrToBoolSlice splits a string to []bool by the specified separator.
// For example, split "1,0,true,false" to []bool{true,false,true,false} by comma.
func SplitStrToBoolSlice(s string, sep string) []bool {
	arr := str.Split(s, sep)
	dst := make([]bool, len(arr))
	for i, v := range arr {
		dst[i] = cast.ToBool(v)
	}
	return dst
}
