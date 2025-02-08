// A viable globally unique ID generation method.
// It is recommended to use industry-proven ID generation schemes, such as Universally Unique identifiers (UUID), Snowflake algorithm, etc.,
// rather than implement it yourself.
package id

import (
	"encoding/hex"
	"net"
	"os"
	"sync/atomic"
	"time"
)

var (
	localIP     net.IP
	processID   uint16
	sequenceNum uint32
)

func init() {
	// Get the local IP address.
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				localIP = ipnet.IP.To4()
				break
			}
		}
	}
	if localIP == nil {
		panic("can't get local IP address")
	}

	// Get process ID.
	processID = uint16(os.Getpid())
}

// NewId generates a globally unique ID in thread safety mode.
// The generation rule is
// [4 bytes timestamp] + [4 bytes local IP] + [2 bytes process ID] + [4 bytes auto-incremented sequence number] + [2 bytes reserved, set to 0]
// All numerical segments use the BigEndian order.
func NewId() string {
	ts := uint32(time.Now().Unix())

	// Auto-incremented sequence number.
	// If overflow occurs, seq will become 0.
	seq := atomic.AddUint32(&sequenceNum, 1)

	// Assemble the ID.
	idBytes := make([]byte, 16)
	idBytes[0] = byte(ts >> 24)
	idBytes[1] = byte(ts >> 16)
	idBytes[2] = byte(ts >> 8)
	idBytes[3] = byte(ts)

	copy(idBytes[4:8], localIP.To4())

	idBytes[8] = byte(processID >> 8)
	idBytes[9] = byte(processID)

	idBytes[10] = byte(seq >> 24)
	idBytes[11] = byte(seq >> 16)
	idBytes[12] = byte(seq >> 8)
	idBytes[13] = byte(seq)

	// Reserved bytes set to 0.
	idBytes[14] = 0
	idBytes[15] = 0

	return hex.EncodeToString(idBytes)
}
