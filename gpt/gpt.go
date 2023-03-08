package gpt

import (
	"context"
	"fmt"
	"os"
	"tf2-rcon/network"
	"tf2-rcon/utils"
	"time"

	openai "github.com/sashabaranov/go-openai"
)

// Connect to openai API
var client, clientAvailable = openAIConnect()

// SelfCommandMap is a map of functions for chat-commands that only you are allowed to execute
var SelfCommandMap = map[string]func(args string){
	// Stuff follows the : are only function pointers not function calls
	// Ask gpt API and print reponse
	"!gpt": func(args string) {
		fmt.Println(args)
		Ask(args)
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
		GetInsult(args)
	},
}

// OtherUsersCommandMap is a map of functions for chat-commands that everyone (but you) is allowed to execute
var OtherUsersCommandMap = map[string]func(args string){
	// Stuff follows the : are only function pointers not function calls
	// Ask gpt API and print reponse
	"!gpt": func(args string) {
		fmt.Println(args)
		Ask(args)
	},
	// Roast someone
	"!roast": func(args string) {
		time.Sleep(1000 * time.Millisecond)
		GetInsult(args)
	},
}

// openAIConnect connects to openai API and returns an instance of the client
func openAIConnect() (*openai.Client, bool) {

	// Get apikey from env
	openAiApikey := os.Getenv("OPENAI_APIKEY")

	// Check if apikey is available, error if not
	if openAiApikey == "" {
		// utils.ErrorHandler(errors.New("Apikey is not set! (env: *OPENAI_APIKEY*)"))
		fmt.Println("Key Not Set")
		return nil, false
	}

	// Create client from lib and request "answer" to "question"
	client := openai.NewClient(openAiApikey)

	return client, true
}

// Ask asks GPT the given question, make request to openai API
func Ask(question string) {

	// Check if client is available
	if !clientAvailable {
		fmt.Println("GPT Client not available")
		return
	}

	fmt.Println("!gpt - requesting to api with Q:", question)

	// execute request and proceed with result or error
	fmt.Println("!gpt - requesting:", question)

	// Send request to openai API
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Answer in one sentence: " + question,
				},
			},
		},
	)

	// Check for error
	if err != nil {
		utils.ErrorHandler(err)
	}

	// Return Content node, remove empty lines from it beforehand
	responses := utils.RemoveEmptyLines(resp.Choices[0].Message.Content)

	fmt.Println("!gpt - requesting:", question, "- Response:", responses)

	// Split the original string into chunks of 121 characters, the overall chat-say limit is 126, subtract any chars needed for prefix
	// Have at max 2 interations cause we dont want to spam chat
	chunk := ""
	for i := 0; i < len(responses); i += 121 {
		end := i + 121

		if end > len(responses) {
			end = len(responses)
		}

		chunk = responses[i:end]

		// If no the 1st try, delay 1000 ms cause else we may get supressed
		if i != 0 {
			time.Sleep(1000 * time.Millisecond)

			network.RconExecute("say \"GPT> " + chunk + "\"")
			break // only execute this once, we dont want to spam

		}

		// on first run only delay 500 ms
		time.Sleep(500 * time.Millisecond)
		network.RconExecute("say \"GPT> " + chunk + "\"")
	}
}
