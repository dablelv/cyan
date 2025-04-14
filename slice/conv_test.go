package slice

import (
	"testing"

	"github.com/dablelv/cyan/internal/utest"
)

func TestGroupBy(t *testing.T) {
	type User struct {
		Name  string
		Class int
	}

	// no same class.
	{
		assert := utest.NewAssert(t, "TestGroupBy_NoSameClass")
		users := []User{{"alice", 1}, {"bob", 2}, {"foo", 3}}
		getKey := func(user User) int {
			return user.Class
		}
		want := map[int][]User{
			1: {
				{Name: "alice", Class: 1},
			},
			2: {
				{Name: "bob", Class: 2},
			},
			3: {
				{Name: "foo", Class: 3},
			},
		}
		m := GroupBy(users, getKey)
		assert.Equal(m, want)
	}

	// has same class.
	{
		assert := utest.NewAssert(t, "TestGroupBy_HasSameClass")
		users := []User{{"alice", 1}, {"bob", 1}, {"foo", 3}}
		getKey := func(user User) int {
			return user.Class
		}
		want := map[int][]User{
			1: {
				{Name: "alice", Class: 1},
				{Name: "bob", Class: 1},
			},
			3: {
				{Name: "foo", Class: 3},
			},
		}
		m := GroupBy(users, getKey)
		assert.Equal(m, want)
	}

	// no item.
	{
		assert := utest.NewAssert(t, "TestGroupBy_NoItem")
		users := []User{}
		getKey := func(user User) int {
			return user.Class
		}
		want := map[int][]User{}
		m := GroupBy(users, getKey)
		assert.Equal(m, want)
	}
}
