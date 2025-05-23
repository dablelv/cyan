// Package cmp can compare two numbers of any type.
// In addition, you can also compare two version using this package.
package cmp

import (
	"cmp"
	"reflect"
)

type CMPRSLT int8

const (
	INCMP CMPRSLT = iota - 2
	LT
	EQ
	GT
)

// Cmp compares two numeric values or strings size.
// Deprecated: as of Go 1.21, plz use standard lib func cmp.Compare.
func Cmp[T cmp.Ordered](l, r T) CMPRSLT {
	switch {
	case l < r:
		return LT
	case l == r:
		return EQ
	default:
		return GT
	}
}

// Compare compares two numeric values or strings size.
// The result is INCMP(incomparable), LT(less than), EQ(equal) or GT(greater than).
// Deprecated: as of Go 1.18, plz use the generic func Cmp.
func Compare(lhs, rhs any) CMPRSLT {
	if !isComparable(lhs, rhs) {
		return INCMP
	}

	lhsVal := reflect.ValueOf(lhs)
	rhsVal := reflect.ValueOf(rhs)

	switch lhsVal.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		switch {
		case lhsVal.Int() < rhsVal.Int():
			return LT
		case lhsVal.Int() == rhsVal.Int():
			return EQ
		case lhsVal.Int() > rhsVal.Int():
			return GT
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		switch {
		case lhsVal.Uint() < rhsVal.Uint():
			return LT
		case lhsVal.Uint() == rhsVal.Uint():
			return EQ
		case lhsVal.Uint() > rhsVal.Uint():
			return GT
		}
	case reflect.Float32, reflect.Float64:
		switch {
		case lhsVal.Float() < rhsVal.Float():
			return LT
		case lhsVal.Float() == rhsVal.Float():
			return EQ
		case lhsVal.Float() > rhsVal.Float():
			return GT
		}
	case reflect.String:
		switch {
		case lhsVal.String() < rhsVal.String():
			return LT
		case lhsVal.String() == rhsVal.String():
			return EQ
		case lhsVal.String() > rhsVal.String():
			return GT
		}
	}
	return INCMP
}

// CompareLT reports left hand side value whether less than right hand side value.
func CompareLT(lhs, rhs any) bool {
	return Compare(lhs, rhs) == LT
}

// CompareLE reports left hand side value whether less than or equal right hand side value.
func CompareLE(lhs, rhs any) bool {
	r := Compare(lhs, rhs)
	return r == LT || r == EQ
}

// CompareEQ reports left hand side value whether equal right hand side value.
func CompareEQ(lhs, rhs any) bool {
	return Compare(lhs, rhs) == EQ
}

// CompareGT reports left hand side value whether greater than right hand side value.
func CompareGT(lhs, rhs any) bool {
	return Compare(lhs, rhs) == GT
}

// CompareGE reports left hand side value whether greater than or equal right hand side value.
func CompareGE(lhs, rhs any) bool {
	r := Compare(lhs, rhs)
	return r == GT || r == EQ
}

func isComparable(lhs, rhs any) bool {
	l := reflect.ValueOf(lhs)
	r := reflect.ValueOf(rhs)
	return l.Kind() == r.Kind()
}
