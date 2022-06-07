package utils

import (
	"errors"
	"fmt"
	"os"
)

// Custom errors
var (
	errDirectoryCreation        = errors.New("FAILED TO CREATE DIRECTORIES")
	errGetDirectory             = errors.New("FAILED TO GET THE CURRENT DIRECTORY")
	errPurgeDirectory           = errors.New("FAILED TO PURGE THE TMP DIRECTORY")
	errCopyFile                 = errors.New("FAILED TO COPY DOCKER-COMPSE-PROD.YML")
	errCloneRepo                = errors.New("FAILED TO CLONE THE REPOSITORY")
	errDockerCheck              = errors.New("DOCKER IS NOT INSTALLED")
	ErrSetupCheck               = errors.New("SETUP CHECK FAILED")
	errDockerComposeRun         = errors.New("FAILED TO RUN DOCKER-COMPOSE")
	errReviewsNotEmpty          = errors.New("REVIEWS DIRECTORY IS NOT EMPTY")
	errMissingSourceFiles       = errors.New("MISSING SOURCE FILES")
	errInputScrapMode           = errors.New("INVALID SCRAP MODE")
	errInputConcurrency         = errors.New("INVALID CONCURRENCY VALUE")
	errDockerComposeYmlNotFound = errors.New("DOCKER-COMPSE-PROD.YML NOT FOUND")
	errValueReplace             = errors.New("FAILED TO REPLACE VALUE")
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
