package cond

// If simulates the conditional operator used in C, c++ and Java.
func If(condition bool, trueVal, falseVal any) any {
	if condition {
		return trueVal
	}
	return falseVal
}
