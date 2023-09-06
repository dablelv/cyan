package os

import (
	"os"
	"runtime"
	"strings"
	"testing"

	"github.com/dablelv/cyan/internal"
)

func TestIsWin(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsWin")

	is := IsWin()
	if runtime.GOOS == "windows" {
		assert.Equal(true, is)
	} else {
		assert.Equal(false, is)
	}
}

func TestIsMac(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsMac")

	is := IsMac()
	if runtime.GOOS == "darwin" {
		assert.Equal(true, is)
	} else {
		assert.Equal(false, is)
	}
}

func TestIsLinux(t *testing.T) {
	assert := internal.NewAssert(t, "IsLinux")

	is := IsLinux()
	if runtime.GOOS == "linux" {
		assert.Equal(true, is)
	} else {
		assert.Equal(false, is)
	}
}

func TestIsSupportColor(t *testing.T) {
	assert := internal.NewAssert(t, "IsSupportColor")

	support := IsSupportColor()
	envTerm := os.Getenv("TERM")
	if strings.Contains(envTerm, "xterm") || strings.Contains(envTerm, "screen") {
		assert.Equal(true, support)
	}
	if os.Getenv("ConEmuANSI") == "ON" {
		assert.Equal(true, support)
	}
	if os.Getenv("ANSICON") != "" {
		assert.Equal(true, support)
	}
}

func TestIsSupport256Color(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsSupport256Color")

	support := IsSupport256Color()
	if strings.Contains(os.Getenv("TERM"), "256color") {
		assert.Equal(true, support)
	} else {
		assert.Equal(false, support)
	}
}

func TestIsSupportTrueColor(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsSupportTrueColor")

	support := IsSupportTrueColor()
	if strings.Contains(os.Getenv("COLORTERM"), "truecolor") {
		assert.Equal(true, support)
	} else {
		assert.Equal(false, support)
	}
}
