package slice

import (
	"reflect"
	"testing"

	"github.com/dablelv/cyan/internal/utest"
)

func TestMakeSlice(t *testing.T) {
	type args struct {
		val  int
		size []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"make specified length int slice with same value",
			args{1, []int{3}},
			[]int{1, 1, 1},
		},
		{
			"make specified length and capacity int slice with same value",
			args{1, []int{3, 3}},
			[]int{1, 1, 1},
		},
		{
			"err: input param is invalid",
			args{1, []int{}},
			[]int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Make(tt.args.val, tt.args.size...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistinct(t *testing.T) {
	type args[T comparable] struct {
		data []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "nil",
			args: args[int]{
				data: nil,
			},
			want: nil,
		}, {
			name: "empty",
			args: args[int]{
				data: []int{},
			},
			want: []int{},
		}, {
			name: "normal",
			args: args[int]{
				data: []int{1, 2, 3},
			},
			want: []int{1, 2, 3},
		}, {
			name: "normal 1",
			args: args[int]{
				data: []int{1, 1, 2, 2, 3, 3},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Distinct(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Distinct() = %v, want %v", got, tt.want)
			}
		})
	}
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
