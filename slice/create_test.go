package slice

import (
	"reflect"
	"testing"
)

func TestInsertSliceE(t *testing.T) {
	type args struct {
		slice interface{}
		index int
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name:    "insert string to the first pos",
			args:    args{[]string{"bar", "baz"}, 0, "foo"},
			want:    []string{"foo", "bar", "baz"},
			wantErr: false,
		},
		{
			name:    "insert string to the last pos",
			args:    args{[]string{"foo", "bar"}, 2, "baz"},
			want:    []string{"foo", "bar", "baz"},
			wantErr: false,
		},
		{
			name:    "insert string failed because pos overflow",
			args:    args{[]string{"foo", "bar"}, 3, "baz"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := InsertSliceE(tt.args.slice, tt.args.index, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("InsertSliceE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertSliceE() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertIntSlice(t *testing.T) {
	type args struct {
		src   []int
		index int
		value int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "insert int to the first pos",
			args: args{[]int{1, 2}, 0, 0},
			want: []int{0, 1, 2},
		},
		{
			name: "insert int to the last pos",
			args: args{[]int{1, 2}, 2, 3},
			want: []int{1, 2, 3},
		},
		{
			name: "insert int failed becacuse the pos overflow",
			args: args{[]int{1, 2}, 3, 3},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertIntSlice(tt.args.src, tt.args.index, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertIntSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
