package cond

import (
	"reflect"
	"testing"
)

func TestIf(t *testing.T) {
	type args struct {
		condition bool
		trueVal   any
		falseVal  any
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			"return true",
			args{
				1 > 0,
				true,
				false,
			},
			true,
		},
		{
			"return false",
			args{
				1 < 0,
				true,
				false,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := If(tt.args.condition, tt.args.trueVal, tt.args.falseVal); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("If() = %v, want %v", got, tt.want)
			}
		})
	}
}
