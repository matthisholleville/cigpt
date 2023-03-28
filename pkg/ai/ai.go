package ai

import (
	"context"
	"errors"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

const (
	default_prompt = "Behave like a GitlabCI expert. My pipeline is failing. Here are the logs: %s. Explain to me in detail the solution in %s."
)

type OpenAIClient struct {
	client   *openai.Client
	language string
}

func (c *OpenAIClient) Configure(token string, language string) error {
	client := openai.NewClient(token)
	if client == nil {
		return errors.New("error creating openai client.")
	}
	c.client = client
	c.language = language
	return nil
}

func (c *OpenAIClient) GetCompletion(ctx context.Context, prompt string) (string, error) {
	resp, err := c.client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "user",
				Content: fmt.Sprintf(default_prompt, c.language, prompt),
			},
		},
	})
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}
