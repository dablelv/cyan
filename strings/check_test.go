package strings

import (
	"testing"
)

func TestIsNumeric(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "empty string",
			input: "",
			want:  false,
		},
		{
			name:  "positive integer",
			input: "12345",
			want:  true,
		},
		{
			name:  "negative integer",
			input: "-12345",
			want:  true,
		},
		{
			name:  "positive decimal",
			input: "0.4581",
			want:  true,
		},
		{
			name:  "negative decimal",
			input: "-0.4581",
			want:  true,
		},
		{
			name:  "invalid number 1",
			input: "0.45.81",
			want:  false,
		},
		{
			name:  "invalid number 2",
			input: "0.45,81",
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNumeric(tt.input); got != tt.want {
				t.Errorf("IsNumeric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsEmail(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "empty string",
			input: "",
			want:  false,
		},
		{
			name:  "valid email",
			input: "10000@qq.com",
			want:  true,
		},
		{
			name:  "chinese mailbox",
			input: "我的@你的.com",
			want:  true,
		},
		{
			name:  "invalid email 1",
			input: "foo@bar@.com",
			want:  false,
		},
		{
			name:  "invalid email 2",
			input: "123456",
			want:  false,
		},
		{
			name:  "invalid email 3",
			input: "foo",
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmail(tt.input); got != tt.want {
				t.Errorf("IsEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsChineseIdCode(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "empty string",
			input: "",
			want:  false,
		},
		{
			name:  "length incorrect",
			input: "44030619810825051",
			want:  false,
		},
		{
			name:  "birthday incorrect",
			input: "440306888808250512",
			want:  false,
		},
		{
			name:  "check code incorrect",
			input: "440306198108250511",
			want:  false,
		},
		{
			name:  "valid id number",
			input: "440306198108250515",
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsChineseIdCode(tt.input); got != tt.want {
				t.Errorf("IsChineseIdCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
