package str

import (
	"bytes"
	"strings"
	"unicode"
)

// Capitalize capitalizes all the delimiter separated words in a string. Only the first letter of each word is changed.
// The delimiters represent a set of characters understood to separate words.The first string character and the first
// non-delimiter character after a delimiter will be capitalized.Capitalization uses the Unicode title case,
// normally equivalent to upper case.
func Capitalize(s string, delimiters ...rune) string {
	if len(delimiters) == 0 {
		delimiters = []rune{' '}
	}

	buffer := []rune(s)
	capitalizeNext := true
	for i := 0; i < len(buffer); i++ {
		ch := buffer[i]
		if isDelimiter(ch, delimiters...) {
			capitalizeNext = true
		} else if capitalizeNext {
			buffer[i] = unicode.ToTitle(ch)
			capitalizeNext = false
		}
	}
	return string(buffer)
}

// CapitalizeFully converts all the delimiter separated words in a string into capitalized words, that is each word is made up of a
// titlecase character and then a series of lowercase characters. The delimiters represent a set of characters understood to separate words.
// The first string character and the first non-delimiter character after a delimiter will be capitalized.
// Capitalization uses the Unicode title case, normally equivalent to upper case.
func CapitalizeFully(s string, delimiters ...rune) string {
	if len(delimiters) == 0 {
		delimiters = []rune{' '}
	}

	s = strings.ToLower(s)
	return Capitalize(s, delimiters...)
}

// Uncapitalize uncapitalizes all the whitespace separated words in a string. Only the first letter of each word is changed.
// The delimiters represent a set of characters understood to separate words. The first string character and the first non-delimiter
// character after a delimiter will be uncapitalized. Whitespace is defined by unicode.IsSpace(char).
func Uncapitalize(str string, delimiters ...rune) string {
	if len(delimiters) == 0 {
		delimiters = []rune{' '}
	}

	buffer := []rune(str)
	uncapitalizeNext := true
	for i := 0; i < len(buffer); i++ {
		ch := buffer[i]
		if isDelimiter(ch, delimiters...) {
			uncapitalizeNext = true
		} else if uncapitalizeNext {
			buffer[i] = unicode.ToLower(ch)
			uncapitalizeNext = false
		}
	}
	return string(buffer)
}

// Initials extracts the initial letters from each word in the string. The first letter of the string and all first
// letters after the defined delimiters are returned as a new string. Their case is not changed. If the delimiters
// parameter is excluded, then Whitespace is used. Whitespace is defined by unicode.IsSpacea(char).
func Initials(s string, delimiters ...rune) string {
	if len(delimiters) == 0 {
		delimiters = []rune{' '}
	}

	var buf bytes.Buffer
	lastWasGap := true
	buffer := []rune(s)
	for i := 0; i < len(buffer); i++ {
		ch := rune(s[i])
		if isDelimiter(ch, delimiters...) {
			lastWasGap = true
		} else if lastWasGap {
			buf.WriteRune(ch)
			lastWasGap = false
		}
	}
	return buf.String()
}

func isDelimiter(r rune, delimiters ...rune) bool {
	if delimiters == nil {
		return unicode.IsSpace(r)
	}
	for _, delimiter := range delimiters {
		if r == delimiter {
			return true
		}
	}
	return false
}
