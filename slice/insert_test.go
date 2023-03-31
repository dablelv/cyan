package slice

import (
	"reflect"
	"testing"
)

func TestInsertIntSlice(t *testing.T) {
	type args struct {
		src   []int
		index int
		value int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "insert int to the first pos",
			args: args{[]int{1, 2}, 0, 0},
			want: []int{0, 1, 2},
		},
		{
			name: "insert int to the last pos",
			args: args{[]int{1, 2}, 2, 3},
			want: []int{1, 2, 3},
		},
		{
			name: "insert int failed becacuse the pos overflow",
			args: args{[]int{1, 2}, 3, 3},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertIntSlice(tt.args.src, tt.args.index, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertIntSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertInt8Slice(t *testing.T) {
	type args struct {
		src []int8
		i   int
		v   int8
	}
	tests := []struct {
		name string
		args args
		want []int8
	}{
		{
			name: "insert int8 to the first pos",
			args: args{[]int8{1, 2}, 0, 0},
			want: []int8{0, 1, 2},
		},
		{
			name: "insert int8 to the last pos",
			args: args{[]int8{1, 2}, 2, 3},
			want: []int8{1, 2, 3},
		},
		{
			name: "insert int8 failed becacuse the pos overflow",
			args: args{[]int8{1, 2}, 3, 3},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertInt8Slice(tt.args.src, tt.args.i, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertInt8Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertInt16Slice(t *testing.T) {
	type args struct {
		src []int16
		i   int
		v   int16
	}
	tests := []struct {
		name string
		args args
		want []int16
	}{
		{
			name: "insert int16 to the first pos",
			args: args{[]int16{1, 2}, 0, 0},
			want: []int16{0, 1, 2},
		},
		{
			name: "insert int16 to the last pos",
			args: args{[]int16{1, 2}, 2, 3},
			want: []int16{1, 2, 3},
		},
		{
			name: "insert int16 failed becacuse the pos overflow",
			args: args{[]int16{1, 2}, 3, 3},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertInt16Slice(tt.args.src, tt.args.i, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertInt16Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertInt32Slice(t *testing.T) {
	type args struct {
		src []int32
		i   int
		v   int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "insert int32 to the first pos",
			args: args{[]int32{1, 2}, 0, 0},
			want: []int32{0, 1, 2},
		},
		{
			name: "insert int32 to the last pos",
			args: args{[]int32{1, 2}, 2, 3},
			want: []int32{1, 2, 3},
		},
		{
			name: "insert int32 failed becacuse the pos overflow",
			args: args{[]int32{1, 2}, 3, 3},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertInt32Slice(tt.args.src, tt.args.i, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertInt32Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertInt64Slice(t *testing.T) {
	type args struct {
		src []int64
		i   int
		v   int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "insert int64 to the first pos",
			args: args{[]int64{1, 2}, 0, 0},
			want: []int64{0, 1, 2},
		},
		{
			name: "insert int64 to the last pos",
			args: args{[]int64{1, 2}, 2, 3},
			want: []int64{1, 2, 3},
		},
		{
			name: "insert int64 failed becacuse the pos overflow",
			args: args{[]int64{1, 2}, 3, 3},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertInt64Slice(tt.args.src, tt.args.i, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertInt64Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertUintSlice(t *testing.T) {
	type args struct {
		src []uint
		i   int
		v   uint
	}
	tests := []struct {
		name string
		args args
		want []uint
	}{
		{
			name: "insert uint to the first pos",
			args: args{[]uint{1, 2}, 0, 0},
			want: []uint{0, 1, 2},
		},
		{
			name: "insert uint to the last pos",
			args: args{[]uint{1, 2}, 2, 3},
			want: []uint{1, 2, 3},
		},
		{
			name: "insert uint failed becacuse the pos overflow",
			args: args{[]uint{1, 2}, 3, 3},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertUintSlice(tt.args.src, tt.args.i, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertUintSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertUint8Slice(t *testing.T) {
	type args struct {
		src []uint8
		i   int
		v   uint8
	}
	tests := []struct {
		name string
		args args
		want []uint8
	}{
		{
			name: "insert uint8 to the first pos",
			args: args{[]uint8{1, 2}, 0, 0},
			want: []uint8{0, 1, 2},
		},
		{
			name: "insert uint8 to the last pos",
			args: args{[]uint8{1, 2}, 2, 3},
			want: []uint8{1, 2, 3},
		},
		{
			name: "insert uint8 failed becacuse the pos overflow",
			args: args{[]uint8{1, 2}, 3, 3},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertUint8Slice(tt.args.src, tt.args.i, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertUint8Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertUint16Slice(t *testing.T) {
	type args struct {
		src []uint16
		i   int
		v   uint16
	}
	tests := []struct {
		name string
		args args
		want []uint16
	}{
		{
			name: "insert uint16 to the first pos",
			args: args{[]uint16{1, 2}, 0, 0},
			want: []uint16{0, 1, 2},
		},
		{
			name: "insert uint16 to the last pos",
			args: args{[]uint16{1, 2}, 2, 3},
			want: []uint16{1, 2, 3},
		},
		{
			name: "insert uint16 failed becacuse the pos overflow",
			args: args{[]uint16{1, 2}, 3, 3},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertUint16Slice(tt.args.src, tt.args.i, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertUint16Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertUint32Slice(t *testing.T) {
	type args struct {
		src []uint32
		i   int
		v   uint32
	}
	tests := []struct {
		name string
		args args
		want []uint32
	}{
		{
			name: "insert uint32 to the first pos",
			args: args{[]uint32{1, 2}, 0, 0},
			want: []uint32{0, 1, 2},
		},
		{
			name: "insert uint32 to the last pos",
			args: args{[]uint32{1, 2}, 2, 3},
			want: []uint32{1, 2, 3},
		},
		{
			name: "insert uint32 failed becacuse the pos overflow",
			args: args{[]uint32{1, 2}, 3, 3},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertUint32Slice(tt.args.src, tt.args.i, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertUint32Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertUint64Slice(t *testing.T) {
	type args struct {
		src []uint64
		i   int
		v   uint64
	}
	tests := []struct {
		name string
		args args
		want []uint64
	}{
		{
			name: "insert uint64 to the first pos",
			args: args{[]uint64{1, 2}, 0, 0},
			want: []uint64{0, 1, 2},
		},
		{
			name: "insert uint64 to the last pos",
			args: args{[]uint64{1, 2}, 2, 3},
			want: []uint64{1, 2, 3},
		},
		{
			name: "insert uint64 failed becacuse the pos overflow",
			args: args{[]uint64{1, 2}, 3, 3},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertUint64Slice(tt.args.src, tt.args.i, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertUint64Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertStrSlice(t *testing.T) {
	type args struct {
		src []string
		i   int
		v   string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "insert string to the first pos",
			args: args{[]string{"foo", "bar"}, 0, "baz"},
			want: []string{"baz", "foo", "bar"},
		},
		{
			name: "insert string to the last pos",
			args: args{[]string{"foo", "bar"}, 2, "baz"},
			want: []string{"foo", "bar", "baz"},
		},
		{
			name: "insert string failed becacuse the pos overflow",
			args: args{[]string{"foo", "bar"}, 3, "baz"},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertStrSlice(tt.args.src, tt.args.i, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertStrSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertSliceE(t *testing.T) {
	type args struct {
		slice any
		index int
		value any
	}
	tests := []struct {
		name    string
		args    args
		want    any
		wantErr bool
	}{
		{
			name:    "insert string to the first pos",
			args:    args{[]string{"bar", "baz"}, 0, "foo"},
			want:    []string{"foo", "bar", "baz"},
			wantErr: false,
		},
		{
			name:    "insert string to the last pos",
			args:    args{[]string{"foo", "bar"}, 2, "baz"},
			want:    []string{"foo", "bar", "baz"},
			wantErr: false,
		},
		{
			name:    "insert string failed because pos overflow",
			args:    args{[]string{"foo", "bar"}, 3, "baz"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := InsertSliceE(tt.args.slice, tt.args.index, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("InsertSliceE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertSliceE() got = %v, want %v", got, tt.want)
			}
		})
	}
}
