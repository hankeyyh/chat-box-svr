package main

import (
	"fmt"
	"net"
	"github.com/hankeyyh/chat-box-svr/util/log"
	"github.com/hankeyyh/chat-box-svr/conf"
	"github.com/hankeyyh/chat-box-svr/pb/chat"
	"google.golang.org/grpc"
	"github.com/hankeyyh/chat-box-svr/service"
)



func main() {
	serverConf := conf.DefaultConf.ServerConf
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", serverConf.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	chat.RegisterChatServiceServer(s, &service.ChatService{})
	log.Infof("server listening at %v", ln.Addr())
	if err := s.Serve(ln); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}