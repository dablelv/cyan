package slice

import (
	"reflect"
	"testing"
)

func TestUniqueSliceE(t *testing.T) {
	type args struct {
		slice any
	}
	tests := []struct {
		name    string
		args    args
		want    any
		wantErr bool
	}{
		{
			name:    "unique string slice",
			args:    args{[]string{"foo", "bar", "bar", "baz"}},
			want:    []string{"foo", "bar", "baz"},
			wantErr: false,
		},
		{
			name:    "unique zero length string slice",
			args:    args{[]string{}},
			want:    []string{},
			wantErr: false,
		},
		{
			name:    "unique uint64 slice",
			args:    args{[]uint64{1, 2, 2, 3}},
			want:    []uint64{1, 2, 3},
			wantErr: false,
		},
		{
			name:    "unique zero length uint64 slice",
			args:    args{[]uint64{}},
			want:    []uint64{},
			wantErr: false,
		},
		{
			name:    "input isn't a slice",
			args:    args{nil},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UniqueSliceE(tt.args.slice)
			if (err != nil) != tt.wantErr {
				t.Errorf("UniqueSliceE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueSliceE() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueStrSlice(t *testing.T) {
	type args struct {
		src []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "unique string slice",
			args: args{[]string{"foo", "bar", "bar", "baz"}},
			want: []string{"foo", "bar", "baz"},
		},
		{
			name: "unique zero length string slice",
			args: args{[]string{}},
			want: []string{},
		},
		{
			name: "unique nil string slice",
			args: args{nil},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueStrSlice(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueStrSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
