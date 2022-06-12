package util

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

// RandIntRange returns a non-negative truly random int number in [min, max).
func RandIntRange(min, max int) int {
	if min < 0 || max <= min {
		return 0
	}
	return RandIntn(max-min) + min
}

// RandIntSlice returns a non-negative truly random int slice of the specified length with the value range in [0,max).
func RandIntSlice(l int, max int) []int {
	slice := make([]int, l)
	for i := 0; i < l; i++ {
		slice[i] = RandIntn(max)
	}
	return slice
}

// RandByteSlice gets a truly random byte slice of the specified length with the value range in [0x00,0xff].
func RandByteSlice(l uint32) []byte {
	slice := make([]byte, l)
	for i := range slice {
		slice[i] = byte(RandInt() % 256)
	}
	return slice
}

const (
	enAlphabet       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	enAlphabetLen    = len(enAlphabet)
	lowerEnLetter    = "abcdefghijklmnopqrstuvwxyz"
	lowerEnLetterLen = len(lowerEnLetter)
)

// RandStr returns a truly random string of the specified length and the contents of which are composed of upper and
// lower case letters.
func RandStr(l uint64) string {
	b := make([]byte, l)
	for i := range b {
		b[i] = enAlphabet[RandIntn(enAlphabetLen)]
	}
	return string(b)
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
