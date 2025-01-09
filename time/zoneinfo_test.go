package time

import (
	"testing"
	"time"

	"github.com/dablelv/cyan/internal/utest"
)

func TestGetLocByUtcOffset(t *testing.T) {
	{
		assert := utest.NewAssert(t, "UTC")
		loc := GetLocByUtcOffset(0)
		assert.Equal(loc, time.UTC)
	}

	{
		assert := utest.NewAssert(t, "UTC+08:00")
		offset := time.Hour * 8
		want := time.FixedZone("UTC+08:00", 8*60*60)
		loc := GetLocByUtcOffset(offset)
		assert.Equal(loc, want)
	}

	{
		assert := utest.NewAssert(t, "UTC+08:44")
		offset := time.Hour*8 + time.Minute*44
		want := time.FixedZone("UTC+08:44", 8*60*60+44*60)
		loc := GetLocByUtcOffset(offset)
		assert.Equal(loc, want)
	}

	{
		assert := utest.NewAssert(t, "UTC-08:44")
		offset := -(time.Hour*8 + time.Minute*44)
		want := time.FixedZone("UTC-08:44", -(8*60*60 + 44*60))
		loc := GetLocByUtcOffset(offset)
		assert.Equal(loc, want)
	}
}
