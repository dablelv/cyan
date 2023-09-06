package crypto

import (
	"testing"

	"github.com/dablelv/cyan/internal"
)

func TestAESCBCEncrypt(t *testing.T) {
	assert := internal.NewAssert(t, "TestAESCBCEncrypt")

	// AES key length must be 16、24 or 32 bytes to select AES-128, AES-192, or AES-256.
	key := []byte("12345678abcdefgh")
	p := []byte("plaintext")
	c, _ := AESCBCEncrypt(p, key)
	assert.Equal([]byte{3, 174, 205, 132, 61, 209, 6, 35, 90, 50, 1, 186, 29, 57, 188, 45}, c)
}

func TestAESCBCDecrypt(t *testing.T) {
	assert := internal.NewAssert(t, "TestAESCBCDecrypt")

	// AES key length must be 16、24 or 32 bytes to select AES-128, AES-192, or AES-256.
	key := []byte("12345678abcdefgh")
	c := []byte{3, 174, 205, 132, 61, 209, 6, 35, 90, 50, 1, 186, 29, 57, 188, 45}
	p, _ := AESCBCDecrypt(c, key)
	assert.Equal([]byte("plaintext"), p)
}

func TestBase64AESCBCEncrypt(t *testing.T) {
	assert := internal.NewAssert(t, "TestBase64AESCBCEncrypt")

	p := []byte("plaintext")
	key := []byte("12345678abcdefgh")
	c, _ := Base64AESCBCEncrypt(p, key)
	assert.Equal("A67NhD3RBiNaMgG6HTm8LQ==", c)

	p = []byte("plaintext")
	key = []byte("12345678abcdefghi")
	_, err := Base64AESCBCEncrypt(p, key)
	assert.IsNotNil(err)
}

func TestBase64AESCBCDecrypt(t *testing.T) {
	assert := internal.NewAssert(t, "TestBase64AESCBCDecrypt")

	c := "A67NhD3RBiNaMgG6HTm8LQ=="
	key := []byte("12345678abcdefgh")
	p, _ := Base64AESCBCDecrypt(c, key)
	assert.Equal([]byte("plaintext"), p)

	key = []byte("12345678abcdefgi")
	_, err := Base64AESCBCDecrypt(c, key)
	assert.IsNotNil(err)

	c = "A67NhD3RBiNaMgG6HTm8LQ=="
	key = []byte("12345678abcdefghi")
	_, err = Base64AESCBCDecrypt(c, key)
	assert.IsNotNil(err)

	c = "A67NhD3RBiNaMgG6HTm8LQ=+="
	key = []byte("12345678abcdefgh")
	_, err = Base64AESCBCDecrypt(c, key)
	assert.IsNotNil(err)
}
