package slice

import (
	"reflect"
	"testing"
)

func TestJoinSliceWithSepE(t *testing.T) {
	type args struct {
		slice any
		sep   string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "join string slice",
			args:    args{[]string{"foo", "bar", "baz"}, ","},
			want:    "foo,bar,baz",
			wantErr: false,
		},
		{
			name:    "join string array",
			args:    args{[3]string{"foo", "bar", "baz"}, ","},
			want:    "foo,bar,baz",
			wantErr: false,
		},
		{
			name:    "join int slice",
			args:    args{[]int{1, 2, 3}, ","},
			want:    "1,2,3",
			wantErr: false,
		},
		{
			name:    "join int array",
			args:    args{[3]int{1, 2, 3}, ","},
			want:    "1,2,3",
			wantErr: false,
		},
		{
			name:    "join nil",
			args:    args{nil, ","},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := JoinSliceWithSepE(tt.args.slice, tt.args.sep)
			if (err != nil) != tt.wantErr {
				t.Errorf("JoinSliceWithSepE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("JoinSliceWithSepE() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRandomSliceElem(t *testing.T) {
	type args struct {
		i any
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			name: "Get a random element from a slice of length one",
			args: args{[]int{1}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRandomSliceElem(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRandomSliceElem() = %v, want %v", got, tt.want)
			}
		})
	}
}
