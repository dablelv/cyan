package conv

import (
	"reflect"
	"testing"
)

func TestToAnyBool(t *testing.T) {
	type args struct {
		a any
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"1 to bool",
			args{1},
			true,
		},
		{
			"0 to bool",
			args{0},
			false,
		},
		{
			"nil to bool",
			args{nil},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToAny[bool](tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToAny() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToAnyInt(t *testing.T) {
	type args struct {
		a any
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"true to int",
			args{true},
			1,
		},
		{
			"false to int",
			args{false},
			0,
		},
		{
			"string 1 to int",
			args{"1"},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToAny[int](tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToAny() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToAnyInt8(t *testing.T) {
	type args struct {
		a any
	}
	tests := []struct {
		name string
		args args
		want int8
	}{
		{
			"true",
			args{true},
			1,
		},
		{
			"false",
			args{false},
			0,
		},
		{
			"string 1",
			args{"1"},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToAny[int8](tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToAny() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToAnyInt16(t *testing.T) {
	type args struct {
		a any
	}
	tests := []struct {
		name string
		args args
		want int16
	}{
		{
			"true",
			args{true},
			1,
		},
		{
			"false",
			args{false},
			0,
		},
		{
			"string 1",
			args{"1"},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToAny[int16](tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToAny() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToAnyInt32(t *testing.T) {
	type args struct {
		a any
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			"true",
			args{true},
			1,
		},
		{
			"false",
			args{false},
			0,
		},
		{
			"string 1",
			args{"1"},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToAny[int32](tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToAny() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToAnyInt64(t *testing.T) {
	type args struct {
		a any
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			"true",
			args{true},
			1,
		},
		{
			"false",
			args{false},
			0,
		},
		{
			"string 1",
			args{"1"},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToAny[int64](tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToAny() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToAnyUint(t *testing.T) {
	type args struct {
		a any
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			"true",
			args{true},
			1,
		},
		{
			"false",
			args{false},
			0,
		},
		{
			"string 1",
			args{"1"},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToAny[uint](tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToAny() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToAnyUint8(t *testing.T) {
	type args struct {
		a any
	}
	tests := []struct {
		name string
		args args
		want uint8
	}{
		{
			"true",
			args{true},
			1,
		},
		{
			"false",
			args{false},
			0,
		},
		{
			"string 1",
			args{"1"},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToAny[uint8](tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToAny() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToAnyUint16(t *testing.T) {
	type args struct {
		a any
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{
			"true",
			args{true},
			1,
		},
		{
			"false",
			args{false},
			0,
		},
		{
			"string 1",
			args{"1"},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToAny[uint16](tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToAny() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToAnyUint32(t *testing.T) {
	type args struct {
		a any
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			"true",
			args{true},
			1,
		},
		{
			"false",
			args{false},
			0,
		},
		{
			"string 1",
			args{"1"},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToAny[uint32](tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToAny() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToAnyUint64(t *testing.T) {
	type args struct {
		a any
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			"true",
			args{true},
			1,
		},
		{
			"false",
			args{false},
			0,
		},
		{
			"string 1",
			args{"1"},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToAny[uint64](tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToAny() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToAnyFloat32(t *testing.T) {
	type args struct {
		a any
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			"true",
			args{true},
			1.0,
		},
		{
			"false",
			args{false},
			0.0,
		},
		{
			"string 1",
			args{"1"},
			1.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToAny[float32](tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToAny() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToAnyFloat64(t *testing.T) {
	type args struct {
		a any
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"true",
			args{true},
			1.0,
		},
		{
			"false",
			args{false},
			0.0,
		},
		{
			"string 1",
			args{"1"},
			1.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToAny[float64](tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToAny() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToAnyString(t *testing.T) {
	type args struct {
		a any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"int 8",
			args{8},
			"8",
		},
		{
			"float64 8.31",
			args{8.31},
			"8.31",
		},
		{
			"nil",
			args{nil},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToAny[string](tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToAny() = %v, want %v", got, tt.want)
			}
		})
	}
}
