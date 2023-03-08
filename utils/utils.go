package utils

import (
	"errors"
	"fmt"
	"math/big"
	"os"
	"os/user"
	"regexp"
	"runtime"
	"strings"

	"github.com/nxadm/tail"
)

// Custom (error) messages
var (
	ErrMissingRconHost                = errors.New("TF2 Not Running / RCON Not Enabled")
	steam3IDRegEx                     = `\[U:[0-9]:\d{6,11}\]`
	steam3AccIDRegEx                  = `\d{6,11}`
	userNameRegEx                     = `\[U:\d:\d+\]\s+\d{2}:\d{2}\s+`
	rconNameCommandGetPlayerNameRegex = `"name" = "([^"]+)"`
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

// LogPathDection
func LogPathDection() string {

	// Get TF2 log path from env variable
	tf2LogPath := os.Getenv("TF2_LOGPATH")

	// Auto detect log path if env variable is not set
	if tf2LogPath == "" {

		// Get operating system name
		osName := runtime.GOOS
		fmt.Println("OS: ", osName)
		switch osName {

		case "windows":
			tf2LogPath = "C:\\Program Files (x86)\\Steam\\steamapps\\common\\Team Fortress 2\\tf\\console.log"
			fmt.Println("Windows Detected. Log Path Defaulting to: ", tf2LogPath)

		case "darwin":
			tf2LogPath = "/Users/Shared/Steam/steamapps/common/Team\\ Fortress\\ 2/tf/console.log"
			fmt.Println("macOS Detected. Log Path Defaulting to: ", tf2LogPath)
			os.Exit(0)

		case "linux":

			// Get current os user name
			user, err := user.Current()
			if err != nil {
				ErrorHandler(err)
			}
			osUSerName := user.Username

			fmt.Println("OS User: ", osUSerName)
			tf2LogPath = `/home/` + osUSerName + `/.local/share/Steam/steamapps/common/Team Fortress 2/tf/console.log`
			fmt.Printf("Linux Detected. Log Path Defaulting to: \n%s\n", tf2LogPath)

		default:
			fmt.Println("OS: ", osName)
			fmt.Println("Custom Log Path Not Provided or OS Not Supported Yet")
			os.Exit(0)
		}
	}

	return tf2LogPath
}

// TailLog tails the tf2 log file
func TailLog(tf2LogPath string) *tail.Tail {

	// Tail tf2 console log
	t, err := tail.TailFile(
		tf2LogPath,
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

// Steam3IDFindString returns the string that matches the given regex
func Steam3IDFindString(text string) string {
	re := regexp.MustCompile(steam3IDRegEx)
	return re.FindString(text)
}

// Steam3IDToSteam64 converts a steam3 id to a steam64 id
func Steam3IDToSteam64(givenSteam3ID string) int64 {
	fmt.Println("Steam3ID: ", givenSteam3ID)
	re := regexp.MustCompile(steam3AccIDRegEx)
	baseSteamID, _ := new(big.Int).SetString("76561197960265728", 0)
	steam3ID, _ := new(big.Int).SetString(re.FindString(givenSteam3ID), 0)
	steam64ID := new(big.Int).Add(baseSteamID, steam3ID)
	num := steam64ID.Int64()

	return num
}

// PlayerNameMatcher returns a boolean indicating if the given string matches the regex
func PlayerNameMatcher(text string) bool {
	re := regexp.MustCompile(userNameRegEx)
	return re.MatchString(text)
}

// CommandMatcher returns a boolean indicating if the given string matches the regex
func CommandMatcher(playerName string, text string) bool {
	re := regexp.MustCompile(playerName + ` :\s{1,2}!\w+`)
	return re.MatchString(text)
}

// FindCurrentPlayerName returns the string that matches the given regex
func FindCurrentPlayerName(text string) string {
	re := regexp.MustCompile(rconNameCommandGetPlayerNameRegex)
	return re.FindString(text)
}

// RemoveEmptyLines takes the supplied content and filter out empty lines, then return resulting string
func RemoveEmptyLines(content string) string {
	lines := strings.Split(content, "\n")
	filtered := make([]string, 0, len(lines))

	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			filtered = append(filtered, line)
		}
	}

	return strings.Join(filtered, "\n")
}

// GetCommandAndArgs sokts supplied func-argument (from rcon log) into command and argument, argument can be empty if there's none
func GetCommandAndArgs(content string) (string, string) {

	// Find the index of the next space character
	index := strings.IndexByte(content, ' ')

	// No whitespace found, everything is a command, there are no arguments
	if index == -1 {
		return strings.TrimSuffix(strings.TrimSuffix(content, "\n"), "\r"), ""
	}

	// argument found, return both command and arg
	commands := strings.TrimSuffix(strings.TrimSuffix(content[0:index], "\n"), "\r")
	arguments := strings.TrimSuffix(strings.TrimSuffix(content[index:], "\n"), "\r")

	return commands, arguments
}

func AddPlayer(players *[]string, elem string) {
	if !SliceContains(*players, elem) {
		*players = append(*players, elem)
		fmt.Println("adding:", elem, *players)
	}
}

// SliceContains checks if a slice contains a specific element
func SliceContains(slice []string, elem string) bool {
	for _, s := range slice {
		if s == elem {
			return true
		}
	}
	return false
}

// ExtractUsername extracts the username from the supplied string
func ExtractUsername(in string) string {
	re := regexp.MustCompile(`(\w+)" \[U:\d:[0-9]+\]`)
	match := re.FindStringSubmatch(in)

	if len(match) > 1 {
		return match[1]
	}

	return ""
}

// Check if supplied argument *in* is a chatline, if so, return: <true>, <the player that said it>, <what did he say>
func GetChatSay(players []string, in string) (bool, string, string) {

	for _, player := range players {
		// check if we found a player saying that in our playerlist
		if len(in) > len(player)+5 && in[0:len(player)] == player && in[len(player)+1:len(player)+2] == ":" { // %TODO, +1 is probably not right
			fmt.Println("IsChatSay - found player that said:", in[len(player)+4:])
			fmt.Println("IsChatSay - found player that said, player:", player)
			return true, TrimCommon(player), TrimCommon(in[len(player)+4:])
		}

		// detect dead playertalk
		// +6 is the len of string "*DEAD* "
		if len(in) > len(player)+5+7 && in[0:len(player)+7] == "*DEAD* "+player && in[len(player)+7+1:len(player)+7+2] == ":" { // %TODO, +1 is probably not right
			fmt.Println("IsChatSay (dead) - found player that said:", in[len(player)+4+7:])
			fmt.Println("IsChatSay (dead) - found player that said, player:", player)
			return true, TrimCommon(player), TrimCommon(in[len(player)+4+7:])
		}
	}

	return false, "", ""
}

// TrimCommon trims the common line endings from a string
func TrimCommon(in string) string {
	return strings.TrimSuffix(strings.TrimSuffix(in, "\n"), "\r")
}

// GetPlayerNameFromLine extracts the playername from the supplied string
func GetPlayerNameFromLine(in string) string {

	re := regexp.MustCompile(`# + \d+ "(.*)" +\[U:\d:\d+\] +[0-9:]+ + \d+ + \d+ (active|spawning)`)
	match := re.FindStringSubmatch(in)

	if len(match) > 1 {
		return match[1]
	}

	return ""
}

// Old shit

// WinTf2LogPath      string = "C:\\Program Files (x86)\\Steam\\steamapps\\common\\Team Fortress 2\\tf\\console.log"
// downMessage = [6]string{"Algo7 Down", "Algo7 Temporarily Unavailable", "Algo7 Waiting to Respawn", "Got smoked. Be right back", "Bruh...", "-.-"}
// critMessage = [5]string{"Nice crit", "Gaben has blessed you with a crit", "Random crits are fair and balanced", "Darn it, crits are always good", "Crit'd"}
// userNameRegExOld          = `#\s\s\s\s[0-9][0-9][0-9]\s"*(.*?)"`

// // pickRandomMessageIndex returns a random index of the messages array
// func pickRandomMessageIndex(min int, max int) int {
// 	rand.Seed(time.Now().UnixNano())
// 	return rand.Intn(max-min+1) + min
// }

// // PickRandomMessage returns a random message from the messages array depending on the given message type
// func PickRandomMessage(msgType string) string {

// 	msg := ""

// 	switch msgType {
// 	case "down":
// 		msgIndex := pickRandomMessageIndex(0, len(downMessage)-1)
// 		msg = downMessage[msgIndex]
// 	case "crit":
// 		msgIndex := pickRandomMessageIndex(0, len(critMessage)-1)
// 		msg = critMessage[msgIndex]
// 	}

// 	return msg
// }
