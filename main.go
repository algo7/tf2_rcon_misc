package main

import (
	"fmt"
	"strings"
	"tf2-rcon/network"
	"tf2-rcon/utils"
)

func main() {

	// Get the rcon host s
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

		if strings.Contains(line.Text, "Lobby updated") {
			network.RconExecute(conn, "status")
		}

		if utils.Steam3IdMatcher(line.Text) {
			// fmt.Println(line.Text)
			fmt.Println(utils.Steam3IdFindString(line.Text))
		}

	}

}
