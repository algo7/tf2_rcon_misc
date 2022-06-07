package network

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"tf2-rcon/utils"
	"time"

	"github.com/gorcon/rcon"
)

func scanPort(protocol, hostname string, port int) bool {

	fmt.Printf("Scanning: %s\n", hostname)
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)
	utils.ErrorHandler(err)
	defer conn.Close()
	return true
}

func getHostInfo() []string {
	// Get host name
	host, err := os.Hostname()
	utils.ErrorHandler(err)

	// Get host's ipv4 and ipv6 addresses
	addrs, err := net.LookupIP(host)
	utils.ErrorHandler(err)

	// Slice to hold ipv4 and ipv6 addresses
	var ips []string

	// Loop through the addresses and keep only ipv4 addresses
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			ips = append(ips, ipv4.String())
		}
	}

	return ips
}

// DetermineRconHost determines the rcon host to connect to
func DetermineRconHost() string {

	var rconHost string = "Nothing"

	// Scan all the ip address opened rcon port and return the ip addr with an opened rcon port
	for _, ip := range getHostInfo() {
		open := scanPort("tcp", ip, 27015)
		if open {
			rconHost = ip
			break
		}
	}

	return rconHost
}

// RconExecute executes a rcon command
func RconExecute(conn *rcon.Conn, command string) {
	fmt.Println("Executing: " + command)
	_, err := conn.Execute(command)
	utils.ErrorHandler(err)
}
