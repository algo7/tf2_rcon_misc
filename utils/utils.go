package utils

import (
	"errors"
	"fmt"
	"math/big"
	"math/rand"
	"os"
	"regexp"
	"time"

	"github.com/nxadm/tail"
)

// Custom errors
var (
	ErrMissingRconHost        = errors.New("TF2 Not Running / RCON Not Enabled")
	WinTf2LogPath      string = "C:\\Program Files (x86)\\Steam\\steamapps\\common\\Team Fortress 2\\tf\\console.log"
	downMessage               = [6]string{"Algo7 Down", "Algo7 Temporarily Unavailable", "Algo7 Waiting to Respawn", "Got smoked. Be right back", "Bruh...", "-.-"}
	critMessage               = [5]string{"Nice crit", "Gaben has blessed you with a crit", "Random crits are fair and balanced", "Darn it, crits are always good", "Crit'd"}
	steam3IDRegEx             = `\[U:[0-9]:\d{8,11}\]`
	steam3AccIDRegEx          = `\d{8,11}`
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

// pickRandomMessageIndex returns a random index of the messages array
func pickRandomMessageIndex(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

// PickRandomMessage returns a random message from the messages array depending on the given message type
func PickRandomMessage(msgType string) string {

	msg := ""

	switch msgType {
	case "down":
		msgIndex := pickRandomMessageIndex(0, len(downMessage)-1)
		msg = downMessage[msgIndex]
	case "crit":
		msgIndex := pickRandomMessageIndex(0, len(critMessage)-1)
		msg = critMessage[msgIndex]
	}

	return msg
}

// TailLog tails the tf2 log file
func TailLog() *tail.Tail {
	// Tail tf2 console log
	t, err := tail.TailFile(
		WinTf2LogPath,
		tail.Config{
			MustExist: true,
			Follow:    true,
			Poll:      true,
		})

	ErrorHandler(err)

	return t
}

// Steam3IDMatcher returns a boolean indicating if the given string matches the regex
func Steam3IDMatcher(text string) bool {
	re := regexp.MustCompile(steam3IDRegEx)
	return re.MatchString(text)
}

// Steam3IDFindString returns the string the matches the given regex
func Steam3IDFindString(text string) string {
	re := regexp.MustCompile(steam3IDRegEx)
	return re.FindString(text)
}

// Steam3IDToSteam64 converts a steam3 id to a steam64 id
func Steam3IDToSteam64(givenSteam3ID string) int64 {
	re := regexp.MustCompile(steam3AccIDRegEx)
	baseSteamID, _ := new(big.Int).SetString("76561197960265728", 0)
	steam3ID, _ := new(big.Int).SetString(re.FindString(givenSteam3ID), 0)
	steam64ID := new(big.Int).Add(baseSteamID, steam3ID)
	num := steam64ID.Int64()

	return num
}
