package cmp

import (
	"reflect"
	"testing"

	"github.com/dablelv/cyan/internal"
	"golang.org/x/exp/constraints"
)

func TestCmpInt(t *testing.T) {
	type args[T constraints.Ordered] struct {
		lhs T
		rhs T
	}
	ints := []struct {
		name string
		args args[int]
		want CMPRSLT
	}{
		{"int compare lt", args[int]{888, 889}, LT},
		{"int compare eq", args[int]{888, 888}, EQ},
		{"int compare gt", args[int]{889, 888}, GT},
	}
	f64s := []struct {
		name string
		args args[float64]
		want CMPRSLT
	}{
		{"float compare lt", args[float64]{88.8, 88.9}, LT},
		{"float compare eq", args[float64]{88.8, 88.8}, EQ},
		{"float compare gt", args[float64]{88.9, 88.8}, GT},
	}
	strs := []struct {
		name string
		args args[string]
		want CMPRSLT
	}{
		{"string compare lt", args[string]{"abc", "b"}, LT},
		{"string compare eq", args[string]{"abc", "abc"}, EQ},
		{"string compare gt", args[string]{"b", "abc"}, GT},
	}
	for _, tt := range ints {
		t.Run(tt.name, func(t *testing.T) {
			if got := Cmp(tt.args.lhs, tt.args.rhs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cmp() = %v, want %v", got, tt.want)
			}
		})
	}
	for _, tt := range f64s {
		t.Run(tt.name, func(t *testing.T) {
			if got := Cmp(tt.args.lhs, tt.args.rhs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cmp() = %v, want %v", got, tt.want)
			}
		})
	}
	for _, tt := range strs {
		t.Run(tt.name, func(t *testing.T) {
			if got := Cmp(tt.args.lhs, tt.args.rhs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cmp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompare(t *testing.T) {
	type args struct {
		lhs any
		rhs any
	}
	tests := []struct {
		name string
		args args
		want CMPRSLT
	}{
		{"incomparable", args{888, "abc"}, INCMP},
		{"int compare lt", args{888, 889}, LT},
		{"int compare eq", args{888, 888}, EQ},
		{"int compare gt", args{889, 888}, GT},
		{"uint compare lt", args{uint(888), uint(889)}, LT},
		{"uint compare eq", args{uint(888), uint(888)}, EQ},
		{"uint compare gt", args{uint(889), uint(888)}, GT},
		{"float compare lt", args{88.8, 88.9}, LT},
		{"float compare eq", args{88.8, 88.8}, EQ},
		{"float compare gt", args{88.9, 88.8}, GT},
		{"string compare lt", args{"abc", "b"}, LT},
		{"string compare eq", args{"abc", "abc"}, EQ},
		{"string compare gt", args{"b", "abc"}, GT},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Compare(tt.args.lhs, tt.args.rhs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMin(t *testing.T) {
	assert := internal.NewAssert(t, "TestMin")

	assert.Equal(1, Min(1, 2))
	assert.Equal(1.1, Min(2.2, 1.1))
	assert.Equal("abc", Min("abc", "cba"))
}

func TestMax(t *testing.T) {
	assert := internal.NewAssert(t, "TestMax")

	assert.Equal(2, Max(1, 2))
	assert.Equal(2.2, Max(2.2, 1.1))
	assert.Equal("cba", Max("abc", "cba"))
}
