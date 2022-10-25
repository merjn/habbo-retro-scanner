package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strings"
)

var createConnectionCh = make(chan string, 30)
var verifyCh = make(chan net.Conn, 30)

func process() {
	for ipAddr := range createConnectionCh {
		addr := fmt.Sprintf("%s:3000", ipAddr)

		if conn, err := Connect(addr); err == nil {
			verifyCh <- conn
		}
	}
}

func verify() {
	for conn := range verifyCh {
		// Store the IP so that we can check.
		ip := conn.RemoteAddr().String()

		isHabboServer, err := VerifyHabboServer(conn)

		// Done with it - close.
		conn.Close()

		if err != nil {
			continue
		}

		if isHabboServer {
			log.Printf("%s is a Habbo server", ip)
		}
	}
}

func init() {
	// Create n instances of process
	fmt.Println(cap(createConnectionCh))
	for i := 0; i < cap(createConnectionCh); i++ {
		go process()
	}

	// Create n instances of verify
	for i := 0; i < cap(verifyCh); i++ {
		go verify()
	}
}

func main() {
	fmt.Println("Habbo retro scanner")
	ipsFromFile, _ := ioutil.ReadFile("ips.txt")

	cleanedIps := strings.Split(strings.TrimSpace(string(ipsFromFile)), "\n")
	for _, ip := range cleanedIps {
		createConnectionCh <- ip
	}
}
