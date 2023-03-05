package gpt

import (
	"context"
	"errors"
	"fmt"
	"os"
	"tf2-rcon/utils"

	openai "github.com/sashabaranov/go-openai"
)

func IsAvailable() bool {
	openAiApikey := os.Getenv("OPENAI_APIKEY")

	return openAiApikey != ""
}

func Ask(question string) (string, error) {
	if !IsAvailable() {
		return "", errors.New("Apikey is not set! (env: *OPENAI_APIKEY*)")
	}

	fmt.Println("!gpt - requesting to api with Q:", question)

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

	if err != nil {
		return "", err
	}

	return utils.RemoveEmptyLines(resp.Choices[0].Message.Content), nil
}
