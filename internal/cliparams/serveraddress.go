package cliparams

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
)

type ServerAddress string

var (
	errInvalidScheme        = errors.New("should be \"address:port\"")
	errInvalidRange         = errors.New("should be in range 1-65535")
	errInvalidListenAddress = errors.New("should be valid listen address")
)

func (address *ServerAddress) UnmarshalText(text []byte) error {
	return address.Set(string(text))
}

func (address *ServerAddress) String() string {
	return fmt.Sprint(*address)
}

func (address *ServerAddress) Set(value string) error {
	splitted := strings.Split(value, ":")
	if len(splitted) == 0 {
		return errInvalidScheme
	}

	port := splitted[len(splitted)-1]
	if !isValidPort(port) {
		return errInvalidRange
	}

	hostname := strings.TrimSuffix(value, ":"+port)
	if !isValidHostname(hostname) {
		return errInvalidListenAddress
	}

	*address = ServerAddress(value)

	return nil
}

func isValidPort(port string) bool {
	p, err := strconv.Atoi(port)
	if err != nil {
		return false
	}

	if p < 1 || p > 65535 {
		return false
	}

	return true
}

func isValidHostname(hostname string) bool {
	if hostname == "" || hostname == "localhost" {
		return true
	}
	//// IPv4
	if net.ParseIP(hostname) != nil {
		return true
	}
	//// IPv6
	if !strings.HasPrefix(hostname, "[") || !strings.HasSuffix(hostname, "]") {
		return false
	}

	serverIPv6Address := hostname[1 : len(hostname)-1]

	return net.ParseIP(serverIPv6Address) != nil
}
