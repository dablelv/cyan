package strings

import (
	"strings"

	"github.com/dablelv/cyan/enum"
)

// EmailDesensitization desensitizes an email address by partially masking the local part (before @) for privacy protection.
// If local part (before @) has ≤3 characters: return entire local part with "***".
// Otherwise shows first 3 characters of local part followed by "***".
// Always preserves the domain part (after @) unchanged.
//
// Examples:
// EmailDesensitization("")          -> ""
// EmailDesensitization("invalid")   -> "invalid"
// EmailDesensitization("a@x.com")   -> "a***@x.com"
// EmailDesensitization("ab@x.com")  -> "ab***@x.com"
// EmailDesensitization("abc@x.com") -> "abc***@x.com"
// EmailDesensitization("abcd@x.com") -> "abc***@x.com"
func EmailDesensitization(email string) string {
	if email == "" {
		return ""
	}
	if !strings.Contains(email, "@") {
		return email
	}
	parts := strings.Split(email, "@")
	local := []rune(parts[0])
	if len(local) <= 3 {
		return parts[0] + "***@" + parts[1]
	}
	return string(local[0:3]) + "***@" + parts[1]
}

// PhoneDesensitization desensitizes mobile phone number.
func PhoneDesensitization(region, number string) string {
	var (
		phoneLen      = len(number)
		before, after int
	)
	switch region {
	case enum.PhoneRegionMainland:
		before = 3
		after = 4
	case enum.PhoneRegionHongKong, enum.PhoneRegionMacao:
		before = 2
		after = 2
	case enum.PhoneRegionTaiwan:
		before = 2
		after = 3
	default:
		before = 2
		after = 2
	}
	if phoneLen < before+after {
		return number
	}
	return string([]rune(number)[0:before]) + "****" + string([]rune(number)[phoneLen-after:phoneLen])
}
