package pkg

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
)

type Client struct {
	openaiClient *openai.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		openaiClient: openai.NewClient(apiKey),
	}
}

func (c *Client) SummarizeTweets(tweets []string) (string, error) {
	// Combine all tweets into a single string
	input := ""
	for _, tweet := range tweets {
		input += tweet + "\n"
	}
	input += "\nHere is the summary of the above tweets in 200 characters or less:\n"
	ctx := context.Background()

	req := openai.CompletionRequest{
		Model:     openai.GPT3TextDavinci003,
		MaxTokens: 300,
		Prompt:    input,
	}
	resp, err := c.openaiClient.CreateCompletion(ctx, req)
	if err != nil {
		fmt.Printf("Completion error: %v\n", err)
	}
	fmt.Println("\tOpenAi response: %v: ", resp.Choices[0].Text)

	return resp.Choices[0].Text, nil
}
