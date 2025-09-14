package service

import (
	"context"
	"os"
	"sync"

	"github.com/hankeyyh/chat-box-svr/aiclient"
	"github.com/hankeyyh/chat-box-svr/pb/chat"
	"github.com/hankeyyh/chat-box-svr/util/log"
	"google.golang.org/grpc"
)

type ChatService struct {
	chat.UnimplementedChatServiceServer
}

func (s *ChatService) ChatStream(req *chat.ChatRequest, stream grpc.ServerStreamingServer[chat.ChatResponse]) error {
	log.Debugf("ChatStream, request: %v", req)

	ctx := context.Background()

	aiClient, err := aiclient.NewGeminiClient(ctx, 
		aiclient.WithAPIKey(os.Getenv("GEMINI_API_KEY")), 
		aiclient.WithBaseURL("https://gemini-apihankeyyh.work"),
		aiclient.WithModel("gemini-2.5-flash"),
		aiclient.WithTemperature(0.5),
	)
	if err != nil {
		log.Errorf("failed to create ai client: %v", err)
		return err
	}
	ch, err := aiClient.ChatStream(ctx, req.GetText())
	if err != nil {
		log.Errorf("failed to create chat stream: %v", err)
		return err
	}

	var once sync.Once

	for chunk := range ch {
		once.Do(func() {
			log.Debugf("ChatStream, start streaming")
		})
		stream.Send(&chat.ChatResponse{
			Code: 0,
			Message: "success",
			Data: &chat.ChatResponseData{
				Text: chunk.Text,
			},
		})
	}

	return nil
}