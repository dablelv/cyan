package math

import (
	"testing"

	"github.com/dablelv/cyan/internal"
)

func TestRandIntRange(t *testing.T) {
	assert := internal.NewAssert(t, "TestRandIntRange")
	type args struct {
		min int
		max int
	}
	tests := []struct {
		name string
		args args
	}{
		{"only get random number 0", args{0, 1}},
		{"get random number in range 0-2", args{0, 3}},
		{"get random number in range 10-100", args{10, 101}},
	}
	for _, tt := range tests {
		got := RandIntRange(tt.args.min, tt.args.max)
		valid := got >= tt.args.min && got < tt.args.max
		assert.Equal(true, valid)
	}
	got := RandIntRange(1, 1)
	assert.Equal(0, got)
}

func TestRandIntn(t *testing.T) {
	assert := internal.NewAssert(t, "TestRandIntn")
	got := RandIntn(0)
	assert.Equal(0, got)
}

func TestRandInts(t *testing.T) {
	assert := internal.NewAssert(t, "TestRandIntSlice")
	s := RandInts(3, 1)
	assert.Equal([]int{0, 0, 0}, s)
}

func TestRandBytes(t *testing.T) {
	assert := internal.NewAssert(t, "TestRandByteSlice")
	s := RandBytes(8)
	assert.NotEqual([]int{0, 0, 0, 0, 0, 0, 0, 0}, s)
}
