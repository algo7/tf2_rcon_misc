package network

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func scanPort(protocol, hostname string, port int) bool {

	fmt.Printf("Scanning: %s\n", hostname)
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func getHostInfo() []string {

	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	var ips []string
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			ips = append(ips, ipv4.String())
		}
	}

	return ips
}

func determineRconHost() string {

	var rconHost string = "Nothing"

	for _, ip := range getHostInfo() {
		open := scanPort("tcp", ip, 27015)
		if open {
			rconHost = ip
			break
		}
	}

	return rconHost
}
