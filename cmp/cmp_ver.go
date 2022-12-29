// This source file contains some functions that can handle version number comparison.
// Your version number should be MAJOR.MINOR.PATCH e.g. 1.0.0.
// For more version number details refer to Semantic Versioning 2.0.0(https://semver.org/).
// Note all of the functions don't support pre-release version, e.g. 1.0.0-alpha.
// If your version number separator isn't dot or the length you want to compare isn't three,
// please specify separator and length.

package cmp

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/dablelv/go-huge-util/str"
)

// VerGTVer determines whether version0 greater than version1.
// Version number like MAJOR.MINOR.PATCH, e.g. 2.1.1 > 2.1.0 > 2.0.0 > 1.0.0.
func VerGTVer(ver0, ver1 string) (bool, error) {
	return VerGTVerMore(ver0, ver1, ".", 3)
}

// VerLTVer determines whether version0 less than version1.
// Version number like MAJOR.MINOR.PATCH, e.g. 1.0.0 < 2.0.0 < 2.1.0 < 2.1.1.
func VerLTVer(ver0, ver1 string) (bool, error) {
	return VerLTVerMore(ver0, ver1, ".", 3)
}

// VerGEVer determines whether version0 greater than or equal to version1.
// Version number like MAJOR.MINOR.PATCH, e.g. 2.1.1 > 2.1.0 > 2.0.0 > 1.0.0.
func VerGEVer(ver0, ver1 string) (bool, error) {
	return VerGEVerMore(ver0, ver1, ".", 3)
}

// VerLEVer determines whether version0 less than or equal to version1.
// Version number like MAJOR.MINOR.PATCH, e.g. 1.0.0 < 2.0.0 < 2.1.0 < 2.1.1.
func VerLEVer(ver0, ver1 string) (bool, error) {
	return VerLEVerMore(ver0, ver1, ".", 3)
}

// VerGTVerMore reports whether version 0 greater than version 1 with specified version separator and length.
// Version number like Field1.Field2.Field3.Field4..., e.g. 2.1.1.1 > 2.1.0.1 > 2.0.0.0 > 1.0.0.0.
func VerGTVerMore(ver0, ver1, sep string, num int) (bool, error) {
	ver0, ver1 = strings.TrimSpace(ver0), strings.TrimSpace(ver1)
	if ver0 == ver1 {
		return false, nil
	}

	ver0s, ver1s := str.Split(ver0, sep), str.Split(ver1, sep)
	if len(ver0s) != num || len(ver1s) != num {
		return false, errors.New(fmt.Sprintf("version field num is not equal to %v", num))
	}

	var i64Ver0, i64Ver1 int64
	for i := 0; i < num; i++ {
		v0, err0 := strconv.ParseInt(ver0s[i], 10, 64)
		if err0 != nil {
			return false, errors.New("the first version number is ill")
		}
		v1, err1 := strconv.ParseInt(ver1s[i], 10, 64)
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

// VerLTVerMore reports whether version 0 less than version 1 with specified version separator and length.
// Version number like Field1.Field2.Field3,Field4..., e.g. 1.0.0.0 < 2.0.0.0 < 2.1.0.0 < 2.1.1.0.
func VerLTVerMore(ver0, ver1, sep string, num int) (bool, error) {
	ver0, ver1 = strings.TrimSpace(ver0), strings.TrimSpace(ver1)
	if ver0 == ver1 {
		return false, nil
	}

	ver0s, ver1s := str.Split(ver0, sep), str.Split(ver1, sep)
	if len(ver0s) != num || len(ver1s) != num {
		return false, errors.New("version field num is too short")
	}

	var i64Ver0, i64Ver1 int64
	for i := 0; i < num; i++ {
		v0, err0 := strconv.ParseInt(ver0s[i], 10, 64)
		if err0 != nil {
			return false, errors.New("the first version number is ill")
		}
		v1, err1 := strconv.ParseInt(ver1s[i], 10, 64)
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

// VerGEVerMore determines whether version0 greater than or equal to version1 with specified version separator and
// length.
// Version number like Field1.Field2.Field3.Field4..., e.g. 2.1.1.1 >= 2.1.0.1 >= 2.0.0.0 >= 1.0.0.0
func VerGEVerMore(ver0, ver1, sep string, num int) (bool, error) {
	ver0, ver1 = strings.TrimSpace(ver0), strings.TrimSpace(ver1)
	if ver0 == ver1 {
		return true, nil
	}
	return VerGTVerMore(ver0, ver1, sep, num)
}

// VerLEVerMore determines whether version0 less than or equal to version1 with specified version separator and length.
// Version number like Field1.Field2.Field3.Field4..., e.g. 1.0.0.0 <= 2.0.0.0 <= 2.1.0.0 <= 2.1.1.0.
func VerLEVerMore(ver0, ver1, sep string, num int) (bool, error) {
	ver0, ver1 = strings.TrimSpace(ver0), strings.TrimSpace(ver1)
	if ver0 == ver1 {
		return true, nil
	}
	return VerLTVerMore(ver0, ver1, sep, num)
}
