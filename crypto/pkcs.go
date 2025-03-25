package crypto

import (
	"bytes"
	"errors"
)

// PKCS7Padding fills plaintext as an integral multiple of the block length.
// In cryptography, PKCS stands for "Public Key Cryptography Standards".
// These are a group of public key cryptography standards devised and published by RSA Security LLC, starting in the early 1990s.
// PKCS#7 is Cryptographic Message Syntax Standard used to sign and/or encrypt messages under a PKI.
// More information about PKCS#7 please refer to the RFC 2315.
// The PKCS#7 padding method works by adding bytes that all have the same value as the number
// of bytes added to the plaintext. For example, if the last block of plaintext is 3 bytes long
// and the block size is 8 bytes, then 5 bytes of padding are added to make the total length of the plaintext 8 bytes.
// The value of the 5 bytes will be the hex value 0x05, since 5 bytes were added as padding.
func PKCS7Padding(p []byte, blockSize int) []byte {
	pad := blockSize - len(p)%blockSize
	padtext := bytes.Repeat([]byte{byte(pad)}, pad)
	return append(p, padtext...)
}

// PKCS7UnPadding removes padding data from the tail of plaintext.
func PKCS7UnPadding(p []byte) ([]byte, error) {
	l := len(p)
	paddLen := int(p[l-1])
	if paddLen > l {
		return nil, errors.New("data is not illegal")
	}
	return p[:(l - paddLen)], nil
}
