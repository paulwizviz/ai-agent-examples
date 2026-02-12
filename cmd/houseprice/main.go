package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ollama/ollama/api"
)

func main() {
	client, err := api.ClientFromEnvironment()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	req := &api.ChatRequest{
		Model: "llama3.1:8b",
		Messages: []api.Message{
			{
				Role:    "user",
				Content: "Why is Go a good choice for systems programming?",
			},
		},
	}

	respFunc := func(resp api.ChatResponse) error {
		fmt.Print(resp.Message.Content)
		return nil
	}

	err = client.Chat(ctx, req, respFunc)
	if err != nil {
		log.Fatal(err)
	}
}
