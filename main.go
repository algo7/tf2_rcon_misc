package main

import (
	"errors"
	"fmt"
	"strings"
	"tf2-rcon/commands"
	"tf2-rcon/db"
	"tf2-rcon/network"
	"tf2-rcon/utils"
	"time"
)

// Const console message that informs you about forceful autobalance
const teamSwitchMessage = "You have switched to team BLU and will receive 500 experience points at the end of the round for changing teams."

// String slice for caching current players
var players []string

func main() {

	network.Connect()

	if network.IsReady() == false {
		utils.ErrorHandler(errors.New("finally unable to establish rcon-connection"), true)
	}

	// Get the current player name
	res := network.RconExecute("name")
	// res sample => "name" = "Algo7" ( def. "unnamed" )
	//res = "\"name\" = \"atomy\"" // hardcode name for testing
	playerNameRaw := strings.Fields(res)

	if len(playerNameRaw) == 0 {
		utils.ErrorHandler(errors.New("unable to parse empty response to 'name' command"), true)
	}

	playerName := strings.TrimSuffix(strings.TrimPrefix(playerNameRaw[2], `"`), `"`)
	fmt.Println("Player name:", playerName)

	// Get log path
	tf2LogPath := utils.LogPathDection()

	// Empty the log file
	utils.EmptyLog(tf2LogPath)

	// Tail the log
	fmt.Println("Tailing Logfile at:", tf2LogPath)
	t := utils.TailLog(tf2LogPath)

	// Loop through the text of each received line
	for line := range t.Lines {

		// Debug, turn on to print every line we read from file
		//fmt.Printf("[+] %s\n", line.Text)

		// Refresh player list logic
		// Dont assume status headlines as player connects
		if strings.Contains(line.Text, "Lobby updated") || (strings.Contains(line.Text, "connected") && !strings.Contains(line.Text, "uniqueid")) {
			fmt.Println("Executing *status* rcon command after line:", line.Text)
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

			// Create a player struct
			player := db.Player{
				SteamID:   steamID,
				Name:      user,
				UpdatedAt: time.Now().UnixNano(),
			}

			// Add the player to the DB
			db.AddPlayer(player)

			// Add the player to the cache
			utils.AddPlayer(&players, user)

			fmt.Println("SteamID: ", steamID, " UserName: ", user)
		}

		// Command logic - TF2
		isSay, user, text := utils.GetChatSayTF2(players, line.Text)

		if isSay && text != "" && string(text[0]) == "!" {

			chat := db.Chat{
				Message:   text,
				UpdatedAt: time.Now().UnixNano(),
			}

			db.AddChat(chat)

			commands.HandleUserSay(text, user, playerName)
		} else {
			// Command logic - Dystopia
			isSay, user, text = utils.GetChatSayDystopia(players, line.Text)

			if isSay && text != "" && string(text[0]) == "!" {
				commands.HandleUserSay(text, user, playerName)
			}
		}

		// Autobalance comment logic
		if strings.Contains(line.Text, teamSwitchMessage) && utils.IsAutobalanceCommentEnabled() { // when you get team switched forcefully, thank gaben for the bonusxp!
			time.Sleep(1000 * time.Millisecond)
			network.RconExecute("say \"Thanks gaben for bonusxp!\"")
		}

		// Input text is not being parsed since there's no logic for parsing it (yet)
		// fmt.Println("Unknown:", line.Text)

	}
}
