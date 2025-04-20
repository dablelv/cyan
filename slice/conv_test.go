package slice

import (
	"testing"

	"github.com/dablelv/cyan/internal/utest"
)

func TestToMap(t *testing.T) {

	getKey := func(v string) string {
		return "prefix_" + v
	}

	s := []string{"foo", "bar", "baz"}

	assert := utest.NewAssert(t, "Succ")
	m := ToMap(s, getKey)
	assert.Equal(m, map[string]string{
		"prefix_foo": "foo",
		"prefix_bar": "bar",
		"prefix_baz": "baz",
	})
}
