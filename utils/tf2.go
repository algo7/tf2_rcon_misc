package utils

import (
	"errors"
	"log"
	"math/big"
	"os"
	"os/user"
	"runtime"
	"strconv"
	"strings"

	"github.com/nxadm/tail"
	"github.com/trivago/grok"
)

// Global variables
const (
	grokPattern          = `^# +%{NUMBER:userId} %{QS:userName} +\[%{WORD:steamAccType}:%{NUMBER:steamUniverse}:%{NUMBER:steamID32}\] +%{MINUTE}:%{SECOND} +%{NUMBER} +%{NUMBER} +%{WORD}$`
	grokPlayerNamePatten = `%{QS}=%{QS:playerName}\(def\.%{QS}\)%{GREEDYDATA}`
	chatPattern          = `(?:(?:\*DEAD\*(?:\(TEAM\))?)|(?:\(TEAM\)))?\s{1}%{GREEDYDATA:player_name}\s{1}:\s{2}%{GREEDYDATA:message}$`
	commandPattern       = `!${WORD:command}${SPACE}${GREENYDATA:args}`
)

var (
	g            *grok.Grok
	gc           *grok.CompiledGrok
	gPlayerName  *grok.Grok
	gcPlayerName *grok.CompiledGrok
	gChat        *grok.Grok
	gcChat       *grok.CompiledGrok
	gCommands    *grok.Grok
	gcCommands   *grok.CompiledGrok
)

// PlayerInfo is a struct containing all the info we need about a player
type PlayerInfo struct {
	SteamID       int64
	Name          string
	UserID        int
	SteamAccType  string
	SteamUniverse int
}

// ChatInfo is a struct containing all the info we need about a chat message
type ChatInfo struct {
	PlayerName string
	Message    string
}

/**
* Exported functions need to start with a capital letter
**/

// GrokInit initializes and compiles the grok patterns
func GrokInit() {
	// Compile the main grok pattern
	g, _ = grok.New(grok.Config{NamedCapturesOnly: true})
	gc, _ = g.Compile(grokPattern)

	// Compile the player name grok pattern
	gPlayerName, _ = grok.New(grok.Config{NamedCapturesOnly: true})
	gcPlayerName, _ = gPlayerName.Compile(grokPlayerNamePatten)

	// Compile the chat grok pattern
	gChat, _ = grok.New(grok.Config{NamedCapturesOnly: true})
	gcChat, _ = gChat.Compile(chatPattern)

	// Compile the command grok pattern
	gCommands, _ = grok.New(grok.Config{NamedCapturesOnly: true})
	gcCommands, _ = gCommands.Compile(commandPattern)
}

// GrokParse parses the given line with the main grok pattern
func GrokParse(line string) (*PlayerInfo, error) {

	parsed := gc.ParseString(line)

	if len(parsed) == 0 {
		return nil, errors.New("failed to parse line")
	}

	// Parse the steamID32 from the steamID3
	userID, err := strconv.Atoi(parsed["userId"])
	if err != nil {
		return nil, errors.New("failed to parse userID")
	}

	steamUniverse, err := strconv.Atoi(parsed["steamUniverse"])
	if err != nil {
		return nil, errors.New("failed to parse steamUniverse")
	}

	steamID32, err := strconv.ParseInt(parsed["steamID32"], 10, 32)
	if err != nil {
		return nil, errors.New("failed to parse SteamID32")
	}

	playerData := PlayerInfo{
		SteamID:       Steam3IDToSteam64(steamID32),
		Name:          removeQuotes(parsed["userName"]),
		UserID:        userID,
		SteamAccType:  parsed["steamAccType"],
		SteamUniverse: steamUniverse,
	}

	return &playerData, nil
}

// GrokParsePlayerName parses the given line with the playerName grok pattern
func GrokParsePlayerName(rconNameResponse string) (string, error) {
	// Remove all newlines and spaces from the string
	processed := strings.ReplaceAll(strings.ReplaceAll(rconNameResponse, "\n", ""), " ", "")
	playerNameMap := gcPlayerName.ParseString(processed)
	if len(playerNameMap) == 0 {
		return "", errors.New("You player name could not be parsed")
	}
	playerName := removeQuotes(playerNameMap["playerName"])
	return playerName, nil
}

// GrokParseChat parses the given line with the chat grok pattern
func GrokParseChat(line string) (*ChatInfo, error) {

	parsed := gcChat.ParseString(line)

	if len(parsed) == 0 {
		return nil, errors.New("failed to parse chat line")
	}

	playerName := parsed["player_name"]
	message := parsed["message"]
	chatInfo := ChatInfo{
		PlayerName: playerName,
		Message:    message,
	}

	return &chatInfo, nil
}

// GrokParseCommand parses the given line with the command grok pattern
func GrokParseCommand(line string) (string, string, error) {

	parsed := gcCommands.ParseString(line)

	if len(parsed) == 0 {
		return "", "", errors.New("failed to parse command line")
	}

	command := parsed["command"]
	args := parsed["args"]

	return command, args, nil
}

// EmptyLog empties the tf2 log file
func EmptyLog(path string) error {
	err := os.Truncate(path, 0)
	return err
}

// LogPathDection detects the tf2 log path
func LogPathDection() string {

	// Get TF2 log path from env variable
	tf2LogPath := os.Getenv("TF2_LOGPATH")

	// Auto detect log path if env variable is not set
	if tf2LogPath == "" {

		// Get operating system name
		osName := runtime.GOOS
		log.Println("OS: ", osName)
		switch osName {

		case "windows":
			tf2LogPath = "C:\\Program Files (x86)\\Steam\\steamapps\\common\\Team Fortress 2\\tf\\console.log"
			log.Println("Windows Detected. Log Path Defaulting to: ", tf2LogPath)

		case "darwin":
			tf2LogPath = "/Users/Shared/Steam/steamapps/common/Team\\ Fortress\\ 2/tf/console.log"
			log.Println("macOS Detected. Log Path Defaulting to: ", tf2LogPath)
			os.Exit(0)

		case "linux":

			// Get current os user name
			user, err := user.Current()
			if err != nil {
				log.Fatalf("Unable to determin the current OS User: %v", err)
			}
			osUSerName := user.Username

			log.Println("OS User: ", osUSerName)
			tf2LogPath = `/home/` + osUSerName + `/.local/share/Steam/steamapps/common/Team Fortress 2/tf/console.log`
			log.Printf("Linux Detected. Log Path Defaulting to: \n%s\n", tf2LogPath)

		default:
			log.Println("OS: ", osName)
			log.Println("Custom Log Path Not Provided or OS Not Supported Yet")
			os.Exit(0)
		}
	}

	return tf2LogPath
}

// TailLog tails the tf2 log file
func TailLog(tf2LogPath string) (*tail.Tail, error) {

	// Tail tf2 console log
	t, err := tail.TailFile(
		tf2LogPath,
		tail.Config{
			MustExist: true,
			Follow:    true,
			Poll:      true,
		})

	if err != nil {
		return nil, err
	}

	return t, nil
}

// Steam3IDToSteam64 converts a steam3 id to a steam64 id
func Steam3IDToSteam64(steam3ID int64) int64 {

	baseSteamID, _ := new(big.Int).SetString("76561197960265728", 0)
	steam3IDBigInt := big.NewInt(steam3ID)
	steam64ID := new(big.Int).Add(baseSteamID, steam3IDBigInt)
	num := steam64ID.Int64()

	return num
}

// removeQuotes removes all quotes from a string
func removeQuotes(str string) string {
	return strings.ReplaceAll(str, "\"", "")
}
