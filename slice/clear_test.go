package slice

import (
	"reflect"
	"testing"
)

func TestClearZeroBool(t *testing.T) {
	type args struct {
		s []bool
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			"clear bool slice",
			args{[]bool{true, false, true}},
			[]bool{true, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClearZero(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClearZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClearZeroInt(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			"clear int slice",
			args{[]int{1, 2, 0, 3}},
			[]int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClearZero(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClearZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClearZeroStr(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			"clear string slice",
			args{[]string{"foo", "bar", "", "baz"}},
			[]string{"foo", "bar", "baz"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClearZero(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClearZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClearZeroRefMap(t *testing.T) {
	type args struct {
		s []map[string]string
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			"clear string slice",
			args{[]map[string]string{
				{"foo": "foo"},
				nil,
				{"bar": "bar"},
			}},
			[]map[string]string{{"foo": "foo"}, {"bar": "bar"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClearZeroRef(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClearZeroRef() = %v, want %v", got, tt.want)
			}
		})
	}
}
