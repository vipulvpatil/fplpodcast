package openai

import (
	"context"
	"fmt"

	ai "github.com/sashabaranov/go-openai"
)

func completion() {
	c := ai.NewClient("your token")
	ctx := context.Background()

	req := ai.CompletionRequest{
		Model:     ai.GPT4oMini,
		MaxTokens: 5,
		Prompt:    "Lorem ipsum",
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		fmt.Printf("Completion error: %v\n", err)
		return
	}
	fmt.Println(resp.Choices[0].Text)
}
