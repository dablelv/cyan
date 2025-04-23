package strings

import (
	"errors"
	"strings"
)

// FormatThousand formats a numeric string with thousand separator.
// The thousand separator default is comma if not specified.
func FormatThousand(s string, sep ...string) (string, error) {
	if s == "" {
		return s, nil
	}

	l := strings.Count(s, ".")
	if l > 1 {
		return s, errors.New("not legal data")
	}

	positivePart := s
	decimalPart := ""
	if l == 1 {
		substrs := strings.Split(s, ".")
		positivePart = substrs[0]
		decimalPart = substrs[1]
	}
	posPartLen := len(positivePart)
	if posPartLen >= 4 {
		separator := ","
		if len(sep) > 0 {
			separator = sep[0]
		}
		count := (posPartLen - 1) / 3
		for i := 0; i < count; i++ {
			positivePart = positivePart[:posPartLen-(i+1)*3] + separator + positivePart[posPartLen-(i+1)*3:]
		}
	}
	if len(decimalPart) > 0 {
		return positivePart + "." + decimalPart, nil
	}
	return positivePart, nil
}
