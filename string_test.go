package util

import "testing"

func TestJoinStrSkipEmpty(t *testing.T) {
	type args struct {
		sep string
		s   []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"join string",
			args {",", []string{"foo","bar","baz"}},
			"foo,bar,baz",
		},
		{
			"join string and skip the empty string",
			args {",", []string{"","foo","bar","baz"}},
			"foo,bar,baz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JoinStrSkipEmpty(tt.args.sep, tt.args.s...); got != tt.want {
				t.Errorf("JoinStrSkipEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJoinStr(t *testing.T) {
	type args struct {
		sep string
		s   []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"join string",
			args {",", []string{"foo","bar","baz"}},
			"foo,bar,baz",
		},
		{
			"join string and don't skip the empty string",
			args {",", []string{"","foo","bar","baz"}},
			",foo,bar,baz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JoinStr(tt.args.sep, tt.args.s...); got != tt.want {
				t.Errorf("JoinStr() = %v, want %v", got, tt.want)
			}
		})
	}
}