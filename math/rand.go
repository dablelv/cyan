package math

import (
	"math/rand"
	"strings"
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

// RandIntSlice returns a non-negative truly random int specified length slice with the value range in [0,max).
func RandIntSlice(l int, max int) []int {
	s := make([]int, l)
	for i := 0; i < l; i++ {
		s[i] = RandIntn(max)
	}
	return s
}

// RandByteSlice gets a truly random byte slice of the specified length with the value range in [0x00,0xff].
func RandByteSlice(l uint32) []byte {
	s := make([]byte, l)
	for i := range s {
		s[i] = byte(RandIntn(256))
	}
	return s
}

const (
	enAlphabet       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	enAlphabetLen    = len(enAlphabet)
	lowerEnLetter    = "abcdefghijklmnopqrstuvwxyz"
	lowerEnLetterLen = len(lowerEnLetter)
)

// RandStr returns a truly random string of the specified length and the contents of
// which are composed of upper and lowercase letters.
func RandStr(l int) string {
	s := make([]byte, l)
	for i := range s {
		s[i] = enAlphabet[RandIntn(enAlphabetLen)]
	}
	return string(s)
}

// RandLowerStr returns a truly random lowercase string of a specified length.
func RandLowerStr(l uint32) string {
	b := make([]byte, l)
	for i := range b {
		b[i] = lowerEnLetter[RandIntn(lowerEnLetterLen)]
	}
	return string(b)
}

// RandUpperStr returns a truly random uppercase string of a specified length.
func RandUpperStr(l uint32) string {
	return strings.ToUpper(RandLowerStr(l))
}
