package conv

import (
	"reflect"
	"testing"
)

func TestToBoolSetGE(t *testing.T) {
	type args struct {
		i any
	}
	tests := []struct {
		name    string
		args    args
		want    map[bool]struct{}
		wantErr bool
	}{
		{
			name: "to bool set",
			args: args{[]bool{true, false, true}},
			want: map[bool]struct{}{
				true:  {},
				false: {},
			},
			wantErr: false,
		}, {
			name:    "arg is nil",
			args:    args{nil},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "arg is neither a slice nor an array",
			args:    args{"foo"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToSetGE[bool](tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToSetGE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToSetGE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToIntSetGE(t *testing.T) {
	type args struct {
		i any
	}
	tests := []struct {
		name    string
		args    args
		want    map[int]struct{}
		wantErr bool
	}{
		{
			name: "to int set",
			args: args{[]int{1, 2, 3}},
			want: map[int]struct{}{
				1: {},
				2: {},
				3: {},
			},
			wantErr: false,
		},
		{
			name:    "arg is nil",
			args:    args{nil},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "arg is neither a slice nor an array",
			args:    args{"foo"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToSetGE[int](tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToSetGE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToSetGE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToStrSetGE(t *testing.T) {
	type args struct {
		i any
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]struct{}
		wantErr bool
	}{
		{
			name: "to string set",
			args: args{[]string{"foo", "bar", "baz"}},
			want: map[string]struct{}{
				"foo": {},
				"bar": {},
				"baz": {},
			},
			wantErr: false,
		},
		{
			name:    "arg is nil",
			args:    args{nil},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "arg is neither a slice nor an array",
			args:    args{"foo"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToSetGE[string](tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToSetGE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToSetGE() = %v, want %v", got, tt.want)
			}
		})
	}
}
