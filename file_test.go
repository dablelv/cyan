package util

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListDir(t *testing.T) {
	filenames := ListDir(".")
	haveFile := len(filenames) > 0
	assert.Equal(t, true, haveFile, "TestListDir")
}

func TestReadLines(t *testing.T) {
	l := 64*1024 + 1
	veryLongString := RandStr(uint64(l)) // MaxScanTokenSize = 64 * 1024 = 65536 < 70000

	// if line length greater than bufio.MaxScanTokenSize(64 * 1024) need to create a new longer buffer
	if l > bufio.MaxScanTokenSize {
		newScanner := bufio.NewScanner(strings.NewReader(veryLongString))
		newScanner.Buffer([]byte{}, bufio.MaxScanTokenSize*2)
		for newScanner.Scan() {
			t.Log("bufio.Scan() success")
		}
	} else {
		scanner := bufio.NewScanner(strings.NewReader(veryLongString))
		for scanner.Scan() {
			t.Log(scanner.Text())
		}
		if scanner.Err() == bufio.ErrTooLong {
			t.Error(bufio.ErrTooLong)
		}
	}
}

func TestReadLinesV2(t *testing.T) {
	lines, _, _ := ReadLinesV2("./tmp/name")
	t.Logf("%v", lines)
}
