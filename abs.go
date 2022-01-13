package util

// AbsInt8 gets absolute value of int8.
//
// Example 1 : AbsInt8(-5)
// -5 的表示
// 原码 1000,0101
// 反码 1111,1010
// 补码 1111,1011
// 负数在内存中以补码的形式表示
// shifted = n >> 7 = (1111,1011) >> 7 = 1111,1111 = -1(十进制) (负数右移，左补1)
//               1111,1011
// n xor shifted =  ----------- = 0000,0100 = 4(十进制)
//               1111,1111
// (n ^ shifted) - shifted = 4 - (-1) = 5
//
// Example 2:  AbsInt8(5)
// 5 的表示
// 原码 0000,0101
// 正数在内存中以原码的形式表示，而正数与 0 进行异或运算等于本身
// shifted = n >> 7 = 0
//               0000,0101
// n xor shifted =  ----------- = 0000,0101 = 5(十进制)
//               0000,0000
// (n ^ shifted) - shifted = 5 - 0 = 5
func AbsInt8(n int8) int8 {
	// n 右移 7 位
	shifted := n >> 7
	return (n ^ shifted) - shifted
}

// AbsInt16 int16 取绝对值
func AbsInt16(n int16) int16 {
	// n 右移 15 位
	shifted := n >> 15
	return (n ^ shifted) - shifted
}

// AbsInt32 int32 取绝对值
func AbsInt32(n int32) int32 {
	// n 右移 31 位
	shifted := n >> 31
	return (n ^ shifted) - shifted
}

// AbsInt64 int64 取绝对值
func AbsInt64(n int64) int64 {
	// n 右移 63 位
	shifted := n >> 63
	return (n ^ shifted) - shifted
}
