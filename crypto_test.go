package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAESCBCEncrypt(t *testing.T) {
	// AES key length must be 16、24 or 32 bytes to select AES-128, AES-192, or AES-256
	key := []byte("12345678abcdefgh")
	p := []byte("plaintext")
	c, _ := AESCBCEncrypt(p, key)
	assert.Equal(t, []byte{3, 174, 205, 132, 61, 209, 6, 35, 90, 50, 1, 186, 29, 57, 188, 45}, c)
}

func TestAesDecrypt(t *testing.T) {
	// AES key length must be 16、24 or 32 bytes to select AES-128, AES-192, or AES-256
	key := []byte("12345678abcdefgh")
	c := []byte{3, 174, 205, 132, 61, 209, 6, 35, 90, 50, 1, 186, 29, 57, 188, 45}
	p, _ := AESCBCDecrypt(c, key)
	assert.Equal(t, "plaintext", string(p))
}

func TestBase64AESCBCEncrypt(t *testing.T) {
	p := []byte("plaintext")
	key := []byte("12345678abcdefgh")
	c, _ := Base64AESCBCEncrypt(p, key)
	assert.Equal(t, "A67NhD3RBiNaMgG6HTm8LQ==", c)
}

func TestBase64AESCBCDecrypt(t *testing.T) {
	c := "A67NhD3RBiNaMgG6HTm8LQ=="
	key := []byte("12345678abcdefgh")
	p, _ := Base64AESCBCDecrypt(c, key)
	assert.Equal(t, "plaintext", string(p))
}

func TestBase64DESCBCEncrypt(t *testing.T) {
	p := []byte("plaintext")
	key := []byte("12345678")
	c, _ := Base64DESCBCEncrypt(p, key)
	assert.Equal(t, "UZS/y4By6ksePYMBbvZdig==", c)
}

func TestBase64DESCBCDecrypt(t *testing.T) {
	c := "UZS/y4By6ksePYMBbvZdig=="
	key := []byte("12345678")
	p, _ := Base64DESCBCDecrypt(c, key)
	assert.Equal(t, "plaintext", string(p))
}

func TestBase64TriDESCBCEncrypt(t *testing.T) {
	p := []byte("plaintext")
	key := []byte("12345678abcdefgh12345678")
	c, _ := Base64TriDESCBCEncrypt(p, key)
	assert.Equal(t, "dau0DzmDGQbHasZaOvxxwg==", c)
}

func TestBase64TriDESCBCDecrypt(t *testing.T) {
	c := "dau0DzmDGQbHasZaOvxxwg=="
	key := []byte("12345678abcdefgh12345678")
	p, _ := Base64TriDESCBCDecrypt(c, key)
	assert.Equal(t, "plaintext", string(p))
}
