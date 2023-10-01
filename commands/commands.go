package commands

import (
	"fmt"
	"github.com/algo7/tf2_rcon_misc/gpt"
	"github.com/algo7/tf2_rcon_misc/network"
	"strings"
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

// otherUsersCommandMap is a map of functions for chat-commands that everyone (but you) is allowed to execute
var otherUsersCommandMap = map[string]func(args string){
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

// runCommands is a function that runs the commands. The function takes in the text, and a boolean that tells if the user itself called the command or not
func runCommands(text string, isSelf bool) {

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
			cmdFunc := otherUsersCommandMap[command]

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

// HandleUserSay is a function that handles the chat messages and runs the commands if detected
func HandleUserSay(text string, user string, playerName string) {
	fmt.Printf("ChatSay - user: '%s' - text: '%s'\n", user, text)

	switch user {

	case playerName:
		fmt.Println("ChatSay, it is me!", user)
		runCommands(text, true)

	default:
		fmt.Println("ChatSay, it is not me!", user)
		runCommands(text, false)
	}
}
