package slice

import (
	"reflect"
	"testing"
)

func TestReverseInt(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "int slice",
			args: args{[]int{1, 2, 3}},
			want: []int{3, 2, 1},
		},
		{
			name: "empty int slice",
			args: args{[]int{}},
			want: []int{},
		},
		{
			name: "nil int slice",
			args: args{nil},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseIntSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseStr(t *testing.T) {
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
			name: "reverse emplty string slice",
			args: args{[]string{}},
			want: []string{},
		},
		{
			name: "reverse nil string slice",
			args: args{nil},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseStrSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseIntSlice(t *testing.T) {
	type args struct {
		src []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "int slice",
			args: args{[]int{1, 2, 3}},
			want: []int{3, 2, 1},
		},
		{
			name: "empty int slice",
			args: args{[]int{}},
			want: []int{},
		},
		{
			name: "nil int slice",
			args: args{nil},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseIntSlice(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseIntSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseUintSlice(t *testing.T) {
	type args struct {
		src []uint
	}
	tests := []struct {
		name string
		args args
		want []uint
	}{
		{
			name: "uint slice",
			args: args{[]uint{1, 2, 3}},
			want: []uint{3, 2, 1},
		},
		{
			name: "empty uint slice",
			args: args{[]uint{}},
			want: []uint{},
		},
		{
			name: "nil uint slice",
			args: args{nil},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseUintSlice(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseIntSlice() = %v, want %v", got, tt.want)
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
			name: "reverse empty string slice",
			args: args{[]string{}},
			want: []string{},
		},
		{
			name: "reverse nil string slice",
			args: args{nil},
			want: nil,
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

func TestReverseSliceRefE(t *testing.T) {
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
			name:    "reverse empty string slice",
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
			name:    "reverse empty uint64 slice",
			args:    args{[]uint64{}},
			want:    []uint64{},
			wantErr: false,
		},
		{
			name:    "input isn't a slice",
			args:    args{"string"},
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
