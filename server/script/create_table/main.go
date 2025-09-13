package main

import (
	"flag"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/hankeyyh/chat-box-svr/conf"
	"github.com/hankeyyh/chat-box-svr/model"
)

var forceCreate = flag.Bool("force", false, "force create table")

// struct to table
func main() {
	flag.Parse()

	mysqlConf := conf.DefaultConf.MysqlConf
	gormdb, err := gorm.Open(mysql.Open(mysqlConf.GetDsn()))
	if err != nil {
		panic(err)
	}

	// create table if table not exists
	if err = createModel(gormdb, model.AiModel{}); err != nil {
		panic(err)
	}
	if err = createModel(gormdb, model.App{}); err != nil {
		panic(err)
	}
	if err = createModel(gormdb, model.ChatHistory{}); err != nil {
		panic(err)
	}
	if err = createModel(gormdb, model.Session{}); err != nil {
		panic(err)
	}
}

func createModel(db *gorm.DB,  model model.BaseTable) error {
	// create table if table not exists
	if !db.Migrator().HasTable(&model) || *forceCreate {
		if err := db.Migrator().DropTable(&model); err != nil {
			return err
		}
		err := db.AutoMigrate(&model)
		if err != nil {
			return err
		}
		fmt.Printf("Table %s created\n", model.TableName())
	} else {
		fmt.Printf("Table %s already exists\n", model.TableName())
	}
	return nil
}