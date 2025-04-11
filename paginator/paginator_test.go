package paginator

import (
	"testing"

	"github.com/dablelv/cyan/internal/utest"
)

func TestPaginate(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// The result is first page.
	{
		assert := utest.NewAssert(t, "TestPaginate_FirstPage")

		p := New(1, 2)
		r := Paginate(items, p)
		assert.Equal(r, []int{1, 2})
	}

	// The result is middle page.
	{
		assert := utest.NewAssert(t, "TestPaginate_MiddlePage")

		p := New(2, 2)
		r := Paginate(items, p)
		assert.Equal(r, []int{3, 4})
	}

	// The result is last page.
	{
		assert := utest.NewAssert(t, "TestPaginate_LastPage")

		p := New(5, 2)
		r := Paginate(items, p)
		assert.Equal(r, []int{9, 10})
	}

	// The result over the slice.
	{
		assert := utest.NewAssert(t, "TestPaginate_PageOver")

		p := New(6, 2)
		r := Paginate(items, p)
		assert.IsNil(r)
	}
}
