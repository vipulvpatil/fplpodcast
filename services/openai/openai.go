package openai

import (
	"context"
	"fmt"

	ai "github.com/sashabaranov/go-openai"
	"github.com/spf13/viper"
)

func GetAnalysis(transcript string) (string, error) {
	c := ai.NewClient(viper.GetString("OPENAI_API_KEY"))
	ctx := context.Background()

	req := ai.ChatCompletionRequest{
		Model:       ai.GPT4oMini,
		Temperature: 0,
		MaxTokens:   16383,
		Messages: []ai.ChatCompletionMessage{
			{
				Role:    "user",
				Content: fmt.Sprintf("%s\n%s\n%s", PROMPT_START, transcript, PROMPT_END),
			},
		},
		ResponseFormat: &responseFormat,
	}

	resp, err := c.CreateChatCompletion(ctx, req)
	if err != nil {

		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}
