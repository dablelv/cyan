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

// Hash_md5 generates 32 lowercase hexadecimal characters of MD5 checksum
func Hash_md5(s string) string {
	return Hash(s, "md5", false)
}

// Hash_MD5 generates 32 uppercase hexadecimal characters of MD5 checksum
func Hash_MD5(s string) string {
	return Hash(s, "MD5", true)
}

// Hash_sha1 generates 40 lowercase hexadecimal characters of SHA1 checksum
func Hash_sha1(s string) string {
	return Hash(s, "sha1", false)
}

// Hash_SHA1 generates 40 uppercase hexadecimal characters of SHA1 checksum
func Hash_SHA1(s string) string {
	return Hash(s, "SHA1", true)
}

// Hash_sha224 generates 56 lowercase hexadecimal characters of SHA224 checksum
func Hash_sha224(s string) string {
	return Hash(s, "sha224", false)
}

// Hash_SHA224 generates 56 uppercase hexadecimal characters of SHA224 checksum
func Hash_SHA224(s string) string {
	return Hash(s, "SHA224", true)
}

// Hash_sha256 generates 64 lowercase hexadecimal characters of SHA256 checksum
func Hash_sha256(s string) string {
	return Hash(s, "sha256", false)
}

// Hash_SHA256 generates 64 uppercase hexadecimal characters of SHA256 checksum
func Hash_SHA256(s string) string {
	return Hash(s, "SHA256", true)
}

// Hash_sha384 generates 96 lowercase hexadecimal characters of SHA384 checksum
func Hash_sha384(s string) string {
	return Hash(s, "sha384", false)
}

// Hash_SHA384 generates 96 uppercase hexadecimal characters of SHA384 checksum
func Hash_SHA384(s string) string {
	return Hash(s, "SHA384", true)
}

// Hash_sha512 generates 128 lowercase hexadecimal characters of SHA512 checksum
func Hash_sha512(s string) string {
	return Hash(s, "sha512", false)
}

// Hash_SHA512 generates 128 uppercase hexadecimal characters of SHA512 checksum
func Hash_SHA512(s string) string {
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

// HMAC_md5 generates 32 lowercase hexadecimal characters authentication code based on MD5
func HMAC_md5(str, key string) string {
	return HMAC(str, key, "md5", false)
}

// HMAC_MD5 generates 32 uppercase hexadecimal characters authentication code based on MD5
func HMAC_MD5(str, key string) string {
	return HMAC(str, key, "MD5", true)
}

// HMAC_sha1 generates 40 lowercase hexadecimal characters authentication code based on SHA1
func HMAC_sha1(str, key string) string {
	return HMAC(str, key, "sha1", false)
}

// HMAC_SHA1 generates 40 uppercase hexadecimal characters authentication code based on SHA1
func HMAC_SHA1(str, key string) string {
	return HMAC(str, key, "SHA1", true)
}

// HMAC_sha224 generates 56 lowercase hexadecimal characters authentication code based on SHA224
func HMAC_sha224(str, key string) string {
	return HMAC(str, key, "sha224", false)
}

// HMAC_SHA224 generates 56 uppercase hexadecimal characters authentication code based on SHA224
func HMAC_SHA224(str, key string) string {
	return HMAC(str, key, "SHA224", true)
}

// HMAC_SHA256Small generates 64 lowercase hexadecimal characters authentication code based on SHA256
func HMAC_sha256(str, key string) string {
	return HMAC(str, key, "sha256", false)
}

// HMAC_SHA256 generates 64 uppercase hexadecimal characters authentication code based on SHA256
func HMAC_SHA256(str, key string) string {
	return HMAC(str, key, "SHA256", true)
}

// HMAC_sha384 generates 96 lowercase hexadecimal characters authentication code based on SHA384
func HMAC_sha384(str, key string) string {
	return HMAC(str, key, "sha384", false)
}

// HMAC_SHA384 generates 96 uppercase hexadecimal characters authentication code based on SHA384
func HMAC_SHA384(str, key string) string {
	return HMAC(str, key, "SHA384", true)
}

// HMAC_sha512 generates 128 lowercase hexadecimal characters authentication code based on SHA512
func HMAC_sha512(str, key string) string {
	return HMAC(str, key, "sha512", false)
}

// HMAC_SHA512 generates 128 uppercase hexadecimal characters authentication code based on SHA512
func HMAC_SHA512(str, key string) string {
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
func AESCBCEncrypt(p, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// AES block size is 16 bytes
	blockSize := block.BlockSize()
	p = PKCS7Padding(p, blockSize)
	ciphertext := make([]byte, len(p))
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	blockMode.CryptBlocks(ciphertext, p)
	return ciphertext, nil
}

// AESCBCDecrypt decrypts cipher text with AES algorithm in CBC mode
// Note that key length must be 16, 24 or 32 bytes to select AES-128, AES-192, or AES-256
func AESCBCDecrypt(c, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// AES block size is 16 bytes
	blockSize := block.BlockSize()
	plaintext := make([]byte, len(c))
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
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

// TripleDESCBCEncrypt encrypts data with 3DES algorithm in CBC mode
// Note that key length must be 24 bytes
func TripleDESCBCEncrypt(src, key []byte) ([]byte, error) {
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

// TripleDESCBCDecrypt decrypts cipher text with 3DES algorithm in CBC mode
// Note that key length must be 24 bytes
func TripleDESCBCDecrypt(src, key []byte) ([]byte, error) {
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

// Base64TripleDESCBCEncrypt encrypts data with 3DES algorithm in CBC mode and encoded by base64
func Base64TripleDESCBCEncrypt(p, key []byte) (string, error) {
	c, err := TripleDESCBCEncrypt(p, key)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(c), nil
}

// Base64TripleDESCBCDecrypt decrypts cipher text encoded by base64 with 3DES algorithm in CBC mode
func Base64TripleDESCBCDecrypt(c string, key []byte) ([]byte, error) {
	oriCipher, err := base64.StdEncoding.DecodeString(c)
	if err != nil {
		return nil, err
	}
	p, err := TripleDESCBCDecrypt(oriCipher, key)
	if err != nil {
		return nil, err
	}
	return p, nil
}
