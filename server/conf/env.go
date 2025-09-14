package conf

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {
	// 加载 .env 文件
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}