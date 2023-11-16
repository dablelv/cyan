package str

import (
	"reflect"
	"testing"

	"github.com/dablelv/cyan/internal"
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

func TestAlphanumericNum(t *testing.T) {
	assert := internal.NewAssert(t, "TestAlphanumericNum")

	assert.Equal(3, AlphanumericNum("108条梁山好汉"))
	assert.Equal(3, AlphanumericNum("一百条梁山man"))
	assert.Equal(6, AlphanumericNum("108条梁山man"))
}

func TestAlphanumericNumV2(t *testing.T) {
	assert := internal.NewAssert(t, "TestAlphanumericNumV2")

	assert.Equal(3, AlphanumericNumV2("108条梁山好汉"))
	assert.Equal(3, AlphanumericNumV2("一百条梁山man"))
	assert.Equal(6, AlphanumericNumV2("108条梁山man"))
}

func TestAlphanumericNumRegExp(t *testing.T) {
	assert := internal.NewAssert(t, "TestAlphanumericNumRegExp")

	assert.Equal(3, AlphanumericNumRegExp("108条梁山好汉"))
	assert.Equal(3, AlphanumericNumRegExp("一百条梁山man"))
	assert.Equal(6, AlphanumericNumRegExp("108条梁山man"))
}

func BenchmarkAlphanumericNum(b *testing.B) {
	for n := 0; n < b.N; n++ {
		AlphanumericNum("108条梁山man")
	}
}

func BenchmarkAlphanumericNumV2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		AlphanumericNumV2("108条梁山man")
	}
}

func BenchmarkGetAlphanumericNumByRegExp(b *testing.B) {
	for n := 0; n < b.N; n++ {
		AlphanumericNumRegExp("108条梁山man")
	}
}

func TestClearWhiteSpace(t *testing.T) {
	assert := internal.NewAssert(t, "TestClearWhiteSpace")

	assert.Equal("abc", ClearWhiteSpace("a b c"))
	assert.Equal("", ClearWhiteSpace("   "))
	assert.Equal("abcd", ClearWhiteSpace("a\rb\nc\td"))
}

func TestIndexOfDiff(t *testing.T) {
	assert := internal.NewAssert(t, "TestIndexOfDiff")

	assert.Equal(-1, IndexOfDiff("abc", "abc"))
	assert.Equal(0, IndexOfDiff("abc", "d"))
	assert.Equal(1, IndexOfDiff("abc", "a_c"))
}

func TestIndexOffset(t *testing.T) {
	assert := internal.NewAssert(t, "TestIndexOffset")

	assert.Equal(3, IndexOffset("abcde", "d", 2))
	assert.Equal(-1, IndexOffset("abcde", "d", 5))
	assert.Equal(3, IndexOffset("abcde", "d", -1))
	assert.Equal(-1, IndexOffset("abcde", "", 0))
	assert.Equal(-1, IndexOffset("abcde", "f", 0))
}

func TestDefault(t *testing.T) {
	assert := internal.NewAssert(t, "TestDefault")

	assert.Equal("abc", Default("abc", "d"))
	assert.Equal("d", Default("", "d"))
}

func TestDefaultIfBlank(t *testing.T) {
	assert := internal.NewAssert(t, "TestDefaultIfBlank")

	assert.Equal("abc", DefaultIfBlank("abc", "d"))
	assert.Equal("d", DefaultIfBlank("", "d"))
	assert.Equal("d", DefaultIfBlank(" ", "d"))
	assert.Equal("d", DefaultIfBlank("\r\n\t", "d"))
}
