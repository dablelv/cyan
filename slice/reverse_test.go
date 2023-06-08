package slice

import (
	"reflect"
	"testing"

	"github.com/dablelv/go-huge-util/internal"
)

func TestReverseInt(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "int slice",
			args: args{[]int{1, 2, 3}},
			want: []int{3, 2, 1},
		},
		{
			name: "empty int slice",
			args: args{[]int{}},
			want: []int{},
		},
		{
			name: "nil int slice",
			args: args{nil},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseIntSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseStr(t *testing.T) {
	type args struct {
		src []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "reverse string slice",
			args: args{[]string{"foo", "bar", "baz"}},
			want: []string{"baz", "bar", "foo"},
		},
		{
			name: "reverse emplty string slice",
			args: args{[]string{}},
			want: []string{},
		},
		{
			name: "reverse nil string slice",
			args: args{nil},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseStrSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseIntRef(t *testing.T) {
	type args struct {
		src []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "int slice",
			args: args{[]int{1, 2, 3}},
			want: []int{3, 2, 1},
		},
		{
			name: "empty int slice",
			args: args{[]int{}},
			want: []int{},
		},
		{
			name: "nil int slice",
			args: args{nil},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseInt(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseInt8(t *testing.T) {
	assert := internal.NewAssert(t, "TestReverseInt8")

	got := ReverseInt8([]int8{1, 2, 3})
	assert.Equal([]int8{3, 2, 1}, got)

	got = ReverseInt8(nil)
	assert.IsNil(got)
}

func TestReverseInt16(t *testing.T) {
	assert := internal.NewAssert(t, "TestReverseInt16")

	got := ReverseInt16([]int16{1, 2, 3})
	assert.Equal([]int16{3, 2, 1}, got)

	got = ReverseInt16(nil)
	assert.IsNil(got)
}

func TestReverseInt32(t *testing.T) {
	assert := internal.NewAssert(t, "TestReverseInt32")

	got := ReverseInt32([]int32{1, 2, 3})
	assert.Equal([]int32{3, 2, 1}, got)

	got = ReverseInt32(nil)
	assert.IsNil(got)
}

func TestReverseInt64(t *testing.T) {
	assert := internal.NewAssert(t, "TestReverseInt8")

	got := ReverseInt64([]int64{1, 2, 3})
	assert.Equal([]int64{3, 2, 1}, got)

	got = ReverseInt64(nil)
	assert.IsNil(got)
}

func TestReverseUint(t *testing.T) {
	assert := internal.NewAssert(t, "TestReverseUint")

	got := ReverseUint([]uint{1, 2, 3})
	assert.Equal([]uint{3, 2, 1}, got)

	got = ReverseUint([]uint{})
	assert.Equal([]uint{}, got)

	got = ReverseUint(nil)
	assert.IsNil(got)
}

func TestReverseUint8(t *testing.T) {
	assert := internal.NewAssert(t, "TestReverseUint8")

	got := ReverseUint8([]uint8{1, 2, 3})
	assert.Equal([]uint8{3, 2, 1}, got)

	got = ReverseUint8(nil)
	assert.IsNil(got)
}

func TestReverseUint16(t *testing.T) {
	assert := internal.NewAssert(t, "TestReverseUint16")

	got := ReverseUint16([]uint16{1, 2, 3})
	assert.Equal([]uint16{3, 2, 1}, got)

	got = ReverseUint16(nil)
	assert.IsNil(got)
}

func TestReverseUint32(t *testing.T) {
	assert := internal.NewAssert(t, "TestReverseUint32")

	got := ReverseUint32([]uint32{1, 2, 3})
	assert.Equal([]uint32{3, 2, 1}, got)

	got = ReverseUint32(nil)
	assert.IsNil(got)
}

func TestReverseUint64(t *testing.T) {
	assert := internal.NewAssert(t, "TestReverseInt8")

	got := ReverseUint64([]uint64{1, 2, 3})
	assert.Equal([]uint64{3, 2, 1}, got)

	got = ReverseUint64(nil)
	assert.IsNil(got)
}

func TestReverseStrSlice(t *testing.T) {
	type args struct {
		src []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "reverse string slice",
			args: args{[]string{"foo", "bar", "baz"}},
			want: []string{"baz", "bar", "foo"},
		},
		{
			name: "reverse empty string slice",
			args: args{[]string{}},
			want: []string{},
		},
		{
			name: "reverse nil string slice",
			args: args{nil},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseStr(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseStrSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseSliceRefE(t *testing.T) {
	type args struct {
		slice any
	}
	tests := []struct {
		name    string
		args    args
		want    any
		wantErr bool
	}{
		{
			name:    "reverse string slice",
			args:    args{[]string{"foo", "bar", "baz"}},
			want:    []string{"baz", "bar", "foo"},
			wantErr: false,
		},
		{
			name:    "reverse empty string slice",
			args:    args{[]string{}},
			want:    []string{},
			wantErr: false,
		},
		{
			name:    "reverse uint64 slice",
			args:    args{[]uint64{1, 2, 3}},
			want:    []uint64{3, 2, 1},
			wantErr: false,
		},
		{
			name:    "reverse empty uint64 slice",
			args:    args{[]uint64{}},
			want:    []uint64{},
			wantErr: false,
		},
		{
			name:    "nil",
			args:    args{nil},
			want:    nil,
			wantErr: false,
		},
		{
			name:    "input isn't a slice",
			args:    args{"string"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReverseE(tt.args.slice)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReverseSliceE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseSliceE() got = %v, want %v", got, tt.want)
			}
		})
	}
}
