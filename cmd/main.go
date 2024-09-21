package main

import (
	"fmt"

	"github.com/hankeyyh/chat-box-svr/dao"
)


func main() {
	aimodel, err := dao.AiModel.GetByName("gpt-4o-mini")
	if err != nil {
		panic(err)
	}
	fmt.Println(aimodel)

	
}