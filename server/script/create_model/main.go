package main

import (
	"github.com/hankeyyh/chat-box-svr/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "model",
		ModelPkgPath: "model/table",
		Mode: gen.WithDefaultQuery,
	})
	gormdb, err := gorm.Open(mysql.Open(conf.DefaultConf.MysqlConf.GetDsn()))
	if err != nil {
		panic(err)
	}
	g.UseDB(gormdb)
	g.ApplyBasic(
		g.GenerateModel("app"),
		g.GenerateModel("ai_model"),
		g.GenerateModel("chat_history"),
		g.GenerateModel("session"),
		g.GenerateModel("session_v2"),
	)
	g.Execute()
}