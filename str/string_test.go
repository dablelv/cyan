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
