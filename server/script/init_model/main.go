package main

import (
	"context"
	"flag"
	"fmt"

	dao_aimodel "github.com/hankeyyh/chat-box-svr/dao/ai_model"
	dao_app "github.com/hankeyyh/chat-box-svr/dao/app"
	"github.com/hankeyyh/chat-box-svr/model/table"
	"github.com/hankeyyh/chat-box-svr/conf"
)

var tableName = flag.String("table_name", "", "table to initialize")

func main() {
	flag.Parse()
	
	switch *tableName {
		case "ai_model":
			initAiModel()
		case "app":
			initApp()
		case "all":
			initAiModel()
			initApp()
		default:
			fmt.Println("table name not found")
			return
	}
}

func initAiModel() {
	err := dao_aimodel.Create(context.Background(), &table.AiModel{
		ID: 1,
		Name: "gemini-2.5-flash",
		Enabled: 1,
	})
	if err != nil {
		fmt.Printf("init ai_model failed, err: %v\n", err)
		return
	}
	fmt.Printf("init ai_model done, err: %v\n", err)
}

func initApp() {
	err := dao_app.Create(context.Background(), &table.App{
		ID: conf.DefaultConf.ChatConf.DefaultAppId,
		ModelID: 1,
		Name:"默认聊天",
		CreatedBy: 1,
		Temperature: 0.5,
		MaxOutputTokens: 2048,
	})
	if err != nil {
		fmt.Printf("init app failed, err: %v\n", err)
		return
	}
	fmt.Printf("init app done, err: %v\n", err)
}
