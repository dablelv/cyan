package util

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap2StrSliceE(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		want1   []string
		wantErr bool
	}{
		{
			name:"map[string]string",
			args:args{map[string]string{"CN":"China", "HK":"Hong Kong","AU":"Australia"}},
			want:[]string{"CN","HK","AU"},
			want1:[]string{"China","Hong Kong","Australia"},
			wantErr: false,
		},
		{
			name:"map[string]int",
			args:args{map[string]int{"CN":86, "HK":852,"AU":61}},
			want:[]string{"CN","HK","AU"},
			want1:[]string{"86","852","61"},
			wantErr: false,
		},
		{
			name:"map[int]string",
			args:args{map[int]string{86:"China", 852:"Hong Kong",61:"Australia"}},
			want:[]string{"86","852","61"},
			want1:[]string{"China","Hong Kong","Australia"},
			wantErr: false,
		},
		{
			name:"map[string]string",
			args:args{map[string]string{"CN":"China", "HK":"Hong Kong","AU":"Australia"}},
			want:[]string{"CN","HK","AU"},
			want1:[]string{"China","Hong Kong","Australia"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := Map2StrSliceE(tt.args.i)
			assert.Equal(t, tt.wantErr, err != nil, tt.name)
			assert.Equal(t, len(tt.want), len(got), tt.name)
			assert.Equal(t, len(tt.want1), len(got1), tt.name)
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
			name:"[]string",
			args:args{[]string{"CN","HK","AU"}},
			want:map[interface{}]struct{}{"CN":struct{}{},"HK":struct{}{},"AU":struct{}{}},
			wantErr: false,
		},
		{
			name:"[3]int",
			args:args{[3]int{86,852,61}},
			want:map[interface{}]struct{}{86:struct{}{},852:struct{}{},61:struct{}{}},
			wantErr: false,
		},
		{
			name:"string",
			args:args{""},
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
			name:"[]string",
			args:args{[]string{"CN","HK","AU"}},
			want:map[string]struct{}{"CN":struct{}{},"HK":struct{}{},"AU":struct{}{}},
			wantErr: false,
		},
		{
			name:"[3]int",
			args:args{[3]int{86,852,61}},
			want:map[string]struct{}{"86":struct{}{},"852":struct{}{},"61":struct{}{}},
			wantErr: false,
		},
		{
			name:"string",
			args:args{""},
			want:nil,
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

func TestToU64MapSet(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		args args
		want map[uint64]struct{}
	}{
		{
			name:"array to map set",
			args:args{[3]uint64{1,2,3}},
			want:map[uint64]struct{}{
				1:struct{}{},
				2:struct{}{},
				3:struct{}{},
			},
		},
		{
			name:"zero length slice to map set",
			args:args{[]uint64{}},
			want:map[uint64]struct{}{
			},
		},
		{
			name:"nil slice to map set",
			args:args{[]uint64(nil)},
			want:map[uint64]struct{}{
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToU64MapSet(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToU64MapSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap2Slice(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name  string
		args  args
		wantK []interface{}
		wantV []interface{}
	}{
		{
			name:"map[string]string to slice ",
			args:args{
				map[string]string{"1":"1","2":"2","3":"3"},
			},
			wantK:[]interface{}{"1","2","3"},
			wantV:[]interface{}{"1","2","3"},
		},
		{
			name:"map[int]string to slice",
			args:args{
				map[int]string{1:"1",2:"2",3:"3"},
			},
			wantK:[]interface{}{1,2,3},
			wantV:[]interface{}{"1","2","3"},
		},
		{
			name:"map[int]int to slice",
			args:args{
				map[int]int{1:1,2:2,3:3},
			},
			wantK:[]interface{}{1,2,3},
			wantV:[]interface{}{1,2,3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotK, gotV := Map2Slice(tt.args.i)
			assert.Equal(t, len(gotK), len(tt.wantK))
			assert.Equal(t, len(gotV), len(tt.wantV))
		})
	}
}