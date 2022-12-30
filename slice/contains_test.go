package slice

import "testing"

func TestIsContains(t *testing.T) {
	type args struct {
		slice  any
		target any
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "string slice contain",
			args: args{slice: []string{"foo", "bar", "baz"}, target: "baz"},
			want: true,
		},
		{
			name: "string array contain",
			args: args{slice: [3]string{"foo", "bar", "baz"}, target: "baz"},
			want: true,
		},
		{
			name: "string slice not contain",
			args: args{slice: []string{"foo", "bar", "baz"}, target: "qux"},
			want: false,
		},
		{
			name: "int slice contain",
			args: args{slice: []int{1, 2, 3}, target: 1},
			want: true,
		},
		{
			name: "int32 slice not contain because type isn't equal",
			args: args{slice: []int32{1, 2, 3}, target: 1},
			want: false,
		},
		{
			name: "nil not contain",
			args: args{slice: nil, target: 1},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.slice, tt.args.target); got != tt.want {
				t.Errorf("IsContains() = %v, want %v", got, tt.want)
			}
		})
	}
}
