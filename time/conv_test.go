package time

import (
	"testing"

	"github.com/dablelv/cyan/internal"
)

// Note that Local timezone may be different so the returned value may be different.
func TestDateTimeToTs(t *testing.T) {
	assert := internal.NewAssert(t, "TestDateTimeToTs")

	got := DateTimeToTs("1970-01-01 00:00:01")
	assert.NotEqual(int64(0), got)
}

func TestTsToDateTime(t *testing.T) {
	assert := internal.NewAssert(t, "TestTsToDateTime")

	got := TsToDateTime(0)
	assert.NotEqual("", got)
}
