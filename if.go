package util

// If simulates the conditional operator used in C, c++ and Java.
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}
