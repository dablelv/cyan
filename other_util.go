package util

import (
	"bytes"
	"encoding/json"
	"unsafe"
)

func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

// IsLittleEndian determines whether the host byte order is little endian
func IsLittleEndian() bool {
	n := 0x1234
	return *(*byte)(unsafe.Pointer(&n)) == 0x34
}

// ToFormattedJSON traverses the golang value to formatted JSON string, such as a struct, map, slice, array and so on
func ToFormattedJSON(obj interface{}) (string, error) {
	bs, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	err = json.Indent(buf, bs, "", "\t")
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
