package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
