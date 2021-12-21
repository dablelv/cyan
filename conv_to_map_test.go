package util

import (
	"reflect"
	"testing"
)

func TestToMapSetStrictE(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name:"string slice to map set",
			args:args{[]string{"foo","bar","baz"}},
			want:map[string]struct{}{
				"foo":struct{}{},
				"bar":struct{}{},
				"baz":struct{}{},
			},
			wantErr: false,
		},
		{
			name:"int slice to map set",
			args:args{[]int{1,2,3, 3}},
			want:map[int]struct{}{
				1:struct{}{},
				2:struct{}{},
				3:struct{}{},
			},
			wantErr: false,
		},
		{
			name:"arg is nil",
			args:args{nil},
			want:nil,
			wantErr: true,
		},
		{
			name:"arg is ill",
			args:args{"foo"},
			want:nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToMapSetStrictE(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToMapSetIfcE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToMapSetIfcE() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToMapSetE(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    map[interface{}]struct{}
		wantErr bool
	}{
		{
			name :"string array to map set",
			args: args{[3]string{"foo","bar","baz"}},
			want:map[interface{}]struct{}{
				"foo": struct{}{},
				"bar":struct{}{},
				"baz": struct{}{},
			},
			wantErr: false,
		},
		{
			name :"string slice to map set",
			args: args{[]string{"foo","bar","baz"}},
			want:map[interface{}]struct{}{
				"foo": struct{}{},
				"bar":struct{}{},
				"baz": struct{}{},
			},
			wantErr: false,
		},
		{
			name:    "int array to  map set",
			args:    args{[3]int{86, 852, 61}},
			want:    map[interface{}]struct{}{86: struct{}{}, 852: struct{}{}, 61: struct{}{}},
			wantErr: false,
		},
		{
			name:    "int slice to  map set",
			args:    args{[]int{86, 852, 61}},
			want:    map[interface{}]struct{}{86: struct{}{}, 852: struct{}{}, 61: struct{}{}},
			wantErr: false,
		},
		{
			name :"param isn't slice or array",
			args: args{"foo"},
			want:nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToMapSetE(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToMapSetE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToMapSetE() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToStrMapSetE(t *testing.T) {
	type args struct {
		i interface{}
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
			want:    map[string]struct{}{"CN": struct{}{}, "HK": struct{}{}, "AU": struct{}{}},
			wantErr: false,
		},
		{
			name:    "[3]int",
			args:    args{[3]int{86, 852, 61}},
			want:    map[string]struct{}{"86": struct{}{}, "852": struct{}{}, "61": struct{}{}},
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
			got, err := ToStrMapSetE(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToStrMapSetE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToStrMapSetE() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint64MapSet(t *testing.T) {
	type args struct {
		i interface{}
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
				1: struct{}{},
				2: struct{}{},
				3: struct{}{},
			},
		},
		{
			name: "uint64 slice to map set",
			args: args{[]uint64{1, 2, 3}},
			want: map[uint64]struct{}{
				1: struct{}{},
				2: struct{}{},
				3: struct{}{},
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
			if got := ToUint64MapSet(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToU64MapSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToStrMapSetStrict(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		args args
		want map[string]struct{}
	}{
		{
			name:"string array to map set",
			args:args{[3]string{"foo","bar","baz"}},
			want:map[string]struct{}{
				"foo":struct{}{},
				"bar":struct{}{},
				"baz":struct{}{},
			},
		},
		{
			name:"string slice to map set",
			args:args{[]string{"foo","bar","baz"}},
			want:map[string]struct{}{
				"foo":struct{}{},
				"bar":struct{}{},
				"baz":struct{}{},
			},
		},
		{
			name:"element type isn't string can't convert to map set",
			args:args{[]uint64{1,2,3}},
			want:nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToStrMapSetStrict(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToStrMapSetStrict() = %v, want %v", got, tt.want)
			}
		})
	}
}