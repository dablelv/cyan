package pointer

import (
	"reflect"
	"testing"

	"github.com/dablelv/cyan/internal/utest"
)

func TestPointer(t *testing.T) {
	type TestStruct struct {
		Field1 int
		Field2 string
	}

	tests := []struct {
		name  string
		input any
	}{
		{"Integer", 42},
		{"Nil Integer", (*int)(nil)},
		{"String", "hello"},
		{"Nil String", (*string)(nil)},
		{"Boolean", true},
		{"Nil Boolean", (*bool)(nil)},
		{"Float", 3.14},
		{"Nil Float", (*float64)(nil)},
		{"Struct", TestStruct{Field1: 1, Field2: "test"}},
		{"Nil Struct", (*TestStruct)(nil)},
		{"Slice", []int{1, 2, 3}},
		{"Nil Slice", (*[]int)(nil)},
		{"Map", map[string]int{"a": 1, "b": 2}},
		{"Nil Map", (*map[string]int)(nil)},
		{"Pointer", new(int)},
	}

	assert := utest.NewAssert(t, "TestPointer")
	for _, test := range tests {
		assert.Equal(true, func() bool {
			result := reflect.ValueOf(Pointer(test.input))
			if result.Kind() != reflect.Ptr {
				t.Logf("%q = got %v, expect pointer", test.name, result.Kind())
				return false
			}

			if !reflect.DeepEqual(result.Elem().Interface(), test.input) {
				t.Logf("%q = got %v, expect value %v", test.name, result.Elem().Interface(), test.input)
				return false
			}

			if !reflect.DeepEqual(Value(Pointer(test.input)), test.input) {
				t.Logf("%q = got %v, expect value %v", test.name, Value(Pointer(test.input)), test.input)
				return false
			}

			return true
		}())
	}
}

func TestNilValue(t *testing.T) {
	assert := utest.NewAssert(t, "TestNilValue")

	t.Run("Integer", func(t *testing.T) {
		assert.Equal(Value[int](nil), int(0))
	})

	t.Run("String", func(t *testing.T) {
		assert.Equal(Value[string](nil), "")
	})

	t.Run("Boolean", func(t *testing.T) {
		assert.Equal(Value[bool](nil), false)
	})

	t.Run("Float", func(t *testing.T) {
		assert.Equal(Value[float64](nil), float64(0.0))
	})

	t.Run("Struct", func(t *testing.T) {
		type TestStruct struct{}
		assert.Equal(Value[TestStruct](nil), TestStruct{})
	})

	t.Run("Slice", func(t *testing.T) {
		assert.Equal(Value[[]byte](nil), ([]byte)(nil))
	})

	t.Run("Map", func(t *testing.T) {
		assert.Equal(Value[map[int]int](nil), (map[int]int)(nil))
	})

	t.Run("Pointer", func(t *testing.T) {
		assert.Equal(Value[*string](nil), (*string)(nil))
	})
}
