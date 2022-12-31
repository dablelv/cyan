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

func TestMinInt8Slice(t *testing.T) {
	type args struct {
		s []int8
	}
	tests := []struct {
		name string
		args args
		want int8
	}{
		{"min int8 slice", args{[]int8{1, 2, 3}}, 1},
		{"empty int8 slice", args{[]int8{}}, 0},
		{"nil", args{nil}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinInt8Slice(tt.args.s); got != tt.want {
				t.Errorf("MinInt8Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinInt16Slice(t *testing.T) {
	type args struct {
		sl []int16
	}
	tests := []struct {
		name string
		args args
		want int16
	}{
		{"min int16 slice", args{[]int16{1, 2, 3}}, 1},
		{"empty int16 slice", args{[]int16{}}, 0},
		{"nil", args{nil}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinInt16Slice(tt.args.sl); got != tt.want {
				t.Errorf("MinInt16Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinInt32Slice(t *testing.T) {
	type args struct {
		sl []int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{"min int32 slice", args{[]int32{1, 2, 3}}, 1},
		{"empty int32 slice", args{[]int32{}}, 0},
		{"nil", args{nil}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinInt32Slice(tt.args.sl); got != tt.want {
				t.Errorf("MinInt32Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinInt64Slice(t *testing.T) {
	type args struct {
		sl []int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"min int64 slice", args{[]int64{1, 2, 3}}, 1},
		{"empty int64 slice", args{[]int64{}}, 0},
		{"nil", args{nil}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinInt64Slice(tt.args.sl); got != tt.want {
				t.Errorf("MinInt64Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinUintSlice(t *testing.T) {
	type args struct {
		sl []uint
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{"min uint slice", args{[]uint{1, 2, 3}}, 1},
		{"empty uint slice", args{[]uint{}}, 0},
		{"nil", args{nil}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinUintSlice(tt.args.sl); got != tt.want {
				t.Errorf("MinUintSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinUint8Slice(t *testing.T) {
	type args struct {
		sl []uint8
	}
	tests := []struct {
		name string
		args args
		want uint8
	}{
		{"min uint8 slice", args{[]uint8{1, 2, 3}}, 1},
		{"empty uint8 slice", args{[]uint8{}}, 0},
		{"nil", args{nil}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinUint8Slice(tt.args.sl); got != tt.want {
				t.Errorf("MinUint8Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinUint16Slice(t *testing.T) {
	type args struct {
		sl []uint16
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{"min uint16 slice", args{[]uint16{1, 2, 3}}, 1},
		{"empty uint16 slice", args{[]uint16{}}, 0},
		{"nil", args{nil}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinUint16Slice(tt.args.sl); got != tt.want {
				t.Errorf("MinUint16Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinUint32Slice(t *testing.T) {
	type args struct {
		sl []uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{"min uint32 slice", args{[]uint32{1, 2, 3}}, 1},
		{"empty uint32 slice", args{[]uint32{}}, 0},
		{"nil", args{nil}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinUint32Slice(tt.args.sl); got != tt.want {
				t.Errorf("MinUint32Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinUint64Slice(t *testing.T) {
	type args struct {
		sl []uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{"min uint64 slice", args{[]uint64{1, 2, 3}}, 1},
		{"empty uint64 slice", args{[]uint64{}}, 0},
		{"nil", args{nil}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinUint64Slice(tt.args.sl); got != tt.want {
				t.Errorf("MinUint64Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinFloat32Slice(t *testing.T) {
	type args struct {
		sl []float32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{"min float32 slice", args{[]float32{1.1, 2.2, 3.3}}, 1.1},
		{"empty float32 slice", args{[]float32{}}, 0},
		{"nil", args{nil}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinFloat32Slice(tt.args.sl); got != tt.want {
				t.Errorf("MinFloat32Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinFloat64Slice(t *testing.T) {
	type args struct {
		sl []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"min float64 slice", args{[]float64{1.1, 2.2, 3.3}}, 1.1},
		{"empty float64 slice", args{[]float64{}}, 0},
		{"nil", args{nil}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinFloat64Slice(tt.args.sl); got != tt.want {
				t.Errorf("MinFloat64Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxIntSlice(t *testing.T) {
	type args struct {
		sl []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"max int slice", args{[]int{1, 2, 3}}, 3},
		{"empty int slice", args{[]int{}}, 0},
		{"nil", args{nil}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxIntSlice(tt.args.sl); got != tt.want {
				t.Errorf("MaxIntSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxInt8Slice(t *testing.T) {
	type args struct {
		sl []int8
	}
	tests := []struct {
		name string
		args args
		want int8
	}{
		{"max int8 slice", args{[]int8{1, 2, 3}}, 3},
		{"empty int8 slice", args{[]int8{}}, 0},
		{"nil", args{nil}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxInt8Slice(tt.args.sl); got != tt.want {
				t.Errorf("MaxInt8Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxInt16Slice(t *testing.T) {
	type args struct {
		sl []int16
	}
	tests := []struct {
		name string
		args args
		want int16
	}{
		{"max int16 slice", args{[]int16{1, 2, 3}}, 3},
		{"empty int16 slice", args{[]int16{}}, 0},
		{"nil", args{nil}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxInt16Slice(tt.args.sl); got != tt.want {
				t.Errorf("MaxInt16Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxInt32Slice(t *testing.T) {
	type args struct {
		sl []int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{"max int32 slice", args{[]int32{1, 2, 3}}, 3},
		{"empty int32 slice", args{[]int32{}}, 0},
		{"nil", args{nil}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxInt32Slice(tt.args.sl); got != tt.want {
				t.Errorf("MaxInt32Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxInt64Slice(t *testing.T) {
	type args struct {
		sl []int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"max int64 slice", args{[]int64{1, 2, 3}}, 3},
		{"empty int64 slice", args{[]int64{}}, 0},
		{"nil", args{nil}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxInt64Slice(tt.args.sl); got != tt.want {
				t.Errorf("MaxInt64Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxUintSlice(t *testing.T) {
	type args struct {
		sl []uint
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{"max uint slice", args{[]uint{1, 2, 3}}, 3},
		{"empty uint slice", args{[]uint{}}, 0},
		{"nil", args{nil}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxUintSlice(tt.args.sl); got != tt.want {
				t.Errorf("MaxUintSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxUint8Slice(t *testing.T) {
	type args struct {
		sl []uint8
	}
	tests := []struct {
		name string
		args args
		want uint8
	}{
		{"max uint8 slice", args{[]uint8{1, 2, 3}}, 3},
		{"empty uint8 slice", args{[]uint8{}}, 0},
		{"nil", args{nil}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxUint8Slice(tt.args.sl); got != tt.want {
				t.Errorf("MaxUint8Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxUint16Slice(t *testing.T) {
	type args struct {
		sl []uint16
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{"max uint16 slice", args{[]uint16{1, 2, 3}}, 3},
		{"empty uint16 slice", args{[]uint16{}}, 0},
		{"nil", args{nil}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxUint16Slice(tt.args.sl); got != tt.want {
				t.Errorf("MaxUint16Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxUint32Slice(t *testing.T) {
	type args struct {
		sl []uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{"max uint32 slice", args{[]uint32{1, 2, 3}}, 3},
		{"empty uint32 slice", args{[]uint32{}}, 0},
		{"nil", args{nil}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxUint32Slice(tt.args.sl); got != tt.want {
				t.Errorf("MaxUint32Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxUint64Slice(t *testing.T) {
	type args struct {
		sl []uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{"max uint64 slice", args{[]uint64{1, 2, 3}}, 3},
		{"empty uint64 slice", args{[]uint64{}}, 0},
		{"nil", args{nil}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxUint64Slice(tt.args.sl); got != tt.want {
				t.Errorf("MaxUint64Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxFloat32Slice(t *testing.T) {
	type args struct {
		sl []float32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{"max float32 slice", args{[]float32{1.1, 2.2, 3.3}}, 3.3},
		{"empty float32 slice", args{[]float32{}}, 0},
		{"nil", args{nil}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxFloat32Slice(tt.args.sl); got != tt.want {
				t.Errorf("MaxFloat32Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxFloat64Slice(t *testing.T) {
	type args struct {
		sl []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"max float64 slice", args{[]float64{1.1, 2.2, 3.3}}, 3.3},
		{"empty float64 slice", args{[]float64{}}, 0},
		{"nil", args{nil}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxFloat64Slice(tt.args.sl); got != tt.want {
				t.Errorf("MaxFloat64Slice() = %v, want %v", got, tt.want)
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
		{"min int", args{[]int{1, 2, 3}}, 1, false},
		{"empty int", args{[]int{}}, nil, true},
		{"not slice", args{""}, nil, true},
		{"nil", args{nil}, nil, true},
		{"min uint", args{[]uint{1, 2, 3}}, uint(1), false},
		{"empty uint", args{[]uint{}}, nil, true},
		{"min float64", args{[]float64{1.0, 2.0, 3.0}}, 1.0, false},
		{"empty float64", args{[]float64{}}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MinSliceE(tt.args.slice)
			if (err != nil) != tt.wantErr {
				t.Errorf("MinSliceE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MinSliceE() = %v, want %v", got, tt.want)
				return
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
		{"max int slice", args{[]int{1, 2, 3}}, 3, false},
		{"empty int slice", args{[]int{}}, nil, true},
		{"not slice", args{""}, nil, true},
		{"nil", args{nil}, nil, true},
		{"max uint slice", args{[]uint{1, 2, 3}}, uint(3), false},
		{"empty uint slice", args{[]uint{}}, nil, true},
		{"max float64 slice", args{[]float64{1.0, 2.0, 3.0}}, 3.0, false},
		{"empty float64 slice", args{[]float64{}}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MaxSliceE(tt.args.slice)
			if (err != nil) != tt.wantErr {
				t.Errorf("MaxSliceE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MaxSliceE() = %v, want %v", got, tt.want)
				return
			}
		})
	}
}
