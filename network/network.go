package network

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/gorcon/rcon"
)

// Global variables
var (
	rconHost       string
	RCONConnection *rcon.Conn
)

const (
	rconPort = 27015
)

// scanPort scans for the given port on the host
func scanPort(protocol, hostname string, port int) bool {

	fmt.Printf("Connecting to: %s:%d\n", hostname, port)
	address := hostname + ":" + strconv.Itoa(port)
	RCONConnection, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil {
		return false
	}

	defer RCONConnection.Close()

	return true
}

// getHostInfo gets the host's internal ip addresses
func getHostInfo() []string {
	// Get host name
	host, err := os.Hostname()
	if err != nil {
		log.Fatalf("Unable to obtain the Hostname: %v", err)
	}

	// Get host's ipv4 and ipv6 addresses
	addrs, err := net.LookupIP(host)
	if err != nil {
		log.Fatalf("Unable to obtain the Host IP Addresses: %v", err)
	}

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
		open := scanPort("tcp", ip, rconPort)
		if open {
			rconHost = ip
			break
		}
	}

	// Check if rconHost is still "Nothing" and error if so
	if rconHost == "Nothing" {
		return ""
	}

	fmt.Printf("Rcon Host: %s:%d\n", rconHost, rconPort)

	return rconHost
}

// rconConnect connects to a rcon host
func rconConnect(rconHost string) *rcon.Conn {

	RCONConnection, err := rcon.Dial(rconHost+":"+strconv.Itoa(rconPort), "123")
	if err != nil {
		log.Printf("Unable to connect to the RCON host: %v", err)
		return nil
	}

	_, err = RCONConnection.Execute("status")
	if err != nil {
		log.Printf("Unable to execute the initial `status` command: %v", err)
		return nil
	}

	log.Println("RCON connection established")

	return RCONConnection
}

// RconExecute executes a rcon command
func RconExecute(command string) string {

	// fmt.Println("Executing: " + command)
	response, err := RCONConnection.Execute(command)

	// Reconnect if the connection is lost (usually when joining a server)
	if err != nil {
		log.Printf("Unable to execute the command: %s because %v", command, err)
		log.Println("Connection failed, retrying...")
		rconConnect(rconHost)
	}

	return response
}

// Connect tries to determine the rcon host and connect to it
func Connect() {

	// Set the loop duration to 5 minutes
	duration := 5 * time.Minute

	// Set the pause interval to 5 seconds
	interval := 5 * time.Second

	// Set the max retries to 20
	maxRetries := 20

	// Try to determine the rcon host for 5 minutes
	// Get the current time
	start := time.Now()
	try := 1

	for time.Since(start) < duration && try <= maxRetries {
		rconHost = determineRconHost()

		if rconHost == "" {
			log.Printf("Rcon host detection failed, retrying, %d/%d tries...\n", try, maxRetries)
			time.Sleep(interval)
		} else {
			break
		}

		try++
	}

	// Try to connect to the rcon host for 5 minutes
	// Get the current time, reset timer
	start = time.Now()
	try = 1

	for time.Since(start) < duration && try <= maxRetries {

		// Connect to the rcon host
		RCONConnection = rconConnect(rconHost)

		if RCONConnection == nil {
			log.Printf("Rcon connection failed, retrying, %d/%d tries...\n", try, maxRetries)
			time.Sleep(interval)
		} else {
			break
		}

		try++
	}
}
