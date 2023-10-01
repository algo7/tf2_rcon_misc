package main

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/algo7/tf2_rcon_misc/network"
	"github.com/algo7/tf2_rcon_misc/utils"
)

// Const console message that informs you about forceful autobalance
const teamSwitchMessage = "You have switched to team BLU and will receive 500 experience points at the end of the round for changing teams."

// Slice of player info cache struct that holds the player info
var playersCache []utils.PlayerInfoCache

func main() {

	// Init the grok patterns
	utils.GrokInit()

	// Connect to the rcon server
	network.Connect()

	if !network.IsReady() {
		utils.ErrorHandler(errors.New("finally unable to establish rcon-connection"), true)
	}

	// Get the current player name
	res := network.RconExecute("name")
	parsedResponse := utils.GrokParsePlayerName(res)
	playerName := parsedResponse["playerName"]

	if len(playerName) == 0 {
		utils.ErrorHandler(errors.New("unable to parse empty response to 'name' command"), true)
	}

	log.Printf("Player Name: %s", playerName)

	// Get log path
	tf2LogPath := utils.LogPathDection()

	// Empty the log file
	utils.EmptyLog(tf2LogPath)

	// Tail the log
	fmt.Println("Tailing Logfile at:", tf2LogPath)
	t := utils.TailLog(tf2LogPath)

	// Loop through the text of each received line
	for line := range t.Lines {

		utils.GrokParse(line.Text)

		// Refresh player list logic
		// Dont assume status headlines as player connects
		if strings.Contains(line.Text, "Lobby updated") || (strings.Contains(line.Text, "connected") && !strings.Contains(line.Text, "uniqueid")) {
			log.Printf("Executing *status* rcon command after line: %s", line.Text)
			// Run the status command when the lobby is updated or a player connects
			network.RconExecute("status")
		}

		// // Save to DB logic
		// if utils.Steam3IDMatcher(line.Text) && utils.GetPlayerNameFromLine(line.Text) != "" {
		// 	// Convert Steam 32 ID to Steam 64 ID
		// 	steamID := utils.Steam3IDToSteam64(utils.Steam3IDFindString(line.Text))

		// 	// Find the player's userName
		// 	user := utils.GetPlayerNameFromLine(line.Text)

		// 	if user == "" {
		// 		fmt.Println("Failed to parse user! line.Text:", line.Text)
		// 	}

		// 	// Create a player struct
		// 	player := db.Player{
		// 		SteamID:   steamID,
		// 		Name:      user,
		// 		UpdatedAt: time.Now().UnixNano(),
		// 	}

		// 	// Add the player to the DB
		// 	db.AddPlayer(player)

		// 	// Player cache logic
		// 	playerInfoCachce := utils.PlayerInfoCache{
		// 		SteamID: steamID,
		// 		Name:    user,
		// 	}

		// 	// Add the player to the cache
		// 	utils.AddPlayerCache(&playersCache, playerInfoCachce)

		// 	fmt.Println("SteamID: ", steamID, " UserName: ", user)
		// }

		// // Command logic - TF2
		// isSay, user, text := utils.GetChatSayTF2(playersCache, line.Text)

		// // Add chat logic. prob better to do this in a separate function
		// if isSay && strings.TrimSpace(text) != "" {
		// 	steamID := utils.GetSteamIDFromPlayerCache(user, playersCache)

		// 	chat := db.Chat{
		// 		SteamID:   steamID,
		// 		Name:      user,
		// 		Message:   text,
		// 		UpdatedAt: time.Now().UnixNano(),
		// 	}

		// 	db.AddChat(chat)
		// }

		// // Command logic - TF2
		// if isSay && strings.TrimSpace(text) != "" && string(text[0]) == "!" {

		// 	commands.HandleUserSay(text, user, playerName)
		// } else {
		// 	// Command logic - Dystopia
		// 	isSay, user, text = utils.GetChatSayDystopia(playersCache, line.Text)

		// 	if isSay && strings.TrimSpace(text) != "" && string(text[0]) == "!" {
		// 		commands.HandleUserSay(text, user, playerName)
		// 	}
		// }

		// // Autobalance comment logic
		// if strings.Contains(line.Text, teamSwitchMessage) && utils.IsAutobalanceCommentEnabled() { // when you get team switched forcefully, thank gaben for the bonusxp!
		// 	time.Sleep(1000 * time.Millisecond)
		// 	network.RconExecute("say \"Thanks gaben for bonusxp!\"")
		// }

		// if utils.IsStatusResponseHostname(line.Text) {
		// 	// Refresh the player cache
		// 	playersCache = []utils.PlayerInfoCache{}
		// }

		// // Input text is not being parsed since there's no logic for parsing it (yet)
		// // fmt.Println("Unknown:", line.Text)

	}
}
