package slice

import (
	"strings"
	"testing"

	"github.com/dablelv/cyan/internal/utest"
)

func TestInsert(t *testing.T) {
	assert := utest.NewAssert(t, "TestInsert")

	s := Insert([]string{"foo", "qux"}, 1, "bar", "baz")
	assert.Equal([]string{"foo", "bar", "baz", "qux"}, s)

	i := Insert([]int{1, 2, 3}, 3, 4, 5)
	assert.Equal([]int{1, 2, 3, 4, 5}, i)

	i = Insert([]int{1, 2, 3}, 0, 0)
	assert.Equal([]int{0, 1, 2, 3}, i)
}

func TestInsertRef(t *testing.T) {
	assert := utest.NewAssert(t, "TestInsertRef")

	r := InsertRef([]string{"bar", "baz"}, 0, "foo")
	assert.Equal([]string{"foo", "bar", "baz"}, r)

	r = InsertRef([]string{"foo", "bar"}, 2, "baz")
	assert.Equal([]string{"foo", "bar", "baz"}, r)

	r = InsertRef([]int{1, 4}, 1, 2, 3)
	assert.Equal([]int{1, 2, 3, 4}, r)
}

func TestDelete(t *testing.T) {
	assert := utest.NewAssert(t, "TestDelete")

	s := Delete([]int{1, 2, 3}, 0)
	assert.Equal([]int{2, 3}, s)

	s = Delete([]int{1, 2, 3}, 2)
	assert.Equal([]int{1, 2}, s)

	s = Delete([]int{1, 2, 3}, 0, 1, 2)
	assert.Equal([]int{}, s)

	r := Delete([]string{"foo", "bar", "baz"}, 0, 1)
	assert.Equal([]string{"baz"}, r)
}

func TestDeleteRef(t *testing.T) {
	assert := utest.NewAssert(t, "TestDeleteRef")

	s := DeleteRef([]int{1, 2, 3}, 0)
	assert.Equal([]int{2, 3}, s)

	s = DeleteRef([]int{1, 2, 3}, 2)
	assert.Equal([]int{1, 2}, s)

	s = DeleteRef([]int{1, 2, 3}, 0, 1, 2)
	assert.Equal([]int{}, s)

	r := DeleteRef([]string{"foo", "bar", "baz"}, 0, 1)
	assert.Equal([]string{"baz"}, r)
}

func TestDeleteElems(t *testing.T) {
	assert := utest.NewAssert(t, "TestDeleteElems")

	s := DeleteElems([]string{"a", "b", "b", "c"}, "b", "c")
	assert.Equal([]string{"a"}, s)

	s = DeleteElems([]string{"a", "b", "b", "c"}, "a", "b", "c")
	assert.Equal([]string{}, s)

	r := DeleteElems([]int{1, 2, 2, 3}, 2)
	assert.Equal([]int{1, 3}, r)
}

func TestDeleteFunc(t *testing.T) {
	{
		// delete the sepecified prefix strings.
		assert := utest.NewAssert(t, "DeleteString")
		s := []string{"foo", "bar", "baz", "qux"}
		r := DeleteFunc(s, func(s string) bool {
			return strings.HasPrefix(s, "b")
		})
		assert.Equal(r, []string{"foo", "qux"})
		assert.Equal(s, []string{"foo", "bar", "baz", "qux"})
	}

	{
		// delete the odd numbers.
		assert := utest.NewAssert(t, "DeleteOdd")
		s := []int{0, 1, 1, 2, 3, 5, 8}
		r := DeleteFunc(s, func(n int) bool {
			return n%2 != 0
		})
		assert.Equal(r, []int{0, 2, 8})
		assert.Equal(s, []int{0, 1, 1, 2, 3, 5, 8})
	}
}

func TestDeleteElemsRef(t *testing.T) {
	assert := utest.NewAssert(t, "TestDeleteElemsRef")

	s := DeleteElemsRef([]string{"a", "b", "b", "c"}, "b", "c")
	assert.Equal([]string{"a"}, s)

	s = DeleteElemsRef([]string{"a", "b", "b", "c"}, "a", "b", "c")
	assert.Equal([]string{}, s)

	r := DeleteElemsRef([]int{1, 2, 2, 3}, 2)
	assert.Equal([]int{1, 3}, r)
}

func TestIndexes(t *testing.T) {
	assert := utest.NewAssert(t, "TestIndexes")

	s := Indexes([]int{1, 2, 3}, 1)
	assert.Equal([]int{0}, s)

	s = Indexes([]int{1, 1, 2, 3}, 1)
	assert.Equal([]int{0, 1}, s)

	s = Indexes([]int{1, 1, 2, 3}, 4)
	assert.IsNil(s)

	s = Indexes([]string{"foo", "bar", "bar", "baz"}, "bar")
	assert.Equal([]int{1, 2}, s)
}

func TestIndexesRef(t *testing.T) {
	assert := utest.NewAssert(t, "IndexesRef")

	s := IndexesRef([]int{1, 2, 3}, 2)
	assert.Equal([]int{1}, s)

	s = IndexesRef([]int{1, 1, 2, 3}, 1)
	assert.Equal([]int{0, 1}, s)

	s = IndexesRef([]int{1, 1, 2, 3}, 4)
	assert.IsNil(s)

	s = IndexesRef("foo", 4)
	assert.IsNil(s)
}

func TestIndexessFunc(t *testing.T) {
	assert := utest.NewAssert(t, "TestIndexessFunc")

	f := func(e int) bool {
		return e <= 1
	}
	s := IndexesFunc([]int{1, 2, 3}, f)
	assert.Equal([]int{0}, s)

	s = IndexesFunc([]int{1, 1, 2, 3}, f)
	assert.Equal([]int{0, 1}, s)

	s = IndexesFunc([]int{2, 3, 4}, f)
	assert.IsNil(s)
}

func TestReplace(t *testing.T) {
	assert := utest.NewAssert(t, "TestReplace")

	r1 := Replace([]string{"fooo", "barr", "baz", "qux"}, 0, 2, "foo", "bar")
	assert.Equal([]string{"foo", "bar", "baz", "qux"}, r1)

	r2 := Replace([]int{1, 2, 4}, 2, 3, 3)
	assert.Equal([]int{1, 2, 3}, r2)
}

func TestReplaceRef(t *testing.T) {
	assert := utest.NewAssert(t, "TestReplaceRef")

	r1 := ReplaceRef([]string{"fooo", "barr", "baz", "qux"}, 0, 2, "foo", "bar")
	assert.Equal([]string{"foo", "bar", "baz", "qux"}, r1)

	r2 := ReplaceRef([]int{1, 2, 4}, 2, 3, 3)
	assert.Equal([]int{1, 2, 3}, r2)
}
