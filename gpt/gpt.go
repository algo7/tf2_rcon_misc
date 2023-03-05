package gpt

import (
	"context"
	"errors"
	"fmt"
	"os"
	"tf2-rcon/utils"

	openai "github.com/sashabaranov/go-openai"
)

// Check if gpt is configured, checking for apikey env being set
func IsAvailable() bool {
	openAiApikey := os.Getenv("OPENAI_APIKEY")

	return openAiApikey != ""
}

// Ask gpt for given question, make request to openai API
func Ask(question string) (string, error) {
	// Check if apikey is available, error if not
	if !IsAvailable() {
		return "", errors.New("Apikey is not set! (env: *OPENAI_APIKEY*)")
	}

	fmt.Println("!gpt - requesting to api with Q:", question)

	// Create client from lib and request "answer" to "question"
	openAiApikey := os.Getenv("OPENAI_APIKEY")
	client := openai.NewClient(openAiApikey)
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
		return "", err
	}

	// Return Content node, remove empty lines from it beforehand
	return utils.RemoveEmptyLines(resp.Choices[0].Message.Content), nil
}
