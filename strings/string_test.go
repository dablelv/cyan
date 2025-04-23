package strings

import (
	"reflect"
	"testing"

	"github.com/dablelv/cyan/internal/utest"
)

func TestHead(t *testing.T) {
	s := "123456"
	{
		assert := utest.NewAssert(t, "Haead-1")
		r := Head(s, -1)
		assert.Equal(r, "")
	}
	{
		assert := utest.NewAssert(t, "Haead0")
		r := Head(s, -1)
		assert.Equal(r, "")
	}
	{
		assert := utest.NewAssert(t, "TestHeadNormal")
		r := Head(s, len(s)-1)
		assert.Equal(r, "12345")
	}
	{
		assert := utest.NewAssert(t, "TestHeadOverflow")
		r := Head(s, len(s)+1)
		assert.Equal(r, s)
	}
}

func TestHeadRunes(t *testing.T) {
	{
		assert := utest.NewAssert(t, "empty string")
		r := HeadRunes("", 1)
		assert.Equal(r, "")
	}
	{
		assert := utest.NewAssert(t, "head english string")
		r := HeadRunes("english", 1)
		assert.Equal(r, "e")
	}
	{
		assert := utest.NewAssert(t, "head chinese string")
		r := HeadRunes("中文", 1)
		assert.Equal(r, "中")
	}
	{
		assert := utest.NewAssert(t, "head chinese string overflow")
		r := HeadRunes("中文", 3)
		assert.Equal(r, "中文")
	}
}

func TestTail(t *testing.T) {
	s := "123456"
	{
		assert := utest.NewAssert(t, "Tail-1")
		r := Tail(s, -1)
		assert.Equal(r, "")
	}
	{
		assert := utest.NewAssert(t, "Tail0")
		r := Tail(s, 0)
		assert.Equal(r, "")
	}
	{
		assert := utest.NewAssert(t, "TailNormal")
		r := Tail(s, 3)
		assert.Equal(r, "456")
	}
	{
		assert := utest.NewAssert(t, "TailOverflow")
		r := Tail(s, len(s)+1)
		assert.Equal(r, s)
	}
}

func TestTailRunes(t *testing.T) {
	{
		assert := utest.NewAssert(t, "empty string")
		r := TailRunes("", 1)
		assert.Equal(r, "")
	}
	{
		assert := utest.NewAssert(t, "tail english string")
		r := TailRunes("english", 1)
		assert.Equal(r, "h")
	}
	{
		assert := utest.NewAssert(t, "tail chinese string")
		r := TailRunes("中文", 1)
		assert.Equal(r, "文")
	}
	{
		assert := utest.NewAssert(t, "tail chinese string overflow")
		r := TailRunes("中文", 3)
		assert.Equal(r, "中文")
	}
}

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
	assert := utest.NewAssert(t, "TestSplitSeps")

	assert.Equal([]string{"foo", "bar", "baz"}, SplitSeps("foo,bar,baz", ","))
	assert.Equal([]string{"foo", "bar", "baz"}, SplitSeps("foo,bar|baz", ",", "|"))
	assert.Equal([]string{"foo", "bar", "baz", "qux"}, SplitSeps("foo,bar|baz qux", ",", "|", " "))
	assert.Equal([]string{"foo", "bar", "baz", "qux"}, SplitSeps("foo,bar|bazSEPqux", ",", "|", "SEP"))
	assert.Equal([]string{"foo,bar|baz"}, SplitSeps("foo,bar|baz"))
	assert.Equal([]string{" ", "x", "y", "z"}, SplitSeps(" xyz", ""))
}

func TestJoinNonEmptyStrs(t *testing.T) {
	assert := utest.NewAssert(t, "TestJoinNonEmptyStrs")

	got := JoinNonEmptyStrs(",", []string{"foo", "bar", "baz"}...)
	assert.Equal("foo,bar,baz", got)

	got = JoinNonEmptyStrs(",", []string{"foo", "bar", "baz"}...)
	assert.Equal("foo,bar,baz", got)
}

func TestJoin(t *testing.T) {
	assert := utest.NewAssert(t, "TestJoin")

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
	assert := utest.NewAssert(t, "TestReverse")
	assert.Equal("foo", Reverse("oof"))
	assert.Equal("123", Reverse("321"))
}

func TestAlphanumericNum(t *testing.T) {
	assert := utest.NewAssert(t, "TestAlphanumericNum")

	assert.Equal(3, AlphanumericNum("108条梁山好汉"))
	assert.Equal(3, AlphanumericNum("一百条梁山man"))
	assert.Equal(6, AlphanumericNum("108条梁山man"))
}

func TestAlphanumericNumV2(t *testing.T) {
	assert := utest.NewAssert(t, "TestAlphanumericNumV2")

	assert.Equal(3, AlphanumericNumV2("108条梁山好汉"))
	assert.Equal(3, AlphanumericNumV2("一百条梁山man"))
	assert.Equal(6, AlphanumericNumV2("108条梁山man"))
}

func TestAlphanumericNumRegExp(t *testing.T) {
	assert := utest.NewAssert(t, "TestAlphanumericNumRegExp")

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
	assert := utest.NewAssert(t, "TestClearWhiteSpace")

	assert.Equal("abc", ClearWhiteSpace("a b c"))
	assert.Equal("", ClearWhiteSpace("   "))
	assert.Equal("abcd", ClearWhiteSpace("a\rb\nc\td"))
}

func TestIndexOfDiff(t *testing.T) {
	assert := utest.NewAssert(t, "TestIndexOfDiff")

	assert.Equal(-1, IndexOfDiff("abc", "abc"))
	assert.Equal(0, IndexOfDiff("abc", "d"))
	assert.Equal(1, IndexOfDiff("abc", "a_c"))
}

func TestIndexOffset(t *testing.T) {
	assert := utest.NewAssert(t, "TestIndexOffset")

	assert.Equal(3, IndexOffset("abcde", "d", 2))
	assert.Equal(-1, IndexOffset("abcde", "d", 5))
	assert.Equal(3, IndexOffset("abcde", "d", -1))
	assert.Equal(-1, IndexOffset("abcde", "", 0))
	assert.Equal(-1, IndexOffset("abcde", "f", 0))
}

func TestDefault(t *testing.T) {
	assert := utest.NewAssert(t, "TestDefault")

	assert.Equal("abc", Default("abc", "d"))
	assert.Equal("d", Default("", "d"))
}

func TestDefaultIfBlank(t *testing.T) {
	assert := utest.NewAssert(t, "TestDefaultIfBlank")

	assert.Equal("abc", DefaultIfBlank("abc", "d"))
	assert.Equal("d", DefaultIfBlank("", "d"))
	assert.Equal("d", DefaultIfBlank(" ", "d"))
	assert.Equal("d", DefaultIfBlank("\r\n\t", "d"))
}

func TestBinToHex(t *testing.T) {
	bin := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}
	expected := string("0102030405060708090a0b0c0d0e0f")

	assert := utest.NewAssert(t, "TestBinToHex")
	assert.Equal(BinToHex(bin), expected)
}

func TestStringToConstBytes(t *testing.T) {
	bytes := []byte{'1', '2', '3'}

	assert := utest.NewAssert(t, "TestStringToConstBytes")
	assert.Equal(StringToConstBytes(string(bytes)), bytes)
}

func TestSplitChunk(t *testing.T) {
	type args struct {
		input string
		sep   string
		size  int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty string",
			args: args{
				input: "",
				sep:   "",
				size:  0,
			},
			want: "",
		},
		{
			name: "separator string is empty",
			args: args{
				input: "123456",
				sep:   "",
				size:  2,
			},
			want: "123456",
		},
		{
			name: "size is 0",
			args: args{
				input: "123456",
				sep:   ",",
				size:  0,
			},
			want: "123456",
		},
		{
			name: "size is greater than string length",
			args: args{
				input: "123456",
				sep:   ",",
				size:  10,
			},
			want: "123456",
		},
		{
			name: "size is equal to string length",
			args: args{
				input: "123456",
				sep:   ",",
				size:  6,
			},
			want: "123456",
		},
		{
			name: "size is less than string length",
			args: args{
				input: "1234567",
				sep:   ",",
				size:  2,
			},
			want: "12,34,56,7",
		},
		{
			name: "just split",
			args: args{
				input: "12345678",
				sep:   ",",
				size:  2,
			},
			want: "12,34,56,78",
		},
		{
			name: "split chinese string",
			args: args{
				input: "你的名字",
				sep:   ",",
				size:  2,
			},
			want: "你的,名字",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitChunk(tt.args.input, tt.args.sep, tt.args.size); got != tt.want {
				t.Errorf("SplitChunk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitChunkWithSpace(t *testing.T) {
	type args struct {
		input string
		size  int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "just split",
			args: args{
				input: "12345678",
				size:  2,
			},
			want: "12 34 56 78",
		},
		{
			name: "split chinese string",
			args: args{
				input: "你的名字",
				size:  2,
			},
			want: "你的 名字",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitChunkWithSpace(tt.args.input, tt.args.size); got != tt.want {
				t.Errorf("SplitChunkWithSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}
