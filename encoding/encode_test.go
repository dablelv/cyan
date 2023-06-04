package encoding

import (
	"testing"

	"github.com/dablelv/go-huge-util/internal"
)

func TestGBKToUTF8Str(t *testing.T) {
	assert := internal.NewAssert(t, "GBKToUTF8Str")
	assert.Equal("中文", GBKToUTF8Str(string([]byte{0xD6, 0xD0, 0xCE, 0xC4})))
}

func TestUTF8ToGBKStr(t *testing.T) {
	assert := internal.NewAssert(t, "UTF8ToGBKStr")
	assert.Equal(string([]byte{0xD6, 0xD0, 0xCE, 0xC4}), UTF8ToGBKStr("中文"))
}
