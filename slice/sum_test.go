package slice

import (
	"testing"
)

func TestSumInt(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "1",
			args: args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			want: 55,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.s); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumSlice(t *testing.T) {
	type args struct {
		s any
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "int slice",
			args: args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			want: 55,
		},
		{
			name: "uint slice",
			args: args{[]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			want: 55,
		},
		{
			name: "float64 slice",
			args: args{[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			want: 55,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumSlice(tt.args.s); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
