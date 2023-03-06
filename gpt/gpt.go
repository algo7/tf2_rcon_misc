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

// Connect to openai API
var client = openAIConnect()

// Commandmap for chat-commands that only you are allowed to execute
var selfCommandMap = map[string]interface{}{

	// ask gpt API and print reponse
	"!gpt": Ask(question),
	// Just a test command
	"!test": network.RconExecute(conn, ("say \"Test confirmed!\"")),
}

// openAIConnect connects to openai API and returns an instance of the client
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
func Ask(question string) string {

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
	responses := utils.RemoveEmptyLines(resp.Choices[0].Message.Content)

	fmt.Println("!gpt - requesting:", question, "- Response:", responses)

	// Split the original string into chunks of 121 characters
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
			return chunk
			// network.RconExecute(conn, ("say \"GPT " + chunk + "\""))
			break // only execute this once, we dont want to spam
		}

		// on first run only delay 500 ms
		time.Sleep(500 * time.Millisecond)
		// network.RconExecute(conn, ("say \"GPT " + chunk + "\""))
	}

	return chunk
}
