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

func TestClearZero(t *testing.T) {
	assert := utest.NewAssert(t, "TestClearZero")

	assert.Equal([]bool{true, true}, ClearZero([]bool{true, false, true}))
	assert.Equal([]int{1, 2, 3}, ClearZero([]int{1, 2, 0, 3}))
	assert.Equal([]string{"foo", "bar", "baz"}, ClearZero([]string{"foo", "bar", "", "baz"}))
}

func TestClearZeroRef(t *testing.T) {
	assert := utest.NewAssert(t, "TestClearZeroRef")

	assert.Equal([]bool{true, true}, ClearZeroRef([]bool{true, false, true}))
	assert.Equal([]int{1, 2, 3}, ClearZeroRef([]int{1, 2, 0, 3}))
	assert.Equal([]string{"foo", "bar", "baz"}, ClearZeroRef([]string{"foo", "bar", "", "baz"}))

	maps := []map[string]string{
		{"foo": "foo"},
		nil,
		{"bar": "bar"},
	}
	assert.Equal([]map[string]string{{"foo": "foo"}, {"bar": "bar"}}, ClearZeroRef(maps))
}

func TestReverse(t *testing.T) {
	assert := utest.NewAssert(t, "TestReverse")

	r := Reverse([]int{1, 2, 3})
	assert.Equal([]int{3, 2, 1}, r)

	r = Reverse([]int{})
	assert.Equal([]int{}, r)

	r = Reverse([]int(nil))
	assert.IsNil(r)

	r1 := Reverse([]uint{1, 2, 3})
	assert.Equal([]uint{3, 2, 1}, r1)

	r2 := Reverse([]string{"foo", "bar", "baz"})
	assert.Equal([]string{"baz", "bar", "foo"}, r2)
}

func TestReverseRef(t *testing.T) {
	assert := utest.NewAssert(t, "TestReverseRef")

	r := ReverseRef([]int{1, 2, 3})
	assert.Equal([]int{3, 2, 1}, r)

	r = ReverseRef([]int{})
	assert.Equal([]int{}, r)

	r = ReverseRef([]int(nil))
	assert.IsNil(r)

	r1 := ReverseRef([]uint{1, 2, 3})
	assert.Equal([]uint{3, 2, 1}, r1)

	r2 := ReverseRef([]string{"foo", "bar", "baz"})
	assert.Equal([]string{"baz", "bar", "foo"}, r2)
}

func TestDistinct(t *testing.T) {
	{
		assert := utest.NewAssert(t, "TestDistinctNil")
		assert.IsNil(Distinct([]int(nil)))
	}

	{
		assert := utest.NewAssert(t, "TestDistinctEmpty")
		assert.Equal(Distinct([]int{}), []int{})
	}

	{
		assert := utest.NewAssert(t, "TesDistinctInt")

		assert.Equal([]int{1, 2, 3}, Distinct([]int{1, 2, 2, 3}))
		assert.Equal([]int{}, Distinct([]int{}))
		assert.IsNil(Distinct([]int(nil)))
	}
}

func TestDistinctRef(t *testing.T) {
	assert := utest.NewAssert(t, "TestDistinctRef")

	assert.Equal([]int{1, 2, 3}, DistinctRef([]int{1, 2, 2, 3}))
	assert.Equal([]int{}, DistinctRef([]int{}))
	assert.Equal([]string{"foo", "bar", "baz"}, DistinctRef([]string{"foo", "bar", "bar", "baz"}))
	assert.Equal([]string{}, DistinctRef([]string{}))
}

func TestDistinctFunc(t *testing.T) {
	type User struct {
		Name  string
		Class int
	}

	// no repeated elements.
	{
		assert := utest.NewAssert(t, "DistinctFunc_no_repeated_element")
		users := []User{{"alice", 1}, {"bob", 2}, {"carl", 3}}
		getKey := func(user User) int {
			return user.Class
		}
		r := DistinctFunc(users, getKey)
		assert.Equal(r, users)
	}

	// has repeated elements.
	{
		assert := utest.NewAssert(t, "DistinctFunc_no_repeated_element")
		users := []User{{"alice", 1}, {"bob", 2}, {"bob", 2}}
		getKey := func(user User) int {
			return user.Class
		}
		r := DistinctFunc(users, getKey)
		assert.Equal(r, []User{{"alice", 1}, {"bob", 2}})
	}
}

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
		assert := utest.NewAssert(t, "NoItem")
		users := []User{}
		getKey := func(user User) int {
			return user.Class
		}
		want := map[int][]User{}
		m := GroupFunc(users, getKey)
		assert.Equal(m, want)
	}
}

func TestExtractField(t *testing.T) {
	type User struct {
		Id   int
		Name string
	}

	{
		assert := utest.NewAssert(t, "Nil")
		fields, err := ExtractField[User, string]([]User(nil), "Name")
		assert.IsNil(err)
		assert.IsNil(fields)
	}

	{
		assert := utest.NewAssert(t, "Empty")
		fields, err := ExtractField[User, string]([]User{}, "Name")
		assert.IsNil(err)
		assert.IsNil(fields)
	}

	users := []User{
		{1, "alice"},
		{2, "bob"},
		{3, "charles"},
	}

	{
		assert := utest.NewAssert(t, "Succ")
		fields, err := ExtractField[User, string](users, "Name")
		assert.IsNil(err)
		assert.Equal(fields, []string{"alice", "bob", "charles"})
	}

	{
		assert := utest.NewAssert(t, "NoneFiled")
		fields, err := ExtractField[User, string](users, "NoneFiled")
		assert.IsNotNil(err)
		assert.IsNil(fields)
	}
}
