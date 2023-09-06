package net

import (
	"encoding/binary"
	"fmt"
	"net"
	"regexp"
	"strconv"
	"unsafe"
)

var (
	nativeEndian binary.ByteOrder

	ipv4PrivateCIDRString = []string{
		"0.0.0.0/8", "10.0.0.0/8", "100.64.0.0/10", "127.0.0.0/8", "169.254.0.0/16",
		"172.16.0.0/12", "192.0.0.0/24", "192.0.2.0/24", "192.88.99.0/24", "192.168.0.0/16",
		"198.18.0.0/15", "198.51.100.0/24", "203.0.113.0/24", "224.0.0.0/4", "240.0.0.0/4", "255.255.255.255/32",
	}
	ipv6PrivateCIDRString = []string{
		"::1/128", "::/128", "64:ff9b::/96", "::ffff:0:0/96", "100::/64", "2001::/23",
		"2001::/32", "2001:2::/48", "2001:db8::/32", "2001:10::/28", "2002::/16", "fc00::/7", "fe80::/10",
		"2001:20::/28", "ff00::/8",
	}

	ipv4PrivateCIDR []*net.IPNet
	ipv6PrivateCIDR []*net.IPNet
)

func init() {
	if nativeEndian == nil {
		var x uint32 = 0x01020304
		if *(*byte)(unsafe.Pointer(&x)) == 0x01 {
			nativeEndian = binary.BigEndian
		} else {
			nativeEndian = binary.LittleEndian
		}
	}

	for _, v := range ipv4PrivateCIDRString {
		_, n, _ := net.ParseCIDR(v)
		ipv4PrivateCIDR = append(ipv4PrivateCIDR, n)
	}

	for _, v := range ipv6PrivateCIDRString {
		_, n, _ := net.ParseCIDR(v)
		ipv6PrivateCIDR = append(ipv6PrivateCIDR, n)
	}
}

// IsReservedIP reports whether ip is private.
// Support ipv4/ipv6, refer rfc6890.
// Return <0 ip is invalid, =0 ip is public, >0 ip is private.
func IsReservedIP(ip string) int {
	addr := net.ParseIP(ip)
	if addr == nil {
		return -1
	}

	if addr.IsLoopback() || addr.IsMulticast() || addr.IsLinkLocalMulticast() || addr.IsLinkLocalUnicast() {
		return 1
	}

	if addr.To4() != nil {
		for _, v := range ipv4PrivateCIDR {
			if v.Contains(addr) {
				return 1
			}
		}
	} else {
		for _, v := range ipv6PrivateCIDR {
			if v.Contains(addr) {
				return 1
			}
		}
	}
	return 0
}

// Swap16 swap a 16 bit value if aren't big endian
func Swap16(i uint16) uint16 {
	return (i&0xff00)>>8 | (i&0xff)<<8
}

// Swap32 swap a 32 bit value if aren't big endian
func Swap32(i uint32) uint32 {
	return (i&0xff000000)>>24 | (i&0xff0000)>>8 | (i&0xff00)<<8 | (i&0xff)<<24
}

// Htons convert uint16 from host byte order to network byte order.
// Network byte order is BigEndian.
func Htons(i uint16) uint16 {
	if GetNativeEndian() == binary.BigEndian {
		return i
	}
	// In big-endian mode, the high position is set to the low address.
	// 0x1234
	// 0x12 0x34
	return Swap16(i)
}

// Htonl convert uint32 from host byte order to network byte order.
// Network byte order is BigEndian.
func Htonl(i uint32) uint32 {
	if GetNativeEndian() == binary.BigEndian {
		return i
	}
	// In big-endian mode, the high position is set to the low address.
	// 0x12345678
	// 0x12 0x34 0x56 0x78
	return Swap32(i)
}

// Ntohs converts uint16 from network byte order to host byte order.
func Ntohs(i uint16) uint16 {
	if GetNativeEndian() == binary.BigEndian {
		return i
	}
	// In small-endian mode, the low position is placed at the low address.
	// 0x1234
	// 0x34 0x12
	return Swap16(i)
}

// Ntohl converts uint32 from network byte order to host byte order.
func Ntohl(i uint32) uint32 {
	if GetNativeEndian() == binary.BigEndian {
		return i
	}
	// In small-endian mode, the low position is placed at the low address.
	// 0x12345678
	// 0x78 0x56 0x34 0x12
	return Swap32(i)
}

// IPv4ToU32 convert ipv4(a.b.c.d) to uint32 in host byte order.
func IPv4ToU32(ip net.IP) uint32 {
	if ip == nil {
		return 0
	}
	a := uint32(ip[12])
	b := uint32(ip[13])
	c := uint32(ip[14])
	d := uint32(ip[15])
	return uint32(a<<24 | b<<16 | c<<8 | d)
}

// U32ToIPv4 converts uint32 to ipv4(a.b.c.d) in host byte order.
func U32ToIPv4(ip uint32) net.IP {
	a := byte((ip >> 24) & 0xFF)
	b := byte((ip >> 16) & 0xFF)
	c := byte((ip >> 8) & 0xFF)
	d := byte(ip & 0xFF)
	return net.IPv4(a, b, c, d)
}

// IPv4StrToU32 converts IPv4 string to uint32 in host byte order.
func IPv4StrToU32(s string) (ip uint32) {
	r := `^(\d{1,3})\.(\d{1,3})\.(\d{1,3})\.(\d{1,3})`
	reg, err := regexp.Compile(r)
	if err != nil {
		return
	}
	ips := reg.FindStringSubmatch(s)
	if ips == nil {
		return
	}

	ip1, _ := strconv.Atoi(ips[1])
	ip2, _ := strconv.Atoi(ips[2])
	ip3, _ := strconv.Atoi(ips[3])
	ip4, _ := strconv.Atoi(ips[4])

	if ip1 > 255 || ip2 > 255 || ip3 > 255 || ip4 > 255 {
		return
	}

	ip += uint32(ip1 * 0x1000000)
	ip += uint32(ip2 * 0x10000)
	ip += uint32(ip3 * 0x100)
	ip += uint32(ip4)
	return
}

// U32ToIPv4Str converts uint32 to IPv4 string in host byte order.
func U32ToIPv4Str(ip uint32) string {
	return fmt.Sprintf("%d.%d.%d.%d", ip>>24, ip<<8>>24, ip<<16>>24, ip<<24>>24)
}

// GetNativeEndian gets byte order for the current system.
func GetNativeEndian() binary.ByteOrder {
	return nativeEndian
}

// IsLittleEndian determines whether the host byte order is little endian.
func IsLittleEndian() bool {
	n := 0x1234
	return *(*byte)(unsafe.Pointer(&n)) == 0x34
}
