package net

import (
	"encoding/binary"
	"net"
	"testing"

	"github.com/dablelv/cyan/internal"
)

func TestIsReservedIP(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"loopback", args{"127.0.0.1"}, 1},
		{"multicast", args{"10.255.255.255"}, 1},
		{"A private", args{"10.0.0.1"}, 1},
		{"B private", args{"172.16.0.1"}, 1},
		{"C private", args{"192.168.0.1"}, 1},
		{"D private", args{"224.0.0.1"}, 1},
		{"E private", args{"240.0.0.1"}, 1},
		{"ipv4 public", args{"1.0.0.1"}, 0},
		{"invalid ipv4", args{"192.168.0.256"}, -1},
		{"ipv6 loopback", args{"::1"}, 1},
		{"ipv6 private", args{"FD00::1"}, 1},
		{"ipv6 public", args{"FE00::1"}, 0},
		{"ipv6 invalid", args{"ff06:::"}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsReservedIP(tt.args.ip); got != tt.want {
				t.Errorf("IsReservedIP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHtons(t *testing.T) {
	assert := internal.NewAssert(t, "TestHtons")

	got := Htons(0x1234)
	if GetNativeEndian() == binary.BigEndian {
		assert.Equal(uint16(0x1234), got)
	} else {
		assert.Equal(uint16(0x3412), got)
	}
}

func TestHtonl(t *testing.T) {
	assert := internal.NewAssert(t, "TestHtonl")

	got := Htonl(0x12345678)
	if GetNativeEndian() == binary.BigEndian {
		assert.Equal(uint32(0x12345678), got)
	} else {
		assert.Equal(uint32(0x78563412), got)
	}
}

func TestNtohs(t *testing.T) {
	assert := internal.NewAssert(t, "TestNtohs")

	got := Ntohs(0x3412)
	if GetNativeEndian() == binary.BigEndian {
		assert.Equal(uint16(0x3412), got)
	} else {
		assert.Equal(uint16(0x1234), got)
	}
}

func TestNtohl(t *testing.T) {
	assert := internal.NewAssert(t, "TestNtohl")

	got := Ntohl(0x78563412)
	if GetNativeEndian() == binary.BigEndian {
		assert.Equal(uint32(0x78563412), got)
	} else {
		assert.Equal(uint32(0x12345678), got)
	}
}

func TestIPv4ToU32(t *testing.T) {
	assert := internal.NewAssert(t, "TestIPv4ToU32")

	got := IPv4ToU32(net.IPv4(127, 0, 0, 1))
	assert.Equal(uint32(0x7f000001), got)
	got = IPv4ToU32(nil)
	assert.Equal(uint32(0x0), got)
}

func TestU32ToIPv4(t *testing.T) {
	assert := internal.NewAssert(t, "TestU32ToIPv4")

	got := U32ToIPv4(0x7f000001)
	assert.Equal(net.IPv4(127, 0, 0, 1), got)
}

func TestIPv4StrToU32(t *testing.T) {
	assert := internal.NewAssert(t, "TestIPv4StrToU32")

	got := IPv4StrToU32("127.0.0.1")
	assert.Equal(uint32(0x7f000001), got)
	got = IPv4StrToU32("127.0.0.256")
	assert.Equal(uint32(0x0), got)
	got = IPv4StrToU32("")
	assert.Equal(uint32(0x0), got)
}

func TestU32ToIPv4Str(t *testing.T) {
	assert := internal.NewAssert(t, "TestU32ToIPv4Str")

	got := U32ToIPv4Str(0x7f000001)
	assert.Equal("127.0.0.1", got)
}

func TestIsLittleEndian(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsLittleEndian")

	got := IsLittleEndian()
	if GetNativeEndian() == binary.BigEndian {
		assert.Equal(false, got)
	} else {
		assert.Equal(true, got)
	}
}
