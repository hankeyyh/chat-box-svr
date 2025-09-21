package main

import (
	"github.com/hankeyyh/chat-box-svr/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:      "model",
		ModelPkgPath: "model/table",
		Mode:         gen.WithDefaultQuery,
	})
	gormdb, err := gorm.Open(mysql.Open(conf.DefaultConf.MysqlConf.GetDsn()))
	if err != nil {
		panic(err)
	}
	g.UseDB(gormdb)
	g.ApplyBasic(
		g.GenerateModel("app", 
			gen.FieldType("id", "uint64"),
			gen.FieldType("model_id", "uint64"),
			gen.FieldType("created_by", "uint64")),

		g.GenerateModel("ai_model", 
			gen.FieldType("id", "uint64"),
			gen.FieldType("max_output_tokens", "uint64")),

		g.GenerateModel("chat_history",
			gen.FieldType("id", "uint64"),
			gen.FieldType("parent_id", "uint64"),
			gen.FieldType("user_id", "uint64"),
			gen.FieldType("session_id", "uint64"),
			gen.FieldType("app_id", "uint64")),
			

		g.GenerateModel("session", 
			gen.FieldType("id", "uint64"),
			gen.FieldType("user_id", "uint64")),

		g.GenerateModel("session_v2",
			gen.FieldType("id", "uint64"),
			gen.FieldType("user_id", "uint64"),
			gen.FieldType("app_id", "uint64")),
	)
	g.Execute()
}
