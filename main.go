package main

import (
	"fmt"
	// "strings"
	"tf2-rcon/network"
	"tf2-rcon/utils"
	// "github.com/nxadm/tail"
)

// Global variables
var (
	winTf2LogPath string = "C:\\Program Files (x86)\\Steam\\steamapps\\common\\Team Fortress 2\\tf\\console.log"
	downMessage          = [6]string{"Algo7 Down", "Algo7 Temporarily Unavailable", "Algo7 Waiting to Respawn", "Got smoked. Be right back", "Bruh...", "-.-"}
	critMessage          = [5]string{"Nice crit", "Gaben has blessed you with a crit", "Random crits are fair and balanced", "Darn it, crits are always good", "Crit'd"}
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
	res := network.RconExecute(conn, "name")
	fmt.Printf("Name: %s\n", res)
	// Empty the log file
	utils.EmptyLog(winTf2LogPath)

	// // Tail tf2 console log
	// t, err := tail.TailFile(
	// 	winTf2LogPath,
	// 	tail.Config{
	// 		MustExist: true,
	// 		Follow:    true,
	// 		Poll:      true,
	// 	})
	// utils.ErrorHandler(err)

	// // Loop through the text of each received line
	// for line := range t.Lines {

	// 	// Function 1
	// 	if strings.Contains(line.Text, "killed Algo7") &&
	// 		!strings.Contains(line.Text, "(crit)") {
	// 		// Send rcon command
	// 		msgIndex := utils.PickRandomMessageIndex(0, len(downMessage)-1)
	// 		network.RconExecute(conn, ("say" + " " + "\"" + downMessage[msgIndex] + "\""))
	// 	}

	// 	// Function 3
	// 	if strings.Contains(line.Text, "killed") &&
	// 		strings.Contains(line.Text, "(crit)") &&
	// 		strings.Contains(line.Text, "Algo7") {
	// 		killer := strings.Split(line.Text, "killed")
	// 		// victim := strings.TrimSpace(strings.Split(killer[1], "with")[0])
	// 		// fmt.Println(line.Text)
	// 		// fmt.Println(killer[0])
	// 		// fmt.Println(victim)
	// 		theKiller := killer[0]
	// 		if theKiller == "Algo7" {
	// 			theKiller = ""
	// 		}
	// 		msgIndex := utils.PickRandomMessageIndex(0, len(critMessage)-1)

	// 		network.RconExecute(conn, ("say" + " " + "\"" + theKiller + critMessage[msgIndex] + "\""))

	// 	}

	// 	fmt.Println(line.Text)

	// }

}
