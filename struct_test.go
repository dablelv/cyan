package util

import (
	"reflect"
	"testing"
)

func TestStruct2Map(t *testing.T) {
	type args struct {
		obj interface{}
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "convert success",
			args: args{
				struct {
					I int
					S string
				}{I: 1, S: "a"}},
			want: map[string]interface{}{"I": 1, "S": "a"},
		},
		{
			name: "convert struct contains unexported fields success",
			args: args{
				struct {
					I int
					S string
					s string
				}{I: 1, S: "a", s: "a"}},
			want: map[string]interface{}{"I": 1, "S": "a"},
		},
		{
			name: "convert empty struct success",
			args: args{struct{}{}},
			want: map[string]interface{}{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Struct2Map(tt.args.obj); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Struct2Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStruct2MapString(t *testing.T) {
	type args struct {
		obj interface{}
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			name: "convert success",
			args: args{
				struct {
					I int
					S string
				}{I: 1, S: "a"}},
			want: map[string]string{"I": "1", "S": "a"},
		},
		{
			name: "convert struct contains unexported fields success",
			args: args{
				struct {
					I int
					S string
					s string
				}{I: 1, S: "a", s: "a"}},
			want: map[string]string{"I": "1", "S": "a"},
		},
		{
			name: "convert empty struct success",
			args: args{struct{}{}},
			want: map[string]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Struct2MapString(tt.args.obj); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Struct2MapString() = %v, want %v", got, tt.want)
			}
		})
	}
}
