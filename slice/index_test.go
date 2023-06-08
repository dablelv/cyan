package slice

import (
	"testing"

	"github.com/dablelv/go-huge-util/internal"
)

func TestGetElemIndexes(t *testing.T) {
	assert := internal.NewAssert(t, "TestGetElemIndexes")

	s := GetElemIndexes([]int{1, 2, 3}, 1)
	assert.Equal([]int{0}, s)

	s = GetElemIndexes([]int{1, 2, 3}, 2)
	assert.Equal([]int{1}, s)

	s = GetElemIndexes([]int{1, 1, 2, 3}, 1)
	assert.Equal([]int{0, 1}, s)

	s = GetElemIndexes([]int{1, 1, 2, 3}, 4)
	assert.IsNil(s)

	s = GetElemIndexes("foo", 4)
	assert.IsNil(s)
}

func TestGetRandomElem(t *testing.T) {
	assert := internal.NewAssert(t, "TestGetRandomElem")

	got := GetRandomElem([]int{1})
	assert.Equal(1, got)

	got = GetRandomElem("foo")
	assert.Equal("foo", got)
}
