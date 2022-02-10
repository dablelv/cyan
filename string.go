package util

import (
	"bytes"
	"strings"
)

// Split replaces strings.Split.
// strings.Split has a giant pit because strings.Split ("", ",") will return a slice with an empty string.
func Split(s, sep string) []string {
	if s == "" {
		return []string{}
	}
	return strings.Split(s, sep)
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

// ReverseStr reverses the specified string without modifying the original string.
func ReverseStr(s string) string {
	rs := []rune(s)
	var r []rune
	for i := len(rs) - 1; i >= 0; i-- {
		r = append(r, rs[i])
	}
	return string(r)
}
