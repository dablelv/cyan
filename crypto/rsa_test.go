package crypto

import (
	"crypto"
	"testing"

	"github.com/dablelv/cyan/internal"
)

func TestRsaEncryptDecrypt(t *testing.T) {
	assert := internal.NewAssert(t, "RsaEncryptDecrypt")

	data := []byte("foo")

	// Encrypt and decrypt data using 2048 bits length key.
	prvkey, pubkey, err := GenRsaKey(2048)
	assert.IsNil(err)

	// Encrypt data.
	cipher, err := RsaEncrypt(pubkey, data)
	assert.IsNil(err)
	assert.Equal(2048/8, len(cipher))

	// Decrypt data.
	plain, err := RsaDecrypt(prvkey, cipher)
	assert.IsNil(err)
	assert.Equal(data, plain)

	// Encrypt and decrypt data using 4096 bits length key.
	prvkey, pubkey, err = GenRsaKey(4096)
	assert.IsNil(err)

	// Encrypt data.
	cipher, err = RsaEncrypt(pubkey, data)
	assert.IsNil(err)
	assert.Equal(4096/8, len(cipher))

	// Decrypt data.
	plain, err = RsaDecrypt(prvkey, cipher)
	assert.IsNil(err)
	assert.Equal(data, plain)

	// Encrypt and decrypt data using 8192 bits length key.
	prvkey, pubkey, err = GenRsaKey(8192)
	assert.IsNil(err)

	// Encrypt data.
	cipher, err = RsaEncrypt(pubkey, data)
	assert.IsNil(err)
	assert.Equal(8192/8, len(cipher))

	// Decrypt data.
	plain, err = RsaDecrypt(prvkey, cipher)
	assert.IsNil(err)
	assert.Equal(data, plain)
}

func TestRsaSignAndVerify(t *testing.T) {
	assert := internal.NewAssert(t, "RsaSignAndVerify")

	msg := []byte("foo")
	prvkey, pubkey, err := GenRsaKey(2048)
	assert.IsNil(err)

	// Using SHA256 to hash msg and then use rsa private key to sign.
	sig, err := RsaSign(prvkey, crypto.SHA256, msg)
	assert.IsNil(err)
	assert.Equal(2048/8, len(sig))

	// Using public key to verify signature.
	err = RsaVerifySign(pubkey, crypto.SHA256, msg, sig)
	assert.IsNil(err)
}
