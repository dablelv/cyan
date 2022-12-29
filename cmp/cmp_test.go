package cmp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompare(t *testing.T) {
	type args struct {
		lhs any
		rhs any
	}
	tests := []struct {
		name string
		args args
		want CMPRES
	}{
		{"incomparable", args{888, "abc"}, INCMP},
		{"int compare", args{888, 889}, LT},
		{"int compare", args{888, 888}, EQ},
		{"int compare", args{889, 888}, GT},
		{"uint compare", args{uint(888), uint(889)}, LT},
		{"uint compare", args{uint(888), uint(888)}, EQ},
		{"uint compare", args{uint(889), uint(888)}, GT},
		{"float compare", args{88.8, 88.9}, LT},
		{"float compare", args{88.8, 88.8}, EQ},
		{"float compare", args{88.9, 88.8}, GT},
		{"string compare", args{"abc", "b"}, LT},
		{"string compare", args{"abc", "abc"}, EQ},
		{"string compare", args{"b", "abc"}, GT},
	}
	for _, tt := range tests {
		got := Compare(tt.args.lhs, tt.args.rhs)
		assert.Equal(t, tt.want, got, tt.name)
	}
}
