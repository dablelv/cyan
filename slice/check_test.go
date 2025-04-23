package slice

import (
	"testing"

	"github.com/dablelv/cyan/internal/utest"
)

func TestContains(t *testing.T) {
	assert := utest.NewAssert(t, "TestContains")

	assert.Equal(true, Contains([]string{"foo", "bar", "baz"}, "bar"))
	assert.Equal(false, Contains([]string{"foo", "bar", "baz"}, "qux"))
	assert.Equal(true, Contains([]int{1, 2, 3}, 1))
}

func TestContainsRef(t *testing.T) {
	assert := utest.NewAssert(t, "TestContainsRef")

	assert.Equal(true, ContainsRef([]string{"foo", "bar", "baz"}, "bar"))
	assert.Equal(false, ContainsRef([]string{"foo", "bar", "baz"}, "qux"))
	assert.Equal(true, ContainsRef([]int{1, 2, 3}, 1))
	assert.Equal(false, ContainsRef([]int32{1, 2, 3}, 1))
}

func TestContainsAll(t *testing.T) {
	{
		// empty elements return true.
		assert := utest.NewAssert(t, "empty elements")
		contains := ContainsAll([]int{1, 2, 3}, nil)
		assert.True(contains)
	}

	{
		// conatains all.
		assert := utest.NewAssert(t, "conatains all")
		contains := ContainsAll([]int{1, 2, 3, 4, 5}, []int{1, 2, 3})
		assert.True(contains)
	}

	{
		// conatains only one element.
		assert := utest.NewAssert(t, "conatains only one element")
		contains := ContainsAll([]int{1, 2, 3, 4, 5}, []int{1, 10})
		assert.False(contains)
	}
}
