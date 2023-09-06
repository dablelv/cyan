package slice

import (
	"testing"

	"github.com/dablelv/cyan/internal"
)

func TestContains(t *testing.T) {
	assert := internal.NewAssert(t, "TestContains")

	assert.Equal(true, Contains([]string{"foo", "bar", "baz"}, "bar"))
	assert.Equal(false, Contains([]string{"foo", "bar", "baz"}, "qux"))
	assert.Equal(true, Contains([]int{1, 2, 3}, 1))
}

func TestContainsRef(t *testing.T) {
	assert := internal.NewAssert(t, "TestContainsRef")

	assert.Equal(true, ContainsRef([]string{"foo", "bar", "baz"}, "bar"))
	assert.Equal(false, ContainsRef([]string{"foo", "bar", "baz"}, "qux"))
	assert.Equal(true, ContainsRef([]int{1, 2, 3}, 1))
	assert.Equal(false, ContainsRef([]int32{1, 2, 3}, 1))
}

func TestClearZero(t *testing.T) {
	assert := internal.NewAssert(t, "TestClearZero")

	assert.Equal([]bool{true, true}, ClearZero([]bool{true, false, true}))
	assert.Equal([]int{1, 2, 3}, ClearZero([]int{1, 2, 0, 3}))
	assert.Equal([]string{"foo", "bar", "baz"}, ClearZero([]string{"foo", "bar", "", "baz"}))
}

func TestClearZeroRef(t *testing.T) {
	assert := internal.NewAssert(t, "TestClearZeroRef")

	assert.Equal([]bool{true, true}, ClearZeroRef([]bool{true, false, true}))
	assert.Equal([]int{1, 2, 3}, ClearZeroRef([]int{1, 2, 0, 3}))
	assert.Equal([]string{"foo", "bar", "baz"}, ClearZeroRef([]string{"foo", "bar", "", "baz"}))

	maps := []map[string]string{
		{"foo": "foo"},
		nil,
		{"bar": "bar"},
	}
	assert.Equal([]map[string]string{{"foo": "foo"}, {"bar": "bar"}}, ClearZeroRef(maps))
}

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

func TestUnique(t *testing.T) {
	assert := internal.NewAssert(t, "TestUnique")

	assert.Equal([]int{1, 2, 3}, Unique([]int{1, 2, 2, 3}))
	assert.Equal([]int{}, Unique([]int{}))
	assert.IsNil(Unique([]int(nil)))
}

func TestUniqueRef(t *testing.T) {
	assert := internal.NewAssert(t, "TestUniqueRef")

	assert.Equal([]int{1, 2, 3}, UniqueRef([]int{1, 2, 2, 3}))
	assert.Equal([]int{}, UniqueRef([]int{}))
	assert.Equal([]string{"foo", "bar", "baz"}, UniqueRef([]string{"foo", "bar", "bar", "baz"}))
	assert.Equal([]string{}, UniqueRef([]string{}))
}
