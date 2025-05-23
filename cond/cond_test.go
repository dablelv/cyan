package cond

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/dablelv/cyan/internal/utest"
)

func TestBool(t *testing.T) {
	assert := utest.NewAssert(t, "TestBool")

	// bool
	assert.Equal(false, Bool(false))
	assert.Equal(true, Bool(true))

	// integer
	assert.Equal(false, Bool(0))
	assert.Equal(true, Bool(1))

	// float
	assert.Equal(false, Bool(0.0))
	assert.Equal(true, Bool(0.1))

	// string
	assert.Equal(false, Bool(""))
	assert.Equal(true, Bool(" "))
	assert.Equal(true, Bool("0"))

	// slice
	var nums []int
	assert.Equal(false, Bool(nums))
	nums = []int{0, 1}
	assert.Equal(true, Bool(nums))

	// map
	assert.Equal(false, Bool(map[string]string{}))
	assert.Equal(true, Bool(map[string]string{"a": "a"}))

	// channel
	var ch chan int
	assert.Equal(false, Bool(ch))
	ch = make(chan int)
	assert.Equal(true, Bool(ch))

	//  interface
	var err error
	assert.Equal(false, Bool(err))
	err = errors.New("error message")
	assert.Equal(true, Bool(err))

	// struct
	assert.Equal(false, Bool(struct{}{}))
	assert.Equal(true, Bool(time.Now()))

	// struct pointer
	ts := struct{}{}
	assert.Equal(false, Bool(ts))
	assert.Equal(true, Bool(&ts))
}

func TestAnd(t *testing.T) {
	assert := utest.NewAssert(t, "TestAnd")
	assert.Equal(false, And(0, 1))
	assert.Equal(false, And(0, ""))
	assert.Equal(false, And(0, "0"))
	assert.Equal(true, And(1, "0"))
}

func TestOr(t *testing.T) {
	assert := utest.NewAssert(t, "TestOr")
	assert.Equal(false, Or(0, ""))
	assert.Equal(true, Or(0, 1))
	assert.Equal(true, Or(0, "0"))
	assert.Equal(true, Or(1, "0"))
}

func TestXor(t *testing.T) {
	assert := utest.NewAssert(t, "TestOr")
	assert.Equal(false, Xor(0, 0))
	assert.Equal(true, Xor(0, 1))
	assert.Equal(true, Xor(1, 0))
	assert.Equal(false, Xor(1, 1))
}

func TestNor(t *testing.T) {
	assert := utest.NewAssert(t, "TestNor")
	assert.Equal(true, Nor(0, 0))
	assert.Equal(false, Nor(0, 1))
	assert.Equal(false, Nor(1, 0))
	assert.Equal(false, Nor(1, 1))
}

func TestXnor(t *testing.T) {
	assert := utest.NewAssert(t, "TestXnor")
	assert.Equal(true, Xnor(0, 0))
	assert.Equal(false, Xnor(0, 1))
	assert.Equal(false, Xnor(1, 0))
	assert.Equal(true, Xnor(1, 1))
}

func TestNand(t *testing.T) {
	assert := utest.NewAssert(t, "TestNand")
	assert.Equal(true, Nand(0, 0))
	assert.Equal(true, Nand(0, 1))
	assert.Equal(true, Nand(1, 0))
	assert.Equal(false, Nand(1, 1))
}

func TestIf(t *testing.T) {
	type args struct {
		condition bool
		trueVal   any
		falseVal  any
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			"return true",
			args{
				1 > 0,
				true,
				false,
			},
			true,
		},
		{
			"return false",
			args{
				1 < 0,
				true,
				false,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IfElse(tt.args.condition, tt.args.trueVal, tt.args.falseVal); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("If() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIfElseF(t *testing.T) {
	type TestStruct struct {
		ID   int
		Name string
	}

	{
		assert := utest.NewAssert(t, "true with string")
		r := IfElseF(true, func() string { return "true" }, func() string { return "false" })
		assert.Equal(r, "true")
	}
	{
		assert := utest.NewAssert(t, "false with string")
		r := IfElseF(false, func() string { return "true" }, func() string { return "false" })
		assert.Equal(r, "false")
	}
	{
		assert := utest.NewAssert(t, "true with struct")
		r := IfElseF(true, func() TestStruct { return TestStruct{ID: 1, Name: "a"} }, func() TestStruct { return TestStruct{ID: 2, Name: "b"} })
		assert.Equal(r, TestStruct{ID: 1, Name: "a"})
	}
	{
		assert := utest.NewAssert(t, "false with struct")
		r := IfElseF(false, func() TestStruct { return TestStruct{ID: 1, Name: "a"} }, func() TestStruct { return TestStruct{ID: 2, Name: "b"} })
		assert.Equal(r, TestStruct{ID: 2, Name: "b"})
	}
}
