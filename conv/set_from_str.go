package conv

import "github.com/dablelv/go-huge-util/str"

// SplitStrToSet convert a string to map set after split
func SplitStrToSet(s string, sep string) map[string]struct{} {
	if s == "" {
		return nil
	}
	m := make(map[string]struct{})
	for _, v := range str.Split(s, sep) {
		m[v] = struct{}{}
	}
	return m
}
