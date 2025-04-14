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
