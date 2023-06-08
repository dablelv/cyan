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
			name: "int8 slice",
			args: args{[]int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			want: 55,
		},
		{
			name: "int16 slice",
			args: args{[]int16{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			want: 55,
		},
		{
			name: "int32 slice",
			args: args{[]int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			want: 55,
		},
		{
			name: "int64 slice",
			args: args{[]int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			want: 55,
		},
		{
			name: "uint slice",
			args: args{[]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			want: 55,
		},
		{
			name: "uint8 slice",
			args: args{[]uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			want: 55,
		},
		{
			name: "uint16 slice",
			args: args{[]uint16{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			want: 55,
		},
		{
			name: "uint32 slice",
			args: args{[]uint32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			want: 55,
		},
		{
			name: "uint64 slice",
			args: args{[]uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			want: 55,
		},
		{
			name: "float32 slice",
			args: args{[]float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			want: 55,
		},
		{
			name: "float64 slice",
			args: args{[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			want: 55,
		},
		{
			name: "string slice invalid",
			args: args{[]string{"foo"}},
			want: 0.0,
		},
		{
			name: "nil",
			args: args{nil},
			want: 0.0,
		},
		{
			name: "",
			args: args{""},
			want: 0.0,
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
