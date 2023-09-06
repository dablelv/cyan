package time

import (
	"testing"
	"time"

	"github.com/dablelv/cyan/internal"
)

func TestTimeCounter(t *testing.T) {
	assert := internal.NewAssert(t, "TestTimeCounter")

	c := NewTimeCounter()

	time.Sleep(time.Second)

	cost := c.GetD()
	assert.LessOrEqual(time.Second, cost)

	sec := c.GetS()
	assert.Equal(int64(1), sec)

	ms := c.GetMs()
	assert.LessOrEqual(int64(time.Second)/int64(time.Millisecond), ms)

	us := c.GetUs()
	assert.LessOrEqual(int64(time.Second)/int64(time.Microsecond), us)

	ns := c.GetNs()
	assert.LessOrEqual(int64(time.Second)/int64(time.Nanosecond), int64(ns))
}

func TestTimeCost(t *testing.T) {
	assert := internal.NewAssert(t, "TestTimeCost")

	cost := TimeCost()
	time.Sleep(time.Second)
	d := cost()
	assert.LessOrEqual(time.Second, d)
	assert.Equal(int64(1), int64(d)/int64(time.Second))
}
