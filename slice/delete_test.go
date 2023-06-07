package slice

import (
	"reflect"
	"testing"

	"github.com/dablelv/go-huge-util/internal"
)

func TestDeleteSliceElems(t *testing.T) {
	type args struct {
		i    any
		elms []any
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			name: "delete success",
			args: args{
				i:    []string{"a", "b", "b", "c"},
				elms: []any{"b", "c"},
			},
			want: []string{"a"},
		},
		{
			name: "delete all",
			args: args{
				i:    []string{"a", "b", "b", "c"},
				elms: []any{"a", "b", "c"},
			},
			want: []string{},
		},
		{
			name: "delete failed because element type is ill",
			args: args{
				i:    []string{"1", "2", "2", "3"},
				elms: []any{2, 3},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteSliceElems(tt.args.i, tt.args.elms...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteSliceElems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteSlice(t *testing.T) {
	type args struct {
		slice   any
		indexes []int
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			name: "delete success",
			args: args{
				slice:   []string{"a", "b", "b", "c"},
				indexes: []int{1, 2},
			},
			want: []string{"a", "c"},
		},
		{
			name: "delete all",
			args: args{
				slice:   []string{"a", "b", "b", "c"},
				indexes: []int{0, 1, 2, 3, 4, 5},
			},
			want: []string{},
		},
		{
			name: "delete failed because the input isn't slice",
			args: args{
				slice:   "a",
				indexes: []int{0, 1, 2, 3, 4, 5},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteSlice(tt.args.slice, tt.args.indexes...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteStrSlice(t *testing.T) {
	type args struct {
		src     []string
		indexes []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "delete success",
			args: args{
				src:     []string{"a", "b", "b", "c"},
				indexes: []int{1, 2},
			},
			want: []string{"a", "c"},
		},
		{
			name: "delete all",
			args: args{
				src:     []string{"a", "b", "b", "c"},
				indexes: []int{0, 1, 2, 3, 4, 5},
			},
			want: []string{},
		},
		{
			name: "not delete",
			args: args{
				src:     []string{"a", "b", "b", "c"},
				indexes: []int{},
			},
			want: []string{"a", "b", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteStrSlice(tt.args.src, tt.args.indexes...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteStrSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteIntSlice(t *testing.T) {
	assert := internal.NewAssert(t, "TestDeleteIntSlice")

	s := DeleteIntSlice([]int{1, 2, 3}, 0)
	assert.Equal([]int{2, 3}, s)
}

func TestDeleteInt8Slice(t *testing.T) {
	assert := internal.NewAssert(t, "TestDeleteInt8Slice")

	s := DeleteInt8Slice([]int8{1, 2, 3}, 0)
	assert.Equal([]int8{2, 3}, s)
}

func TestDeleteInt16Slice(t *testing.T) {
	assert := internal.NewAssert(t, "TestDeleteInt16Slice")

	s := DeleteInt16Slice([]int16{1, 2, 3}, 0)
	assert.Equal([]int16{2, 3}, s)
}

func TestDeleteInt32Slice(t *testing.T) {
	assert := internal.NewAssert(t, "TestDeleteInt32Slice")

	s := DeleteInt32Slice([]int32{1, 2, 3}, 0)
	assert.Equal([]int32{2, 3}, s)
}

func TestDeleteInt64Slice(t *testing.T) {
	assert := internal.NewAssert(t, "TestDeleteInt64Slice")

	s := DeleteInt64Slice([]int64{1, 2, 3}, 0)
	assert.Equal([]int64{2, 3}, s)
}

func TestDeleteUintSlice(t *testing.T) {
	assert := internal.NewAssert(t, "TestDeleteUintSlice")

	s := DeleteUintSlice([]uint{1, 2, 3}, 0)
	assert.Equal([]uint{2, 3}, s)
}

func TestDeleteUint8Slice(t *testing.T) {
	assert := internal.NewAssert(t, "TestDeleteUint8Slice")

	s := DeleteUint8Slice([]uint8{1, 2, 3}, 0)
	assert.Equal([]uint8{2, 3}, s)
}

func TestDeleteUint16Slice(t *testing.T) {
	assert := internal.NewAssert(t, "TestDeleteUint16Slice")

	s := DeleteUint16Slice([]uint16{1, 2, 3}, 0)
	assert.Equal([]uint16{2, 3}, s)
}

func TestDeleteUint32Slice(t *testing.T) {
	assert := internal.NewAssert(t, "TestDeleteUint8Slice")

	s := DeleteUint32Slice([]uint32{1, 2, 3}, 0)
	assert.Equal([]uint32{2, 3}, s)
}

func TestDeleteUint64Slice(t *testing.T) {
	assert := internal.NewAssert(t, "TestDeleteUint64Slice")

	s := DeleteUint64Slice([]uint64{1, 2, 3}, 0)
	assert.Equal([]uint64{2, 3}, s)
}
