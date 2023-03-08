package main

import (
	"fmt"
	"os"
	"strings"
	"tf2-rcon/commands"
	"tf2-rcon/db"
	"tf2-rcon/network"
	"tf2-rcon/utils"
	"time"
)

// Const console message that informs you about forceful autobalance
const teamSwitchMessage = "You have switched to team BLU and will receive 500 experience points at the end of the round for changing teams."

var players []string

func main() {

	// Get the current player name
	res := network.RconExecute("name")
	playerName := strings.Split(res, " ")[2]
	playerName = strings.TrimSuffix(strings.TrimPrefix(playerName, `"`), `"`)
	fmt.Println("Player name:", playerName)

	// Get log path
	tf2LogPath := utils.LogPathDection()

	// Empty the log file
	utils.EmptyLog(tf2LogPath)

	// Tail the log
	t := utils.TailLog(tf2LogPath)

	// Loop through the text of each received line
	for line := range t.Lines {

		// Refresh player list logic
		if strings.Contains(line.Text, "Lobby updated") || strings.Contains(line.Text, "connected") {
			// Run the status command when the lobby is updated or a player connects
			network.RconExecute("status")

			// erase local player storage
			copy(players, []string{})
		}

		// Save to DB logic
		if utils.Steam3IDMatcher(line.Text) && utils.GetPlayerNameFromLine(line.Text) != "" {

			// Convert Steam 32 ID to Steam 64 ID
			steamID := utils.Steam3IDToSteam64(utils.Steam3IDFindString(line.Text))

			// Find the player's userName
			user := utils.GetPlayerNameFromLine(line.Text)

			if user == "" {
				fmt.Println("Failed to parse user! line.Text:", line.Text)
			}

			// Add the player to the DB
			db.AddPlayer(steamID, user)

			// Add the player to the cache
			utils.AddPlayer(&players, user)

			fmt.Println("SteamID: ", steamID, " UserName: ", user)
		}

		// Command logic
		isSay, user, text := utils.GetChatSay(players, line.Text)

		if isSay && text != "" && string(text[0]) == "!" {

			fmt.Printf("ChatSay - user: '%s' - text: '%s'\n", user, text)

			switch user {

			case playerName:
				fmt.Println("ChatSay, it is me!", user)
				commands.RunCommands(text, playerName, false)

			default:
				fmt.Println("ChatSay, it is not me!", user)
				commands.RunCommands(text, playerName, true)
			}
		}

		// To be decided if this is needed
		if len(line.Text) > len(playerName)+5 && line.Text[0:len(playerName)] == playerName { // that's my own say stuff

		}

		// Autobalance comment logic
		if strings.Contains(line.Text, teamSwitchMessage) && IsAutobalanceCommentEnabled() { // when you get team switched forcefully, thank gaben for the bonusxp!
			time.Sleep(1000 * time.Millisecond)
			network.RconExecute("say \"Thanks gaben for bonusxp!\"")
		}

		// Input text is not being parsed since there's no logic for parsing it (yet)
		// fmt.Println("Unknown:", line.Text)

	}
}

// IsAutobalanceCommentEnabled Check if autobalance-response is enabled or not, specified by ENV var
func IsAutobalanceCommentEnabled() bool {
	enabled := os.Getenv("ENABLE_AUTOBALANCE_COMMENT")

	return enabled == "1"
}

// // Function 3
// if strings.Contains(line.Text, "killed") &&
// 	strings.Contains(line.Text, "(crit)") &&
// 	strings.Contains(line.Text, playerName) {

// 	killer := strings.Split(line.Text, "killed")
// 	theKiller := killer[0]

// 	if theKiller == playerName {
// 		theKiller = ""
// 	}

// 	msg := utils.PickRandomMessage("crit")
// 	network.RconExecute(conn, ("say" + " " + "\"" + " " + msg + "\""))

// }

// if utils.Steam3IDMatcher(line.Text) {
// 	steamID := utils.Steam3IDToSteam64(utils.Steam3IDFindString(line.Text))
// 	fmt.Println(steamID)
// 	db.DBAddPlayer(client, steamID)
// }

// if utils.UserNameMatcher(line.Text) {
// 	userName := utils.UserNameFindString(line.Text)
// 	fmt.Println(userName)
// }

// if utils.CommandMatcher(playerName, line.Text) { // that's my own say stuff
// if len(strings.Fields(line.Text)) >= 4 {
// 	command := strings.Fields(line.Text)[2:3][0]
// 	args := strings.Fields(line.Text)[3:4][0]
// 	cmdFunc := gpt.SelfCommandMap[command]
// 	fmt.Println("Command:", command)

// 	// Command is not configured
// 	if cmdFunc == nil {
// 		continue
// 	}

// 	fmt.Print("Args: ", args)

// 	// call func for given command
// 	cmdFunc(args)
// }
// }
