// This file is copied from https://github.com/duke-git/lancet/interanl
// and modified on demand.

// Package utest is for internal unit test.
package utest

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

const (
	compareNotEqual int = iota - 2
	compareLess
	compareEqual
	compareGreater
)

// Assert is a simple implementation of assertion, only for internal usage.
type Assert struct {
	T        *testing.T
	CaseName string
}

// NewAssert return instance of Assert.
func NewAssert(t *testing.T, caseName string) *Assert {
	return &Assert{T: t, CaseName: caseName}
}

// Equal check if expected is equal with actual.
func (a *Assert) Equal(actual, expected any) {
	if compare(actual, expected) != compareEqual {
		makeTestFailed(a.T, a.CaseName, actual, expected)
	}
}

// False check if expected is false.
func (a *Assert) False(actual any) {
	a.Equal(actual, false)
}

// True check if expected is true.
func (a *Assert) True(actual any) {
	a.Equal(actual, true)
}

// NotEqual check if expected is not equal with actual
func (a *Assert) NotEqual(actual, expected any) {
	if compare(actual, expected) == compareEqual {
		expectedInfo := fmt.Sprintf("not %v", expected)
		makeTestFailed(a.T, a.CaseName, expectedInfo, actual)
	}
}

// Greater check if actual is greate than  expected.
func (a *Assert) Greater(actual, expected any) {
	if compare(actual, expected) != compareGreater {
		expectedInfo := fmt.Sprintf("> %v", expected)
		makeTestFailed(a.T, a.CaseName, expectedInfo, actual)
	}
}

// GreaterOrEqual check if expected is greate than or equal with actual
func (a *Assert) GreaterOrEqual(actual, expected any) {
	isGreatOrEqual := compare(actual, expected) == compareGreater || compare(actual, expected) == compareEqual
	if !isGreatOrEqual {
		expectedInfo := fmt.Sprintf(">= %v", expected)
		makeTestFailed(a.T, a.CaseName, expectedInfo, actual)
	}
}

// Less check if expected is less than actual
func (a *Assert) Less(actual, expected any) {
	if compare(actual, expected) != compareLess {
		expectedInfo := fmt.Sprintf("< %v", expected)
		makeTestFailed(a.T, a.CaseName, expectedInfo, actual)
	}
}

// LessOrEqual check if expected is less than or equal with actual
func (a *Assert) LessOrEqual(actual, expected any) {
	isLessOrEqual := compare(actual, expected) == compareLess || compare(actual, expected) == compareEqual
	if !isLessOrEqual {
		expectedInfo := fmt.Sprintf("<= %v", expected)
		makeTestFailed(a.T, a.CaseName, expectedInfo, actual)
	}
}

// IsNil check if value is nil
func (a *Assert) IsNil(v any) {
	if v == nil || (reflect.ValueOf(v).Kind() == reflect.Ptr && reflect.ValueOf(v).IsNil()) {
		return
	}
	if reflect.ValueOf(v).IsNil() {
		return
	}
	makeTestFailed(a.T, a.CaseName, nil, v)
}

// IsNotNil check if value is not nil
func (a *Assert) IsNotNil(v any) {
	if v == nil || (reflect.ValueOf(v).Kind() == reflect.Ptr && reflect.ValueOf(v).IsNil()) {
		makeTestFailed(a.T, a.CaseName, "not nil", v)
	}
}

func (a *Assert) ErrorContains(err error, contains string) {
	if err == nil {
		makeTestFailed(a.T, a.CaseName, err, contains)
	}
	if !strings.Contains(err.Error(), contains) {
		makeTestFailed(a.T, a.CaseName, err, contains)
	}
}

// ShouldBeError asserts that the first argument implements the error interface.
// It also compares the first argument against the second argument if provided
// (which must be an error message string or another error value).
// func (a *Assert) IsError(actual, expected any) string {
// 	if !isError(actual) {
// 		return fmt.Sprintf(shouldBeError, reflect.TypeOf(actual))

// 		makeTestFailed(a.T, a.CaseName, "not nil", v)
// 	}

// 	if len(expected) == 0 {
// 		return success
// 	}

// 	if expected := expected[0]; !isString(expected) && !isError(expected) {
// 		return fmt.Sprintf(shouldBeErrorInvalidComparisonValue, reflect.TypeOf(expected))
// 	}
// 	return ShouldEqual(fmt.Sprint(actual), fmt.Sprint(expected[0]))
// }

// compare x and y return :
// x > y -> 1, x < y -> -1, x == y -> 0, x != y -> -2
func compare(x, y any) int {
	vx := reflect.ValueOf(x)
	vy := reflect.ValueOf(y)

	if vx.Type() != vy.Type() {
		return compareNotEqual
	}

	switch vx.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		xInt := vx.Int()
		yInt := vy.Int()
		if xInt > yInt {
			return compareGreater
		}
		if xInt == yInt {
			return compareEqual
		}
		if xInt < yInt {
			return compareLess
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		xUint := vx.Uint()
		yUint := vy.Uint()
		if xUint > yUint {
			return compareGreater
		}
		if xUint == yUint {
			return compareEqual
		}
		if xUint < yUint {
			return compareLess
		}
	case reflect.Float32, reflect.Float64:
		xFloat := vx.Float()
		yFloat := vy.Float()
		if xFloat > yFloat {
			return compareGreater
		}
		if xFloat == yFloat {
			return compareEqual
		}
		if xFloat < yFloat {
			return compareLess
		}
	case reflect.String:
		xString := vx.String()
		yString := vy.String()
		if xString > yString {
			return compareGreater
		}
		if xString == yString {
			return compareEqual
		}
		if xString < yString {
			return compareLess
		}
	default:
		if reflect.DeepEqual(x, y) {
			return compareEqual
		}
		return compareNotEqual
	}
	return compareNotEqual
}

// logFailedInfo make test failed and log error info.
func makeTestFailed(t *testing.T, caseName string, actual, expected any) {
	_, file, line, _ := runtime.Caller(2)
	errInfo := fmt.Sprintf("Case %v failed. file: %v, line: %v, actual: %v, expected: %v.", caseName, file, line, actual, expected)
	t.Error(errInfo)
	t.FailNow()
}

func isString(value interface{}) bool { _, ok := value.(string); return ok }
func isError(value interface{}) bool  { _, ok := value.(error); return ok }
