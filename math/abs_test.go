package math

import "testing"

func TestAbsInt8(t *testing.T) {
	type args struct {
		n int8
	}
	tests := []struct {
		name string
		args args
		want int8
	}{
		{
			name: "abs(8)",
			args: args{
				n: 8,
			},
			want: 8,
		},
		{
			name: "abs(-8)",
			args: args{
				n: -8,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AbsInt8(tt.args.n); got != tt.want {
				t.Errorf("AbsInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbsInt16(t *testing.T) {
	type args struct {
		n int16
	}
	tests := []struct {
		name string
		args args
		want int16
	}{
		{
			name: "abs(8)",
			args: args{
				n: 8,
			},
			want: 8,
		},
		{
			name: "abs(-8)",
			args: args{
				n: -8,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AbsInt16(tt.args.n); got != tt.want {
				t.Errorf("AbsInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbsInt32(t *testing.T) {
	type args struct {
		n int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "abs(8)",
			args: args{
				n: 8,
			},
			want: 8,
		},
		{
			name: "abs(-8)",
			args: args{
				n: -8,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AbsInt32(tt.args.n); got != tt.want {
				t.Errorf("AbsInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbsInt64(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "abs(8)",
			args: args{
				n: 8,
			},
			want: 8,
		},
		{
			name: "abs(-8)",
			args: args{
				n: -8,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AbsInt64(tt.args.n); got != tt.want {
				t.Errorf("AbsInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}
