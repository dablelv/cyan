package slice

import (
	"testing"

	"github.com/dablelv/cyan/internal/utest"
)

func TestIntersection(t *testing.T) {
	{
		assert := utest.NewAssert(t, "left slice is nil")
		r := Intersection(nil, []string{"foo", "bar", "baz"})
		assert.IsNil(r)
	}
	{
		assert := utest.NewAssert(t, "right slice is nil")
		r := Intersection([]string{"foo", "bar", "baz"}, nil)
		assert.IsNil(r)
	}
	{
		assert := utest.NewAssert(t, "has intersection")
		s1 := []string{"foo", "bar", "baz"}
		r := Intersection(s1, []string{"bar", "baz", "qux"})
		assert.Equal(r, []string{"bar", "baz"})
	}
	{
		assert := utest.NewAssert(t, "no intersection")
		r := Intersection([]string{"foo", "bar", "baz"}, []string{"qux", "quux"})
		assert.Equal(r, []string{})
	}
}

func TestIntersectionFunc(t *testing.T) {
	type User struct {
		Name string
	}

	{
		assert := utest.NewAssert(t, "set1_is_nil")
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
		assert := utest.NewAssert(t, "no_intersection")
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
		assert := utest.NewAssert(t, "has_intersection")
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

func TestDiff(t *testing.T) {
	{
		assert := utest.NewAssert(t, "left slice is nil")
		r := Diff(nil, []string{"foo", "bar", "baz"})
		assert.IsNil(r)
	}
	{
		assert := utest.NewAssert(t, "right slice is nil")
		s1 := []string{"foo", "bar", "baz"}
		r := Diff(s1, nil)
		assert.Equal(r, s1)
	}
	{
		assert := utest.NewAssert(t, "right slice is not nil")
		s1 := []string{"foo", "bar", "baz"}
		r := Diff(s1, []string{"bar", "baz"})
		assert.Equal(r, []string{"foo"})
	}
}

func TestDiffFunc(t *testing.T) {
	type User struct {
		Name string
	}

	{
		assert := utest.NewAssert(t, "TestDiffFunc_set1_is_nil")
		set1 := []User(nil)
		set2 := []User{
			{
				Name: "a",
			},
		}
		getKey := func(s User) string {
			return s.Name
		}
		r := DiffFunc(set1, set2, getKey)
		assert.IsNil(r)
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
