package gpt

import (
	"context"
	"errors"
	"fmt"
	"os"
	"tf2-rcon/network"
	"tf2-rcon/utils"
	"time"

	openai "github.com/sashabaranov/go-openai"
)

// Commandmap for chat-commands that only you are allowed to execute
var selfCommandMap = map[string]func(args string){

	// ask gpt API and print reponse
	"!gpt": func(args string) {

		// execute request and proceed with result or error
		fmt.Println("!gpt - requesting:", args)
		responses, err := Ask(args)
		fmt.Println("!gpt - requesting:", args, "- Response:", responses)

		// Check for error
		if err != nil {
			utils.ErrorHandler(err)
		}

		// Split the original string into chunks of 121 characters
		// Have at max 2 interations cause we dont want to spam chat
		for i := 0; i < len(responses); i += 121 {
			end := i + 121

			if end > len(responses) {
				end = len(responses)
			}

			chunk := responses[i:end]

			// If no the 1st try, delay 1000 ms cause else we may get supressed
			if i != 0 {
				time.Sleep(1000 * time.Millisecond)
				network.RconExecute(conn, ("say \"GPT " + chunk + "\""))
				break // only execute this once, we dont want to spam
			}

			// on first run only delay 500 ms
			time.Sleep(500 * time.Millisecond)
			network.RconExecute(conn, ("say \"GPT " + chunk + "\""))
		}
	},
	// Just a test command
	"!test": func(args string) {
		// 500 ms seems to work often, but not always, so lets be safe and use 1k
		time.Sleep(1000 * time.Millisecond)
		network.RconExecute(conn, ("say \"Test confirmed!\""))
	},
}

func openAIConnect() *openai.Client {

	// Get apikey from env
	openAiApikey := os.Getenv("OPENAI_APIKEY")

	// Check if apikey is available, error if not
	if openAiApikey == "" {
		utils.ErrorHandler(errors.New("Apikey is not set! (env: *OPENAI_APIKEY*)"))
	}

	// Create client from lib and request "answer" to "question"
	client := openai.NewClient(openAiApikey)

	return client
}

// Ask asks GPT the given question, make request to openai API
func Ask(question string) (string, error) {

	// Connect to openai API
	client := openAIConnect()

	fmt.Println("!gpt - requesting to api with Q:", question)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: question,
				},
			},
		},
	)

	// Check for error
	if err != nil {
		utils.ErrorHandler(err)
	}

	// Return Content node, remove empty lines from it beforehand
	return utils.RemoveEmptyLines(resp.Choices[0].Message.Content), nil
}
