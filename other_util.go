package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"strings"
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

// VerGTVer determines whether version0 is greater than version1
// Given a version number is MAJOR.MINOR.PATCH, e.g. 2.1.1 > 2.1.0 > 2.0.0 > 1.0.0
func VerGTVer(ver0, ver1 string) (bool, error) {
	ver0 = strings.TrimSpace(ver0)
	ver1 = strings.TrimSpace(ver1)

	slVer0 := Split(ver0, ".")
	slVer1 := Split(ver1, ".")
	if len(slVer0) != len(slVer1) {
		return false, errors.New("param error")
	}

	var i64Ver0, i64Ver1 int64
	for i := 0; i < len(slVer0); i++ {
		v0, err0 := ToInt64E(slVer0[i])
		if err0 != nil {
			return false, errors.New("the first version number is ill")
		}
		v1, err1 := ToInt64E(slVer1[i])
		if err1 != nil {
			return false, errors.New("the second version number is ill")
		}

		i64Ver0 *= 10
		i64Ver0 += v0
		i64Ver1 *= 10
		i64Ver1 += v1
	}
	return i64Ver0 > i64Ver1, nil
}

// VerGTVer determines whether version0 is less than version1
// Given a version number is MAJOR.MINOR.PATCH, e.g. 1.0.0 < 2.0.0 < 2.1.0 < 2.1.1
func VerLTVer(ver0, ver1 string) (bool, error) {
	ver0 = strings.TrimSpace(ver0)
	ver1 = strings.TrimSpace(ver1)

	slVer0 := Split(ver0, ".")
	slVer1 := Split(ver1, ".")
	if len(slVer0) != len(slVer1) {
		return false, errors.New("param error")
	}

	var i64Ver0, i64Ver1 int64
	for i := 0; i < len(slVer0); i++ {
		v0, err0 := ToInt64E(slVer0[i])
		if err0 != nil {
			return false, errors.New("the first version number is ill")
		}
		v1, err1 := ToInt64E(slVer1[i])
		if err1 != nil {
			return false, errors.New("the second version number is ill")
		}

		i64Ver0 *= 10
		i64Ver0 += v0
		i64Ver1 *= 10
		i64Ver1 += v1
	}
	return i64Ver0 < i64Ver1, nil
}

// VerGEVer determines whether version0 is greater than or equal to version1
// Given a version number is MAJOR.MINOR.PATCH, e.g. 2.1.1 > 2.1.0 > 2.0.0 > 1.0.0
func VerGEVer(ver0, ver1 string) (bool, error) {
	ver0 = strings.TrimSpace(ver0)
	ver1 = strings.TrimSpace(ver1)
	if ver0 == ver1 {
		return true, nil
	}
	return VerGTVer(ver0, ver1)
}

// VerLEVer determines whether version0 is less than or equal to version1
// Given a version number is MAJOR.MINOR.PATCH, e.g. 1.0.0 < 2.0.0 < 2.1.0 < 2.1.1
func VerLEVer(ver0, ver1 string) (bool, error) {
	ver0 = strings.TrimSpace(ver0)
	ver1 = strings.TrimSpace(ver1)
	if ver0 == ver1 {
		return true, nil
	}
	return VerLTVer(ver0, ver1)
}
