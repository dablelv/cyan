package slice

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumSlice(t *testing.T) {
	f32Slice := []float32{1.1, 2.2, 3.3}
	f64 := SumSlice(f32Slice)
	assert.Equal(t, float32(6.6), float32(f64))
}

func TestIsContains(t *testing.T) {
	type args struct {
		slice  any
		target any
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "string slice contain",
			args: args{slice: []string{"foo", "bar", "baz"}, target: "baz"},
			want: true,
		},
		{
			name: "string array contain",
			args: args{slice: [3]string{"foo", "bar", "baz"}, target: "baz"},
			want: true,
		},
		{
			name: "string slice not contain",
			args: args{slice: []string{"foo", "bar", "baz"}, target: "qux"},
			want: false,
		},
		{
			name: "int slice contain",
			args: args{slice: []int{1, 2, 3}, target: 1},
			want: true,
		},
		{
			name: "int32 slice not contain because type isn't equal",
			args: args{slice: []int32{1, 2, 3}, target: 1},
			want: false,
		},
		{
			name: "nil not contain",
			args: args{slice: nil, target: 1},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsContains(tt.args.slice, tt.args.target); got != tt.want {
				t.Errorf("IsContains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJoinSliceWithSepE(t *testing.T) {
	type args struct {
		slice any
		sep   string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "join string slice",
			args:    args{[]string{"foo", "bar", "baz"}, ","},
			want:    "foo,bar,baz",
			wantErr: false,
		},
		{
			name:    "join string array",
			args:    args{[3]string{"foo", "bar", "baz"}, ","},
			want:    "foo,bar,baz",
			wantErr: false,
		},
		{
			name:    "join int slice",
			args:    args{[]int{1, 2, 3}, ","},
			want:    "1,2,3",
			wantErr: false,
		},
		{
			name:    "join int array",
			args:    args{[3]int{1, 2, 3}, ","},
			want:    "1,2,3",
			wantErr: false,
		},
		{
			name:    "join nil",
			args:    args{nil, ","},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := JoinSliceWithSepE(tt.args.slice, tt.args.sep)
			if (err != nil) != tt.wantErr {
				t.Errorf("JoinSliceWithSepE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("JoinSliceWithSepE() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRandomSliceElem(t *testing.T) {
	type args struct {
		i any
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			name: "Get a random element from a slice of length one",
			args: args{[]int{1}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRandomSliceElem(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRandomSliceElem() = %v, want %v", got, tt.want)
			}
		})
	}
}
