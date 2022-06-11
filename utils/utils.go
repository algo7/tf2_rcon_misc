package utils

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// Custom errors
var (
	ErrMissingRconHost = errors.New("TF2 Not Running / RCON Not Enabled")
)

/**
* Exported functions need to start with a capital letter
**/

// ErrorHandler print the err, stop the program if err is not nil, and exit on user input
func ErrorHandler(err error) {
	if err != nil {
		fmt.Println(err)
		fmt.Println("Press Any Key to Exit...")
		fmt.Scanln()
		os.Exit(0)
	}
}

// EmptyLog empties the tf2 log file
func EmptyLog(path string) {
	err := os.Truncate(path, 0)
	ErrorHandler(err)
}

// PickRandomMessageIndex returns a random index of the messages array
func PickRandomMessageIndex(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
