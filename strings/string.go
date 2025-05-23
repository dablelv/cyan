package strings

import (
	"bytes"
	"encoding/hex"
	"reflect"
	"regexp"
	"strings"
	"unicode"
	"unsafe"

	"github.com/dablelv/cyan/conv"
)

const INDEX_NOT_FOUND = -1

// Head returns the first n bytes of the string s if len(s) greater than n,
// or else returns the raw s.
func Head(s string, n int) string {
	switch {
	case n < 0:
		return ""
	case n < len(s):
		return s[:n]
	default:
		return s
	}
}

// HeadRunes returns the first n runes of the string s.
// HeadRunes properly handles string containing multibyte characters by converting
// the string to a slice of runes, ensuring accurate indexing by character rather than by byte.
func HeadRunes(s string, n int) string {
	runes := []rune(s)
	switch {
	case n < 0:
		return ""
	case n < len(runes):
		return string(runes[:n])
	default:
		return s
	}
}

// Tail return the last n length of substring if len(s) greater than n,
// or else return th raw s.
func Tail(s string, n int) string {
	switch {
	case n < 0:
		return ""
	case n < len(s):
		return s[len(s)-n:]
	default:
		return s
	}
}

// TailRunes returns the last n runes of the string s.
// HeadRunes properly handles string containing multibyte characters by converting
// the string to a slice of runes, ensuring accurate indexing by character rather than by byte.
func TailRunes(s string, n int) string {
	runes := []rune(s)
	switch {
	case n < 0:
		return ""
	case n < len(runes):
		return string(runes[len(runes)-n:])
	default:
		return s
	}
}

// Split replaces standard lib strings.Split.
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

// SplitChunk splits the string according to the specified number and separator string into a new string.
func SplitChunk(input, sep string, size int) string {
	if size <= 0 {
		return input
	}

	result := ""
	// Convert the string to a rune slice to support multi-byte characters.
	runeStr := []rune(input)
	for i := 0; i < len(runeStr); i += size {
		end := i + size
		if end > len(runeStr) {
			end = len(runeStr)
		}
		if i == 0 {
			result += string(runeStr[i:end])
		} else {
			result += sep + string(runeStr[i:end])
		}
	}
	return result
}

// SplitChunkWithSpace splits the string according to the specified number using space character into a new string.
func SplitChunkWithSpace(input string, size int) string {
	return SplitChunk(input, " ", size)
}

// JoinNonEmptyStrs concatenates multiple strings to a single string with the specified separator and skips the empty
// string.
func JoinNonEmptyStrs(sep string, s ...string) string {
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

// Join concatenates all elements of Array, Slice or String to a single string with a separator.
func Join(a any, sep string) string {
	s, _ := JoinE(a, sep)
	return s
}

// JoinE concatenates all elements of Array, Slice or String to a single string with a separator
// and returns an error if error occurred.
// E.g. input []int{1, 2, 3} and separator ",", output is a string "1,2,3".
// It panics if a's Kind is not Array, Slice, or String.
func JoinE(a any, sep string) (string, error) {
	v := reflect.ValueOf(a)
	if v.Kind() == reflect.String {
		return JoinE(strings.Split(a.(string), ""), sep)
	}
	var s string
	for i := 0; i < v.Len(); i++ {
		if len(s) > 0 {
			s += sep
		}
		str, err := conv.ToStringE(v.Index(i).Interface())
		if err != nil {
			return "", err
		}
		s += str
	}
	return s, nil
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

// AlphanumericNum returns the alphanumeric number based on the ASCII code value.
// AlphanumericNum should be used first because it has a better performance than AlphanumericNumV2 and AlphanumericNumRegExp.
func AlphanumericNum(s string) int {
	var num int
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

// AlphanumericNumV2 returns the alphanumeric number based on the ASCII code value.
// Because range by rune so the performance is worse than AlphanumericNum.
func AlphanumericNumV2(s string) int {
	var num int
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

// AlphanumericNumRegExp returns the alphanumeric number based on regular expression.
// Note that this function has a poor performance when compared to AlphanumericNum,
// so the AlphanumericNum is recommended.
func AlphanumericNumRegExp(s string) int {
	rNum := regexp.MustCompile(`\d`)
	rLetter := regexp.MustCompile("[a-zA-Z]")
	return len(rNum.FindAllString(s, -1)) + len(rLetter.FindAllString(s, -1))
}

// ClearWhiteSpace deletes all whitespaces from a string as defined by unicode.IsSpace(rune).
// It returns the string without whitespaces.
func ClearWhiteSpace(s string) string {
	var buf bytes.Buffer
	for _, c := range s {
		if !unicode.IsSpace(c) {
			buf.WriteRune(c)
		}
	}
	return buf.String()
}

// IndexOfDiff compares two strings and returns the index at which the strings begin to differ.
// -1 is returned if two strings are equal.
func IndexOfDiff(s1, s2 string) int {
	if s1 == s2 {
		return INDEX_NOT_FOUND
	}
	if s1 == "" || s2 == "" {
		return 0
	}
	var i int
	for i = 0; i < len(s1) && i < len(s2); i++ {
		if s1[i] != s2[i] {
			break
		}
	}
	if i < len(s1) || i < len(s2) {
		return i
	}
	return INDEX_NOT_FOUND
}

// IndexOffset returns the index of the first instance of sub in str, with the search beginning from the
// index start point specified. -1 is returned if sub is not present in str.
// An empty string ("") will return -1 (INDEX_NOT_FOUND).
// A negative start position is treated as zero.
// A start position greater than the string length returns -1.
func IndexOffset(s, sub string, start int) int {
	if start < 0 {
		start = 0
	}

	if start >= len(s) {
		return INDEX_NOT_FOUND
	}

	if IsEmpty(s) || IsEmpty(sub) {
		return INDEX_NOT_FOUND
	}

	partialIndex := strings.Index(s[start:], sub)
	if partialIndex == -1 {
		return INDEX_NOT_FOUND
	}
	return partialIndex + start
}

// Default returns either the passed in string, or if the string is empty, the value of default string.
func Default(s, d string) string {
	if IsEmpty(s) {
		return d
	}
	return s
}

// DefaultIfBlank returns either the passed in string, or if the string is whitespace, empty (""), the value of default string.
func DefaultIfBlank(s, d string) string {
	if IsBlank(s) {
		return d
	}
	return s
}

// BinToHex returns the hexadecimal encoding of src buffer.
func BinToHex(src []byte) string {
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	// No copy operation, better performance than string(dst).
	return ConstBytesToString(dst)
}

// ConstBytesToString transfer bytes to string without copy real buffer.
// You must NOT modify the input bytes, otherwise the output string is NOT constant!
// Call it only if necessary.
func ConstBytesToString(buf []byte) string {
	return *(*string)(unsafe.Pointer(&buf))
}

// StringToConstBytes transfer string to const bytes without copy real buffer.
// You must NOT modify the output bytes, otherwise the input string is modified or process will coredump!!!
// Call it only if necessary.
func StringToConstBytes(s string) (ret []byte) {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&ret))

	sliceHeader.Data = stringHeader.Data
	sliceHeader.Len, sliceHeader.Cap = stringHeader.Len, stringHeader.Len
	return ret
}
