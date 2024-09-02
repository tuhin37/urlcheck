package netcheck

import (
	"net"
	"time"
)

// CheckConnectivity quickly checks internet connectivity using UDP.
func CheckConnectivity() bool {
	// Define the address and port for a quick connectivity check
	address := "8.8.8.8:53" // Using port 53 (DNS) for UDP

	// Set a short timeout for the connection
	timeout := 2 * time.Second

	conn, err := net.DialTimeout("udp", address, timeout)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}
