package conv

import (
	"reflect"
	"testing"
)

func TestToAnyEString(t *testing.T) {
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
			name:    "from any type",
			args:    args{i: any("from any type")},
			want:    "from any type",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToAnyE[string](tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToAnyE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToAnyE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToAnyEBool(t *testing.T) {
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
		{
			name:    "any 1",
			args:    args{i: any(1)},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToAnyE[bool](tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToAnyE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToAnyE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToAnyEInt(t *testing.T) {
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
			args:    args{i: 1.1},
			want:    1,
			wantErr: false,
		},
		{
			name:    "nil",
			args:    args{i: nil},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToAnyE[int](tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToAnyE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToAnyE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToAnyEUint(t *testing.T) {
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
			args:    args{i: 1.1},
			want:    1,
			wantErr: false,
		},
		{
			name:    "nil",
			args:    args{i: nil},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToAnyE[uint](tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToAnyE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToAnyE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToAnyEFloat64(t *testing.T) {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToAnyE[float64](tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToAnyE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToAnyE() = %v, want %v", got, tt.want)
			}
		})
	}
}
