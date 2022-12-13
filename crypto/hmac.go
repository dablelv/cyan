package crypto

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
)

//
// Desc: Hash-based Message Authentication Code functions.
//

// HMACMD5L generates 32 lowercase hexadecimal characters authentication code based on MD5.
func HMACMD5L(str, key string) string {
	return HMAC(str, key, "md5", false)
}

// HMACMD5U generates 32 uppercase hexadecimal characters authentication code based on MD5.
func HMACMD5U(str, key string) string {
	return HMAC(str, key, "MD5", true)
}

// HMACSHA1L generates 40 lowercase hexadecimal characters authentication code based on SHA1.
func HMACSHA1L(str, key string) string {
	return HMAC(str, key, "sha1", false)
}

// HMACSHA1U generates 40 uppercase hexadecimal characters authentication code based on SHA1.
func HMACSHA1U(str, key string) string {
	return HMAC(str, key, "SHA1", true)
}

// HMACSHA224L generates 56 lowercase hexadecimal characters authentication code based on SHA224.
func HMACSHA224L(str, key string) string {
	return HMAC(str, key, "sha224", false)
}

// HMACSHA224U generates 56 uppercase hexadecimal characters authentication code based on SHA224.
func HMACSHA224U(str, key string) string {
	return HMAC(str, key, "SHA224", true)
}

// HMACSHA256L generates 64 lowercase hexadecimal characters authentication code based on SHA256.
func HMACSHA256L(str, key string) string {
	return HMAC(str, key, "sha256", false)
}

// HMACSHA256U generates 64 uppercase hexadecimal characters authentication code based on SHA256.
func HMACSHA256U(str, key string) string {
	return HMAC(str, key, "SHA256", true)
}

// HMACSHA384L generates 96 lowercase hexadecimal characters authentication code based on SHA384.
func HMACSHA384L(str, key string) string {
	return HMAC(str, key, "sha384", false)
}

// HMACSHA384U generates 96 uppercase hexadecimal characters authentication code based on SHA384.
func HMACSHA384U(str, key string) string {
	return HMAC(str, key, "SHA384", true)
}

// HMACSHA512L generates 128 lowercase hexadecimal characters authentication code based on SHA512.
func HMACSHA512L(str, key string) string {
	return HMAC(str, key, "sha512", false)
}

// HMACSHA512U generates 128 uppercase hexadecimal characters authentication code based on SHA512.
func HMACSHA512U(str, key string) string {
	return HMAC(str, key, "SHA512", true)
}

// HMAC generates authentication code based on specified hash algorithm.
// e.g. HMAC("", "", "MD5", false) returns 74e6f7298a9c2d168935f58c001bad88.
func HMAC(str, key string, hashName string, capital bool) string {
	var hm hash.Hash
	switch hashName {
	case "md5", "MD5":
		hm = hmac.New(md5.New, []byte(key))
	case "sha1", "SHA1":
		hm = hmac.New(sha1.New, []byte(key))
	case "sha224", "SHA224":
		hm = hmac.New(sha256.New224, []byte(key))
	case "sha256", "SHA256":
		hm = hmac.New(sha256.New, []byte(key))
	case "sha384", "SHA384":
		hm = hmac.New(sha512.New384, []byte(key))
	case "sha512", "SHA512":
		hm = hmac.New(sha512.New, []byte(key))
	default:
		return ""
	}
	hm.Write([]byte(str))
	if capital {
		return fmt.Sprintf("%X", hm.Sum(nil))
	}
	return fmt.Sprintf("%x", hm.Sum(nil))
}
