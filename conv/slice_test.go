package conv

import (
	"testing"
	"time"

	"github.com/dablelv/cyan/internal/utest"
)

func TestToSlice(t *testing.T) {
	assert := utest.NewAssert(t, "TestToSlice")

	// to bool slice
	assert.Equal([]bool{true, false, true}, ToSlice[bool]([]bool{true, false, true}))
	assert.Equal([]bool{true, false, true}, ToSlice[bool]([3]bool{true, false, true}))
	assert.Equal([]bool{true, false, true}, ToSlice[bool]([]int{1, 0, 1}))
	assert.Equal([]bool{true, false, true}, ToSlice[bool]([3]int{1, 0, 1}))
	assert.Equal([]bool{true, false, true}, ToSlice[bool]([]string{"true", "false", "true"}))
	assert.Equal([]bool{true, false, true}, ToSlice[bool]([3]string{"true", "false", "true"}))
	assert.Equal([]bool(nil), ToSlice[bool]([]string(nil)))
	assert.Equal([]bool(nil), ToSlice[bool](nil))

	bs, err := ToSliceE[bool]([]string{"foo"})
	assert.IsNotNil(err)
	assert.IsNil(bs)
	bs, err = ToSliceE[bool]("foo")
	assert.IsNotNil(err)
	assert.IsNil(bs)

	// to int slice
	assert.Equal([]int{1, 2, 3}, ToSlice[int]([]int{1, 2, 3}))
	assert.Equal([]int{1, 2, 3}, ToSlice[int]([3]int{1, 2, 3}))
	assert.Equal([]int{1}, ToSlice[int](1))
	assert.Equal([]int{1, 2, 3}, ToSlice[int]([]string{"1", "2", "3"}))
	assert.Equal([]int{1, 2, 3}, ToSlice[int]([3]string{"1", "2", "3"}))
	assert.Equal([]int(nil), ToSlice[int]([]string(nil)))
	assert.Equal([]int(nil), ToSlice[int](nil))

	ints, err := ToSliceE[int](complex(3, -5))
	assert.IsNotNil(err)
	assert.IsNil(ints)

	// to uint slice
	assert.Equal([]uint{1, 2, 3}, ToSlice[uint]([]int{1, 2, 3}))
	assert.Equal([]uint{1, 2, 3}, ToSlice[uint]([3]int{1, 2, 3}))
	assert.Equal([]uint{1}, ToSlice[uint](1))
	assert.Equal([]uint{1, 2, 3}, ToSlice[uint]([]string{"1", "2", "3"}))
	assert.Equal([]uint{1, 2, 3}, ToSlice[uint]([3]string{"1", "2", "3"}))
	assert.Equal([]uint(nil), ToSlice[uint]([]string(nil)))
	assert.Equal([]uint(nil), ToSlice[uint](nil))

	uints, err := ToSliceE[uint]("foo")
	assert.IsNotNil(err)
	assert.IsNil(uints)

	// to float64 slice
	assert.Equal([]float64{1, 2, 3}, ToSlice[float64]([]int{1, 2, 3}))
	assert.Equal([]float64{1, 2, 3}, ToSlice[float64]([3]int{1, 2, 3}))
	assert.Equal([]float64{1}, ToSlice[float64](1))
	assert.Equal([]float64{1, 2, 3}, ToSlice[float64]([]string{"1", "2", "3"}))
	assert.Equal([]float64{1, 2, 3}, ToSlice[float64]([3]string{"1", "2", "3"}))
	assert.Equal([]float64(nil), ToSlice[float64]([]string(nil)))
	assert.Equal([]float64(nil), ToSlice[float64](nil))

	f64s, err := ToSliceE[float64]("foo")
	assert.IsNotNil(err)
	assert.IsNil(f64s)

	// to duration slice
	assert.Equal([]time.Duration{1, 2, 3}, ToSlice[time.Duration]([]int{1, 2, 3}))
	assert.Equal([]time.Duration{1, 2, 3}, ToSlice[time.Duration]([3]int{1, 2, 3}))
	assert.Equal([]time.Duration{1}, ToSlice[time.Duration](1))
	assert.Equal([]time.Duration{1, 2, 3}, ToSlice[time.Duration]([]string{"1", "2", "3"}))
	assert.Equal([]time.Duration{1, 2, 3}, ToSlice[time.Duration]([3]string{"1", "2", "3"}))
	assert.Equal([]time.Duration(nil), ToSlice[time.Duration]([]string(nil)))
	assert.Equal([]time.Duration(nil), ToSlice[time.Duration](nil))

	ds, err := ToSliceE[time.Duration]("foo")
	assert.IsNotNil(err)
	assert.IsNil(ds)

	// to string slice
	assert.Equal([]string{"a", "b", "c"}, ToSlice[string]([]string{"a", "b", "c"}))
	assert.Equal([]string{"a", "b", "c"}, ToSlice[string]([3]string{"a", "b", "c"}))
	assert.Equal([]string{"a", "b", "c"}, ToSlice[string]("a b c"))
	assert.Equal([]string{"a"}, ToSlice[string]("a"))
	assert.Equal([]string{"1", "2", "3"}, ToSlice[string]([]int{1, 2, 3}))
	assert.Equal([]string{"1.1", "2.2", "3.3"}, ToSlice[string]([]float64{1.1, 2.2, 3.3}))
	assert.Equal([]string(nil), ToSlice[string]([]string(nil)))
	assert.Equal([]string(nil), ToSlice[string](nil))

	ss, err := ToSliceE[string](complex(3, -5))
	assert.IsNotNil(err)
	assert.IsNil(ss)
}

func TestJsonToSliceE(t *testing.T) {
	{
		assert := utest.NewAssert(t, "success")

		r1, err := JsonToSliceE[[]int]([]byte("[1,2,3]"))
		assert.IsNil(err)
		assert.Equal(r1, []int{1, 2, 3})

		r2, err := JsonToSliceE[[]float64]([]byte("[1.1,2.2,3.3]"))
		assert.IsNil(err)
		assert.Equal(r2, []float64{1.1, 2.2, 3.3})

		r3, err := JsonToSliceE[[]string]([]byte(`["foo","bar","baz"]`))
		assert.IsNil(err)
		assert.Equal(r3, []string{"foo", "bar", "baz"})
	}
	{
		assert := utest.NewAssert(t, "success but zero value")
		r, err := JsonToSliceE[[]int]([]byte(`[1,2,3,"bar"]`))
		assert.IsNotNil(err)
		assert.Equal(r, []int{1, 2, 3, 0})
	}
	{
		assert := utest.NewAssert(t, "error:invalid json")
		r, err := JsonToSliceE[[]int]([]byte(`[1,2,3,"bar"`))
		assert.IsNotNil(err)
		assert.IsNil(r)
	}
}

func TestSplitStrToSlice(t *testing.T) {
	assert := utest.NewAssert(t, "TestSplitStrToSlice")

	// to int slice
	assert.Equal([]int{1, 2, 3}, SplitStrToSlice[int]("1,2,3", ","))
	_, err := SplitStrToSliceE[int]("1,2,foo", ",")
	assert.IsNotNil(err)

	// to uint slice
	assert.Equal([]uint{1, 2, 3}, SplitStrToSlice[uint]("1,2,3", ","))
	_, err = SplitStrToSliceE[int]("1,2,foo", ",")
	assert.IsNotNil(err)

	// to float64 slice
	assert.Equal([]float64{1, 2, 3}, SplitStrToSlice[float64]("1,2,3", ","))
	_, err = SplitStrToSliceE[float64]("1,2,foo", ",")
	assert.IsNotNil(err)

	// to bool slice
	assert.Equal([]bool{true, false, true}, SplitStrToSlice[bool]("1,0,1", ","))
	assert.Equal([]bool{true, false, true}, SplitStrToSlice[bool]("true,false,true", ","))
	_, err = SplitStrToSliceE[bool]("1,0,2", ",")
	assert.IsNotNil(err)
}

func TestMapToSlice(t *testing.T) {
	assert := utest.NewAssert(t, "TestMapToSlice")

	ks, vs := MapToSlice(map[string]string{"1": "1", "2": "2", "3": "3"})
	assert.Equal(3, len(ks.([]string)))
	assert.Equal(3, len(vs.([]string)))

	ks, vs = MapToSlice(map[int]int{1: 1, 2: 2, 3: 3})
	assert.Equal(3, len(ks.([]int)))
	assert.Equal(3, len(vs.([]int)))

	// empty map[int]int to slice
	ks, vs = MapToSlice(map[int]int{})
	assert.Equal(0, len(ks.([]int)))
	assert.Equal(0, len(vs.([]int)))

	// not map failed
	_, _, err := MapToSliceE(1)
	assert.IsNotNil(err)
}
