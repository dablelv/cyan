package slice

import (
	"testing"

	"github.com/dablelv/cyan/internal/utest"
)

func TestIntersectionFunc(t *testing.T) {
	type User struct {
		Name string
	}

	{
		assert := utest.NewAssert(t, "TestIntersectionFunc_set1_is_nil")
		set1 := []User{}
		set2 := []User{
			{
				Name: "a",
			},
		}
		getKey := func(s User) string {
			return s.Name
		}
		r := IntersectionFunc(set1, set2, getKey)
		assert.Equal(r, []User{})
	}

	{
		assert := utest.NewAssert(t, "TestIntersectionFunc_no_intersection")
		set1 := []User{
			{
				Name: "a",
			},
		}
		set2 := []User{
			{
				Name: "b",
			},
		}
		getKey := func(s User) string {
			return s.Name
		}
		r := IntersectionFunc(set1, set2, getKey)
		assert.Equal(r, []User{})
	}

	{
		assert := utest.NewAssert(t, "TestIntersectionFunc_has_intersection")
		set1 := []User{
			{
				Name: "a",
			},
			{
				Name: "b",
			},
			{
				Name: "c",
			},
		}
		set2 := []User{
			{
				Name: "b",
			},
			{
				Name: "c",
			},
			{
				Name: "d",
			},
		}
		getKey := func(s User) string {
			return s.Name
		}
		r := IntersectionFunc(set1, set2, getKey)
		assert.Equal(r, []User{{"b"}, {"c"}})
	}
}

func TestDiffFunc(t *testing.T) {
	type User struct {
		Name string
	}

	{
		assert := utest.NewAssert(t, "TestDiffFunc_set1_is_nil")
		set1 := []User{}
		set2 := []User{
			{
				Name: "a",
			},
		}
		getKey := func(s User) string {
			return s.Name
		}
		r := DiffFunc(set1, set2, getKey)
		assert.Equal(r, []User{})
	}

	{
		assert := utest.NewAssert(t, "TestDiffFunc_no_intersection")
		set1 := []User{
			{
				Name: "a",
			},
		}
		set2 := []User{
			{
				Name: "b",
			},
		}
		getKey := func(s User) string {
			return s.Name
		}
		r := DiffFunc(set1, set2, getKey)
		assert.Equal(r, set1)
	}

	{
		assert := utest.NewAssert(t, "TestDiffFunc_has_intersection")
		set1 := []User{
			{
				Name: "a",
			},
			{
				Name: "b",
			},
			{
				Name: "c",
			},
		}
		set2 := []User{
			{
				Name: "b",
			},
			{
				Name: "c",
			},
			{
				Name: "d",
			},
		}
		getKey := func(s User) string {
			return s.Name
		}
		r := DiffFunc(set1, set2, getKey)
		assert.Equal(r, []User{{"a"}})
	}
}

func TestSymDiffFunc(t *testing.T) {
	type User struct {
		Name string
	}
	getKey := func(s User) string {
		return s.Name
	}

	{
		assert := utest.NewAssert(t, "TestSymDiffFunc_set1_is_nil")
		set1 := []User{}
		set2 := []User{
			{
				Name: "a",
			},
		}
		r := SymDiffFunc(set1, set2, getKey)
		assert.Equal(r, set2)
	}

	{
		assert := utest.NewAssert(t, "TestSymDiffFunc_no_intersection")
		set1 := []User{
			{
				Name: "a",
			},
		}
		set2 := []User{
			{
				Name: "b",
			},
		}
		r := SymDiffFunc(set1, set2, getKey)
		assert.Equal(r, Merge(set1, set2))
	}

	{
		assert := utest.NewAssert(t, "TestSymDiffFunc_has_intersection")
		set1 := []User{
			{
				Name: "a",
			},
			{
				Name: "b",
			},
			{
				Name: "c",
			},
		}
		set2 := []User{
			{
				Name: "b",
			},
			{
				Name: "c",
			},
			{
				Name: "d",
			},
		}
		r := SymDiffFunc(set1, set2, getKey)
		assert.Equal(r, []User{{"a"}, {"d"}})
	}
}
