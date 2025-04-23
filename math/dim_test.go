package math

import (
	"testing"

	"github.com/dablelv/cyan/internal/utest"
)

func TestMin(t *testing.T) {
	assert := utest.NewAssert(t, "TestMin")

	assert.Equal(1, Min(1, 2))
	assert.Equal(1.1, Min(2.2, 1.1))
	assert.Equal("abc", Min("abc", "cba"))
}

func TestMax(t *testing.T) {
	assert := utest.NewAssert(t, "TestMax")

	assert.Equal(2, Max(1, 2))
	assert.Equal(2.2, Max(2.2, 1.1))
	assert.Equal("cba", Max("abc", "cba"))
}
