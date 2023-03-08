package commands

import (
	"fmt"
	"tf2-rcon/gpt"
	"tf2-rcon/network"
	"time"
)

// SelfCommandMap is a map of functions for chat-commands that only you are allowed to execute
var SelfCommandMap = map[string]func(args string){
	// Stuff follows the : are only function pointers not function calls
	// Ask gpt API and print reponse
	"!gpt": func(args string) {
		fmt.Println(args)
		gpt.Ask(args)
	},
	// Just a test command
	"!test": func(args string) {
		fmt.Println("Test command executed!")
		time.Sleep(1000 * time.Millisecond)
		network.RconExecute("say \"Test command executed!. Value:" + args + "\"")
	},
	// Roast someone
	"!roast": func(args string) {
		time.Sleep(1000 * time.Millisecond)
		gpt.GetInsult(args)
	},
}

// OtherUsersCommandMap is a map of functions for chat-commands that everyone (but you) is allowed to execute
var OtherUsersCommandMap = map[string]func(args string){
	// Stuff follows the : are only function pointers not function calls
	// Ask gpt API and print reponse
	"!gpt": func(args string) {
		fmt.Println(args)
		gpt.Ask(args)
	},
	// Roast someone
	"!roast": func(args string) {
		time.Sleep(1000 * time.Millisecond)
		gpt.GetInsult(args)
	},
}

func RunCommands(text string) {

}
