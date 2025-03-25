package crypto

import (
	"testing"

	"github.com/dablelv/cyan/internal/utest"
)

func TestAesCbcEncryptDecrypt(t *testing.T) {
	assert := utest.NewAssert(t, "TestAesCbcEncryptDecrypt")

	key := []byte("12345678abcdefgh")
	p := []byte("plaintext")

	// encrypt
	c, err := AesCbcEncrypt(p, key)
	assert.IsNil(err)

	// decrypt
	plaintext, err := AesCbcDecrypt(c, key)
	assert.IsNil(err)
	assert.Equal(plaintext, p)
}

func TestBase64AesCbcEncryptDecrypt(t *testing.T) {
	assert := utest.NewAssert(t, "TestAesCbcEncryptSame")

	key := []byte("12345678abcdefgh")
	p := []byte("plaintext")

	// encrypt
	c, err := Base64AesCbcEncrypt(p, key)
	assert.IsNil(err)

	// decrypt
	plaintext, err := Base64AesCbcDecrypt(c, key)
	assert.IsNil(err)
	assert.Equal(plaintext, p)
}

func TestAesCbcEncryptSame(t *testing.T) {
	assert := utest.NewAssert(t, "TestAesCbcEncryptSame")

	// AES key length must be 16、24 or 32 bytes to select AES-128, AES-192, or AES-256.
	key := []byte("12345678abcdefgh")
	p := []byte("plaintext")
	c, _ := AesCbcEncryptSame(p, key)
	assert.Equal([]byte{3, 174, 205, 132, 61, 209, 6, 35, 90, 50, 1, 186, 29, 57, 188, 45}, c)
}

func TestAesCbcDecryptSame(t *testing.T) {
	assert := utest.NewAssert(t, "TestAesCbcDecryptSame")

	// AES key length must be 16、24 or 32 bytes to select AES-128, AES-192, or AES-256.
	key := []byte("12345678abcdefgh")
	c := []byte{3, 174, 205, 132, 61, 209, 6, 35, 90, 50, 1, 186, 29, 57, 188, 45}
	p, _ := AesCbcDecryptSame(c, key)
	assert.Equal([]byte("plaintext"), p)
}

func TestBase64AESCBCEncrypt(t *testing.T) {
	assert := utest.NewAssert(t, "TestBase64AESCBCEncrypt")

	// success.
	{
		p := []byte("plaintext")
		key := []byte("12345678abcdefgh")
		c, _ := Base64AesCbcEncryptSame(p, key)
		assert.Equal("A67NhD3RBiNaMgG6HTm8LQ==", c)
	}

	// fail: key is invalid.
	{
		p := []byte("plaintext")
		key := []byte("12345678abcdefghi")
		_, err := Base64AesCbcEncryptSame(p, key)
		assert.IsNotNil(err)
	}
}

func TestBase64AesCbcDecryptSame(t *testing.T) {
	assert := utest.NewAssert(t, "TestBase64AesCbcDecryptSame")

	c := "A67NhD3RBiNaMgG6HTm8LQ=="
	key := []byte("12345678abcdefgh")
	p, _ := Base64AesCbcDecryptSame(c, key)
	assert.Equal([]byte("plaintext"), p)

	key = []byte("12345678abcdefgi")
	_, err := Base64AesCbcDecryptSame(c, key)
	assert.IsNotNil(err)

	c = "A67NhD3RBiNaMgG6HTm8LQ=="
	key = []byte("12345678abcdefghi")
	_, err = Base64AesCbcDecryptSame(c, key)
	assert.IsNotNil(err)

	c = "A67NhD3RBiNaMgG6HTm8LQ=+="
	key = []byte("12345678abcdefgh")
	_, err = Base64AesCbcDecryptSame(c, key)
	assert.IsNotNil(err)
}
