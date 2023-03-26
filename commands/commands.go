package commands

import (
	"fmt"
	"strings"
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
	"!nice": func(args string) {
		time.Sleep(1000 * time.Millisecond)
		gpt.GetCompliment(args)
	},
}

// OtherUsersCommandMap is a map of functions for chat-commands that everyone (but you) is allowed to execute
var OtherUsersCommandMap = map[string]func(args string){
	// Stuff follows the : are only function pointers not function calls
	// Ask gpt API and print reponse
	"!gpt": func(args string) {
		gpt.Ask(args)
	},
	"!nice": func(args string) {
		time.Sleep(1000 * time.Millisecond)
		gpt.GetCompliment(args)
	},
}

// RunCommands is a function that runs the commands. The function takes in the text, and a boolean that tells if the user itself called the command or not
func RunCommands(text string, isSelf bool) {

	// Get the command string, e.g. !gpt
	commandArgsParsed := strings.Fields(text)
	command := commandArgsParsed[0]
	args := strings.Join(commandArgsParsed[1:], " ")

	// when command is too long, we skip
	if len(commandArgsParsed) < 128 {

		// Call different functions from the respective command maps depending on if the user itself called the command or not
		switch isSelf {

		case true:

			fmt.Println("Self Command:", command)
			fmt.Println(" Self Args: ", args)

			// Get the function for the given command
			cmdFunc := SelfCommandMap[command]

			if cmdFunc != nil {
				// Call func for given command
				cmdFunc(args)
			} else {
				// Command is not configured
				fmt.Printf("\nSelfCommand '%s' unconfigured!\n", strings.TrimSuffix(strings.TrimSuffix(command, "\n"), "\r"))
			}

		case false:

			fmt.Println("Other's Command:", command)
			fmt.Println("Other's Args: ", args)

			// Split parsed string into actual !command and arguments
			cmdFunc := OtherUsersCommandMap[command]

			if cmdFunc != nil {
				// Call func for given command
				cmdFunc(args)
			} else {
				// Command is not configured
				fmt.Printf("\nOthersCommand '%s' unconfigured!\n", strings.TrimSuffix(strings.TrimSuffix(command, "\n"), "\r"))
			}
		}
	}
}
