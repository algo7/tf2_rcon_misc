package commands

import (
	"time"

	"github.com/algo7/tf2_rcon_misc/network"
)

// CommandExecuted is a function that executes a given command using the given arguments
func CommandExecuted(command string, args string, callerName string, currentPlayer string) {

	// If the caller is the current player
	if callerName == currentPlayer {
		switch command {
		case "test":
			time.Sleep(1000 * time.Millisecond)
			network.RconExecute("say \"Test command executed!. Value:" + args + "\"")
		case "roast":
			time.Sleep(1000 * time.Millisecond)
			getInsult(args)
		default:
			return
		}
	}

	// If the caller is not the current player
	switch command {
	case "test":
		time.Sleep(1000 * time.Millisecond)
		network.RconExecute("say \"Test command executed!. Value:" + args + "\"")
	default:
		return
	}

}
