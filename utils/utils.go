package utils

import (
	"fmt"
	"os"
)

// Custom error handler
func errorHandler(err error) {
	if err != nil {
		fmt.Println(err)
		fmt.Println("Press Any Key to Exit...")
		fmt.Scanln()
		os.Exit(0)
	}
}
