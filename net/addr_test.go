package net

import (
	"testing"

	"github.com/dablelv/cyan/internal/utest"
)

func TestLocalIPv4(t *testing.T) {
	assert := utest.NewAssert(t, "LocalIPv4")
	ip := LocalIPv4()
	assert.Greater(len(ip), 0)
	t.Logf("LocalIPv4 end, ip is %v", ip.String())
}
