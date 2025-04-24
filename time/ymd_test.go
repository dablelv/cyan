package time

import (
	"testing"
	"time"

	"github.com/dablelv/cyan/internal/utest"
)

func TestTimeToYmd(t *testing.T) {
	assert := utest.NewAssert(t, "TestToYmd")

	l, _ := time.LoadLocation(AsiaShanghai)
	tim, _ := time.ParseInLocation(time.DateTime, "2008-08-08 20:00:00", l)
	ymd := TimeToYmd(tim)
	assert.Equal(ymd, Ymd(20080808))
}

func TestYmd_Sub(t *testing.T) {
	assert := utest.NewAssert(t, "TestYmd_Sub")

	ymd := Ymd(20080808)
	days, err := ymd.Sub(Ymd(20080731))
	assert.IsNil(err)
	assert.Equal(days, 8)
}
