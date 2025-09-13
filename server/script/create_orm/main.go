package main

import (
	"github.com/hankeyyh/chat-box-svr/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// table to struct
func main() {
	mysqlConf := conf.DefaultConf.MysqlConf
	gormdb, err := gorm.Open(mysql.Open(mysqlConf.GetDsn()))
	if err != nil {
		panic(err)
	}
	
	g := gen.NewGenerator(gen.Config{
		OutPath: "dao",
	})
	g.UseDB(gormdb)
	g.ApplyBasic(
		g.GenerateModelAs("chat_history", "ChatHistory2"),
	)
	g.Execute()
}
