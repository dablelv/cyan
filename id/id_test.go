// A viable unique ID generation function.
package id

import (
	"net"
	"testing"
	stdtime "time"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/dablelv/cyan/internal/utest"
	"github.com/dablelv/cyan/time"
)

func TestNewId(t *testing.T) {
	assert := utest.NewAssert(t, "TestNewId")

	// Reset global variables.
	ip := net.IPv4(192, 168, 1, 1)
	pid := uint16(1234)
	seqNo := uint32(0x1)
	tm := stdtime.Date(2025, 2, 8, 12, 34, 56, 0, time.LocAsiaShanghai)

	p := gomonkey.NewPatches()
	p.ApplyGlobalVar(&localIP, ip)
	p.ApplyGlobalVar(&processID, pid)
	p.ApplyGlobalVar(&sequenceNum, seqNo)
	p.ApplyFunc(stdtime.Now, func() stdtime.Time {
		return tm
	})
	defer p.Reset()

	id := NewId()
	assert.Equal(id, "67a6def0c0a8010104d2000000020000")
}
