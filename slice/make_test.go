package slice

import (
	"reflect"
	"testing"
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
