package conv

import huge "github.com/dablelv/go-huge-util"

// SplitStrToSet convert a string to map set after split
func SplitStrToSet(v string, sep string) map[string]struct{} {
	m := make(map[string]struct{})
	for _, v := range huge.Split(v, sep) {
		m[v] = struct{}{}
	}
	return m
}
