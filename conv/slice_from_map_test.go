package conv

import (
	"reflect"
	"testing"

	"github.com/dablelv/cyan/internal"
)

func TestMapToSlice(t *testing.T) {
	assert := internal.NewAssert(t, "TestMapToSlice")

	type args struct {
		i any
	}
	tests := []struct {
		name  string
		args  args
		wantK any
		wantV any
	}{
		{
			name: "map[string]string to slice ",
			args: args{
				map[string]string{"1": "1", "2": "2", "3": "3"},
			},
			wantK: []string{"1", "2", "3"},
			wantV: []string{"1", "2", "3"},
		},
		{
			name: "map[int]string to slice",
			args: args{
				map[int]string{1: "1", 2: "2", 3: "3"},
			},
			wantK: []int{1, 2, 3},
			wantV: []string{"1", "2", "3"},
		},
		{
			name: "map[int]int to slice",
			args: args{
				map[int]int{1: 1, 2: 2, 3: 3},
			},
			wantK: []int{1, 2, 3},
			wantV: []int{1, 2, 3},
		},
		{
			name: "empty map[int]int to slice",
			args: args{
				map[int]int{},
			},
			wantK: []int{},
			wantV: []int{},
		},
		{
			name:  "nil to slice",
			args:  args{nil},
			wantK: nil,
			wantV: nil,
		},
		{
			name:  "not map failed",
			args:  args{1},
			wantK: nil,
			wantV: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotK, gotV := MapToSlice(tt.args.i)
			if gotK != nil && gotV != nil {
				assert.Equal(reflect.ValueOf(gotK).Len(), reflect.ValueOf(tt.wantK).Len())
				assert.Equal(reflect.ValueOf(gotV).Len(), reflect.ValueOf(tt.wantV).Len())
			}
		})
	}
}

func TestMapKeysInt(t *testing.T) {
	type args struct {
		m map[int]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1 2 3",
			args: args{map[int]int{1: 1, 2: 2, 3: 3}},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapKeys(tt.args.m); !reflect.DeepEqual(len(got), len(tt.want)) {
				t.Errorf("MapKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapValsInt(t *testing.T) {
	type args struct {
		m map[int]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1 2 3",
			args: args{map[int]int{1: 1, 2: 2, 3: 3}},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapVals(tt.args.m); !reflect.DeepEqual(len(got), len(tt.want)) {
				t.Errorf("MapKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapKeysValsStrInt(t *testing.T) {
	type args struct {
		m map[string]int
	}
	tests := []struct {
		name  string
		args  args
		want  []string
		want1 []int
	}{
		{
			name:  "1 2 3",
			args:  args{map[string]int{"foo": 1, "bar": 2, "baz": 3}},
			want:  []string{"foo", "bar", "baz"},
			want1: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := MapKeysVals(tt.args.m)
			if !reflect.DeepEqual(len(got), len(tt.want)) {
				t.Errorf("MapKeysVals() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(len(got1), len(tt.want1)) {
				t.Errorf("MapKeysVals() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
