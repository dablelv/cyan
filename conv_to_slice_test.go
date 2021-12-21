package util

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap2Slice(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name  string
		args  args
		wantK interface{}
		wantV interface{}
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotK, gotV := Map2Slice(tt.args.i)
			assert.Equal(t, reflect.ValueOf(gotK).Len(), reflect.ValueOf(tt.wantK).Len())
			assert.Equal(t, reflect.ValueOf(gotV).Len(), reflect.ValueOf(tt.wantV).Len())
		})
	}
}

func TestToUint64Slice(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		args args
		want []uint64
	}{
		{"nil to uint64 slice",
			args{nil},
			nil,
		},
		{
			"uint64 slice to uint64 slice",
			args{[]uint64{1, 2, 3}},
			[]uint64{1, 2, 3},
		},
		{
			"uint64 array to uint64 slice",
			args{[3]uint64{1, 2, 3}},
			[]uint64{1, 2, 3},
		},
		{
			"uint64 value to uint64 slice",
			args{uint64(1)},
			[]uint64{1},
		},
		{
			"string slice to uint64 slice",
			args{[]string{"1", "2", "3"}},
			[]uint64{1, 2, 3},
		},
		{
			"string array to uint64 slice",
			args{[3]string{"1", "2", "3"}},
			[]uint64{1, 2, 3},
		},
		{"string value to uint64 slice",
			args{"1"},
			[]uint64{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUint64Slice(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToUint64Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToStrSlice(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"nil to string slice",
			args{nil},
			nil,
		},
		{
			"string slice to string slice",
			args{[]string{"a", "b", "c"}},
			[]string{"a", "b", "c"},
		},
		{
			"string array to string slice",
			args{[3]string{"a", "b", "c"}},
			[]string{"a", "b", "c"},
		},
		{
			"string value to string slice",
			args{"a"},
			[]string{"a"},
		},
		{
			"string value split by white space character to string slice",
			args{"a b c"},
			[]string{"a", "b", "c"},
		},
		{
			"uint64 slice to string slice",
			args{[]uint64{1, 2, 3}},
			[]string{"1", "2", "3"},
		},
		{
			"uint64 array to string slice",
			args{[3]uint64{1, 2, 3}},
			[]string{"1", "2", "3"},
		},
		{"uint64 value to string slice",
			args{uint64(1)},
			[]string{"1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToStrSlice(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToStrSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint32Slice(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		args args
		want []uint32
	}{
		{"nil to uint32 slice",
			args{nil},
			nil,
		},
		{
			"uint32 slice to uint32 slice",
			args{[]uint32{1, 2, 3}},
			[]uint32{1, 2, 3},
		},
		{
			"uint32 array to uint32 slice",
			args{[3]uint32{1, 2, 3}},
			[]uint32{1, 2, 3},
		},
		{
			"uint32 value to uint32 slice",
			args{uint32(1)},
			[]uint32{1},
		},
		{
			"string slice to uint32 slice",
			args{[]string{"1", "2", "3"}},
			[]uint32{1, 2, 3},
		},
		{
			"string array to uint32 slice",
			args{[3]string{"1", "2", "3"}},
			[]uint32{1, 2, 3},
		},
		{"string value to uint32 slice",
			args{"1"},
			[]uint32{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUint32Slice(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToUint32Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint16Slice(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		args args
		want []uint16
	}{
		{"nil to uint16 slice",
			args{nil},
			nil,
		},
		{
			"uint16 slice to uint16 slice",
			args{[]uint16{1, 2, 3}},
			[]uint16{1, 2, 3},
		},
		{
			"uint16 array to uint16 slice",
			args{[3]uint16{1, 2, 3}},
			[]uint16{1, 2, 3},
		},
		{
			"uint16 value to uint16 slice",
			args{uint16(1)},
			[]uint16{1},
		},
		{
			"string slice to uint16 slice",
			args{[]string{"1", "2", "3"}},
			[]uint16{1, 2, 3},
		},
		{
			"string array to uint16 slice",
			args{[3]string{"1", "2", "3"}},
			[]uint16{1, 2, 3},
		},
		{"string value to uint16 slice",
			args{"1"},
			[]uint16{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUint16Slice(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToUint16Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint8Slice(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		args args
		want []uint8
	}{
		{"nil to uint8 slice",
			args{nil},
			nil,
		},
		{
			"uint8 slice to uint8 slice",
			args{[]uint8{1, 2, 3}},
			[]uint8{1, 2, 3},
		},
		{
			"uint8 array to uint8 slice",
			args{[3]uint8{1, 2, 3}},
			[]uint8{1, 2, 3},
		},
		{
			"uint8 value to uint8 slice",
			args{uint8(1)},
			[]uint8{1},
		},
		{
			"string slice to uint8 slice",
			args{[]string{"1", "2", "3"}},
			[]uint8{1, 2, 3},
		},
		{
			"string array to uint8 slice",
			args{[3]string{"1", "2", "3"}},
			[]uint8{1, 2, 3},
		},
		{"string value to uint8 slice",
			args{"1"},
			[]uint8{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUint8Slice(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToUint8Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt64Slice(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{"nil to int64 slice",
			args{nil},
			nil,
		},
		{
			"int64 slice to int64 slice",
			args{[]int64{1, 2, 3}},
			[]int64{1, 2, 3},
		},
		{
			"int64 array to int64 slice",
			args{[3]int64{1, 2, 3}},
			[]int64{1, 2, 3},
		},
		{
			"int64 value to int64 slice",
			args{int64(1)},
			[]int64{1},
		},
		{
			"string slice to int64 slice",
			args{[]string{"1", "2", "3"}},
			[]int64{1, 2, 3},
		},
		{
			"string array to int64 slice",
			args{[3]string{"1", "2", "3"}},
			[]int64{1, 2, 3},
		},
		{"string value to int64 slice",
			args{"1"},
			[]int64{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt64Slice(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToInt64Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt32Slice(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{"nil to int32 slice",
			args{nil},
			nil,
		},
		{
			"int32 slice to int32 slice",
			args{[]int32{1, 2, 3}},
			[]int32{1, 2, 3},
		},
		{
			"int32 array to int32 slice",
			args{[3]int32{1, 2, 3}},
			[]int32{1, 2, 3},
		},
		{
			"int32 value to int32 slice",
			args{int32(1)},
			[]int32{1},
		},
		{
			"string slice to int32 slice",
			args{[]string{"1", "2", "3"}},
			[]int32{1, 2, 3},
		},
		{
			"string array to int32 slice",
			args{[3]string{"1", "2", "3"}},
			[]int32{1, 2, 3},
		},
		{"string value to int32 slice",
			args{"1"},
			[]int32{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt32Slice(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToInt32Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt16Slice(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		args args
		want []int16
	}{
		{"nil to int16 slice",
			args{nil},
			nil,
		},
		{
			"int16 slice to int16 slice",
			args{[]int16{1, 2, 3}},
			[]int16{1, 2, 3},
		},
		{
			"int16 array to int16 slice",
			args{[3]int16{1, 2, 3}},
			[]int16{1, 2, 3},
		},
		{
			"int16 value to int16 slice",
			args{int16(1)},
			[]int16{1},
		},
		{
			"string slice to int16 slice",
			args{[]string{"1", "2", "3"}},
			[]int16{1, 2, 3},
		},
		{
			"string array to int16 slice",
			args{[3]string{"1", "2", "3"}},
			[]int16{1, 2, 3},
		},
		{"string value to int16 slice",
			args{"1"},
			[]int16{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt16Slice(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToInt16Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt8Slice(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		args args
		want []int8
	}{
		{"nil to int8 slice",
			args{nil},
			nil,
		},
		{
			"int8 slice to int8 slice",
			args{[]int8{1, 2, 3}},
			[]int8{1, 2, 3},
		},
		{
			"int8 array to int8 slice",
			args{[3]int8{1, 2, 3}},
			[]int8{1, 2, 3},
		},
		{
			"int8 value to int8 slice",
			args{int8(1)},
			[]int8{1},
		},
		{
			"string slice to int8 slice",
			args{[]string{"1", "2", "3"}},
			[]int8{1, 2, 3},
		},
		{
			"string array to int8 slice",
			args{[3]string{"1", "2", "3"}},
			[]int8{1, 2, 3},
		},
		{"string value to int8 slice",
			args{"1"},
			[]int8{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt8Slice(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToInt8Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUintSlice(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		args args
		want []uint
	}{
		{"nil to uint slice",
			args{nil},
			nil,
		},
		{
			"uint slice to uint slice",
			args{[]uint{1, 2, 3}},
			[]uint{1, 2, 3},
		},
		{
			"uint array to uint slice",
			args{[3]uint{1, 2, 3}},
			[]uint{1, 2, 3},
		},
		{
			"uint value to uint slice",
			args{uint(1)},
			[]uint{1},
		},
		{
			"string slice to uint slice",
			args{[]string{"1", "2", "3"}},
			[]uint{1, 2, 3},
		},
		{
			"string array to uint slice",
			args{[3]string{"1", "2", "3"}},
			[]uint{1, 2, 3},
		},
		{"string value to uint slice",
			args{"1"},
			[]uint{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUintSlice(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToUintSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToIntSlice(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"nil to int slice",
			args{nil},
			[]int{},
		},
		{
			"int slice to int slice",
			args{[]int{1, 2, 3}},
			[]int{1, 2, 3},
		},
		{
			"int array to int slice",
			args{[3]int{1, 2, 3}},
			[]int{1, 2, 3},
		},
		{
			"int value to int slice fail",
			args{int(1)},
			[]int{},
		},
		{
			"string slice to int slice",
			args{[]string{"1", "2", "3"}},
			[]int{1, 2, 3},
		},
		{
			"string array to int slice",
			args{[3]string{"1", "2", "3"}},
			[]int{1, 2, 3},
		},
		{"string value to int slice fail",
			args{"1"},
			[]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToIntSlice(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToIntSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToBoolSlice(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{"nil to int slice",
			args{nil},
			[]bool{},
		},
		{
			"bool slice to bool slice",
			args{[]bool{true, false, true}},
			[]bool{true, false, true},
		},
		{
			"bool array to bool slice",
			args{[3]bool{true, false, true}},
			[]bool{true, false, true},
		},
		{
			"string slice to bool slice",
			args{[]string{"true", "false", "true"}},
			[]bool{true, false, true},
		},
		{
			"string array to bool slice",
			args{[3]string{"true", "false", "true"}},
			[]bool{true, false, true},
		},
		{
			"int slice to bool slice",
			args{[]int{1, 0, 1}},
			[]bool{true, false, true},
		},
		{
			"int array to bool slice",
			args{[3]int{1, 0, 1}},
			[]bool{true, false, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToBoolSlice(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToBoolSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
