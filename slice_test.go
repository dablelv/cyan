package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseIntSlice(t *testing.T) {
	intSlice := []int{1, 2, 3}
	assert.Equal(t, []int{3, 2, 1}, ReverseIntSlice(intSlice))
}

func TestSumSlice(t *testing.T) {
	f32Slice := []float32{1.1, 2.2, 3.3}
	f64 := SumSlice(f32Slice)
	assert.Equal(t, float32(6.6), float32(f64))
}

func TestMinSliceE(t *testing.T) {
	type args struct {
		slice interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{"err int", args{""}, nil, true},
		{"zero int", args{[]int{}}, nil, false},
		{"min int", args{[]int{1, 2, 3}}, 1, false},
		{"err uint", args{""}, nil, true},
		{"zero uint", args{[]uint{}}, nil, false},
		{"min uint", args{[]uint{1, 2, 3}}, uint(1), false},
		{"err float", args{""}, nil, true},
		{"zero float", args{[]float64{}}, nil, false},
		{"min float", args{[]float64{1.0, 2.0, 3.0}}, 1.0, false},
	}
	for _, tt := range tests {
		got, err := MinSliceE(tt.args.slice)
		assert.Equal(t, tt.wantErr, err != nil, tt.name)
		assert.Equal(t, tt.want, got, tt.name)
	}
}

func TestMaxSliceE(t *testing.T) {
	type args struct {
		slice interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{"err int", args{""}, nil, true},
		{"zero int", args{[]int{}}, nil, false},
		{"max int", args{[]int{1, 2, 3}}, 3, false},
		{"err uint", args{""}, nil, true},
		{"zero uint", args{[]uint{}}, nil, false},
		{"max uint", args{[]uint{1, 2, 3}}, uint(3), false},
		{"err float", args{""}, nil, true},
		{"zero float", args{[]float64{}}, nil, false},
		{"max float", args{[]float64{1.0, 2.0, 3.0}}, 3.0, false},
	}
	for _, tt := range tests {
		got, err := MaxSliceE(tt.args.slice)
		assert.Equal(t, tt.wantErr, err != nil, tt.name)
		assert.Equal(t, tt.want, got, tt.name)
	}
}
