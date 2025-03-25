package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
)

//
// Desc: symmetric aes encryption and decryption functions.
//

// AesCbcEncrypt encrypts data with AES algorithm in CBC mode.
// Note that key length must be 16, 24 or 32 bytes to select AES-128, AES-192, or AES-256.
func AesCbcEncrypt(p, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	p = PKCS7Padding(p, block.BlockSize())
	ciphertext := make([]byte, aes.BlockSize+len(p))
	iv := ciphertext[:aes.BlockSize]
	if _, err = rand.Read(iv); err != nil {
		return nil, err
	}

	blockMode := cipher.NewCBCEncrypter(block, iv)
	blockMode.CryptBlocks(ciphertext[aes.BlockSize:], p)
	return ciphertext, nil
}

// AesCbcDecrypt decrypts cipher text with AES algorithm in CBC mode.
// Note that key length must be 16, 24 or 32 bytes to select AES-128, AES-192, or AES-256.
func AesCbcDecrypt(ciphertext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	plaintext := make([]byte, len(ciphertext))
	blockMode := cipher.NewCBCDecrypter(block, iv)
	blockMode.CryptBlocks(plaintext, ciphertext)

	return PKCS7UnPadding(plaintext)
}

// Base64AesCbcEncrypt encrypts data with AES algorithm in CBC mode and then encode using base64.
func Base64AesCbcEncrypt(p, key []byte) (string, error) {
	c, err := AesCbcEncrypt(p, key)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(c), nil
}

// Base64AesCbcDecrypt decrypts cipher text encoded by base64 with AES algorithm in CBC mode.
func Base64AesCbcDecrypt(c string, key []byte) ([]byte, error) {
	oriCipher, err := base64.StdEncoding.DecodeString(c)
	if err != nil {
		return nil, err
	}
	p, err := AesCbcDecrypt(oriCipher, key)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// AesCbcEncryptSame encrypts data with AES algorithm in CBC mode.
// Note that key length must be 16, 24 or 32 bytes to select AES-128, AES-192, or AES-256.
// This function using the iv from key has security risks so is not recommended to be used.
// If you want to generate the same ciphertext you can use this func.
func AesCbcEncryptSame(p, key []byte) ([]byte, error) {
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

// AesCbcDecryptSame decrypts cipher text with AES algorithm in CBC mode.
// Note that key length must be 16, 24 or 32 bytes to select AES-128, AES-192, or AES-256.
func AesCbcDecryptSame(c, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	plaintext := make([]byte, len(c))
	blockMode := cipher.NewCBCDecrypter(block, key[:block.BlockSize()])
	blockMode.CryptBlocks(plaintext, c)

	return PKCS7UnPadding(plaintext)
}

// Base64AesCbcEncryptSame encrypts data with AES algorithm in CBC mode and then encode using base64.
func Base64AesCbcEncryptSame(p, key []byte) (string, error) {
	c, err := AesCbcEncryptSame(p, key)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(c), nil
}

// Base64AesCbcDecryptSame decrypts cipher text encoded by base64 with AES algorithm in CBC mode.
func Base64AesCbcDecryptSame(c string, key []byte) ([]byte, error) {
	oriCipher, err := base64.StdEncoding.DecodeString(c)
	if err != nil {
		return nil, err
	}
	p, err := AesCbcDecryptSame(oriCipher, key)
	if err != nil {
		return nil, err
	}
	return p, nil
}
