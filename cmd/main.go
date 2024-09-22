package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hankeyyh/chat-box-svr/conf"
	"github.com/hankeyyh/chat-box-svr/logic"
)

func main() {
	serverConf := conf.DefaultConf.ServerConf
	addr := fmt.Sprintf(":%d", serverConf.Port)

	http.HandleFunc("/model/list", logic.HandleGetFormRequest(logic.ModelList))
	http.HandleFunc("/app/public-list", logic.HandleGetFormRequest(logic.AppPublicList))
	http.HandleFunc("/app/private-list", logic.HandleGetFormRequest(logic.AppPrivateList))
	http.HandleFunc("/app/detail", logic.HandleGetFormRequest(logic.AppDetail))
	http.HandleFunc("/app/upsert", logic.HandlePostJsonRequest(logic.AppUpsert))
	http.HandleFunc("/app/release", logic.HandlePostJsonRequest(logic.AppRelease))
	http.HandleFunc("/app/unrelease", logic.HandlePostJsonRequest(logic.AppUnrelease))
	http.HandleFunc("/app/chat-list", logic.HandleGetFormRequest(logic.AppChatList))
	http.HandleFunc("/app/chat", logic.AppChat)

	fmt.Printf("Server started at %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
	fmt.Println("Server closed")
}
