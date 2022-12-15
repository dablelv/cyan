package str

import "testing"

func TestJoinStrSkipEmpty(t *testing.T) {
	type args struct {
		sep string
		s   []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"join string",
			args{",", []string{"foo", "bar", "baz"}},
			"foo,bar,baz",
		},
		{
			"join string and skip the empty string",
			args{",", []string{"", "foo", "bar", "baz"}},
			"foo,bar,baz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JoinStrSkipEmpty(tt.args.sep, tt.args.s...); got != tt.want {
				t.Errorf("JoinStrSkipEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJoinStr(t *testing.T) {
	type args struct {
		sep string
		s   []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"join string",
			args{",", []string{"foo", "bar", "baz"}},
			"foo,bar,baz",
		},
		{
			"join string and don't skip the empty string",
			args{",", []string{"", "foo", "bar", "baz"}},
			",foo,bar,baz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JoinStr(tt.args.sep, tt.args.s...); got != tt.want {
				t.Errorf("JoinStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAlphanumericNumByASCII(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "包含数字",
			args: args{"108条梁山好汉"},
			want: 3,
		},
		{
			name: "包含字母",
			args: args{"一百条梁山man"},
			want: 3,
		},
		{
			name: "包含数字与字母",
			args: args{"108条梁山man"},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAlphanumericNumByASCII(tt.args.s); got != tt.want {
				t.Errorf("GetAlphanumericNumByASCII() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAlphanumericNumByASCIIV2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "包含数字",
			args: args{"108条梁山好汉"},
			want: 3,
		},
		{
			name: "包含字母",
			args: args{"一百条梁山man"},
			want: 3,
		},
		{
			name: "包含数字与字母",
			args: args{"108条梁山man"},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAlphanumericNumByASCIIV2(tt.args.s); got != tt.want {
				t.Errorf("GetAlphanumericNumByASCII() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAlphanumericNumByRegExp(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "包含数字",
			args: args{"108条梁山好汉"},
			want: 3,
		},
		{
			name: "包含字母",
			args: args{"一百条梁山man"},
			want: 3,
		},
		{
			name: "包含数字与字母",
			args: args{"108条梁山man"},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAlphanumericNumByRegExp(tt.args.s); got != tt.want {
				t.Errorf("GetAlphanumericNumByRegExp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkGetAlphanumericNumByASCII(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetAlphanumericNumByASCII("108条梁山man")
	}
}

func BenchmarkGetAlphanumericNumByASCIIV2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetAlphanumericNumByASCIIV2("108条梁山man")
	}
}

func BenchmarkGetAlphanumericNumByRegExp(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetAlphanumericNumByRegExp("108条梁山man")
	}
}
