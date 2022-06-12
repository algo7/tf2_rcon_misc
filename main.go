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

		// Function 1
		if strings.Contains(line.Text, "killed Algo7") && !strings.Contains(line.Text, "(crit)") {
			// Send rcon command
			msgIndex := utils.PickRandomMessageIndex(0, len(downMessage)-1)
			network.RconExecute(conn, ("say" + " " + "\"" + downMessage[msgIndex] + "\""))
		}

		// Function 3
		if strings.Contains(line.Text, "killed") &&
			strings.Contains(line.Text, "(crit)") &&
			strings.Contains(line.Text, "Algo7") {
			killer := strings.Split(line.Text, "killed")
			// victim := strings.TrimSpace(strings.Split(killer[1], "with")[0])
			// fmt.Println(line.Text)
			// fmt.Println(killer[0])
			// fmt.Println(victim)
			theKiller := killer[0]
			if theKiller == playerName {
				theKiller = ""
			}
			msgIndex := utils.PickRandomMessageIndex(0, len(critMessage)-1)

			network.RconExecute(conn, ("say" + " " + "\"" + theKiller + critMessage[msgIndex] + "\""))

		}

		fmt.Println(line.Text)

	}

}
