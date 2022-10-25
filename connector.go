package main

import (
	"net"
	"time"
)

// Connect tries to connect to the TCP server.
func Connect(addr string) (net.Conn, error) {
	conn, err := net.DialTimeout("tcp", addr, 500*time.Millisecond)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
