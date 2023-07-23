package str

import (
	"reflect"
	"testing"

	"github.com/dablelv/go-huge-util/internal"
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
	assert := internal.NewAssert(t, "TestSplitSeps")

	assert.Equal([]string{"foo", "bar", "baz"}, SplitSeps("foo,bar,baz", ","))
	assert.Equal([]string{"foo", "bar", "baz"}, SplitSeps("foo,bar|baz", ",", "|"))
	assert.Equal([]string{"foo", "bar", "baz", "qux"}, SplitSeps("foo,bar|baz qux", ",", "|", " "))
	assert.Equal([]string{"foo", "bar", "baz", "qux"}, SplitSeps("foo,bar|bazSEPqux", ",", "|", "SEP"))
	assert.Equal([]string{"foo,bar|baz"}, SplitSeps("foo,bar|baz"))
	assert.Equal([]string{" ", "x", "y", "z"}, SplitSeps(" xyz", ""))
}

func TestJoinNonEmptyStrs(t *testing.T) {
	assert := internal.NewAssert(t, "TestJoinNonEmptyStrs")

	got := JoinNonEmptyStrs(",", []string{"foo", "bar", "baz"}...)
	assert.Equal("foo,bar,baz", got)

	got = JoinNonEmptyStrs(",", []string{"foo", "bar", "baz"}...)
	assert.Equal("foo,bar,baz", got)
}

func TestJoin(t *testing.T) {
	assert := internal.NewAssert(t, "TestJoin")

	got := Join([]string{"foo", "bar", "baz"}, ",")
	assert.Equal("foo,bar,baz", got)

	got = Join([3]string{"foo", "bar", "baz"}, ",")
	assert.Equal("foo,bar,baz", got)

	got = Join([]int{1, 2, 3}, ",")
	assert.Equal("1,2,3", got)

	got = Join([3]int{1, 2, 3}, ",")
	assert.Equal("1,2,3", got)

	got = Join("foo", ",")
	assert.Equal("f,o,o", got)

	got = Join([]struct{}{{}}, ",")
	assert.Equal("", got)
}

func TestReverse(t *testing.T) {
	assert := internal.NewAssert(t, "TestReverse")
	assert.Equal("foo", Reverse("oof"))
	assert.Equal("123", Reverse("321"))
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
