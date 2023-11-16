package encoding

import (
	"testing"

	"github.com/dablelv/cyan/internal"
)

func TestGBKToUTF8Str(t *testing.T) {
	assert := internal.NewAssert(t, "GBKToUTF8Str")
	assert.Equal("中文", GbkToUtf8Str(string([]byte{0xD6, 0xD0, 0xCE, 0xC4})))
}

func TestUTF8ToGBKStr(t *testing.T) {
	assert := internal.NewAssert(t, "UTF8ToGBKStr")
	assert.Equal(string([]byte{0xD6, 0xD0, 0xCE, 0xC4}), Utf8ToGbkStr("中文"))
}
