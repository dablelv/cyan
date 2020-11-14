package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseIntSlice(t *testing.T) {
	intSlice := []int{1, 2, 3}
	assert.Equal(t, []int{3, 2, 1}, ReverseIntSlice(intSlice))
}

func TestSumSlice(t *testing.T) {
	f32Slice := []float32{1.1, 2.2, 3.3}
	f64 := SumSlice(f32Slice)
	assert.Equal(t, float32(6.6), float32(f64))
}
