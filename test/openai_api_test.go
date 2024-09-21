package test

import (
	"context"
	"errors"
	"fmt"
	"io"
	"testing"

	"github.com/sashabaranov/go-openai"
)

var key = "sk-mRPg1ANRkqBB3JAxEb3185899519477dB896B558203e6183"
var baseUrl = "https://api.b3n.fun/v1"

func TestStream(t *testing.T) {
	conf := openai.DefaultConfig(key)
	conf.BaseURL = baseUrl

	client := openai.NewClientWithConfig(conf)
	req := openai.ChatCompletionRequest{
		Model: "gpt-4o-mini",
		MaxTokens: 20,
		Messages: []openai.ChatCompletionMessage{
			{
				Role: openai.ChatMessageRoleAssistant,
				Content: "Lorem ipsum",
			},
		},
		Stream: true,
	}
	stream, err := client.CreateChatCompletionStream(context.Background(), req)
	if err != nil {
		panic(err)
	}
	defer stream.Close()
	
	fmt.Println("Stream response: ")
	for {
		rsp, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			return
		}
		if err != nil {
			panic(err)
		}
		fmt.Println(rsp.Choices[0].Delta.Content)
	}
}