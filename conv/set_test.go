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
			"bool slice",
			args{[]bool{true, true, false, false}},
			map[bool]struct{}{true: {}, false: {}},
			false,
		},
		{
			"int slice",
			args{[]int{1, 1, 1}},
			map[bool]struct{}{true: {}},
			false,
		},
		{
			"string slice",
			args{[]string{"false", "0", "FALSE"}},
			map[bool]struct{}{false: {}},
			false,
		},
		{
			"string slice failed",
			args{[]string{"foo", "1", "2", "2"}},
			nil,
			true,
		},
		{
			"arg isn't slice and array",
			args{"foo"},
			nil,
			true,
		},
		{
			name:    "arg is nil",
			args:    args{nil},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToBoolSet(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToBoolSet() = %v, want %v", got, tt.want)
			}
			got, err := ToBoolSetE(tt.args.i)
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
			"int slice",
			args{[]int{0, 1, 2, 2}},
			map[int]struct{}{0: {}, 1: {}, 2: {}},
			false,
		},
		{
			"bool slice",
			args{[]bool{false, true, true}},
			map[int]struct{}{0: {}, 1: {}},
			false,
		},
		{
			"string slice",
			args{[]string{"0", "1", "2", "2"}},
			map[int]struct{}{0: {}, 1: {}, 2: {}},
			false,
		},
		{
			"string slice failed",
			args{[]string{"foo", "1", "2", "2"}},
			nil,
			true,
		},
		{
			"arg isn't slice and array",
			args{"foo"},
			nil,
			true,
		},
		{
			"arg is nil",
			args{nil},
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToIntSet(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToIntSet() = %v, want %v", got, tt.want)
			}
			got, err := ToIntSetE(tt.args.i)
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

func TestToInt8SetE(t *testing.T) {
	type args struct {
		a any
	}
	tests := []struct {
		name    string
		args    args
		want    map[int8]struct{}
		wantErr bool
	}{
		{
			"int8 slice to int8 set",
			args{[]int8{0, 1, 2, 2}},
			map[int8]struct{}{0: {}, 1: {}, 2: {}},
			false,
		},
		{
			"bool slice to int8 set",
			args{[]bool{false, true, true}},
			map[int8]struct{}{0: {}, 1: {}},
			false,
		},
		{
			"string slice to int8 set",
			args{[]string{"0", "1", "2", "2"}},
			map[int8]struct{}{0: {}, 1: {}, 2: {}},
			false,
		},
		{
			"string slice to int8 set failed",
			args{[]string{"foo", "1", "2", "2"}},
			nil,
			true,
		},
		{
			"arg isn't slice and array",
			args{"foo"},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt8Set(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToInt8Set() = %v, want %v", got, tt.want)
			}
			got, err := ToInt8SetE(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToInt8SetE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToInt8SetE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt16SetE(t *testing.T) {
	type args struct {
		a any
	}
	tests := []struct {
		name    string
		args    args
		want    map[int16]struct{}
		wantErr bool
	}{
		{
			"int16 slice to int16 set",
			args{[]int16{0, 1, 2, 2}},
			map[int16]struct{}{0: {}, 1: {}, 2: {}},
			false,
		},
		{
			"bool slice to int16 set",
			args{[]bool{false, true, true}},
			map[int16]struct{}{0: {}, 1: {}},
			false,
		},
		{
			"string slice to int16 set",
			args{[]string{"0", "1", "2", "2"}},
			map[int16]struct{}{0: {}, 1: {}, 2: {}},
			false,
		},
		{
			"string slice to int16 set failed",
			args{[]string{"foo", "1", "2"}},
			nil,
			true,
		},
		{
			"arg isn't slice and array",
			args{"foo"},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt16Set(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToInt16Set() = %v, want %v", got, tt.want)
			}
			got, err := ToInt16SetE(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToInt16SetE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToInt16SetE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt32SetE(t *testing.T) {
	type args struct {
		a any
	}
	tests := []struct {
		name    string
		args    args
		want    map[int32]struct{}
		wantErr bool
	}{
		{
			"int32 slice to int32 set",
			args{[]int32{0, 1, 2, 2}},
			map[int32]struct{}{0: {}, 1: {}, 2: {}},
			false,
		},
		{
			"bool slice to int32 set",
			args{[]bool{false, true, true}},
			map[int32]struct{}{0: {}, 1: {}},
			false,
		},
		{
			"string slice to int32 set",
			args{[]string{"0", "1", "2", "2"}},
			map[int32]struct{}{0: {}, 1: {}, 2: {}},
			false,
		},
		{
			"string slice to int32 set failed",
			args{[]string{"foo", "1", "2"}},
			nil,
			true,
		},
		{
			"arg isn't slice and array",
			args{"foo"},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt32Set(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToInt32Set() = %v, want %v", got, tt.want)
			}
			got, err := ToInt32SetE(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToInt32SetE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToInt32SetE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt64SetE(t *testing.T) {
	type args struct {
		a any
	}
	tests := []struct {
		name    string
		args    args
		want    map[int64]struct{}
		wantErr bool
	}{
		{
			"int64 slice to int64 set",
			args{[]int64{0, 1, 2, 2}},
			map[int64]struct{}{0: {}, 1: {}, 2: {}},
			false,
		},
		{
			"bool slice to int64 set",
			args{[]bool{false, true, true}},
			map[int64]struct{}{0: {}, 1: {}},
			false,
		},
		{
			"string slice to int64 set",
			args{[]string{"0", "1", "2", "2"}},
			map[int64]struct{}{0: {}, 1: {}, 2: {}},
			false,
		},
		{
			"string slice to int64 set failed",
			args{[]string{"foo", "1", "2"}},
			nil,
			true,
		},
		{
			"arg isn't slice and array",
			args{"foo"},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt64Set(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToInt64Set() = %v, want %v", got, tt.want)
			}
			got, err := ToInt64SetE(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToInt64SetE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToInt64SetE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUintSetE(t *testing.T) {
	type args struct {
		a any
	}
	tests := []struct {
		name    string
		args    args
		want    map[uint]struct{}
		wantErr bool
	}{
		{
			"uint slice to int set",
			args{[]uint{0, 1, 2, 2}},
			map[uint]struct{}{0: {}, 1: {}, 2: {}},
			false,
		},
		{
			"bool slice to uint set",
			args{[]bool{false, true, true}},
			map[uint]struct{}{0: {}, 1: {}},
			false,
		},
		{
			"string slice to uint set",
			args{[]string{"0", "1", "2", "2"}},
			map[uint]struct{}{0: {}, 1: {}, 2: {}},
			false,
		},
		{
			"string slice to uint set failed",
			args{[]string{"foo", "1", "2"}},
			nil,
			true,
		},
		{
			"arg isn't slice and array",
			args{"foo"},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToUintSet(tt.args.a)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToUintSet() = %v, want %v", got, tt.want)
			}
			got, err := ToUintSetE(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToUintSetE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToUintSetE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint8SetE(t *testing.T) {
	type args struct {
		a any
	}
	tests := []struct {
		name    string
		args    args
		want    map[uint8]struct{}
		wantErr bool
	}{
		{
			"uint8 slice to uint8 set",
			args{[]uint8{0, 1, 2, 2}},
			map[uint8]struct{}{0: {}, 1: {}, 2: {}},
			false,
		},
		{
			"bool slice to uint8 set",
			args{[]bool{false, true, true}},
			map[uint8]struct{}{0: {}, 1: {}},
			false,
		},
		{
			"string slice to uint8 set",
			args{[]string{"0", "1", "2", "2"}},
			map[uint8]struct{}{0: {}, 1: {}, 2: {}},
			false,
		},
		{
			"string slice to uint8 set failed",
			args{[]string{"foo", "1", "2"}},
			nil,
			true,
		},
		{
			"arg isn't slice and array",
			args{"foo"},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUint8Set(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToUint8Set() = %v, want %v", got, tt.want)
			}
			got, err := ToUint8SetE(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToUint8SetE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToUint8SetE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint16SetE(t *testing.T) {
	type args struct {
		a any
	}
	tests := []struct {
		name    string
		args    args
		want    map[uint16]struct{}
		wantErr bool
	}{
		{
			"uint16 slice to uint16 set",
			args{[]uint16{0, 1, 2, 2}},
			map[uint16]struct{}{0: {}, 1: {}, 2: {}},
			false,
		},
		{
			"bool slice to uint16 set",
			args{[]bool{false, true, true}},
			map[uint16]struct{}{0: {}, 1: {}},
			false,
		},
		{
			"string slice to uint16 set",
			args{[]string{"0", "1", "2", "2"}},
			map[uint16]struct{}{0: {}, 1: {}, 2: {}},
			false,
		},
		{
			"string slice to uint16 set failed",
			args{[]string{"foo", "1", "2"}},
			nil,
			true,
		},
		{
			"arg isn't slice and array",
			args{"foo"},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUint16Set(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToUint16Set() = %v, want %v", got, tt.want)
			}
			got, err := ToUint16SetE(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToUint16SetE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToUint16SetE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint32SetE(t *testing.T) {
	type args struct {
		a any
	}
	tests := []struct {
		name    string
		args    args
		want    map[uint32]struct{}
		wantErr bool
	}{
		{
			"uint32 slice to uint32 set",
			args{[]uint32{0, 1, 2, 2}},
			map[uint32]struct{}{0: {}, 1: {}, 2: {}},
			false,
		},
		{
			"bool slice to uint32 set",
			args{[]bool{false, true, true}},
			map[uint32]struct{}{0: {}, 1: {}},
			false,
		},
		{
			"string slice to uint32 set",
			args{[]string{"0", "1", "2", "2"}},
			map[uint32]struct{}{0: {}, 1: {}, 2: {}},
			false,
		},
		{
			"string slice to uint32 set failed",
			args{[]string{"foo", "1", "2"}},
			nil,
			true,
		},
		{
			"arg isn't slice and array",
			args{"foo"},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUint32Set(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToUint32Set() = %v, want %v", got, tt.want)
			}
			got, err := ToUint32SetE(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToUint32SetE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToUint32SetE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint64SetE(t *testing.T) {
	type args struct {
		a any
	}
	tests := []struct {
		name    string
		args    args
		want    map[uint64]struct{}
		wantErr bool
	}{
		{
			"uint64 array to map set",
			args{[3]uint64{1, 2, 3}},
			map[uint64]struct{}{1: {}, 2: {}, 3: {}},
			false,
		},
		{
			"uint64 slice to map set",
			args{[]uint64{1, 2, 3}},
			map[uint64]struct{}{1: {}, 2: {}, 3: {}},
			false,
		},
		{
			"string slice",
			args{[]string{"0", "1", "2", "2"}},
			map[uint64]struct{}{0: {}, 1: {}, 2: {}},
			false,
		},
		{
			"string slice failed",
			args{[]string{"foo", "1", "2"}},
			nil,
			true,
		},
		{
			"zero length uint64 array",
			args{[0]uint64{}},
			map[uint64]struct{}{},
			false,
		},
		{
			"zero length uint64 slice",
			args{[]uint64{}},
			map[uint64]struct{}{},
			false,
		},
		{
			"nil slice to map set",
			args{[]uint64(nil)},
			nil,
			false,
		},
		{
			"arg isn't slice and array",
			args{"foo"},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUint64Set(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToU64Set() = %v, want %v", got, tt.want)
			}
			got, err := ToUint64SetE(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToUint64SetE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToUint64SetE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToFloat32SetE(t *testing.T) {
	type args struct {
		a any
	}
	tests := []struct {
		name    string
		args    args
		want    map[float32]struct{}
		wantErr bool
	}{
		{
			"float32 slice to float32 set",
			args{[]float32{0, 1, 2, 2}},
			map[float32]struct{}{0: {}, 1: {}, 2: {}},
			false,
		},
		{
			"bool slice to float32 set",
			args{[]bool{false, true, true}},
			map[float32]struct{}{0: {}, 1: {}},
			false,
		},
		{
			"string slice to float32 set",
			args{[]string{"0", "1.1", "2.2", "2.2"}},
			map[float32]struct{}{0: {}, 1.1: {}, 2.2: {}},
			false,
		},
		{
			"string slice to float32 set failed",
			args{[]string{"foo", "1", "2"}},
			nil,
			true,
		},
		{
			"arg isn't slice and array",
			args{"foo"},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToFloat32Set(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToFloat32Set() = %v, want %v", got, tt.want)
			}
			got, err := ToFloat32SetE(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToFloat32SetE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToFloat32SetE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToFloat64SetE(t *testing.T) {
	type args struct {
		a any
	}
	tests := []struct {
		name    string
		args    args
		want    map[float64]struct{}
		wantErr bool
	}{
		{
			"float64 slice to float64 set",
			args{[]float64{0, 1, 2, 2}},
			map[float64]struct{}{0: {}, 1: {}, 2: {}},
			false,
		},
		{
			"bool slice to float64 set",
			args{[]bool{false, true, true}},
			map[float64]struct{}{0: {}, 1: {}},
			false,
		},
		{
			"string slice to float64 set",
			args{[]string{"0", "1.1", "2.2", "2.2"}},
			map[float64]struct{}{0: {}, 1.1: {}, 2.2: {}},
			false,
		},
		{
			"string slice to float64 set failed",
			args{[]string{"foo", "1", "2"}},
			nil,
			true,
		},
		{
			"arg isn't slice and array",
			args{"foo"},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToFloat64Set(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToFloat64Set() = %v, want %v", got, tt.want)
			}
			got, err := ToFloat64SetE(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToFloat64SetE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToFloat64SetE() = %v, want %v", got, tt.want)
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
			"[]string",
			args{[]string{"CN", "HK", "AU"}},
			map[string]struct{}{"CN": {}, "HK": {}, "AU": {}},
			false,
		},
		{
			"bool slice",
			args{[]bool{false, true, true}},
			map[string]struct{}{"false": {}, "true": {}},
			false,
		},
		{
			"int slice",
			args{[]int{0, 1, 2, 2}},
			map[string]struct{}{"0": {}, "1": {}, "2": {}},
			false,
		},
		{
			name:    "[3]int",
			args:    args{[3]int{86, 852, 61}},
			want:    map[string]struct{}{"86": {}, "852": {}, "61": {}},
			wantErr: false,
		},
		{
			"float32 slice",
			args{[]float32{0, 1.1, 2.2}},
			map[string]struct{}{"0": {}, "1.1": {}, "2.2": {}},
			false,
		},
		{
			"struct{} slice",
			args{[]struct{}{{}}},
			nil,
			true,
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
			if got := ToStrSet(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToStrSet() = %v, want %v", got, tt.want)
			}
			got, err := ToStrSetE(tt.args.i)
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
