// This source file contains some functions that can handle version comparison.
// Your version number should be MAJOR.MINOR.PATCH e.g. 1.0.0.
// For more version number details refer to Semantic Versioning 2.0.0(https://semver.org/).
// Note all of the functions don't support pre-release version, e.g. 1.0.0-alpha.
// If your version number separator isn't dot or the length you want to compare isn't three,
// please specify separator and length.
//
// If your version is strictly conformed to the Semantic Versioning 2.0.0, you should
// refer to the https://github.com/hashicorp/go-version.

package cmp

import (
	"errors"
	"strconv"
	"strings"
)

// VerGTVer checks whether one version is greater than another version.
// Version like MAJOR.MINOR.PATCH, e.g. 2.1.1 > 2.1.0 > 2.0.0 > 1.0.0.
func VerGTVer(v0, v1 string) (bool, error) {
	r, err := VerCmp(v0, v1, ".")
	if err != nil {
		return false, err
	}
	return r == GT, nil
}

// VerGEVer checks whether one version is greater than or equal to another version.
// Version like MAJOR.MINOR.PATCH, e.g. 2.1.1 > 2.1.0 > 2.0.0 > 1.0.0.
func VerGEVer(v0, v1 string) (bool, error) {
	r, err := VerCmp(v0, v1, ".")
	if err != nil {
		return false, err
	}
	return r == GT || r == EQ, nil
}

// VerLTVer checks whether one version is less than another version.
// Version like MAJOR.MINOR.PATCH, e.g. 1.0.0 < 2.0.0 < 2.1.0 < 2.1.1.
func VerLTVer(v0, v1 string) (bool, error) {
	r, err := VerCmp(v0, v1, ".")
	if err != nil {
		return false, err
	}
	return r == LT, nil
}

// VerLEVer checks whether one version is less than or equal to another version.
// Version like MAJOR.MINOR.PATCH, e.g. 1.0.0 < 2.0.0 < 2.1.0 < 2.1.1.
func VerLEVer(v0, v1 string) (bool, error) {
	r, err := VerCmp(v0, v1, ".")
	if err != nil {
		return false, err
	}
	return r == LT || r == EQ, nil
}

// VerCmp compares two version using specified separator.
// Version like Field1.Field2.Field3,Field4..., e.g. 1.0.0.0 < 2.0.0.0 < 2.1.0.0 < 2.1.1.0.
func VerCmp(v0, v1, sep string) (CMPRSLT, error) {
	if v0 == v1 {
		return EQ, nil
	}
	sl1 := strings.Split(v0, sep)
	sl2 := strings.Split(v1, sep)
	if len(sl1) != len(sl2) {
		return INCMP, errors.New("version format is inconsistent")
	}
	for i := 0; i < len(sl1); i++ {
		n1, err := strconv.Atoi(sl1[i])
		if err != nil {
			return INCMP, err
		}
		n2, err := strconv.Atoi(sl2[i])
		if err != nil {
			return INCMP, err
		}
		if n1 < n2 {
			return LT, nil
		}
		if n1 > n2 {
			return GT, nil
		}
	}
	return EQ, nil
}
