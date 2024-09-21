package main

import (
	"fmt"

	"github.com/hankeyyh/chat-box-svr/conf"
	"github.com/hankeyyh/chat-box-svr/model"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

type AiModelQuery interface {
	// SELECT * FROM @@table WHERE name = @name
	GetByName(name string) ([]gen.T, error)
	// SELECT * FROM @@table
	All() ([]gen.T, error)
}

type AppQuery interface {
	// SELECT * FROM @@table WHERE name = @name
	GetByName(name string) ([]gen.T, error)
	// SELECT * FROM @@table WHERE model_id = @modelId
	GetByModelID(modelId int) ([]gen.T, error)
	// SEECT * FROM @@table WHERE created_by = @createdBy
	GetByAuthor(createdBy string) ([]gen.T, error)
	// SELECT * FROM @@table WHERE is_public = 1
	AllPublic() ([]gen.T, error)
}

func main() {
	mysqlConf := conf.DefaultConf.MysqlConf
	gormdb, err := gorm.Open(mysql.Open(mysqlConf.GetDsn()))
	if err != nil {
		panic(err)
	}
	g := gen.NewGenerator(gen.Config{
		OutPath: "dao",
		Mode: gen.WithoutContext|gen.WithDefaultQuery|gen.WithQueryInterface,
	})

	// create table if table not exists
	if err = createModelAndApi(gormdb, g, model.AiModel{}, func(AiModelQuery){}); err != nil {
		panic(err)
	}
	if err = createModelAndApi(gormdb, g, model.App{}, func(AppQuery){}); err != nil {
		panic(err)
	}

	// fill data
	// if err = fillAiModel(gormdb); err != nil {
	// 	panic(err)
	// }
}

func createModelAndApi(db *gorm.DB, gen *gen.Generator, model model.BaseTable, fc interface{}) error {
	// create table if table not exists
	if !db.Migrator().HasTable(&model) {
		err := db.AutoMigrate(&model)
		if err != nil {
			return err
		}
		fmt.Printf("Table %s created\n", model.TableName())
	} else {
		fmt.Printf("Table %s already exists\n", model.TableName())
	}

	// create api
	gen.UseDB(db)
	gen.ApplyBasic(model)
	gen.ApplyInterface(fc, model)
	gen.Execute()
	return nil
}

func fillAiModel(db *gorm.DB) error {
	// fill data
	err := db.Exec(`
		INSERT INTO ai_model (name, enabled, max_output_token)
		VALUES ('gpt-4o-mini', 1, 100), ('gpt-4o', 1, 100);
	`).Error
	if err != nil {
		return err
	}
	return nil
}