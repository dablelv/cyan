package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"errors"
	"fmt"
	"hash"
)

//
// part 0: hash functions
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

//
// part 1: Hash-based Message Authentication Code functions
//

// HMACMD5L generates 32 lowercase hexadecimal characters authentication code based on MD5
func HMACMD5L(str, key string) string {
	return HMAC(str, key, "md5", false)
}

// HMACMD5U generates 32 uppercase hexadecimal characters authentication code based on MD5
func HMACMD5U(str, key string) string {
	return HMAC(str, key, "MD5", true)
}

// HMACSHA1L generates 40 lowercase hexadecimal characters authentication code based on SHA1
func HMACSHA1L(str, key string) string {
	return HMAC(str, key, "sha1", false)
}

// HMACSHA1U generates 40 uppercase hexadecimal characters authentication code based on SHA1
func HMACSHA1U(str, key string) string {
	return HMAC(str, key, "SHA1", true)
}

// HMACSHA224L generates 56 lowercase hexadecimal characters authentication code based on SHA224
func HMACSHA224L(str, key string) string {
	return HMAC(str, key, "sha224", false)
}

// HMACSHA224U generates 56 uppercase hexadecimal characters authentication code based on SHA224
func HMACSHA224U(str, key string) string {
	return HMAC(str, key, "SHA224", true)
}

// HMACSHA256L generates 64 lowercase hexadecimal characters authentication code based on SHA256
func HMACSHA256L(str, key string) string {
	return HMAC(str, key, "sha256", false)
}

// HMACSHA256U generates 64 uppercase hexadecimal characters authentication code based on SHA256
func HMACSHA256U(str, key string) string {
	return HMAC(str, key, "SHA256", true)
}

// HMACSHA384L generates 96 lowercase hexadecimal characters authentication code based on SHA384
func HMACSHA384L(str, key string) string {
	return HMAC(str, key, "sha384", false)
}

// HMACSHA384U generates 96 uppercase hexadecimal characters authentication code based on SHA384
func HMACSHA384U(str, key string) string {
	return HMAC(str, key, "SHA384", true)
}

// HMACSHA512L generates 128 lowercase hexadecimal characters authentication code based on SHA512
func HMACSHA512L(str, key string) string {
	return HMAC(str, key, "sha512", false)
}

// HMACSHA512U generates 128 uppercase hexadecimal characters authentication code based on SHA512
func HMACSHA512U(str, key string) string {
	return HMAC(str, key, "SHA512", true)
}

// HMAC generates authentication code based on specified hash algorithm
// e.g. HMAC("", "", "MD5", false) returns 74e6f7298a9c2d168935f58c001bad88
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

//
// part 2 symmetric encryption and decryption functions
//

// PKCS7Padding fills plaintext as an integral multiple of the block length
func PKCS7Padding(p []byte, blockSize int) []byte {
	pad := blockSize - len(p)%blockSize
	padtext := bytes.Repeat([]byte{byte(pad)}, pad)
	return append(p, padtext...)
}

// PKCS7UnPadding removes padding data from the tail of plaintext
func PKCS7UnPadding(p []byte) ([]byte, error) {
	length := len(p)
	paddLen := int(p[length-1])
	if paddLen > length {
		return nil, errors.New("data is not illegal")
	}
	return p[:(length - paddLen)], nil
}

// AESCBCEncrypt encrypts data with AES algorithm in CBC mode
// Note that key length must be 16, 24 or 32 bytes to select AES-128, AES-192, or AES-256
// Note that AES block size is 16 bytes
func AESCBCEncrypt(p, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	p = PKCS7Padding(p, block.BlockSize())
	ciphertext := make([]byte, len(p))
	blockMode := cipher.NewCBCEncrypter(block, key[:block.BlockSize()])
	blockMode.CryptBlocks(ciphertext, p)
	return ciphertext, nil
}

// AESCBCDecrypt decrypts cipher text with AES algorithm in CBC mode
// Note that key length must be 16, 24 or 32 bytes to select AES-128, AES-192, or AES-256
// Note that AES block size is 16 bytes
func AESCBCDecrypt(c, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	plaintext := make([]byte, len(c))
	blockMode := cipher.NewCBCDecrypter(block, key[:block.BlockSize()])
	blockMode.CryptBlocks(plaintext, c)
	return PKCS7UnPadding(plaintext)
}

// Base64AESCBCEncrypt encrypts data with AES algorithm in CBC mode and encoded by base64
func Base64AESCBCEncrypt(p, key []byte) (string, error) {
	c, err := AESCBCEncrypt(p, key)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(c), nil
}

// Base64AESCBCDecrypt decrypts cipher text encoded by base64 with AES algorithm in CBC mode
func Base64AESCBCDecrypt(c string, key []byte) ([]byte, error) {
	oriCipher, err := base64.StdEncoding.DecodeString(c)
	if err != nil {
		return nil, err
	}
	p, err := AESCBCDecrypt(oriCipher, key)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// DESCBCEncrypt encrypts data with DES algorithm in CBC mode
// Note that key length must be 8 bytes
func DESCBCEncrypt(p, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	p = PKCS7Padding(p, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	c := make([]byte, len(p))
	blockMode.CryptBlocks(c, p)
	return c, nil
}

// DESCBCDecrypt decrypts cipher text with DES algorithm in CBC mode
// Note that key length must be 8 bytes
func DESCBCDecrypt(c, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	p := make([]byte, len(c))
	blockMode.CryptBlocks(p, c)
	return PKCS7UnPadding(p)
}

// Base64DESCBCEncrypt encrypts data with DES algorithm in CBC mode and encoded by base64
func Base64DESCBCEncrypt(p, key []byte) (string, error) {
	c, err := DESCBCEncrypt(p, key)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(c), nil
}

// Base64DESCBCDecrypt decrypts cipher text encoded by base64 with DES algorithm in CBC mode
func Base64DESCBCDecrypt(c string, key []byte) ([]byte, error) {
	oriCipher, err := base64.StdEncoding.DecodeString(c)
	if err != nil {
		return nil, err
	}
	p, err := DESCBCDecrypt(oriCipher, key)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// TriDESCBCEncrypt encrypts data with 3DES algorithm in CBC mode
// Note that key length must be 24 bytes
func TriDESCBCEncrypt(src, key []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	src = PKCS7Padding(src, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:block.BlockSize()])
	dst := make([]byte, len(src))
	blockMode.CryptBlocks(dst, src)
	return dst, nil
}

// TriDESCBCDecrypt decrypts cipher text with 3DES algorithm in CBC mode
// Note that key length must be 24 bytes
func TriDESCBCDecrypt(src, key []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key[:block.BlockSize()])
	dst := make([]byte, len(src))
	blockMode.CryptBlocks(dst, src)
	dst, err = PKCS7UnPadding(dst)
	return dst, err
}

// Base64TriDESCBCEncrypt encrypts data with 3DES algorithm in CBC mode and encoded by base64
func Base64TriDESCBCEncrypt(p, key []byte) (string, error) {
	c, err := TriDESCBCEncrypt(p, key)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(c), nil
}

// Base64TriDESCBCDecrypt decrypts cipher text encoded by base64 with 3DES algorithm in CBC mode
func Base64TriDESCBCDecrypt(c string, key []byte) ([]byte, error) {
	oriCipher, err := base64.StdEncoding.DecodeString(c)
	if err != nil {
		return nil, err
	}
	p, err := TriDESCBCDecrypt(oriCipher, key)
	if err != nil {
		return nil, err
	}
	return p, nil
}
