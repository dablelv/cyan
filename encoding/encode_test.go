package encoding

import (
	"testing"
)

func TestGBK2UTF8Str(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"中文",
			args{string([]byte{0xD6, 0xD0, 0xCE, 0xC4})},
			"中文",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GBK2UTF8Str(tt.args.s); got != tt.want {
				t.Errorf("GBK2UTF8Str() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUTF82GBKStr(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"中文",
			args{"中文"},
			string([]byte{0xD6, 0xD0, 0xCE, 0xC4}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UTF82GBKStr(tt.args.s); got != tt.want {
				t.Errorf("UTF82GBKStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
