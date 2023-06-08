package str

import (
	"bytes"
	"regexp"
	"strings"
)

// Split replaces std lib strings.Split.
// strings.Split has a giant pit because strings.Split ("", ",") will return a slice with an empty string.
func Split(s, sep string) []string {
	if s == "" {
		return nil
	}
	return strings.Split(s, sep)
}

// SplitSeps splits string into substring slice by multiple string separators.
// If you want to specify multiple string separators by regexp,
// please refer to `func (*Regexp) Split` in standard library regexp package.
func SplitSeps(s string, seps ...string) []string {
	if len(seps) == 0 {
		return []string{s}
	}

	result := strings.Split(s, seps[0])
	for _, sep := range seps[1:] {
		var temp []string
		for _, r := range result {
			temp = append(temp, strings.Split(r, sep)...)
		}
		result = temp
	}
	return result
}

// JoinStrSkipEmpty concatenates multiple strings to a single string with the specified separator and skips the empty
// string.
func JoinStrSkipEmpty(sep string, s ...string) string {
	var buf bytes.Buffer
	for _, v := range s {
		if v == "" {
			continue
		}
		if buf.Len() > 0 {
			buf.WriteString(sep)
		}
		buf.WriteString(v)
	}
	return buf.String()
}

// JoinStr concatenates multiple strings to a single string with the specified separator.
// Note that JoinStr doesn't skip the empty string.
func JoinStr(sep string, s ...string) string {
	var buf bytes.Buffer
	for i, v := range s {
		if i != 0 {
			buf.WriteString(sep)
		}
		buf.WriteString(v)
	}
	return buf.String()
}

// Reverse reverses string without modifying the original string.
func Reverse(s string) string {
	rs := []rune(s)
	r := make([]rune, len(rs))
	for i := len(rs) - 1; i >= 0; i-- {
		r[len(rs)-1-i] = rs[i]
	}
	return string(r)
}

// GetAlphanumericNumByASCII gets the alphanumeric number based on the ASCII code value.
// Note that this function has a better performance than GetAlphanumericNumByRegExp, so this function is recommended.
func GetAlphanumericNumByASCII(s string) int {
	num := int(0)
	for i := 0; i < len(s); i++ {
		switch {
		case 48 <= s[i] && s[i] <= 57: // digits
			fallthrough
		case 65 <= s[i] && s[i] <= 90: // uppercase letters
			fallthrough
		case 97 <= s[i] && s[i] <= 122: // lowercase letters
			num++
		default:
		}
	}
	return num
}

// GetAlphanumericNumByASCIIV2 gets the alphanumeric number based on the ASCII code value.
// Because range by rune so the performance is worse than GetAlphanumericNumByASCII.
func GetAlphanumericNumByASCIIV2(s string) int {
	num := int(0)
	for _, c := range s {
		switch {
		case '0' <= c && c <= '9':
			fallthrough
		case 'a' <= c && c <= 'z':
			fallthrough
		case 'A' <= c && c <= 'Z':
			num++
		default:
		}
	}
	return num
}

// GetAlphanumericNumByRegExp gets the alphanumeric number based on regular expression.
// Note that this function has a poor performance when compared to GetAlphanumericNumByASCII,
// so the GetAlphanumericNumByASCII is recommended.
func GetAlphanumericNumByRegExp(s string) int {
	rNum := regexp.MustCompile(`\d`)
	rLetter := regexp.MustCompile("[a-zA-Z]")
	return len(rNum.FindAllString(s, -1)) + len(rLetter.FindAllString(s, -1))
}
