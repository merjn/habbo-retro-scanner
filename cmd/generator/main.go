package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
)

func Hosts(cidr []string) {
	f, err := os.OpenFile("ips.txt", os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer f.Close()

	for _, cur := range cidr {
		ip, ipnet, err := net.ParseCIDR(cur)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(cur)

		for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
			str := fmt.Sprintf("%s\n", ip.String())
			f.WriteString(str)
		}
	}

}

//  http://play.golang.org/p/m8TNTtygK0
func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func main() {
	data, err := ioutil.ReadFile("ovh.txt")
	if err != nil {
		log.Fatal(err)
	}

	subs := strings.Split(strings.TrimSpace(string(data)), "\n")

	Hosts(subs)
}
