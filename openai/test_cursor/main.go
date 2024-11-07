package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

func main() {
	// Initialize OpenAI client
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	assistantName := "My Assistant"
	assistantDescription := "这是一个测试助理"
	assistantInstructions := "You are a helpful assistant that answers questions clearly and concisely."
	// Create an assistant
	assistant, err := client.CreateAssistant(context.Background(), openai.AssistantRequest{
		Name:         &assistantName,
		Description:  &assistantDescription,
		Model:        "gpt-4-1106-preview",
		Instructions: &assistantInstructions,
	})
	if err != nil {
		fmt.Printf("Error creating assistant: %v\n", err)
		return
	}

	// Create a thread
	thread, err := client.CreateThread(context.Background(), openai.ThreadRequest{})
	if err != nil {
		fmt.Printf("Error creating thread: %v\n", err)
		return
	}

	// Start interactive loop
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\nEnter your question (or 'quit' to exit): ")
		if !scanner.Scan() {
			break
		}

		userInput := scanner.Text()
		if strings.ToLower(userInput) == "quit" {
			break
		}

		// Add message to thread
		_, err := client.CreateMessage(context.Background(), thread.ID, openai.MessageRequest{
			Role:    "user",
			Content: userInput,
		})
		if err != nil {
			fmt.Printf("Error creating message: %v\n", err)
			continue
		}

		// Run the assistant
		run, err := client.CreateRun(context.Background(), thread.ID, openai.RunRequest{
			AssistantID: assistant.ID,
		})
		if err != nil {
			fmt.Printf("Error creating run: %v\n", err)
			continue
		}

		// Wait for completion
		for {
			run, err = client.RetrieveRun(context.Background(), thread.ID, run.ID)
			if err != nil {
				fmt.Printf("Error retrieving run: %v\n", err)
				break
			}
			if run.Status == "completed" {
				break
			}
		}

		// Get messages (including assistant's response)
		messages, err := client.ListMessages(context.Background(), thread.ID, nil)
		if err != nil {
			fmt.Printf("Error listing messages: %v\n", err)
			continue
		}

		// Print the latest assistant response
		for _, msg := range messages.Messages {
			if msg.Role == "assistant" {
				fmt.Printf("\nAssistant: %s\n", msg.Content[0].Text.Value)
				break
			}
		}
	}
}
