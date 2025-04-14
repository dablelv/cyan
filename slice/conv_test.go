package slice

import (
	"testing"

	"github.com/dablelv/cyan/internal/utest"
)

func TestGroupFunc(t *testing.T) {
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
		m := GroupFunc(users, getKey)
		assert.Equal(m, want)
	}

	// has same class.
	{
		assert := utest.NewAssert(t, "TestGroupFunc_HasSameClass")
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
		m := GroupFunc(users, getKey)
		assert.Equal(m, want)
	}

	// no item.
	{
		assert := utest.NewAssert(t, "TestGroupFunc_NoItem")
		users := []User{}
		getKey := func(user User) int {
			return user.Class
		}
		want := map[int][]User{}
		m := GroupFunc(users, getKey)
		assert.Equal(m, want)
	}
}
