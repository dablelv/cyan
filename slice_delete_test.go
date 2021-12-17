package util

import (
	"reflect"
	"testing"
)

func TestDeleteSliceElms(t *testing.T) {
	type args struct {
		i    interface{}
		elms []interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name:"delete success",
			args:args{
				i:[]string{"a","b","b","c"},
				elms: []interface{}{"b", "c"},
			},
			want:[]string{"a"},
		},
		{
			name:"delete all",
			args:args{
				i:[]string{"a","b","b","c"},
				elms: []interface{}{"a","b","c"},
			},
			want:[]string{},
		},
		{
			name:"delete failed because element type is ill",
			args:args{
				i:[]string{"1","2","2","3"},
				elms: []interface{}{2, 3},
			},
			want:nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteSliceElms(tt.args.i, tt.args.elms...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteSliceElms() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteSlice(t *testing.T) {
	type args struct {
		slice   interface{}
		indexes []int
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name:"delete success",
			args:args{
				slice:[]string{"a","b","b","c"},
				indexes: []int{1,2},
			},
			want:[]string{"a","c"},
		},
		{
			name:"delete all",
			args:args{
				slice:[]string{"a","b","b","c"},
				indexes: []int{0,1,2,3,4,5},
			},
			want:[]string{},
		},
		{
			name:"delete failed because the input isn't slice",
			args:args{
				slice:"a",
				indexes: []int{0,1,2,3,4,5},
			},
			want:nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteSlice(tt.args.slice, tt.args.indexes...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteStrSlice(t *testing.T) {
	type args struct {
		src     []string
		indexes []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name:"delete success",
			args:args{
				src:[]string{"a","b","b","c"},
				indexes: []int{1,2},
			},
			want:[]string{"a","c"},
		},
		{
			name:"delete all",
			args:args{
				src:[]string{"a","b","b","c"},
				indexes: []int{0,1,2,3,4,5},
			},
			want:[]string{},
		},
		{
			name:"not delete",
			args:args{
				src:[]string{"a","b","b","c"},
				indexes: []int{},
			},
			want:[]string{"a","b","b","c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteStrSlice(tt.args.src, tt.args.indexes...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteStrSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}