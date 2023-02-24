package main

import (
	"fmt"
	"strings"
	"tf2-rcon/db"
	"tf2-rcon/network"
	"tf2-rcon/utils"
)

func main() {

	// Connect to the DB
	client := db.Connect()

	// Get the rcon host
	rconHost := network.DetermineRconHost()
	if rconHost == "Nothing" {
		utils.ErrorHandler(utils.ErrMissingRconHost)
	}

	fmt.Printf("Rcon Host: %s\n", rconHost)

	// Connect to the rcon host
	conn := network.RconConnect(rconHost)

	// Get the current player name
	res := network.RconExecute(conn, "name")
	playerName := strings.Split(res, " ")[2]
	playerName = strings.TrimSuffix(strings.TrimPrefix(playerName, `"`), `"`)

	// Empty the log file
	utils.EmptyLog(utils.WinTf2LogPath)

	// Tail the log
	t := utils.TailLog()

	// Loop through the text of each received line
	for line := range t.Lines {

		// Run the status command when the lobby is updated or a player connects
		if strings.Contains(line.Text, "Lobby updated") || strings.Contains(line.Text, "connected") {
			network.RconExecute(conn, "status")
		}

		// Match all the players' steamID and name from the output of the status command
		if utils.Steam3IDMatcher(line.Text) && utils.PlayerNameMatcher(line.Text) {

			// Potential logic for parsing the status command output
			// // Sample input with player name containing spaces and nested quotes
			// input := `32 "Saigial ""D""arun" [U:1:60866573] 07:46 64 0 active`

			// // Split the input string into an array using spaces as the delimiter
			// fields := strings.Fields(input)
			// fmt.Println(fields)

			// playerName := strings.Join(fields[1:len(fields)-5], " ")
			// fmt.Println(playerName)

			// Convert Steam 32 ID to Steam 64 ID
			steamID := utils.Steam3IDToSteam64(utils.Steam3IDFindString(line.Text))

			// Find the player's userName
			userName := utils.PlayerNameFindString(line.Text)

			// Add the player to the DB
			db.AddPlayer(client, steamID, userName)

			fmt.Println("SteamID: ", steamID, " UserName: ", userName)
		}

	}

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
