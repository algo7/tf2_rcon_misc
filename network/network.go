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

// Get the rcon host
var rconHost = determineRconHost()

// Connect to the rcon host
var conn = rconConnect(rconHost)

// scanPort scans for the given port on the host
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

// determineRconHost determines the rcon host to connect to
func determineRconHost() string {

	var rconHost string = "Nothing"

	// Scan all the ip address opened rcon port and return the ip addr with an opened rcon port
	for _, ip := range getHostInfo() {
		open := scanPort("tcp", ip, 27015)
		if open {
			rconHost = ip
			break
		}
	}

	// Check if rconHost is still "Nothing" and error if so
	if rconHost == "Nothing" {
		utils.ErrorHandler(utils.ErrMissingRconHost)
	}

	fmt.Printf("Rcon Host: %s\n", rconHost)

	return rconHost
}

// rconConnect connects to a rcon host
func rconConnect(rconHost string) *rcon.Conn {

	conn, err := rcon.Dial(rconHost+":27015", "123")
	utils.ErrorHandler(err)

	_, err = conn.Execute("status")
	utils.ErrorHandler(err)

	fmt.Println("Connected")

	return conn
}

// RconExecute executes a rcon command
func RconExecute(command string) string {

	// fmt.Println("Executing: " + command)
	response, err := conn.Execute(command)
	utils.ErrorHandler(err)

	return response
}
