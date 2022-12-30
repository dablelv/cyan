package conv

import (
	"reflect"
	"testing"
	"time"
)

func TestToBoolE(t *testing.T) {
	type args struct {
		i any
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:    "true",
			args:    args{i: "true"},
			want:    true,
			wantErr: false,
		},
		{
			name:    "false",
			args:    args{i: "false"},
			want:    false,
			wantErr: false,
		},
		{
			name:    "True",
			args:    args{i: "True"},
			want:    true,
			wantErr: false,
		},
		{
			name:    "False",
			args:    args{i: "False"},
			want:    false,
			wantErr: false,
		},
		{
			name:    "1",
			args:    args{i: 1},
			want:    true,
			wantErr: false,
		},
		{
			name:    "0",
			args:    args{i: 0},
			want:    false,
			wantErr: false,
		},
		{
			name:    "nil",
			args:    args{i: nil},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToBoolE(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToBoolE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToBoolE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToIntE(t *testing.T) {
	type args struct {
		i any
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "true",
			args:    args{i: true},
			want:    1,
			wantErr: false,
		},
		{
			name:    "false",
			args:    args{i: false},
			want:    0,
			wantErr: false,
		},
		{
			name:    "1",
			args:    args{i: 1},
			want:    1,
			wantErr: false,
		},
		{
			name:    "0",
			args:    args{i: 0},
			want:    0,
			wantErr: false,
		},
		{
			name:    "1.1",
			args:    args{i: 1.1},
			want:    1,
			wantErr: false,
		},
		{
			name:    "1.00",
			args:    args{i: "1.00"},
			want:    1,
			wantErr: false,
		},
		{
			name:    "nil",
			args:    args{i: nil},
			want:    0,
			wantErr: false,
		},
		{
			name:    "1.01",
			args:    args{i: "1.01"},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToIntE(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToIntE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToIntE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt8E(t *testing.T) {
	type args struct {
		i any
	}
	tests := []struct {
		name    string
		args    args
		want    int8
		wantErr bool
	}{
		{
			name:    "true",
			args:    args{i: true},
			want:    1,
			wantErr: false,
		},
		{
			name:    "false",
			args:    args{i: false},
			want:    0,
			wantErr: false,
		},
		{
			name:    "1",
			args:    args{i: 1},
			want:    1,
			wantErr: false,
		},
		{
			name:    "0",
			args:    args{i: 0},
			want:    0,
			wantErr: false,
		},
		{
			name:    "1.1",
			args:    args{i: 1.1},
			want:    1,
			wantErr: false,
		},
		{
			name:    "1.00",
			args:    args{i: "1.00"},
			want:    1,
			wantErr: false,
		},
		{
			name:    "nil",
			args:    args{i: nil},
			want:    0,
			wantErr: false,
		},
		{
			name:    "1.01",
			args:    args{i: "1.01"},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToInt8E(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToInt8E() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToInt8E() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt16E(t *testing.T) {
	type args struct {
		i any
	}
	tests := []struct {
		name    string
		args    args
		want    int16
		wantErr bool
	}{
		{
			name:    "true",
			args:    args{i: true},
			want:    1,
			wantErr: false,
		},
		{
			name:    "false",
			args:    args{i: false},
			want:    0,
			wantErr: false,
		},
		{
			name:    "1",
			args:    args{i: 1},
			want:    1,
			wantErr: false,
		},
		{
			name:    "0",
			args:    args{i: 0},
			want:    0,
			wantErr: false,
		},
		{
			name:    "1.1",
			args:    args{i: 1.1},
			want:    1,
			wantErr: false,
		},
		{
			name:    "1.00",
			args:    args{i: "1.00"},
			want:    1,
			wantErr: false,
		},
		{
			name:    "nil",
			args:    args{i: nil},
			want:    0,
			wantErr: false,
		},
		{
			name:    "1.01",
			args:    args{i: "1.01"},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToInt16E(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToInt16E() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToInt16E() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt32E(t *testing.T) {
	type args struct {
		i any
	}
	tests := []struct {
		name    string
		args    args
		want    int32
		wantErr bool
	}{
		{
			name:    "true",
			args:    args{i: true},
			want:    1,
			wantErr: false,
		},
		{
			name:    "false",
			args:    args{i: false},
			want:    0,
			wantErr: false,
		},
		{
			name:    "1",
			args:    args{i: 1},
			want:    1,
			wantErr: false,
		},
		{
			name:    "0",
			args:    args{i: 0},
			want:    0,
			wantErr: false,
		},
		{
			name:    "1.1",
			args:    args{i: 1.1},
			want:    1,
			wantErr: false,
		},
		{
			name:    "1.00",
			args:    args{i: "1.00"},
			want:    1,
			wantErr: false,
		},
		{
			name:    "nil",
			args:    args{i: nil},
			want:    0,
			wantErr: false,
		},
		{
			name:    "1.01",
			args:    args{i: "1.01"},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToInt32E(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToInt32E() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToInt32E() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt64E(t *testing.T) {
	type args struct {
		i any
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:    "true",
			args:    args{i: true},
			want:    1,
			wantErr: false,
		},
		{
			name:    "false",
			args:    args{i: false},
			want:    0,
			wantErr: false,
		},
		{
			name:    "1",
			args:    args{i: 1},
			want:    1,
			wantErr: false,
		},
		{
			name:    "0",
			args:    args{i: 0},
			want:    0,
			wantErr: false,
		},
		{
			name:    "1.1",
			args:    args{i: 1.1},
			want:    1,
			wantErr: false,
		},
		{
			name:    "1.00",
			args:    args{i: "1.00"},
			want:    1,
			wantErr: false,
		},
		{
			name:    "nil",
			args:    args{i: nil},
			want:    0,
			wantErr: false,
		},
		{
			name:    "1.01",
			args:    args{i: "1.01"},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToInt64E(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToInt64E() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToInt64E() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUintE(t *testing.T) {
	type args struct {
		i any
	}
	tests := []struct {
		name    string
		args    args
		want    uint
		wantErr bool
	}{
		{
			name:    "true",
			args:    args{i: true},
			want:    1,
			wantErr: false,
		},
		{
			name:    "false",
			args:    args{i: false},
			want:    0,
			wantErr: false,
		},
		{
			name:    "1",
			args:    args{i: 1},
			want:    1,
			wantErr: false,
		},
		{
			name:    "0",
			args:    args{i: 0},
			want:    0,
			wantErr: false,
		},
		{
			name:    "1.1",
			args:    args{i: 1.1},
			want:    1,
			wantErr: false,
		},
		{
			name:    "1.00",
			args:    args{i: "1.00"},
			want:    1,
			wantErr: false,
		},
		{
			name:    "nil",
			args:    args{i: nil},
			want:    0,
			wantErr: false,
		},
		{
			name:    "1.01",
			args:    args{i: "1.01"},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToUintE(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToUintE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToUintE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint8E(t *testing.T) {
	type args struct {
		i any
	}
	tests := []struct {
		name    string
		args    args
		want    uint8
		wantErr bool
	}{
		{
			name:    "true",
			args:    args{i: true},
			want:    1,
			wantErr: false,
		},
		{
			name:    "false",
			args:    args{i: false},
			want:    0,
			wantErr: false,
		},
		{
			name:    "1",
			args:    args{i: 1},
			want:    1,
			wantErr: false,
		},
		{
			name:    "0",
			args:    args{i: 0},
			want:    0,
			wantErr: false,
		},
		{
			name:    "1.1",
			args:    args{i: 1.1},
			want:    1,
			wantErr: false,
		},
		{
			name:    "1.00",
			args:    args{i: "1.00"},
			want:    1,
			wantErr: false,
		},
		{
			name:    "nil",
			args:    args{i: nil},
			want:    0,
			wantErr: false,
		},
		{
			name:    "1.01",
			args:    args{i: "1.01"},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToUint8E(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToUint8E() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToUint8E() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint16E(t *testing.T) {
	type args struct {
		i any
	}
	tests := []struct {
		name    string
		args    args
		want    uint16
		wantErr bool
	}{
		{
			name:    "true",
			args:    args{i: true},
			want:    1,
			wantErr: false,
		},
		{
			name:    "false",
			args:    args{i: false},
			want:    0,
			wantErr: false,
		},
		{
			name:    "1",
			args:    args{i: 1},
			want:    1,
			wantErr: false,
		},
		{
			name:    "0",
			args:    args{i: 0},
			want:    0,
			wantErr: false,
		},
		{
			name:    "1.1",
			args:    args{i: 1.1},
			want:    1,
			wantErr: false,
		},
		{
			name:    "1.00",
			args:    args{i: "1.00"},
			want:    1,
			wantErr: false,
		},
		{
			name:    "nil",
			args:    args{i: nil},
			want:    0,
			wantErr: false,
		},
		{
			name:    "1.01",
			args:    args{i: "1.01"},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToUint16E(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToUint16E() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToUint16E() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint32E(t *testing.T) {
	type args struct {
		i any
	}
	tests := []struct {
		name    string
		args    args
		want    uint32
		wantErr bool
	}{
		{
			name:    "true",
			args:    args{i: true},
			want:    1,
			wantErr: false,
		},
		{
			name:    "false",
			args:    args{i: false},
			want:    0,
			wantErr: false,
		},
		{
			name:    "1",
			args:    args{i: 1},
			want:    1,
			wantErr: false,
		},
		{
			name:    "0",
			args:    args{i: 0},
			want:    0,
			wantErr: false,
		},
		{
			name:    "1.1",
			args:    args{i: 1.1},
			want:    1,
			wantErr: false,
		},
		{
			name:    "1.00",
			args:    args{i: "1.00"},
			want:    1,
			wantErr: false,
		},
		{
			name:    "nil",
			args:    args{i: nil},
			want:    0,
			wantErr: false,
		},
		{
			name:    "1.01",
			args:    args{i: "1.01"},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToUint32E(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToUint32E() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToUint32E() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint64E(t *testing.T) {
	type args struct {
		i any
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		{
			name:    "true",
			args:    args{i: true},
			want:    1,
			wantErr: false,
		},
		{
			name:    "false",
			args:    args{i: false},
			want:    0,
			wantErr: false,
		},
		{
			name:    "1",
			args:    args{i: 1},
			want:    1,
			wantErr: false,
		},
		{
			name:    "0",
			args:    args{i: 0},
			want:    0,
			wantErr: false,
		},
		{
			name:    "1.1",
			args:    args{i: 1.1},
			want:    1,
			wantErr: false,
		},
		{
			name:    "1.00",
			args:    args{i: "1.00"},
			want:    1,
			wantErr: false,
		},
		{
			name:    "nil",
			args:    args{i: nil},
			want:    0,
			wantErr: false,
		},
		{
			name:    "1.01",
			args:    args{i: "1.01"},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToUint64E(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToUint64E() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToUint64E() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToFloat32E(t *testing.T) {
	type args struct {
		i any
	}
	tests := []struct {
		name    string
		args    args
		want    float32
		wantErr bool
	}{
		{
			name:    "true",
			args:    args{i: true},
			want:    1.0,
			wantErr: false,
		},
		{
			name:    "false",
			args:    args{i: false},
			want:    0.0,
			wantErr: false,
		},
		{
			name:    "1",
			args:    args{i: 1},
			want:    1.0,
			wantErr: false,
		},
		{
			name:    "0",
			args:    args{i: 0},
			want:    0.0,
			wantErr: false,
		},
		{
			name:    "1.1",
			args:    args{i: 1.1},
			want:    1.1,
			wantErr: false,
		},
		{
			name:    "string 1.00",
			args:    args{i: "1.00"},
			want:    1.0,
			wantErr: false,
		},
		{
			name:    "nil",
			args:    args{i: nil},
			want:    0.0,
			wantErr: false,
		},
		{
			name:    "error",
			args:    args{i: "error"},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToFloat32E(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToFloat32E() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToFloat32E() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToFloat64E(t *testing.T) {
	type args struct {
		i any
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name:    "true",
			args:    args{i: true},
			want:    1.0,
			wantErr: false,
		},
		{
			name:    "false",
			args:    args{i: false},
			want:    0.0,
			wantErr: false,
		},
		{
			name:    "1",
			args:    args{i: 1},
			want:    1.0,
			wantErr: false,
		},
		{
			name:    "0",
			args:    args{i: 0},
			want:    0.0,
			wantErr: false,
		},
		{
			name:    "1.1",
			args:    args{i: 1.1},
			want:    1.1,
			wantErr: false,
		},
		{
			name:    "string 1.00",
			args:    args{i: "1.00"},
			want:    1.0,
			wantErr: false,
		},
		{
			name:    "nil",
			args:    args{i: nil},
			want:    0,
			wantErr: false,
		},
		{
			name:    "error",
			args:    args{i: "error"},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToFloat64E(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToFloat64E() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToFloat64E() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToStringE(t *testing.T) {
	type args struct {
		i any
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "foo",
			args:    args{i: "foo"},
			want:    "foo",
			wantErr: false,
		},
		{
			name:    "8",
			args:    args{i: 8},
			want:    "8",
			wantErr: false,
		},
		{
			name:    "8.31",
			args:    args{i: 8.31},
			want:    "8.31",
			wantErr: false,
		},
		{
			name:    "one time",
			args:    args{i: []byte("one time")},
			want:    "one time",
			wantErr: false,
		},
		{
			name:    "nil",
			args:    args{i: nil},
			want:    "",
			wantErr: false,
		},
		{
			name:    "error",
			args:    args{i: []int{}},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToStringE(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToStringE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToStringE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToDurationE(t *testing.T) {
	type args struct {
		i any
	}
	tests := []struct {
		name    string
		args    args
		want    time.Duration
		wantErr bool
	}{
		{
			name:    "duration",
			args:    args{i: time.Second},
			want:    time.Second,
			wantErr: false,
		},
		{
			name:    "int",
			args:    args{i: 8},
			want:    time.Duration(8),
			wantErr: false,
		},
		{
			name:    "uint",
			args:    args{i: 8},
			want:    time.Duration(8),
			wantErr: false,
		},
		{
			name:    "float",
			args:    args{i: 8.31},
			want:    time.Duration(8),
			wantErr: false,
		},
		{
			name:    "string",
			args:    args{i: "string"},
			want:    time.Duration(0),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotD, err := ToDurationE(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToDurationE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotD, tt.want) {
				t.Errorf("ToDurationE() = %v, want %v", gotD, tt.want)
			}
		})
	}
}
