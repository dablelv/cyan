package math

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandInt returns a non-negative truly random int number.
func RandInt() int {
	return rand.Int()
}

// RandIntn returns a non-negative truly random int number in [0,n).
func RandIntn(n int) int {
	if n <= 0 {
		return 0
	}
	return rand.Intn(n)
}

// RandIntRange returns a non-negative truly random int number in [min,max).
func RandIntRange(min, max int) int {
	if min < 0 || max <= min {
		return 0
	}
	return RandIntn(max-min) + min
}

// RandInts returns a non-negative truly random int specified length slice with the value range in [0,max).
func RandInts(l, max int) []int {
	s := make([]int, l)
	for i := 0; i < l; i++ {
		s[i] = RandIntn(max)
	}
	return s
}

// RandBytes gets a truly random byte slice of the specified length with the value range in [0x00,0xff].
func RandBytes(l uint32) []byte {
	s := make([]byte, l)
	for i := range s {
		s[i] = byte(RandIntn(256))
	}
	return s
}
