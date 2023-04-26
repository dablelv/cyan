package crypto

import (
	"testing"

	"github.com/dablelv/go-huge-util/internal"
)

func TestBase64DESCBCEncrypt(t *testing.T) {
	assert := internal.NewAssert(t, "TestBase64DESCBCEncrypt")

	p := []byte("plaintext")
	key := []byte("12345678")
	c, _ := Base64DESCBCEncrypt(p, key)
	assert.Equal("UZS/y4By6ksePYMBbvZdig==", c)
}

func TestBase64DESCBCDecrypt(t *testing.T) {
	assert := internal.NewAssert(t, "TestBase64DESCBCDecrypt")

	c := "UZS/y4By6ksePYMBbvZdig=="
	key := []byte("12345678")
	p, _ := Base64DESCBCDecrypt(c, key)
	assert.Equal("plaintext", string(p))
}

func TestBase64TriDESCBCEncrypt(t *testing.T) {
	assert := internal.NewAssert(t, "TestBase64TriDESCBCEncrypt")

	p := []byte("plaintext")
	key := []byte("12345678abcdefgh12345678")
	c, _ := Base64TriDESCBCEncrypt(p, key)
	assert.Equal("dau0DzmDGQbHasZaOvxxwg==", c)
}

func TestBase64TriDESCBCDecrypt(t *testing.T) {
	assert := internal.NewAssert(t, "TestBase64TriDESCBCDecrypt")

	c := "dau0DzmDGQbHasZaOvxxwg=="
	key := []byte("12345678abcdefgh12345678")
	p, _ := Base64TriDESCBCDecrypt(c, key)
	assert.Equal("plaintext", string(p))
}
