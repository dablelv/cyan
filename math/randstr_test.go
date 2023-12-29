package math

import (
	"testing"

	"github.com/dablelv/cyan/internal"
)

func TestRandStr(t *testing.T) {
	assert := internal.NewAssert(t, "TestRandStr")
	s := RandStr(4)
	assert.Equal(4, len(s))
}

func TestRandLowerStr(t *testing.T) {
	assert := internal.NewAssert(t, "TestRandLowerStr")
	s := RandLowerStr(4)
	assert.Equal(4, len(s))
}

func TestRandUpperStr(t *testing.T) {
	assert := internal.NewAssert(t, "TestRandUpperStr")
	s := RandUpperStr(4)
	assert.Equal(4, len(s))
}

func TestRandAlphanumeric(t *testing.T) {
	assert := internal.NewAssert(t, "TestRandAlphanumeric")
	s := RandAlphanumeric(4)
	assert.Equal(4, len(s))
}

func TestRandLowerAlphanumeric(t *testing.T) {
	assert := internal.NewAssert(t, "TestRandLowerAlphanumeric")
	s := RandLowerAlphanumeric(4)
	assert.Equal(4, len(s))
}
