package slice

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseIntSlice(t *testing.T) {
	intSlice := []int{1, 2, 3}
	assert.Equal(t, []int{3, 2, 1}, ReverseIntSlice(intSlice))
}

func TestReverseSliceE(t *testing.T) {
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
			name:    "reverse string slice",
			args:    args{[]string{"foo", "bar", "baz"}},
			want:    []string{"baz", "bar", "foo"},
			wantErr: false,
		},
		{
			name:    "reverse zero length string slice",
			args:    args{[]string{}},
			want:    []string{},
			wantErr: false,
		},
		{
			name:    "reverse uint64 slice",
			args:    args{[]uint64{1, 2, 3}},
			want:    []uint64{3, 2, 1},
			wantErr: false,
		},
		{
			name:    "reverse zero length uint64 slice",
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
			got, err := ReverseSliceE(tt.args.slice)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReverseSliceE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseSliceE() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseStrSlice(t *testing.T) {
	type args struct {
		src []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "reverse string slice",
			args: args{[]string{"foo", "bar", "baz"}},
			want: []string{"baz", "bar", "foo"},
		},
		{
			name: "reverse zero length string slice",
			args: args{[]string{}},
			want: []string{},
		},
		{
			name: "reverse nil string slice",
			args: args{[]string{}},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseStrSlice(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseStrSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
