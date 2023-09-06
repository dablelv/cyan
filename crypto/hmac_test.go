package crypto

import (
	"testing"

	"github.com/dablelv/cyan/internal"
)

func TestHMACMd5(t *testing.T) {
	assert := internal.NewAssert(t, "HMACMd5")
	assert.Equal("31b6db9e5eb4addb42f1a6ca07367adc", HMACMd5("foo", "bar"))
}

func TestHMACMD5(t *testing.T) {
	assert := internal.NewAssert(t, "HMACMD5")
	assert.Equal("31B6DB9E5EB4ADDB42F1A6CA07367ADC", HMACMD5("foo", "bar"))
}

func TestHMACSha1(t *testing.T) {
	assert := internal.NewAssert(t, "HMACSha1")
	assert.Equal("85d155c55ed286a300bd1cf124de08d87e914f3a", HMACSha1("foo", "bar"))
}

func TestHMACSHA1(t *testing.T) {
	assert := internal.NewAssert(t, "HMACSHA1")
	assert.Equal("85D155C55ED286A300BD1CF124DE08D87E914F3A", HMACSHA1("foo", "bar"))
}

func TestHMACSha224(t *testing.T) {
	assert := internal.NewAssert(t, "HMACSha224")
	assert.Equal("d7f508375f4b5b1c236d2df1b850de2474a913644876705e62bd78cc", HMACSha224("foo", "bar"))
}

func TestHMACSHA224(t *testing.T) {
	assert := internal.NewAssert(t, "HMACSHA224")
	assert.Equal("D7F508375F4B5B1C236D2DF1B850DE2474A913644876705E62BD78CC", HMACSHA224("foo", "bar"))
}

func TestHMACSha256(t *testing.T) {
	assert := internal.NewAssert(t, "HMACSha256")
	assert.Equal("147933218aaabc0b8b10a2b3a5c34684c8d94341bcf10a4736dc7270f7741851", HMACSha256("foo", "bar"))
}

func TestHMACSHA256(t *testing.T) {
	assert := internal.NewAssert(t, "HMACSHA256")
	assert.Equal("147933218AAABC0B8B10A2B3A5C34684C8D94341BCF10A4736DC7270F7741851", HMACSHA256("foo", "bar"))
}

func TestHMACSha384(t *testing.T) {
	assert := internal.NewAssert(t, "HMACSha384")
	assert.Equal("1d9070d07cb7746e0664cccc6cec1fa996dc7f46368982acfa2095ee8d73fe25b5b6e32279900cdb0fd372a3654e41c5", HMACSha384("foo", "bar"))
}

func TestHMACSHA384(t *testing.T) {
	assert := internal.NewAssert(t, "HMACSHA384")
	assert.Equal("1D9070D07CB7746E0664CCCC6CEC1FA996DC7F46368982ACFA2095EE8D73FE25B5B6E32279900CDB0FD372A3654E41C5", HMACSHA384("foo", "bar"))
}

func TestHMACSha512(t *testing.T) {
	assert := internal.NewAssert(t, "HMACSha512")
	assert.Equal("24257d7210582a65c731ec55159c8184cc24c02489453e58587f71f44c23a2d61b4b72154a89d17b2d49448a8452ea066f4fc56a2bcead45c088572ffccdb3d8", HMACSha512("foo", "bar"))
}

func TestHMACSHA512(t *testing.T) {
	assert := internal.NewAssert(t, "HMACSHA512")
	assert.Equal("24257D7210582A65C731EC55159C8184CC24C02489453E58587F71F44C23A2D61B4B72154A89D17B2D49448A8452EA066F4FC56A2BCEAD45C088572FFCCDB3D8", HMACSHA512("foo", "bar"))
}
