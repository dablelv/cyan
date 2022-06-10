package util

import (
	"unsafe"
)

// If simulates the conditional operator used in C, c++ and Java.
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

// IsLittleEndian determines whether the host byte order is little endian.
func IsLittleEndian() bool {
	n := 0x1234
	return *(*byte)(unsafe.Pointer(&n)) == 0x34
}
