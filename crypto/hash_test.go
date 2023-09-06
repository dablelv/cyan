package crypto

import (
	"testing"

	"github.com/dablelv/cyan/internal"
)

func TestMd5(t *testing.T) {
	assert := internal.NewAssert(t, "Md5")
	assert.Equal("d41d8cd98f00b204e9800998ecf8427e", Md5(""))
	assert.Equal("acbd18db4cc2f85cedef654fccc4a4d8", Md5("foo"))
}

func TestMD5(t *testing.T) {
	assert := internal.NewAssert(t, "MD5")
	assert.Equal("D41D8CD98F00B204E9800998ECF8427E", MD5(""))
	assert.Equal("ACBD18DB4CC2F85CEDEF654FCCC4A4D8", MD5("foo"))
}

func TestSha1(t *testing.T) {
	assert := internal.NewAssert(t, "Sha1")
	assert.Equal("da39a3ee5e6b4b0d3255bfef95601890afd80709", Sha1(""))
	assert.Equal("0beec7b5ea3f0fdbc95d0dd47f3c5bc275da8a33", Sha1("foo"))
}

func TestSHA1(t *testing.T) {
	assert := internal.NewAssert(t, "SHA1")
	assert.Equal("DA39A3EE5E6B4B0D3255BFEF95601890AFD80709", SHA1(""))
}

func TestSha224(t *testing.T) {
	assert := internal.NewAssert(t, "Sha224")
	assert.Equal("d14a028c2a3a2bc9476102bb288234c415a2b01f828ea62ac5b3e42f", Sha224(""))
}

func TestSHA224(t *testing.T) {
	assert := internal.NewAssert(t, "SHA224")
	assert.Equal("D14A028C2A3A2BC9476102BB288234C415A2B01F828EA62AC5B3E42F", SHA224(""))
}

func TestSha256(t *testing.T) {
	assert := internal.NewAssert(t, "Sha256")
	assert.Equal("e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855", Sha256(""))
}

func TestSHA256(t *testing.T) {
	assert := internal.NewAssert(t, "SHA256")
	assert.Equal("E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855", SHA256(""))
}

func TestSha384(t *testing.T) {
	assert := internal.NewAssert(t, "Sha384")
	assert.Equal("38b060a751ac96384cd9327eb1b1e36a21fdb71114be07434c0cc7bf63f6e1da274edebfe76f65fbd51ad2f14898b95b", Sha384(""))
}

func TestSHA384(t *testing.T) {
	assert := internal.NewAssert(t, "SHA384")
	assert.Equal("38B060A751AC96384CD9327EB1B1E36A21FDB71114BE07434C0CC7BF63F6E1DA274EDEBFE76F65FBD51AD2F14898B95B", SHA384(""))
}

func TestSha512(t *testing.T) {
	assert := internal.NewAssert(t, "Sha384")
	assert.Equal("cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e", Sha512(""))
}

func TestSHA512(t *testing.T) {
	assert := internal.NewAssert(t, "SHA512")
	assert.Equal("CF83E1357EEFB8BDF1542850D66D8007D620E4050B5715DC83F4A921D36CE9CE47D0D13C5D85F2B0FF8318D2877EEC2F63B931BD47417A81A538327AF927DA3E", SHA512(""))
}
