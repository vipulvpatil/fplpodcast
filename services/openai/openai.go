package openai

import (
	"context"
	"fmt"

	ai "github.com/sashabaranov/go-openai"
	"github.com/spf13/viper"
)

const PROMPT = `I need an analysis of the given podcast transcript. Podcast description is as follows.
Welcome to the FPL Pod. Join Kelly Somers, Julien Laurens and Sam Bonfield ahead of each Gameweek as they invite special guests to discuss and digress about the world of Fantasy Premier League. After each Gameweek, two Fantasy Premier League experts will discuss all the major talking points and analyse the key decisions in Off the Bench.
I need to do the following analysis.

Identify every player mentioned. For each player, provide a summary of the views of individual speakers. Also try to identify each individual speakers current team setup.`

func GetAnalysis(transcript string) (string, error) {
	c := ai.NewClient(viper.GetString("OPENAI_API_KEY"))
	ctx := context.Background()

	req := ai.ChatCompletionRequest{
		Model:       ai.GPT4oMini,
		Temperature: 0,
		MaxTokens:   16383,
		TopP:        1,
		Messages: []ai.ChatCompletionMessage{
			ai.ChatCompletionMessage{
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
