package slice

import (
	"testing"

	"github.com/dablelv/cyan/internal"
)

func TestRandomElem(t *testing.T) {
	assert := internal.NewAssert(t, "TestRandomElem")

	got := RandomElem([]int{1})
	assert.Equal(1, got)

	got = RandomElem([]string{"foo"})
	assert.Equal("foo", got)
}

func TestMin(t *testing.T) {
	assert := internal.NewAssert(t, "TestMin")

	assert.Equal(1, Min([]int{1, 2, 3}))
	assert.Equal(0, Min([]int{}))
	assert.Equal(uint(1), Min([]uint{1, 2, 3}))
	assert.Equal(uint(0), Min([]uint{}))
	assert.Equal(1.0, Min([]float64{1.0, 2.0, 3.0}))
	assert.Equal(0.0, Min([]float64{}))
	assert.Equal(0.0, Min([]float64(nil)))
}

func TestMax(t *testing.T) {
	assert := internal.NewAssert(t, "TestMax")

	assert.Equal(3, Max([]int{1, 2, 3}))
	assert.Equal(0, Max([]int{}))
	assert.Equal(uint(3), Max([]uint{1, 2, 3}))
	assert.Equal(uint(0), Max([]uint{}))
	assert.Equal(3.0, Max([]float64{1.0, 2.0, 3.0}))
	assert.Equal(0.0, Max([]float64{}))
	assert.Equal(0.0, Max([]float64(nil)))
}

func TestMinRef(t *testing.T) {
	assert := internal.NewAssert(t, "TestMaxRef")

	assert.Equal(1, MinRef([]int{1, 2, 3}))
	assert.IsNil(MinRef([]int{}))
	assert.Equal(int8(1), MinRef([]int8{1, 2, 3}))
	assert.Equal(int16(1), MinRef([]int16{1, 2, 3}))
	assert.Equal(int32(1), MinRef([]int32{1, 2, 3}))
	assert.Equal(int64(1), MinRef([]int64{1, 2, 3}))
	assert.Equal(uint(1), MinRef([]uint{1, 2, 3}))
	assert.Equal(uint(1), MinRef([]uint{1, 2, 3}))
	assert.Equal(uint8(1), MinRef([]uint8{1, 2, 3}))
	assert.Equal(uint16(1), MinRef([]uint16{1, 2, 3}))
	assert.Equal(uint32(1), MinRef([]uint32{1, 2, 3}))
	assert.Equal(uint64(1), MinRef([]uint64{1, 2, 3}))
	assert.Equal(float32(1.0), MinRef([]float32{1.0, 2.0, 3.0}))
	assert.Equal(1.0, MinRef([]float64{1.0, 2.0, 3.0}))
}

func TestMaxRef(t *testing.T) {
	assert := internal.NewAssert(t, "TestMaxRef")

	assert.Equal(3, MaxRef([]int{1, 2, 3}))
	assert.IsNil(MaxRef([]int{}))
	assert.Equal(int8(3), MaxRef([]int8{1, 2, 3}))
	assert.Equal(int16(3), MaxRef([]int16{1, 2, 3}))
	assert.Equal(int32(3), MaxRef([]int32{1, 2, 3}))
	assert.Equal(int64(3), MaxRef([]int64{1, 2, 3}))
	assert.Equal(uint(3), MaxRef([]uint{1, 2, 3}))
	assert.Equal(uint8(3), MaxRef([]uint8{1, 2, 3}))
	assert.Equal(uint16(3), MaxRef([]uint16{1, 2, 3}))
	assert.Equal(uint32(3), MaxRef([]uint32{1, 2, 3}))
	assert.Equal(uint64(3), MaxRef([]uint64{1, 2, 3}))
	assert.Equal(float32(3.0), MaxRef([]float32{1.0, 2.0, 3.0}))
	assert.Equal(3.0, MaxRef([]float64{1.0, 2.0, 3.0}))
}

func TestSum(t *testing.T) {
	assert := internal.NewAssert(t, "TestSum")

	assert.Equal(float64(55), Sum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	assert.Equal(float64(55), Sum([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
}

func TestSumRef(t *testing.T) {
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
			name: "",
			args: args{""},
			want: 0.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumRef(tt.args.s); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
