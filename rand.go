package util

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GetRandInt gets a real non-negative random int number
func GetRandInt() int {
	return rand.Int()
}

// GetRandIntn gets a real non-negative random int number in [0,n)
func GetRandIntn(n int) int {
	if n <= 0 {
		return 0
	}
	return rand.Intn(n)
}

// GetRandIntRange gets a real non-negative random int number in [min, max)
func GetRandIntRange(min, max int) int {
	if min < 0 || max <= min {
		return 0
	}
	return GetRandIntn(max-min) + min
}

// GetRandByteSlice gets a real random byte slice of the specified length, with the value range of 0x00-0xff for each byte
func GetRandByteSlice(length uint32) []byte {
	b := make([]byte, length)
	for i := range b {
		b[i] = byte(GetRandInt() % 256)
	}
	return b
}

const (
	englishAlphabet           = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	englishAlphabetLen        = len(englishAlphabet)
	lowercaseEnglishLetter    = "abcdefghijklmnopqrstuvwxyz"
	lowercaseEnglishLetterLen = len(lowercaseEnglishLetter)
)

// GetRandStr Gets a real random string of a specified length and the contents of which are composed of upper and lower case letters
func GetRandStr(length uint32) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = englishAlphabet[GetRandIntn(englishAlphabetLen)]
	}
	return string(b)
}

// GetRandLowerAlphaStr gets a real random lowercase string of a specified length
func GetRandLowerStr(length uint32) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = lowercaseEnglishLetter[GetRandIntn(lowercaseEnglishLetterLen)]
	}
	return string(b)
}

// GetRandLowerAlphaStr gets a real random uppercase string of a specified length
func GetRandUpperStr(length uint32) string {
	return strings.ToUpper(GetRandLowerStr(length))
}
