package main

import (
	"flag"
	"fmt"

	"github.com/hankeyyh/chat-box-svr/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var tableName = flag.String("table_name", "", "table to initialize")

func main() {
	flag.Parse()
	
	mysqlConf := conf.DefaultConf.MysqlConf
	gormdb, err := gorm.Open(mysql.Open(mysqlConf.GetDsn()))
	if err != nil {
		panic(err)
	}
	switch *tableName {
		case "ai_model":
			initAiModel(gormdb)
		case "app":
			initApp(gormdb)
		case "all":
			initAiModel(gormdb)
			initApp(gormdb)
		default:
			fmt.Println("table name not found")
			return
	}
}

func initAiModel(db *gorm.DB) {
	// fill data
	err := db.Exec(`
		INSERT INTO ai_model (name, enabled, max_output_token)
		VALUES ('gpt-4o-mini', 1, 100), ('gpt-4o', 1, 100);
	`).Error
	fmt.Printf("init ai_model done, err: %v\n", err)
}

func initApp(db *gorm.DB) {
	err := db.Exec(`
		INSERT INTO app (model_id, name, created_by, introduction, prompt, is_public)
		VALUES (1, '变量命名专家', 'yuhanyang', '擅长生成变量名和函数名', '角色\n你是一个英语纯熟的计算机程序员。你的主要特长是根据功能描述为用户产生变量名或函数名。\n\n技能\n技能 1: 生成变量名\n细读用户提供的功描述。\n根据描述选取关键词，转化成英文（如果用户提供的是非英文描述）\n基于这些关键词，构建符合命名规范的变量名。示例格式：\n=====\n<!---->\n变量名: <variable name>\n====\n\n技能 2: 生成函数名\n细读用户提供的功描述。\n取出描述中的动作或动词部分，转化成英文（如果用户提供的是非英文描述）\n根据这些关键词，构建符合规范的函数名。示例格式：\n=====\n<!---->\n函数名: <function name>\n=====\n\n限制\n只解答与变量命名和函数命名相关的问题。如果用户提问其他问题，不进行回答。\n使用与原始提示一致的语言进行回答。\n使用用户使用的语言进行回答。\n直接以优化的提示开始你的回答。', 1)
	`).Error
	fmt.Printf("init app done, err: %v\n", err)
}
