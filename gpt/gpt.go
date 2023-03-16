package gpt

import (
	"context"
	"fmt"
	"os"
	"tf2-rcon/network"
	"tf2-rcon/utils"
	"time"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

// Connect to openai API
var client, clientAvailable = openAIConnect()

// openAIConnect connects to openai API and returns an instance of the client
func openAIConnect() (*openai.Client, bool) {

	// Get apikey from env
	openAiApikey := os.Getenv("OPENAI_APIKEY")

	// Check if apikey is available, error if not
	if openAiApikey == "" {
		// utils.ErrorHandler(errors.New("Apikey is not set! (env: *OPENAI_APIKEY*)"))
		fmt.Println("OpenAI API Key Key Not Set")
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
					Content: "Always limit your response to maximum of 121 characters, try to formulate a short answer that always fits that criteria. The question is: " + question,
				},
			},
			MaxTokens: int(121),
		},
	)

	// Check for error
	if err != nil {
		utils.ErrorHandler(err)
	}

	// Return Content node, remove empty lines from it beforehand
	responseText := utils.RemoveEmptyLines(resp.Choices[0].Message.Content)
	
	fmt.Println("!gpt - requesting:", question, "- Response:", responseText)
	
	// Remove any newlines and limit response to 121 characters
	responseText = strings.Replace(responseText, "\n", " ", -1)
	if len(responseText) > 121 {
		responseText = strings.TrimSpace(responseText)[:121]
		lastSpace := strings.LastIndex(responseText, " ")
		
		if lastSpace != -1 {
			responseText = responseText[:lastSpace] + "..."
		} else {
			responseText += "..."
		}		
	} else {
		responseText = strings.TrimSpace(responseText)
	}

	// on first run only delay 500 ms
	time.Sleep(500 * time.Millisecond)
	network.RconExecute("say \"GPT> " + responseText + "\"")
}
