package str

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	type args struct {
		s   string
		sep string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"case 1",
			args{"a,b,c", ","},
			[]string{"a", "b", "c"},
		},
		{
			"case 2",
			args{" xyz ", ""},
			[]string{" ", "x", "y", "z", " "},
		},
		{
			"case 3",
			args{"", ","},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Split(tt.args.s, tt.args.sep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Split() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitSeps(t *testing.T) {
	type args struct {
		s    string
		seps []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"case 1",
			args{"foo,bar,baz", []string{","}},
			[]string{"foo", "bar", "baz"},
		},
		{
			"case 2",
			args{"foo,bar|baz", []string{",", "|"}},
			[]string{"foo", "bar", "baz"},
		},
		{
			"case 3",
			args{"foo,bar|baz qux", []string{",", "|", " "}},
			[]string{"foo", "bar", "baz", "qux"},
		},
		{
			"case 4",
			args{"foo,bar|bazSEPqux", []string{",", "|", "SEP"}},
			[]string{"foo", "bar", "baz", "qux"},
		},
		{
			"case 5",
			args{"foo,bar|baz", []string{}},
			[]string{"foo,bar|baz"},
		},
		{
			"case 6",
			args{" xyz", []string{""}},
			[]string{" ", "x", "y", "z"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitSeps(tt.args.s, tt.args.seps...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitSeps() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
