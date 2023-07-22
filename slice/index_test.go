package slice

import (
	"testing"

	"github.com/dablelv/go-huge-util/internal"
)

func TestIndexes(t *testing.T) {
	assert := internal.NewAssert(t, "TestIndexes")

	s := Indexes([]int{1, 2, 3}, 2)
	assert.Equal([]int{1}, s)

	s = Indexes([]int{1, 1, 2, 3}, 1)
	assert.Equal([]int{0, 1}, s)

	s = Indexes([]int{1, 1, 2, 3}, 4)
	assert.IsNil(s)

	s = Indexes("foo", 4)
	assert.IsNil(s)
}

func TestIndices(t *testing.T) {
	assert := internal.NewAssert(t, "TestIndices")

	s := Indices([]int{1, 2, 3}, 1)
	assert.Equal([]int{0}, s)

	s = Indices([]int{1, 1, 2, 3}, 1)
	assert.Equal([]int{0, 1}, s)

	s = Indices([]int{1, 1, 2, 3}, 4)
	assert.IsNil(s)
}

func TestIndicesFunc(t *testing.T) {
	assert := internal.NewAssert(t, "TestIndicesFunc")

	f := func(e int) bool {
		return e <= 1
	}
	s := IndicesFunc([]int{1, 2, 3}, f)
	assert.Equal([]int{0}, s)

	s = IndicesFunc([]int{1, 1, 2, 3}, f)
	assert.Equal([]int{0, 1}, s)

	s = IndicesFunc([]int{2, 3, 4}, f)
	assert.IsNil(s)
}

func TestRandomElem(t *testing.T) {
	assert := internal.NewAssert(t, "TestRandomElem")

	got := RandomElem([]int{1})
	assert.Equal(1, got)

	got = RandomElem("foo")
	assert.Equal("foo", got)
}
