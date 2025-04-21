package maps

import (
	"testing"

	"github.com/dablelv/cyan/internal/utest"
)

func TestKeys(t *testing.T) {
	assert := utest.NewAssert(t, "TestMapKeys")

	// to ints
	ks := Keys(map[int]int{1: 1, 2: 2, 3: 3})
	assert.Equal(3, len(ks))
}

func TestSortedKeys(t *testing.T) {
	assert := utest.NewAssert(t, "TestMapSortedKeys")

	// to strings
	ks := SortedKeys(map[string]int{"a": 1, "b": 2, "c": 3})
	assert.Equal(ks, []string{"a", "b", "c"})
}

func TestVals(t *testing.T) {
	assert := utest.NewAssert(t, "TestMapVals")

	// to ints
	vs := Vals(map[int]int{1: 1, 2: 2, 3: 3})
	assert.Equal(3, len(vs))
}

func TestMapSortedVals(t *testing.T) {
	assert := utest.NewAssert(t, "TestMapSortedVals")

	// to strings
	vs := SortedVals(map[int]string{1: "a", 2: "b", 3: "c"})
	assert.Equal(vs, []string{"a", "b", "c"})
}

func TestMapKeyVals(t *testing.T) {
	assert := utest.NewAssert(t, "TestMapKeyVals")

	ks, vs := KeyVals(map[string]int{"foo": 1, "bar": 2, "baz": 3})
	assert.Equal(3, len(ks))
	assert.Equal(3, len(vs))
}

func TestMerge(t *testing.T) {
	{
		assert := utest.NewAssert(t, "TestMerge_one_map")
		m := map[string]int{"foo": 1, "bar": 2, "baz": 3}
		r := Merge(m)
		assert.Equal(r, m)
	}

	{
		assert := utest.NewAssert(t, "TestMerge_two_map")
		m1 := map[string]int{"foo": 1, "bar": 2, "baz": 3}
		m2 := map[string]int{"qux": 4}
		r := Merge(m1, m2)
		assert.Equal(r, map[string]int{"foo": 1, "bar": 2, "baz": 3, "qux": 4})
	}

	{
		assert := utest.NewAssert(t, "TestMerge_three_map")
		m1 := map[string]int{"foo": 1, "bar": 2, "baz": 3}
		m2 := map[string]int{"qux": 4}
		m3 := map[string]int{"foo": 0}
		r := Merge(m1, m2, m3)
		assert.Equal(r, map[string]int{"foo": 0, "bar": 2, "baz": 3, "qux": 4})
	}
}

func TestContainsAllKeys(t *testing.T) {
	{
		assert := utest.NewAssert(t, "exist")
		m := map[string]int{"foo": 1, "bar": 2, "baz": 3}
		b := ContainsAllKeys(m, []string{"foo", "bar", "baz"})
		assert.True(b)
	}

	{
		assert := utest.NewAssert(t, "not_exist")
		m := map[string]int{"foo": 1, "bar": 2, "baz": 3}
		b := ContainsAllKeys(m, []string{"foo", "bar", "baz", "qux"})
		assert.False(b)
	}
}

func TestContainsAnyKey(t *testing.T) {
	{
		assert := utest.NewAssert(t, "exist")
		m := map[string]int{"foo": 1, "bar": 2, "baz": 3}
		b := ContainsAnyKey(m, []string{"foo", "bar", "baz"})
		assert.True(b)
	}

	{
		assert := utest.NewAssert(t, "not_exist")
		m := map[string]int{"foo": 1, "bar": 2, "baz": 3}
		b := ContainsAnyKey(m, []string{"qux"})
		assert.False(b)
	}
}
