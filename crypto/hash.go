package crypto

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
)

//
// Desc: Hash functions.
//

// MD5L generates 32 lowercase hexadecimal characters of MD5 checksum
func MD5L(s string) string {
	return Hash(s, "md5", false)
}

// MD5U generates 32 uppercase hexadecimal characters of MD5 checksum
func MD5U(s string) string {
	return Hash(s, "MD5", true)
}

// SHA1L generates 40 lowercase hexadecimal characters of SHA1 checksum
func SHA1L(s string) string {
	return Hash(s, "sha1", false)
}

// SHA1U generates 40 uppercase hexadecimal characters of SHA1 checksum
func SHA1U(s string) string {
	return Hash(s, "SHA1", true)
}

// SHA224U generates 56 lowercase hexadecimal characters of SHA224 checksum
func SHA224L(s string) string {
	return Hash(s, "sha224", false)
}

// SHA224U generates 56 uppercase hexadecimal characters of SHA224 checksum
func SHA224U(s string) string {
	return Hash(s, "SHA224", true)
}

// SHA256L generates 64 lowercase hexadecimal characters of SHA256 checksum
func SHA256L(s string) string {
	return Hash(s, "sha256", false)
}

// SHA256U generates 64 uppercase hexadecimal characters of SHA256 checksum
func SHA256U(s string) string {
	return Hash(s, "SHA256", true)
}

// SHA384L generates 96 lowercase hexadecimal characters of SHA384 checksum
func SHA384L(s string) string {
	return Hash(s, "sha384", false)
}

// SHA384U generates 96 uppercase hexadecimal characters of SHA384 checksum
func SHA384U(s string) string {
	return Hash(s, "SHA384", true)
}

// SHA512L generates 128 lowercase hexadecimal characters of SHA512 checksum
func SHA512L(s string) string {
	return Hash(s, "sha512", false)
}

// SHA512U generates 128 uppercase hexadecimal characters of SHA512 checksum
func SHA512U(s string) string {
	return Hash(s, "SHA512", true)
}

// Hash generates a checksum of the specified hash algorithm
// e.g. Hash("", "MD5", "false") returns d41d8cd98f00b204e9800998ecf8427e
func Hash(s string, hashName string, capital bool) string {
	var h hash.Hash
	switch hashName {
	case "md5", "MD5":
		h = md5.New()
	case "sha1", "SHA1":
		h = sha1.New()
	case "sha224", "SHA224":
		h = sha256.New224()
	case "sha256", "SHA256":
		h = sha256.New()
	case "sha384", "SHA384":
		h = sha512.New384()
	case "sha512", "SHA512":
		h = sha512.New()
	default:
		return ""
	}
	h.Write([]byte(s))
	if capital {
		return fmt.Sprintf("%X", h.Sum(nil))
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}
