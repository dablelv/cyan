package slice

import (
	"testing"

	"github.com/dablelv/go-huge-util/internal"
)

func TestRandomElem(t *testing.T) {
	assert := internal.NewAssert(t, "TestRandomElem")

	got := RandomElem([]int{1})
	assert.Equal(1, got)

	got = RandomElem("foo")
	assert.Equal("foo", got)
}
