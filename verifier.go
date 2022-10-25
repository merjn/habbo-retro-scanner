package main

import (
	"bytes"
	"errors"
	"net"
)

// ErrNoPongResponse will be returned if we didn't get a pong response from the server. It is therefore
// not a Habbo server.
var ErrNoPongResponse = errors.New("didn't get pong response from server - not a habbo retro")

// PongPackets contains a byte-representation of the pong event.
var PongPacket = []byte{0, 0, 0, 52, 9, 186, 0, 15, 115, 116, 111, 114, 101, 100, 109, 97, 99, 104, 105, 110, 101, 105, 100, 0, 17, 99, 108, 105, 101, 110, 116, 102, 105, 110, 103, 101, 114, 112, 114, 105, 110, 116, 0, 12, 99, 97, 112, 97, 98, 105, 108, 105, 116, 105, 101, 115}

// PongExpectedResponse contains a byte-representation of the expected header.
var PongExpectedResponse = []byte{0, 0, 0, 68, 5, 208, 0, 64}

// VerifyHabboServer checks if the connection is a real Habbo server.
func VerifyHabboServer(conn net.Conn) (bool, error) {
	_, err := conn.Write(PongPacket)
	if err != nil {
		return false, err
	}

	// Allocate 8 bytes for the response.
	serverResponse := make([]byte, 8)
	_, err = conn.Read(serverResponse)
	if err != nil {
		return false, err
	}

	if !bytes.Equal(serverResponse, PongExpectedResponse) {
		return false, ErrNoPongResponse
	}

	return true, nil
}
