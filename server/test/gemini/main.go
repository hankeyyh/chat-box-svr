package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/genai"
	"google.golang.org/protobuf/proto"
)

func main() {
	// 加载 .env 文件
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("GEMINI_API_KEY is not set")
	} else {
		fmt.Printf("GEMINI_API_KEY is %s\n", apiKey)
	}

	ctx := context.Background()
	// The client gets the API key from the environment variable `GEMINI_API_KEY`.
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: apiKey,
		HTTPOptions: genai.HTTPOptions{
			BaseURL: "https://gemini-apihankeyyh.work",
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	stream := client.Models.GenerateContentStream(
		ctx,
		"gemini-2.5-flash",
		genai.Text("讲一个睡前故事"),
		&genai.GenerateContentConfig{
			ThinkingConfig: &genai.ThinkingConfig{
				ThinkingBudget: proto.Int32(0), // Disables thinking
			},
		},
	)

	for chunk := range stream {
		part := chunk.Candidates[0].Content.Parts[0]
		fmt.Print(part.Text)
	}
}
