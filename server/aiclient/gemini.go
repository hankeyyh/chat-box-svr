package aiclient

import (
	"context"

	"github.com/hankeyyh/chat-box-svr/util/log"
	"google.golang.org/genai"
)

type GeminiClient struct {
	client *genai.Client
	option *Option
}

func NewGeminiClient(ctx context.Context, opts ...OptionFunc) (*GeminiClient, error) {
	option := &Option{}
	for _, opt := range opts {
		opt(option)
	}
	if option.Model == "" {
		option.Model = "gemini-2.5-flash"
	}
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: option.APIKey,
		HTTPOptions: genai.HTTPOptions{
			BaseURL: option.BaseURL,
		},
	})
	if err != nil {
		return nil, err
	}
	return &GeminiClient{
		client: client,
		option: option,
	}, nil
}

func (c *GeminiClient) buildContentConfig() *genai.GenerateContentConfig {
	result := &genai.GenerateContentConfig{}
	if c.option.ThinkingBudget > 0 {
		result.ThinkingConfig = &genai.ThinkingConfig{
			ThinkingBudget: &c.option.ThinkingBudget,
		}
	}
	if c.option.Temperature > 0 {
		result.Temperature = &c.option.Temperature
	}
	return result
} 

func (c *GeminiClient) ChatStream(ctx context.Context, text string) (<-chan *StreamResponse, error) {
	streamConf := c.buildContentConfig()
	chat, err := c.client.Chats.Create(ctx, c.option.Model, streamConf, nil)
	if err != nil {
		log.Errorf("failed to create chat: %v", err)
		return nil, err
	}
	part := genai.Part{Text: text}
	p := make([]genai.Part, 1)
	p[0] = part

	// send chat message
	result := make(chan *StreamResponse)
	go func() {
		defer close(result)
		for geminiResult, err := range chat.SendMessageStream(ctx, p...) {
			if err != nil {
				log.Errorf("failed to send chat message: %v", err)
			}
			response := &StreamResponse{
				Text: geminiResult.Text(),
			}
			select {
			case <-ctx.Done():
				return 
			case result <- response:
			}
		}
	}()

	return result, nil
}