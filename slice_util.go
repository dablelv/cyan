package util

import (
	"errors"
	"reflect"
)

func UniqueIntSlice(src []int) []int {
	tmp, _ := UniqueSlice(src)
	dst := []int{}
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, ToInt(v))
	}
	return dst
}

func UniqueInt8Slice(src []int8) []int8 {
	tmp, _ := UniqueSlice(src)
	dst := []int8{}
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, ToInt8(v))
	}
	return dst
}

func UniqueInt16Slice(src []int16) []int16 {
	tmp, _ := UniqueSlice(src)
	dst := []int16{}
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, ToInt16(v))
	}
	return dst
}

func UniqueInt32Slice(src []int32) []int32 {
	tmp, _ := UniqueSlice(src)
	dst := []int32{}
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, ToInt32(v))
	}
	return dst
}

func UniqueInt64Slice(src []int64) []int64 {
	tmp, _ := UniqueSlice(src)
	dst := []int64{}
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, ToInt64(v))
	}
	return dst
}

func UniqueUintSlice(src []uint) []uint {
	tmp, _ := UniqueSlice(src)
	dst := []uint{}
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, ToUint(v))
	}
	return dst
}

func UniqueUint8Slice(src []uint8) []uint8 {
	tmp, _ := UniqueSlice(src)
	dst := []uint8{}
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, ToUint8(v))
	}
	return dst
}

func UniqueUint16Slice(src []uint16) []uint16 {
	tmp, _ := UniqueSlice(src)
	dst := []uint16{}
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, ToUint16(v))
	}
	return dst
}

func UniqueUint32Slice(src []uint32) []uint32 {
	tmp, _ := UniqueSlice(src)
	dst := []uint32{}
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, ToUint32(v))
	}
	return dst
}

func UniqueUint64Slice(src []uint64) []uint64 {
	tmp, _ := UniqueSlice(src)
	dst := []uint64{}
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, ToUint64(v))
	}
	return dst
}

func UniqueByteSlice(src []byte) []byte {
	tmp, _ := UniqueSlice(src)
	dst := []byte{}
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, ToByte(v))
	}
	return dst
}

func UniqueStrSlice(src []string) []string {
	tmp, _ := UniqueSlice(src)
	dst := []string{}
	for _, v := range tmp.([]interface{}) {
		dst = append(dst, ToString(v))
	}
	return dst
}

// UniqueSlice delete repeated elements
func UniqueSlice(slice interface{}) (interface{}, error) {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return nil, errors.New("param isn't a slice")
	}

	var dst []interface{}
	m := make(map[interface{}]bool)
	for i := 0; i < v.Len(); i++ {
		if _, ok := m[v.Index(i).Interface()]; !ok {
			dst = append(dst, v.Index(i).Interface())
			m[v.Index(i).Interface()] = true
		}
	}
	return dst, nil
}
