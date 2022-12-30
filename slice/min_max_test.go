package slice

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinSlice(t *testing.T) {
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
			if got := MinSlice(tt.args.s.([]float64)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinSliceE(t *testing.T) {
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
		got, err := MinSliceE(tt.args.slice.([]int))
		assert.Equal(t, tt.wantErr, err != nil, tt.name)
		assert.Equal(t, tt.want, got, tt.name)
	}
}

func TestMaxSlice(t *testing.T) {
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
			if got := MaxSlice(tt.args.s.([]float64)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MinSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxSliceE(t *testing.T) {
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
		got, err := MaxSliceE(tt.args.slice.([]int))
		assert.Equal(t, tt.wantErr, err != nil, tt.name)
		assert.Equal(t, tt.want, got, tt.name)
	}
}
