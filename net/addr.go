package net

import (
	"fmt"
	"net"
)

var localIPv4 = preferredIPv4()

// LocalIPv4 returns the cached local IPv4 address of the machine.
// Note: If the machine's IP address is changed, the process needs to be restarted
// for the new IP to be recognized.
func LocalIPv4() net.IP {
	return localIPv4
}

// GetLocalIpv4s get all local ipv4 addresses.
func GetLocalIPv4s() ([]net.IP, error) {
	var ips []net.IP
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			ips = append(ips, ipNet.IP)
		}
	}

	return ips, nil
}

// GetOutboundIPv4 get the reachable IPv4 address on the external network (via UDP connection).
func GetOutboundIPv4() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP, nil
}

// GetPreferredIPv4 get the preferred IPv4 address (excluding loopback and link-local)
func GetPreferredIPv4() (net.IP, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			// Skip the unenabled and loopback interfaces.
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)
			if ok && ipNet.IP.To4() != nil && !ipNet.IP.IsLinkLocalUnicast() {
				return ipNet.IP, nil
			}
		}
	}
	return nil, fmt.Errorf("no suitable IPv4 address found")
}

// preferredIPv4 get the preferred IPv4 address.
// If no suitable IPv4 address found, return loopback address 127.0.0.1.
func preferredIPv4() net.IP {
	ipv4, err := GetPreferredIPv4()
	if err == nil {
		return ipv4
	}
	// May has loopback interface only.
	return net.ParseIP("127.0.0.1")
}
