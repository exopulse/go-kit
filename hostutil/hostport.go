package hostutil

import (
	"net"
	"strings"
)

// HostPort encapsulates host and port.
type HostPort struct {
	Host string
	Port string
}

// NewHostPort composes HostPort from a specified address, port and a default port.
// Supported formats for address are:
//   - host
//   - host:port
//   - :port
//   - host:
//   - [ipv6-host]:port (e.g. [::1]:8080)
//   - bare IPv6 host (e.g. ::1, 2001:db8::1)
//
// The method inserts specified defaultPort if port is omitted in address provided.
// The method panics if defaultPort is not specified.
// If address is empty, method will return address in form of ":defaultPort".
func NewHostPort(address, port, defaultPort string) HostPort {
	if defaultPort == "" {
		panic("missing default port")
	}

	if port == "" {
		port = defaultPort
	}

	address = strings.TrimSpace(address)

	if address == "" {
		return HostPort{Port: port}
	}

	if host, addrPort, err := net.SplitHostPort(address); err == nil {
		if addrPort == "" {
			addrPort = port
		}

		return HostPort{Host: host, Port: addrPort}
	}

	return HostPort{Host: address, Port: port}
}

// String implements Stringer interface.
func (h HostPort) String() string {
	return net.JoinHostPort(h.Host, h.Port)
}
