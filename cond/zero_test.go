package cond

import (
	"testing"
)

func TestIsZeroBool(t *testing.T) {
	type args struct {
		v bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"false is zero",
			args{false},
			true,
		},
		{
			"true isn't zero",
			args{true},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsZero(tt.args.v); got != tt.want {
				t.Errorf("IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsZeroInt(t *testing.T) {
	type args struct {
		v int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"0 is zero",
			args{0},
			true,
		},
		{
			"1 isn't zero",
			args{1},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsZero(tt.args.v); got != tt.want {
				t.Errorf("IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsZeroStr(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			`"" is zero`,
			args{""},
			true,
		},
		{
			`"foo" isn't zero"`,
			args{"foo"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsZero(tt.args.v); got != tt.want {
				t.Errorf("IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsZeroRefMap(t *testing.T) {
	type args struct {
		v map[string]string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"nil is zero",
			args{nil},
			true,
		},
		{
			"empty is not zero",
			args{map[string]string{}},
			false,
		},
		{
			"nonempty map is not zero",
			args{map[string]string{"foo": "foo"}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsZeroRef(tt.args.v); got != tt.want {
				t.Errorf("IsZeroRef() = %v, want %v", got, tt.want)
			}
		})
	}
}
