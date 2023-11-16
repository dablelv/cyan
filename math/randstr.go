package math

import "strings"

const (
	alphabetic       = "abcdefghijklmnopqrstuvwxyz"
	alphabeticLen    = len(alphabetic)
	alphabeticMix    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphabeticMixLen = len(alphabeticMix)
	alphanumeric     = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	alphanumericLen  = len(alphanumeric)
)

// RandLowerStr returns a truly random lowercase string of a specified length.
func RandLowerStr(l uint32) string {
	b := make([]byte, l)
	for i := range b {
		b[i] = alphabetic[RandIntn(alphabeticLen)]
	}
	return string(b)
}

// RandUpperStr returns a truly random uppercase string of a specified length.
func RandUpperStr(l uint32) string {
	return strings.ToUpper(RandLowerStr(l))
}

// RandStr returns a truly random string of the specified length and the contents of
// which are composed of upper and lowercase letters.
func RandStr(l int) string {
	s := make([]byte, l)
	for i := range s {
		s[i] = alphabeticMix[RandIntn(alphabeticMixLen)]
	}
	return string(s)
}

// RandAlphanumeric returns a truly random string of the specified length and the contents of
// which are composed of alphanumeric letters.
func RandAlphanumeric(l int) string {
	s := make([]byte, l)
	for i := range s {
		s[i] = alphanumeric[RandIntn(alphanumericLen)]
	}
	return string(s)
}
