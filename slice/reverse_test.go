package slice

import (
	"testing"

	"github.com/dablelv/go-huge-util/internal"
)

func TestReverse(t *testing.T) {
	assert := internal.NewAssert(t, "TestReverse")

	r := Reverse([]int{1, 2, 3})
	assert.Equal([]int{3, 2, 1}, r)

	r = Reverse([]int{})
	assert.Equal([]int{}, r)

	r = Reverse([]int(nil))
	assert.IsNil(r)

	r1 := Reverse([]uint{1, 2, 3})
	assert.Equal([]uint{3, 2, 1}, r1)

	r2 := Reverse([]string{"foo", "bar", "baz"})
	assert.Equal([]string{"baz", "bar", "foo"}, r2)
}

func TestReverseRef(t *testing.T) {
	assert := internal.NewAssert(t, "TestReverseRef")

	r := ReverseRef([]int{1, 2, 3})
	assert.Equal([]int{3, 2, 1}, r)

	r = ReverseRef([]int{})
	assert.Equal([]int{}, r)

	r = ReverseRef([]int(nil))
	assert.IsNil(r)

	r1 := ReverseRef([]uint{1, 2, 3})
	assert.Equal([]uint{3, 2, 1}, r1)

	r2 := ReverseRef([]string{"foo", "bar", "baz"})
	assert.Equal([]string{"baz", "bar", "foo"}, r2)
}
