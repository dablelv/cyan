package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerGTVer(t *testing.T) {
	res0, err0 := VerGTVer("1.0.5", "1.0.4")
	res1, err1 := VerGTVer("1.0.5", "2.0.4")
	assert.NoError(t, err0)
	assert.Equal(t, true, res0)
	assert.NoError(t, err1)
	assert.Equal(t, false, res1)
}

func TestVerLTVer(t *testing.T) {
	res0, err0 := VerLTVer("1.0.4","1.0.5")
	res1, err1 := VerLTVer("1.0.5", "1.0.4")
	assert.NoError(t, err0)
	assert.Equal(t, true, res0)
	assert.NoError(t, err1)
	assert.Equal(t, false, res1)
}