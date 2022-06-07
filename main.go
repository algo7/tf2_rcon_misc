package main

import (
	"tf2-rcon/network"
	"tf2-rcon/utils"
)

// Global variables
var (
	winTf2LogPath string = "C:\\Program Files (x86)\\Steam\\steamapps\\common\\Team Fortress 2\\tf\\console.log"
)

func main() {
	rconHost := network.DetermineRconHost()
	if rconHost == "Nothing" {
		utils.ErrorHandler(utils.ErrMissingRconHost)
	}

}
