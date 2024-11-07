package main

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"os"
)

var client = openai.NewClient(os.Getenv("OPENAI_API_KEY"))
var ctx = context.Background()

func create_assistants() {
	assistantName := "客服机器人"
	assistantDescription := "这是一个测试助理"
	assistantInstructions := "你是一个编剧，擅长模仿各大导演的台词，说话风格。幽默风趣为主"

	assistant, err := client.CreateAssistant(ctx, openai.AssistantRequest{
		Name:         &assistantName,
		Description:  &assistantDescription,
		Model:        openai.GPT4TurboPreview,
		Instructions: &assistantInstructions,
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("%+v\n", assistant)
	}
}

func create_thread() (response Thread) {
	response, err := client.CreateThread(ctx, openai.ThreadRequest{
		Messages: []openai.ThreadMessage{
			{
				Role:    openai.ThreadMessageRoleUser,
				Content: "Hello, World!",
			},
		},
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("success")
	}
	return response
}

func create_run(assistant_id string, thread_id string) {
	run, err := client.CreateRun(ctx, thread_id, openai.RunRequest{
		AssistantID: assistant_id,
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("%+v\n", run)
	}
}

func main() {

	create_assistants()
}
