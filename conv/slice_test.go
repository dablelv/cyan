package conv

import (
	"reflect"
	"testing"
	"time"
)

func TestToBoolSliceE(t *testing.T) {
	type args struct {
		a any
	}
	tests := []struct {
		name    string
		args    args
		want    []bool
		wantErr bool
	}{
		{
			"bool slice",
			args{[]bool{true, false, true}},
			[]bool{true, false, true},
			false,
		},
		{
			"bool array",
			args{[3]bool{true, false, true}},
			[]bool{true, false, true},
			false,
		},
		{
			"string slice",
			args{[]string{"true", "false", "true"}},
			[]bool{true, false, true},
			false,
		},
		{
			"string array",
			args{[3]string{"true", "false", "true"}},
			[]bool{true, false, true},
			false,
		},
		{
			"nil string slice",
			args{[]string(nil)},
			nil,
			false,
		},
		{
			"int slice",
			args{[]int{1, 0, 1}},
			[]bool{true, false, true},
			false,
		},
		{
			"int array",
			args{[3]int{1, 0, 1}},
			[]bool{true, false, true},
			false,
		},
		{
			"nil",
			args{nil},
			nil,
			false,
		},
		{
			"string slice failed",
			args{[]string{"foo"}},
			nil,
			true,
		},
		{
			"string failed",
			args{"a"},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToBoolSlice(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToBoolSlice() = %v, want %v", got, tt.want)
			}
			got, err := ToBoolSliceE(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToBoolSliceE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToBoolSliceE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToIntSliceE(t *testing.T) {
	type args struct {
		a any
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			"int slice",
			args{[]int{1, 2, 3}},
			[]int{1, 2, 3},
			false,
		},
		{
			"int array",
			args{[3]int{1, 2, 3}},
			[]int{1, 2, 3},
			false,
		},
		{
			"int",
			args{1},
			[]int{1},
			false,
		},
		{
			"string slice",
			args{[]string{"1", "2", "3"}},
			[]int{1, 2, 3},
			false,
		},
		{
			"string array",
			args{[3]string{"1", "2", "3"}},
			[]int{1, 2, 3},
			false,
		},
		{
			"nil",
			args{nil},
			nil,
			false,
		},
		{
			"complex64 3-5i failed",
			args{complex(3, -5)},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToIntSlice(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToIntSlice() = %v, want %v", got, tt.want)
			}
			got, err := ToIntSliceE(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToIntSliceE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToIntSliceE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt8SliceE(t *testing.T) {
	type args struct {
		a any
	}
	tests := []struct {
		name    string
		args    args
		want    []int8
		wantErr bool
	}{
		{
			"nil",
			args{nil},
			nil,
			false,
		},
		{
			"int8 slice",
			args{[]int8{1, 2, 3}},
			[]int8{1, 2, 3},
			false,
		},
		{
			"int8 array",
			args{[3]int8{1, 2, 3}},
			[]int8{1, 2, 3},
			false,
		},
		{
			"int8 value",
			args{int8(1)},
			[]int8{1},
			false,
		},
		{
			"string slice",
			args{[]string{"1", "2", "3"}},
			[]int8{1, 2, 3},
			false,
		},
		{
			"string array",
			args{[3]string{"1", "2", "3"}},
			[]int8{1, 2, 3},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt8Slice(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToInt8Slice() = %v, want %v", got, tt.want)
			}
			got, err := ToInt8SliceE(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToInt8SliceE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToInt8SliceE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt16SliceE(t *testing.T) {
	type args struct {
		a any
	}
	tests := []struct {
		name    string
		args    args
		want    []int16
		wantErr bool
	}{
		{
			"nil",
			args{nil},
			nil,
			false,
		},
		{
			"int16 slice",
			args{[]int16{1, 2, 3}},
			[]int16{1, 2, 3},
			false,
		},
		{
			"int16 array",
			args{[3]int16{1, 2, 3}},
			[]int16{1, 2, 3},
			false,
		},
		{
			"int16 value",
			args{int16(1)},
			[]int16{1},
			false,
		},
		{
			"string slice",
			args{[]string{"1", "2", "3"}},
			[]int16{1, 2, 3},
			false,
		},
		{
			"string array",
			args{[3]string{"1", "2", "3"}},
			[]int16{1, 2, 3},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt16Slice(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToInt16Slice() = %v, want %v", got, tt.want)
			}
			got, err := ToInt16SliceE(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToInt16SliceE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToInt16SliceE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt32SliceE(t *testing.T) {
	type args struct {
		a any
	}
	tests := []struct {
		name    string
		args    args
		want    []int32
		wantErr bool
	}{
		{
			"int32 slice to int32 slice",
			args{[]int32{1, 2, 3}},
			[]int32{1, 2, 3},
			false,
		},
		{
			"int32 array to int32 slice",
			args{[3]int32{1, 2, 3}},
			[]int32{1, 2, 3},
			false,
		},
		{
			"int32 value to int32 slice",
			args{int32(1)},
			[]int32{1},
			false,
		},
		{
			"string slice to int32 slice",
			args{[]string{"1", "2", "3"}},
			[]int32{1, 2, 3},
			false,
		},
		{
			"string array to int32 slice",
			args{[3]string{"1", "2", "3"}},
			[]int32{1, 2, 3},
			false,
		},
		{
			"nil",
			args{nil},
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt32Slice(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToInt32Slice() = %v, want %v", got, tt.want)
			}
			got, err := ToInt32SliceE(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToInt32SliceE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToInt32SliceE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt64SliceE(t *testing.T) {
	type args struct {
		a any
	}
	tests := []struct {
		name    string
		args    args
		want    []int64
		wantErr bool
	}{
		{
			"int64 slice",
			args{[]int64{1, 2, 3}},
			[]int64{1, 2, 3},
			false,
		},
		{
			"int64 array",
			args{[3]int64{1, 2, 3}},
			[]int64{1, 2, 3},
			false,
		},
		{
			"int64 value",
			args{int64(1)},
			[]int64{1},
			false,
		},
		{
			"string slice",
			args{[]string{"1", "2", "3"}},
			[]int64{1, 2, 3},
			false,
		},
		{
			"string array",
			args{[3]string{"1", "2", "3"}},
			[]int64{1, 2, 3},
			false,
		},
		{
			"nil",
			args{nil},
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt64Slice(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToInt64Slice() = %v, want %v", got, tt.want)
			}
			got, err := ToInt64SliceE(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToInt64SliceE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToInt64SliceE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUintSliceE(t *testing.T) {
	type args struct {
		a any
	}
	tests := []struct {
		name    string
		args    args
		want    []uint
		wantErr bool
	}{
		{
			"uint slice",
			args{[]uint{1, 2, 3}},
			[]uint{1, 2, 3},
			false,
		},
		{
			"uint array",
			args{[3]uint{1, 2, 3}},
			[]uint{1, 2, 3},
			false,
		},
		{
			"string slice",
			args{[]string{"1", "2", "3"}},
			[]uint{1, 2, 3},
			false,
		},
		{
			"string array",
			args{[3]string{"1", "2", "3"}},
			[]uint{1, 2, 3},
			false,
		},
		{
			"nil",
			args{nil},
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUintSlice(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToUintSlice() = %v, want %v", got, tt.want)
			}
			got, err := ToUintSliceE(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToUintSliceE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToUintSliceE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToByteSliceE(t *testing.T) {
	type args struct {
		a any
	}
	tests := []struct {
		name    string
		args    args
		want    []uint8
		wantErr bool
	}{
		{
			"nil",
			args{nil},
			nil,
			false,
		},
		{
			"uint8 slice",
			args{[]uint8{1, 2, 3}},
			[]uint8{1, 2, 3},
			false,
		},
		{
			"uint8 array",
			args{[3]uint8{1, 2, 3}},
			[]uint8{1, 2, 3},
			false,
		},
		{
			"uint8 value",
			args{uint8(1)},
			[]uint8{1},
			false,
		},
		{
			"string slice",
			args{[]string{"1", "2", "3"}},
			[]uint8{1, 2, 3},
			false,
		},
		{
			"string array",
			args{[3]string{"1", "2", "3"}},
			[]uint8{1, 2, 3},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToByteSlice(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToByteSlice() = %v, want %v", got, tt.want)
			}
			got, err := ToByteSliceE(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToByteSliceE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToByteSliceE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint16SliceE(t *testing.T) {
	type args struct {
		a any
	}
	tests := []struct {
		name    string
		args    args
		want    []uint16
		wantErr bool
	}{
		{
			"uint16 slice",
			args{[]uint16{1, 2, 3}},
			[]uint16{1, 2, 3},
			false,
		},
		{
			"uint16 array",
			args{[3]uint16{1, 2, 3}},
			[]uint16{1, 2, 3},
			false,
		},
		{
			"uint16 value",
			args{uint16(1)},
			[]uint16{1},
			false,
		},
		{
			"string slice",
			args{[]string{"1", "2", "3"}},
			[]uint16{1, 2, 3},
			false,
		},
		{
			"string array",
			args{[3]string{"1", "2", "3"}},
			[]uint16{1, 2, 3},
			false,
		},
		{
			"nil",
			args{nil},
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUint16Slice(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToUint16Slice() = %v, want %v", got, tt.want)
			}
			got, err := ToUint16SliceE(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToUint16SliceE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToUint16SliceE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint32SliceE(t *testing.T) {
	type args struct {
		a any
	}
	tests := []struct {
		name    string
		args    args
		want    []uint32
		wantErr bool
	}{
		{
			"uint32 slice",
			args{[]uint32{1, 2, 3}},
			[]uint32{1, 2, 3},
			false,
		},
		{
			"uint32 array",
			args{[3]uint32{1, 2, 3}},
			[]uint32{1, 2, 3},
			false,
		},
		{
			"uint32 value",
			args{uint32(1)},
			[]uint32{1},
			false,
		},
		{
			"string slice",
			args{[]string{"1", "2", "3"}},
			[]uint32{1, 2, 3},
			false,
		},
		{
			"string array",
			args{[3]string{"1", "2", "3"}},
			[]uint32{1, 2, 3},
			false,
		},
		{
			"nil",
			args{nil},
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUint32Slice(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToUint32Slice() = %v, want %v", got, tt.want)
			}
			got, err := ToUint32SliceE(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToUint32SliceE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToUint32SliceE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint64SliceE(t *testing.T) {
	type args struct {
		a any
	}
	tests := []struct {
		name    string
		args    args
		want    []uint64
		wantErr bool
	}{
		{"nil", args{nil}, nil, false},
		{"uint64 slice", args{[]uint64{1, 2, 3}}, []uint64{1, 2, 3}, false},
		{"uint64 array", args{[3]uint64{1, 2, 3}}, []uint64{1, 2, 3}, false},
		{"string slice", args{[]string{"1", "2", "3"}}, []uint64{1, 2, 3}, false},
		{"string array", args{[3]string{"1", "2", "3"}}, []uint64{1, 2, 3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUint64Slice(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToUint64Slice() = %v, want %v", got, tt.want)
			}
			got, err := ToUint64SliceE(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToUint64SliceE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToUint64SliceE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToDurationSliceE(t *testing.T) {
	type args struct {
		a any
	}
	tests := []struct {
		name    string
		args    args
		want    []time.Duration
		wantErr bool
	}{
		{"nil", args{nil}, nil, false},
		{"string slice", args{[]string{"1", "2", "3"}}, []time.Duration{1, 2, 3}, false},
		{"string array", args{[3]string{"1", "2", "3"}}, []time.Duration{1, 2, 3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToDurationSlice(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToDurationSlice() = %v, want %v", got, tt.want)
			}
			got, err := ToDurationSliceE(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToDurationSliceE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToDurationSliceE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToStrSliceE(t *testing.T) {
	type args struct {
		a any
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"nil", args{nil}, nil, false},
		{"string slice", args{[]string{"a", "b", "c"}}, []string{"a", "b", "c"}, false},
		{"string array", args{[3]string{"a", "b", "c"}}, []string{"a", "b", "c"}, false},
		{"string value", args{"a"}, []string{"a"}, false},
		{"string value separated by white space character", args{"a b c"}, []string{"a", "b", "c"}, false},
		{"uint64 slice", args{[]uint64{1, 2, 3}}, []string{"1", "2", "3"}, false},
		{"uint64 array", args{[3]uint64{1, 2, 3}}, []string{"1", "2", "3"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToStrSlice(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToStrSlice() = %v, want %v", got, tt.want)
			}
			got, err := ToStrSliceE(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToStrSliceE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToStrSliceE() = %v, want %v", got, tt.want)
			}
		})
	}
}
