package conv

import (
	"reflect"
	"testing"
)

func TestToBoolSetE(t *testing.T) {
	type args struct {
		i any
	}
	tests := []struct {
		name    string
		args    args
		want    map[bool]struct{}
		wantErr bool
	}{
		{
			name: "to bool set",
			args: args{[]bool{true, false, true}},
			want: map[bool]struct{}{
				true:  {},
				false: {},
			},
			wantErr: false,
		}, {
			name:    "arg is nil",
			args:    args{nil},
			want:    nil,
			wantErr: false,
		},
		{
			name:    "arg is neither a slice nor an array",
			args:    args{"foo"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToSetE[bool](tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToSetGE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToSetGE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToIntSetE(t *testing.T) {
	type args struct {
		i any
	}
	tests := []struct {
		name    string
		args    args
		want    map[int]struct{}
		wantErr bool
	}{
		{
			name: "to int set",
			args: args{[]int{1, 2, 3}},
			want: map[int]struct{}{
				1: {},
				2: {},
				3: {},
			},
			wantErr: false,
		},
		{
			name:    "arg is nil",
			args:    args{nil},
			want:    nil,
			wantErr: false,
		},
		{
			name:    "arg is neither a slice nor an array",
			args:    args{"foo"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToSetE[int](tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToSetGE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToSetGE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToStrSetGE(t *testing.T) {
	type args struct {
		i any
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]struct{}
		wantErr bool
	}{
		{
			name: "to string set",
			args: args{[]string{"foo", "bar", "baz"}},
			want: map[string]struct{}{
				"foo": {},
				"bar": {},
				"baz": {},
			},
			wantErr: false,
		},
		{
			name:    "arg is nil",
			args:    args{nil},
			want:    nil,
			wantErr: false,
		},
		{
			name:    "arg is neither a slice nor an array",
			args:    args{"foo"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToSetE[string](tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToSetGE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToSetGE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToBoolSet(t *testing.T) {
	type args struct {
		i any
	}
	tests := []struct {
		name string
		args args
		want map[bool]struct{}
	}{
		{
			"bool slice",
			args{[]bool{true, true, false, false}},
			map[bool]struct{}{true: {}, false: {}},
		},
		{
			"int slice",
			args{[]int{1, 1, 1}},
			map[bool]struct{}{true: {}},
		},
		{
			"string slice",
			args{[]string{"false", "0", "FALSE"}},
			map[bool]struct{}{false: {}},
		},
		{
			"string slice failed",
			args{[]string{"foo", "1", "2", "2"}},
			nil,
		},
		{
			"arg isn't slice and array",
			args{"foo"},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToBoolSet(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToBoolSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToIntSet(t *testing.T) {
	type args struct {
		i any
	}
	tests := []struct {
		name string
		args args
		want map[int]struct{}
	}{
		{
			"int slice",
			args{[]int{0, 1, 2, 2}},
			map[int]struct{}{0: {}, 1: {}, 2: {}},
		},
		{
			"bool slice",
			args{[]bool{false, true, true}},
			map[int]struct{}{0: {}, 1: {}},
		},
		{
			"string slice",
			args{[]string{"0", "1", "2", "2"}},
			map[int]struct{}{0: {}, 1: {}, 2: {}},
		},
		{
			"string slice failed",
			args{[]string{"foo", "1", "2", "2"}},
			nil,
		},
		{
			"arg isn't slice and array",
			args{"foo"},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToIntSet(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToIntSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt8Set(t *testing.T) {
	type args struct {
		i any
	}
	tests := []struct {
		name string
		args args
		want map[int8]struct{}
	}{
		{
			"int8 slice to int8 set",
			args{[]int8{0, 1, 2, 2}},
			map[int8]struct{}{0: {}, 1: {}, 2: {}},
		},
		{
			"bool slice to int8 set",
			args{[]bool{false, true, true}},
			map[int8]struct{}{0: {}, 1: {}},
		},
		{
			"string slice to int8 set",
			args{[]string{"0", "1", "2", "2"}},
			map[int8]struct{}{0: {}, 1: {}, 2: {}},
		},
		{
			"string slice to int8 set failed",
			args{[]string{"foo", "1", "2", "2"}},
			nil,
		},
		{
			"arg isn't slice and array",
			args{"foo"},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt8Set(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToInt8Set() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt16Set(t *testing.T) {
	type args struct {
		i any
	}
	tests := []struct {
		name string
		args args
		want map[int16]struct{}
	}{
		{
			"int16 slice to int16 set",
			args{[]int16{0, 1, 2, 2}},
			map[int16]struct{}{0: {}, 1: {}, 2: {}},
		},
		{
			"bool slice to int16 set",
			args{[]bool{false, true, true}},
			map[int16]struct{}{0: {}, 1: {}},
		},
		{
			"string slice to int16 set",
			args{[]string{"0", "1", "2", "2"}},
			map[int16]struct{}{0: {}, 1: {}, 2: {}},
		},
		{
			"string slice to int16 set failed",
			args{[]string{"foo", "1", "2"}},
			nil,
		},
		{
			"arg isn't slice and array",
			args{"foo"},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt16Set(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToInt16Set() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt32Set(t *testing.T) {
	type args struct {
		i any
	}
	tests := []struct {
		name string
		args args
		want map[int32]struct{}
	}{
		{
			"int32 slice to int32 set",
			args{[]int32{0, 1, 2, 2}},
			map[int32]struct{}{0: {}, 1: {}, 2: {}},
		},
		{
			"bool slice to int32 set",
			args{[]bool{false, true, true}},
			map[int32]struct{}{0: {}, 1: {}},
		},
		{
			"string slice to int32 set",
			args{[]string{"0", "1", "2", "2"}},
			map[int32]struct{}{0: {}, 1: {}, 2: {}},
		},
		{
			"string slice to int32 set failed",
			args{[]string{"foo", "1", "2"}},
			nil,
		},
		{
			"arg isn't slice and array",
			args{"foo"},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt32Set(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToInt32Set() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt64Set(t *testing.T) {
	type args struct {
		i any
	}
	tests := []struct {
		name string
		args args
		want map[int64]struct{}
	}{
		{
			"int64 slice to int64 set",
			args{[]int64{0, 1, 2, 2}},
			map[int64]struct{}{0: {}, 1: {}, 2: {}},
		},
		{
			"bool slice to int64 set",
			args{[]bool{false, true, true}},
			map[int64]struct{}{0: {}, 1: {}},
		},
		{
			"string slice to int64 set",
			args{[]string{"0", "1", "2", "2"}},
			map[int64]struct{}{0: {}, 1: {}, 2: {}},
		},
		{
			"string slice to int64 set failed",
			args{[]string{"foo", "1", "2"}},
			nil,
		},
		{
			"arg isn't slice and array",
			args{"foo"},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt64Set(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToInt64Set() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUintSet(t *testing.T) {
	type args struct {
		i any
	}
	tests := []struct {
		name string
		args args
		want map[uint]struct{}
	}{
		{
			"uint slice to int set",
			args{[]uint{0, 1, 2, 2}},
			map[uint]struct{}{0: {}, 1: {}, 2: {}},
		},
		{
			"bool slice to uint set",
			args{[]bool{false, true, true}},
			map[uint]struct{}{0: {}, 1: {}},
		},
		{
			"string slice to uint set",
			args{[]string{"0", "1", "2", "2"}},
			map[uint]struct{}{0: {}, 1: {}, 2: {}},
		},
		{
			"string slice to uint set failed",
			args{[]string{"foo", "1", "2"}},
			nil,
		},
		{
			"arg isn't slice and array",
			args{"foo"},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUintSet(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToUintSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint8Set(t *testing.T) {
	type args struct {
		i any
	}
	tests := []struct {
		name string
		args args
		want map[uint8]struct{}
	}{
		{
			"uint8 slice to uint8 set",
			args{[]uint8{0, 1, 2, 2}},
			map[uint8]struct{}{0: {}, 1: {}, 2: {}},
		},
		{
			"bool slice to uint8 set",
			args{[]bool{false, true, true}},
			map[uint8]struct{}{0: {}, 1: {}},
		},
		{
			"string slice to uint8 set",
			args{[]string{"0", "1", "2", "2"}},
			map[uint8]struct{}{0: {}, 1: {}, 2: {}},
		},
		{
			"string slice to uint8 set failed",
			args{[]string{"foo", "1", "2"}},
			nil,
		},
		{
			"arg isn't slice and array",
			args{"foo"},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUint8Set(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToUint8Set() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint16Set(t *testing.T) {
	type args struct {
		i any
	}
	tests := []struct {
		name string
		args args
		want map[uint16]struct{}
	}{
		{
			"uint16 slice to uint16 set",
			args{[]uint16{0, 1, 2, 2}},
			map[uint16]struct{}{0: {}, 1: {}, 2: {}},
		},
		{
			"bool slice to uint16 set",
			args{[]bool{false, true, true}},
			map[uint16]struct{}{0: {}, 1: {}},
		},
		{
			"string slice to uint16 set",
			args{[]string{"0", "1", "2", "2"}},
			map[uint16]struct{}{0: {}, 1: {}, 2: {}},
		},
		{
			"string slice to uint16 set failed",
			args{[]string{"foo", "1", "2"}},
			nil,
		},
		{
			"arg isn't slice and array",
			args{"foo"},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUint16Set(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToUint16Set() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint32Set(t *testing.T) {
	type args struct {
		i any
	}
	tests := []struct {
		name string
		args args
		want map[uint32]struct{}
	}{
		{
			"uint32 slice to uint32 set",
			args{[]uint32{0, 1, 2, 2}},
			map[uint32]struct{}{0: {}, 1: {}, 2: {}},
		},
		{
			"bool slice to uint32 set",
			args{[]bool{false, true, true}},
			map[uint32]struct{}{0: {}, 1: {}},
		},
		{
			"string slice to uint32 set",
			args{[]string{"0", "1", "2", "2"}},
			map[uint32]struct{}{0: {}, 1: {}, 2: {}},
		},
		{
			"string slice to uint32 set failed",
			args{[]string{"foo", "1", "2"}},
			nil,
		},
		{
			"arg isn't slice and array",
			args{"foo"},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUint32Set(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToUint32Set() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint64Set(t *testing.T) {
	type args struct {
		i any
	}
	tests := []struct {
		name string
		args args
		want map[uint64]struct{}
	}{
		{
			name: "uint64 array to map set",
			args: args{[3]uint64{1, 2, 3}},
			want: map[uint64]struct{}{1: {}, 2: {}, 3: {}},
		},
		{
			name: "uint64 slice to map set",
			args: args{[]uint64{1, 2, 3}},
			want: map[uint64]struct{}{1: {}, 2: {}, 3: {}},
		},
		{
			"string slice",
			args{[]string{"0", "1", "2", "2"}},
			map[uint64]struct{}{0: {}, 1: {}, 2: {}},
		},
		{
			"string slice failed",
			args{[]string{"foo", "1", "2"}},
			nil,
		},
		{
			name: "zero length uint64 array to map set",
			args: args{[0]uint64{}},
			want: map[uint64]struct{}{},
		},
		{
			name: "zero length uint64 slice to map set",
			args: args{[]uint64{}},
			want: map[uint64]struct{}{},
		},
		{
			name: "nil slice to map set",
			args: args{[]uint64(nil)},
			want: map[uint64]struct{}{},
		},
		{
			"arg isn't slice and array",
			args{"foo"},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUint64Set(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToU64Set() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToFloat32Set(t *testing.T) {
	type args struct {
		i any
	}
	tests := []struct {
		name string
		args args
		want map[float32]struct{}
	}{
		{
			"float32 slice to float32 set",
			args{[]float32{0, 1, 2, 2}},
			map[float32]struct{}{0: {}, 1: {}, 2: {}},
		},
		{
			"bool slice to float32 set",
			args{[]bool{false, true, true}},
			map[float32]struct{}{0: {}, 1: {}},
		},
		{
			"string slice to float32 set",
			args{[]string{"0", "1.1", "2.2", "2.2"}},
			map[float32]struct{}{0: {}, 1.1: {}, 2.2: {}},
		},
		{
			"string slice to float32 set failed",
			args{[]string{"foo", "1", "2"}},
			nil,
		},
		{
			"arg isn't slice and array",
			args{"foo"},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToFloat32Set(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToFloat32Set() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToFloat64Set(t *testing.T) {
	type args struct {
		i any
	}
	tests := []struct {
		name string
		args args
		want map[float64]struct{}
	}{
		{
			"float64 slice to float64 set",
			args{[]float64{0, 1, 2, 2}},
			map[float64]struct{}{0: {}, 1: {}, 2: {}},
		},
		{
			"bool slice to float64 set",
			args{[]bool{false, true, true}},
			map[float64]struct{}{0: {}, 1: {}},
		},
		{
			"string slice to float64 set",
			args{[]string{"0", "1.1", "2.2", "2.2"}},
			map[float64]struct{}{0: {}, 1.1: {}, 2.2: {}},
		},
		{
			"string slice to float64 set failed",
			args{[]string{"foo", "1", "2"}},
			nil,
		},
		{
			"arg isn't slice and array",
			args{"foo"},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToFloat64Set(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToFloat64Set() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToStrSet(t *testing.T) {
	type args struct {
		i any
	}
	tests := []struct {
		name string
		args args
		want map[string]struct{}
	}{
		{
			"string slice to string set",
			args{[]string{"foo", "bar", "baz", "baz"}},
			map[string]struct{}{"foo": {}, "bar": {}, "baz": {}},
		},
		{
			"bool slice to string set",
			args{[]bool{false, true, true}},
			map[string]struct{}{"false": {}, "true": {}},
		},
		{
			"int slice to string set",
			args{[]int{0, 1, 2, 2}},
			map[string]struct{}{"0": {}, "1": {}, "2": {}},
		},
		{
			"float32 slice to string set",
			args{[]float32{0, 1.1, 2.2}},
			map[string]struct{}{"0": {}, "1.1": {}, "2.2": {}},
		},
		{
			"struct{} slice to string set",
			args{[]struct{}{{}}},
			nil,
		},
		{
			"arg isn't slice and array",
			args{"foo"},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToStrSet(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToStrSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToStrSetE(t *testing.T) {
	type args struct {
		i any
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]struct{}
		wantErr bool
	}{
		{
			name:    "[]string",
			args:    args{[]string{"CN", "HK", "AU"}},
			want:    map[string]struct{}{"CN": {}, "HK": {}, "AU": {}},
			wantErr: false,
		},
		{
			name:    "[3]int",
			args:    args{[3]int{86, 852, 61}},
			want:    map[string]struct{}{"86": {}, "852": {}, "61": {}},
			wantErr: false,
		},
		{
			name:    "string",
			args:    args{""},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToStrSetE(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToStrSetE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToStrSetE() got = %v, want %v", got, tt.want)
			}
		})
	}
}
