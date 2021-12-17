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
