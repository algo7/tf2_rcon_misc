package commands

import (
	"time"

	"github.com/algo7/tf2_rcon_misc/network"
)

// otherUsersCommandMap is a map of functions for chat-commands that everyone (but you) is allowed to execute
var otherUsersCommandMap = map[string]func(args string){
	// Stuff follows the : are only function pointers not function calls
	"!nice": func(args string) {
		time.Sleep(1000 * time.Millisecond)
		getCompliment(args)
	},
}

// CommandExecuted is a function that executes a given command using the given arguments
func CommandExecuted(command string, args string) {

	switch command {
	case "test":
		time.Sleep(1000 * time.Millisecond)
		network.RconExecute("say \"Test command executed!. Value:" + args + "\"")
	case "roast":
		time.Sleep(1000 * time.Millisecond)
		getInsult(args)
	}
}
