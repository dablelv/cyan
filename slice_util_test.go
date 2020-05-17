package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseIntSlice(t *testing.T) {
	intSlice := []int{1, 2, 3}
	assert.Equal(t, []int{3, 2, 1}, ReverseIntSlice(intSlice))
}
