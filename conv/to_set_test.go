package conv

import (
	"reflect"
	"testing"
)

func TesttoSetE(t *testing.T) {
	type args struct {
		i any
	}
	tests := []struct {
		name    string
		args    args
		want    map[any]struct{}
		wantErr bool
	}{
		{
			name: "string array to map set",
			args: args{[3]string{"foo", "bar", "baz"}},
			want: map[any]struct{}{
				"foo": {},
				"bar": {},
				"baz": {},
			},
			wantErr: false,
		},
		{
			name: "string slice to map set",
			args: args{[]string{"foo", "bar", "baz"}},
			want: map[any]struct{}{
				"foo": {},
				"bar": {},
				"baz": {},
			},
			wantErr: false,
		},
		{
			name:    "int array to  map set",
			args:    args{[3]int{86, 852, 61}},
			want:    map[any]struct{}{86: {}, 852: {}, 61: {}},
			wantErr: false,
		},
		{
			name:    "int slice to map set",
			args:    args{[]int{86, 852, 61}},
			want:    map[any]struct{}{86: {}, 852: {}, 61: {}},
			wantErr: false,
		},
		{
			name:    "arg isn't slice and array",
			args:    args{"foo"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := toSetE(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("toSetE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toSetE() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToStrSetE(t *testing.T) {
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
			name:    "[]string",
			args:    args{[]string{"CN", "HK", "AU"}},
			want:    map[string]struct{}{"CN": {}, "HK": {}, "AU": {}},
			wantErr: false,
		},
		{
			name:    "[3]int",
			args:    args{[3]int{86, 852, 61}},
			want:    map[string]struct{}{"86": {}, "852": {}, "61": {}},
			wantErr: false,
		},
		{
			name:    "string",
			args:    args{""},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToStrSetE(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToStrSetE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToStrSetE() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint64Set(t *testing.T) {
	type args struct {
		i any
	}
	tests := []struct {
		name string
		args args
		want map[uint64]struct{}
	}{
		{
			name: "uint64 array to map set",
			args: args{[3]uint64{1, 2, 3}},
			want: map[uint64]struct{}{
				1: {},
				2: {},
				3: {},
			},
		},
		{
			name: "uint64 slice to map set",
			args: args{[]uint64{1, 2, 3}},
			want: map[uint64]struct{}{
				1: {},
				2: {},
				3: {},
			},
		},
		{
			name: "zero length uint64 array to map set",
			args: args{[0]uint64{}},
			want: map[uint64]struct{}{},
		},
		{
			name: "zero length uint64 slice to map set",
			args: args{[]uint64{}},
			want: map[uint64]struct{}{},
		},
		{
			name: "nil slice to map set",
			args: args{[]uint64(nil)},
			want: map[uint64]struct{}{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUint64Set(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToU64Set() = %v, want %v", got, tt.want)
			}
		})
	}
}
