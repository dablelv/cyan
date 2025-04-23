package strings

import (
	"regexp"
	"strconv"
	"time"
	"unicode"
)

const (
	// Email regular expression.
	emailRegexString = "^(?:(?:(?:(?:[a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(?:\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|(?:(?:\\x22)(?:(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(?:\\x20|\\x09)+)?(?:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(\\x20|\\x09)+)?(?:\\x22))))@(?:(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$"

	// Integers or floating-point numbers (including positive and negative signs).
	numericRegexString = "^[-+]?[0-9]+(?:\\.[0-9]+)?$"

	// Chinese id card number.
	idCodeRegexString = "^[1-9]\\d{5}(19|20)\\d{2}(0[1-9]|1[0-2])(0[1-9]|[12]\\d|3[01])\\d{3}[\\dX]$"
)

var (
	emailRegex   = regexp.MustCompile(emailRegexString)
	numericRegex = regexp.MustCompile(numericRegexString)
	idCodeRegex  = regexp.MustCompile(idCodeRegexString)
)

func IsNumeric(input string) bool {
	return numericRegex.MatchString(input)
}

func IsEmail(input string) bool {
	return emailRegex.MatchString(input)
}

// IsChineseIdCode checks string is chinese id card number.
func IsChineseIdCode(input string) bool {
	if len(input) != 18 {
		return false
	}

	// Check basic format.
	if !idCodeRegex.MatchString(input) {
		return false
	}

	// Check if the birthday is valid.
	birthday := input[6:14]
	_, err := time.Parse("20060102", birthday)
	if err != nil {
		return false
	}

	// Calculate the checksum.
	weights := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	checkCode := []byte{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'}
	var sum int
	for i := 0; i < 17; i++ {
		var num int
		num, err = strconv.Atoi(string(input[i]))
		if err != nil {
			return false
		}
		sum += num * weights[i]
	}
	mod := sum % 11
	return checkCode[mod] == input[17]
}

// IsEmpty checks if a string is empty ("").
func IsEmpty(s string) bool {
	return s == ""
}

// IsBlank checks if a string is whitespace or empty ("").
func IsBlank(s string) bool {
	if s == "" {
		return true
	}
	for _, v := range s {
		if !unicode.IsSpace(v) {
			return false
		}
	}
	return true
}
