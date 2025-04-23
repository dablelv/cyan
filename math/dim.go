//
// This file name dim.go means
//

package math

import "cmp"

// Min compare two values and return smaller value.
// Deprecated: as of go 1.21, please use builtin func min.
func Min[T cmp.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

// Max compare two values and return larger value.
// Deprecated: as of go 1.21, please use builtin func max.
func Max[T cmp.Ordered](x, y T) T {
	if x > y {
		return x
	}
	return y
}
