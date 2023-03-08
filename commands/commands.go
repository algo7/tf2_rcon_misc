package commands

import (
	"fmt"
	"strings"
	"tf2-rcon/gpt"
	"tf2-rcon/network"
	"tf2-rcon/utils"
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

// RunCommands is a function that runs the commands. The function takes in the text, the playername and a boolean if the user itself called the command or not
func RunCommands(text string, playerName string, isSelf bool) {

	// Get the command string, e.g. !gpt
	completeCommand := text[len(playerName)+4:]
	fmt.Println("Command:", completeCommand)

	// when command is too long, we skip
	if len(completeCommand) > 128 {
		return
	}

	// Split parsed string into actual !command and arguments
	command, args := utils.GetCommandAndArgs(completeCommand)

	// Call different functions from the respective command maps depending on if the user itself called the command or not
	switch isSelf {

	case true:

		fmt.Println("Self Command:", command)
		fmt.Print(" Self Args: ", args)

		// Get the function for the given command
		cmdFunc := SelfCommandMap[command]

		// Command is not configured
		if cmdFunc == nil {
			fmt.Printf("Command '%s' unconfigured!\n", strings.TrimSuffix(strings.TrimSuffix(command, "\n"), "\r"))
			return
		}

		// Call func for given command
		cmdFunc(args)

	case false:

		fmt.Println("Other's Command:", command)
		fmt.Print("Other's Args: ", args)

		// Split parsed string into actual !command and arguments
		command, args := utils.GetCommandAndArgs(completeCommand)
		cmdFunc := OtherUsersCommandMap[command]

		// Command is not configured
		if cmdFunc == nil {
			fmt.Printf("Command '%s' unconfigured!\n", strings.TrimSuffix(strings.TrimSuffix(command, "\n"), "\r"))
			return
		}

		// Call func for given command
		cmdFunc(args)

	}
}
