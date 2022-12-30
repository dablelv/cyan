package slice

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	type args struct {
		s any
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{"nil float64 slice", args{[]float64(nil)}, 0.0},
		{"empty float64 slice", args{[]float64{}}, 0.0},
		{"min float64 slice", args{[]float64{1.0, 2.0, 3.0}}, 1.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.args.s.([]float64)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinE(t *testing.T) {
	type args struct {
		slice any
	}
	tests := []struct {
		name    string
		args    args
		want    any
		wantErr bool
	}{
		{"err nil int", args{[]int(nil)}, 0, true},
		{"err empty int", args{[]int{}}, 0, true},
		{"min int", args{[]int{1, 2, 3}}, 1, false},
	}
	for _, tt := range tests {
		got, err := MinE(tt.args.slice.([]int))
		assert.Equal(t, tt.wantErr, err != nil, tt.name)
		assert.Equal(t, tt.want, got, tt.name)
	}
}

func TestMax(t *testing.T) {
	type args struct {
		s any
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{"nil float64 slice", args{[]float64{}}, 0.0},
		{"empty float64 slice", args{[]float64{}}, 0.0},
		{"max float64 slice", args{[]float64{1.0, 2.0, 3.0}}, 3.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.s.([]float64)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MinSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxE(t *testing.T) {
	type args struct {
		slice any
	}
	tests := []struct {
		name    string
		args    args
		want    any
		wantErr bool
	}{
		{"nil int slice", args{[]int(nil)}, 0, true},
		{"empty int slice", args{[]int{}}, 0, true},
		{"max int", args{[]int{1, 2, 3}}, 3, false},
	}
	for _, tt := range tests {
		got, err := MaxE(tt.args.slice.([]int))
		assert.Equal(t, tt.wantErr, err != nil, tt.name)
		assert.Equal(t, tt.want, got, tt.name)
	}
}

func TestMinIntSlice(t *testing.T) {
	type args struct {
		sl []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"min int", args{[]int{1, 2, 3}}, 1},
		{"empty int", args{[]int{}}, 0},
		{"nil int", args{nil}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinIntSlice(tt.args.sl); got != tt.want {
				t.Errorf("MinIntSlice() = %v, want %v", got, tt.want)
			}
		})
	}
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
		{"min int", args{[]int{1, 2, 3}}, 1, false},
		{"empty int", args{[]int{}}, nil, false},
		{"not slice", args{""}, nil, true},
		{"empty uint", args{[]uint{}}, nil, false},
		{"min uint", args{[]uint{1, 2, 3}}, uint(1), false},
		{"empty float", args{[]float64{}}, nil, false},
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
		{"zero uint", args{[]uint{}}, nil, false},
		{"max uint", args{[]uint{1, 2, 3}}, uint(3), false},
		{"zero float", args{[]float64{}}, nil, false},
		{"max float", args{[]float64{1.0, 2.0, 3.0}}, 3.0, false},
	}
	for _, tt := range tests {
		got, err := MaxSliceE(tt.args.slice)
		assert.Equal(t, tt.wantErr, err != nil, tt.name)
		assert.Equal(t, tt.want, got, tt.name)
	}
}
