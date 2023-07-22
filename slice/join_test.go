package slice

import (
	"testing"

	"github.com/dablelv/go-huge-util/internal"
)

func TestJoin(t *testing.T) {
	assert := internal.NewAssert(t, "TestJoinWithSep")

	got := Join([]int{1}, "")
	assert.Equal("1", got)
}

func TestJoinE(t *testing.T) {
	assert := internal.NewAssert(t, "TestJoinE")

	got := Join([]string{"foo", "bar", "baz"}, ",")
	assert.Equal("foo,bar,baz", got)

	got = Join([3]string{"foo", "bar", "baz"}, ",")
	assert.Equal("foo,bar,baz", got)

	got = Join([]int{1, 2, 3}, ",")
	assert.Equal("1,2,3", got)

	got = Join([3]int{1, 2, 3}, ",")
	assert.Equal("1,2,3", got)

	got = Join([]struct{}{{}}, ",")
	assert.Equal("", got)

	got = Join(nil, ",")
	assert.Equal("", got)
}
