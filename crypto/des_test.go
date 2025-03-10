package crypto

import (
	"testing"

	"github.com/dablelv/cyan/internal/utest"
)

func TestBase64DESCBCEncrypt(t *testing.T) {
	assert := utest.NewAssert(t, "TestBase64DESCBCEncrypt")

	// plaintext length is less than 8 bytes.
	{
		p := []byte("plainte")
		key := []byte("12345678")
		c, _ := Base64DESCBCEncrypt(p, key)
		assert.Equal(c, "fpQVUCN7zx4=")
	}

	// plaintext length is greater than 8 bytes.
	{
		p := []byte("plaintext")
		key := []byte("12345678")
		c, _ := Base64DESCBCEncrypt(p, key)
		assert.Equal("UZS/y4By6ksePYMBbvZdig==", c)

		key = []byte("123456789")
		_, err := Base64DESCBCEncrypt(p, key)
		assert.IsNotNil(err)
	}

	// plaintext length is equal to 8 bytes.
	{
		p := []byte("plaintex")
		key := []byte("12345678")
		c, _ := Base64DESCBCEncrypt(p, key)
		assert.Equal(c, "UZS/y4By6kuFP7DXTccMJQ==")

		key = []byte("123456789")
		_, err := Base64DESCBCEncrypt(p, key)
		assert.IsNotNil(err)
	}
}

func TestBase64DESCBCDecrypt(t *testing.T) {
	assert := utest.NewAssert(t, "TestBase64DESCBCDecrypt")

	c := "UZS/y4By6ksePYMBbvZdig=="
	key := []byte("12345678")
	p, _ := Base64DESCBCDecrypt(c, key)
	assert.Equal("plaintext", string(p))

	key = []byte("123456789")
	_, err := Base64DESCBCDecrypt(c, key)
	assert.IsNotNil(err)

	c = "UZS/y4By6ksePYMBbvZdig=+="
	_, err = Base64DESCBCDecrypt(c, key)
	assert.IsNotNil(err)
}

func TestBase64TriDESCBCEncrypt(t *testing.T) {
	assert := utest.NewAssert(t, "TestBase64TriDESCBCEncrypt")

	p := []byte("plaintext")
	key := []byte("12345678abcdefgh12345678")
	c, _ := Base64TriDESCBCEncrypt(p, key)
	assert.Equal("dau0DzmDGQbHasZaOvxxwg==", c)

	key = []byte("12345678abcdefgh123456789")
	_, err := Base64TriDESCBCEncrypt(p, key)
	assert.IsNotNil(err)
}

func TestBase64TriDESCBCDecrypt(t *testing.T) {
	assert := utest.NewAssert(t, "TestBase64TriDESCBCDecrypt")

	c := "dau0DzmDGQbHasZaOvxxwg=="
	key := []byte("12345678abcdefgh12345678")
	p, _ := Base64TriDESCBCDecrypt(c, key)
	assert.Equal("plaintext", string(p))

	key = []byte("12345678abcdefgh123456789")
	_, err := Base64TriDESCBCDecrypt(c, key)
	assert.IsNotNil(err)

	c = "dau0DzmDGQbHasZaOvxxwg=+="
	_, err = Base64TriDESCBCDecrypt(c, key)
	assert.IsNotNil(err)
}
