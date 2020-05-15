package util

import (
	"bytes"
	"strings"
)

func StrToMapTrue(v string, sep string) map[string]bool {
	m := make(map[string]bool)
	for _, v := range Split(v, sep) {
		m[v] = true
	}
	return m
}

func StrToMapFalse(v string, sep string) map[string]bool {
	m := make(map[string]bool)
	for _, v := range Split(v, sep) {
		m[v] = false
	}
	return m
}

func StrToIntSlice(s string, sep string) []int {
	arr := Split(s, sep)
	dst := make([]int, len(arr))
	for i, v := range arr {
		dst[i] = ToInt(v)
	}
	return dst
}

func StrToInt8Slice(s string, sep string) []int8 {
	arr := Split(s, sep)
	dst := make([]int8, len(arr))
	for i, v := range arr {
		dst[i] = ToInt8(v)
	}
	return dst
}

func StrToInt16Slice(s string, sep string) []int16 {
	arr := Split(s, sep)
	dst := make([]int16, len(arr))
	for i, v := range arr {
		dst[i] = ToInt16(v)
	}
	return dst
}

func StrToInt32Slice(s string, sep string) []int32 {
	arr := Split(s, sep)
	dst := make([]int32, len(arr))
	for i, v := range arr {
		dst[i] = ToInt32(v)
	}
	return dst
}

func StrToInt64Slice(s string, sep string) []int64 {
	arr := Split(s, sep)
	dst := make([]int64, len(arr))
	for i, v := range arr {
		dst[i] = ToInt64(v)
	}
	return dst
}

func StrToUintSlice(s string, sep string) []uint {
	arr := Split(s, sep)
	dst := make([]uint, len(arr))
	for i, v := range arr {
		dst[i] = ToUint(v)
	}
	return dst
}

func StrToUint8Slice(s string, sep string) []uint8 {
	arr := Split(s, sep)
	dst := make([]uint8, len(arr))
	for i, v := range arr {
		dst[i] = ToUint8(v)
	}
	return dst
}

func StrToByteSlice(s string, sep string) []byte {
	arr := Split(s, sep)
	dst := make([]byte, len(arr))
	for i, v := range arr {
		dst[i] = byte(ToUint8(v))
	}
	return dst
}

func StrToUint16Slice(s string, sep string) []uint16 {
	arr := Split(s, sep)
	dst := make([]uint16, len(arr))
	for i, v := range arr {
		dst[i] = ToUint16(v)
	}
	return dst
}

func StrToUint32Slice(s string, sep string) []uint32 {
	arr := Split(s, sep)
	dst := make([]uint32, len(arr))
	for i, v := range arr {
		dst[i] = ToUint32(v)
	}
	return dst
}

func StrToUint64Slice(s string, sep string) []uint64 {
	arr := Split(s, sep)
	dst := make([]uint64, len(arr))
	for i, v := range arr {
		dst[i] = ToUint64(v)
	}
	return dst
}

func StrToFloat32Slice(s string, sep string) []float32 {
	arr := Split(s, sep)
	dst := make([]float32, len(arr))
	for i, v := range arr {
		dst[i] = ToFloat32(v)
	}
	return dst
}

func StrToFloat64Slice(s string, sep string) []float64 {
	arr := Split(s, sep)
	dst := make([]float64, len(arr))
	for i, v := range arr {
		dst[i] = ToFloat64(v)
	}
	return dst
}

// Split replace strings.Split
// strings.Split has a giant pit because strings.Split ("", ",") will return a slice with an empty string
func Split(s, sep string) []string {
	if s == "" {
		return []string{}
	}
	return strings.Split(s, sep)
}

func JoinSkipEmpty(sep string, a ...string) string {
	var buf bytes.Buffer
	for i := 0; i < len(a); i++ {
		if a[i] == "" {
			continue
		}

		if i != 0 {
			buf.WriteString(sep)
		}
		buf.WriteString(a[i])
	}
	return buf.String()
}

func JoinNoSkipEmpty(sep string, a ...string) string {
	var buf bytes.Buffer
	for i := 0; i < len(a); i++ {
		if i != 0 {
			buf.WriteString(sep)
		}
		buf.WriteString(a[i])
	}
	return buf.String()
}
