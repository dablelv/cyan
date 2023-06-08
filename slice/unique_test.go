package slice

import (
	"reflect"
	"testing"

	"github.com/dablelv/go-huge-util/internal"
)

func TestUnique(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "unique int slice",
			args: args{[]int{1, 2, 3, 2}},
			want: []int{1, 2, 3},
		},
		{
			name: "unique empty int slice",
			args: args{[]int{}},
			want: []int{},
		},
		{
			name: "unique nil int slice",
			args: args{nil},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Unique(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unique() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueInt(t *testing.T) {
	assert := internal.NewAssert(t, "TestUniqueInt")

	got := UniqueInt([]int{1, 2, 2, 3})
	assert.Equal([]int{1, 2, 3}, got)
}

func TestUniqueInt8(t *testing.T) {
	assert := internal.NewAssert(t, "TestUniqueInt8")

	got := UniqueInt8([]int8{1, 2, 2, 3})
	assert.Equal([]int8{1, 2, 3}, got)
}

func TestUniqueInt16(t *testing.T) {
	assert := internal.NewAssert(t, "TestUniqueInt16")

	got := UniqueInt16([]int16{1, 2, 2, 3})
	assert.Equal([]int16{1, 2, 3}, got)
}

func TestUniqueInt32(t *testing.T) {
	assert := internal.NewAssert(t, "TestUniqueInt32")

	got := UniqueInt32([]int32{1, 2, 2, 3})
	assert.Equal([]int32{1, 2, 3}, got)
}

func TestUniqueInt64(t *testing.T) {
	assert := internal.NewAssert(t, "TestUniqueInt64")

	got := UniqueInt64([]int64{1, 2, 2, 3})
	assert.Equal([]int64{1, 2, 3}, got)
}

func TestUniqueUint(t *testing.T) {
	assert := internal.NewAssert(t, "TestUniqueUint")

	got := UniqueUint([]uint{1, 2, 2, 3})
	assert.Equal([]uint{1, 2, 3}, got)
}

func TestUniqueUint8(t *testing.T) {
	assert := internal.NewAssert(t, "TestUniqueUint8")

	got := UniqueUint8([]uint8{1, 2, 2, 3})
	assert.Equal([]uint8{1, 2, 3}, got)
}

func TestUniqueUint16(t *testing.T) {
	assert := internal.NewAssert(t, "TestUniqueUint16")

	got := UniqueUint16([]uint16{1, 2, 2, 3})
	assert.Equal([]uint16{1, 2, 3}, got)
}

func TestUniqueUint32(t *testing.T) {
	assert := internal.NewAssert(t, "TestUniqueUint32")

	got := UniqueUint32([]uint32{1, 2, 2, 3})
	assert.Equal([]uint32{1, 2, 3}, got)
}

func TestUniqueUint64(t *testing.T) {
	assert := internal.NewAssert(t, "TestUniqueUint64")

	got := UniqueUint64([]uint64{1, 2, 2, 3})
	assert.Equal([]uint64{1, 2, 3}, got)
}

func TestUniqueFloat32(t *testing.T) {
	assert := internal.NewAssert(t, "TestUniqueFloat32")

	got := UniqueFloat32([]float32{1, 2, 2, 3})
	assert.Equal([]float32{1, 2, 3}, got)
}

func TestUniqueFloat64(t *testing.T) {
	assert := internal.NewAssert(t, "TestUniqueFloat64")

	got := UniqueFloat64([]float64{1, 2, 2, 3})
	assert.Equal([]float64{1, 2, 3}, got)
}

func TestUniqueStr(t *testing.T) {
	type args struct {
		src []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "unique string slice",
			args: args{[]string{"foo", "bar", "bar", "baz"}},
			want: []string{"foo", "bar", "baz"},
		},
		{
			name: "unique zero length string slice",
			args: args{[]string{}},
			want: []string{},
		},
		{
			name: "unique nil string slice",
			args: args{nil},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueStr(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueStrSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueE(t *testing.T) {
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
			name:    "unique string slice",
			args:    args{[]string{"foo", "bar", "bar", "baz"}},
			want:    []string{"foo", "bar", "baz"},
			wantErr: false,
		},
		{
			name:    "unique zero length string slice",
			args:    args{[]string{}},
			want:    []string{},
			wantErr: false,
		},
		{
			name:    "unique uint64 slice",
			args:    args{[]uint64{1, 2, 2, 3}},
			want:    []uint64{1, 2, 3},
			wantErr: false,
		},
		{
			name:    "unique zero length uint64 slice",
			args:    args{[]uint64{}},
			want:    []uint64{},
			wantErr: false,
		},
		{
			name:    "input isn't a slice",
			args:    args{nil},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UniqueE(tt.args.slice)
			if (err != nil) != tt.wantErr {
				t.Errorf("UniqueSliceE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueSliceE() got = %v, want %v", got, tt.want)
			}
		})
	}
}
