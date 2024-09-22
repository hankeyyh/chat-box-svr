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
	// SELECT * FROM @@table WHERE id = @id LIMIT 1
	GetByID(id int) ([]gen.T, error)
	// SELECT * FROM @@table
	All() ([]gen.T, error)
}

type AppQuery interface {
	// SELECT * FROM @@table WHERE name = @name
	GetByName(name string) ([]gen.T, error)
	// SELECT * FROM @@table WHERE model_id = @modelId
	GetByModelID(modelId int) ([]gen.T, error)
	// SELECT * FROM @@table WHERE id = @id LIMIT 1
	GetByID(id int) ([]gen.T, error)
	// SEECT * FROM @@table WHERE created_by = @createdBy
	GetByAuthor(createdBy string) ([]gen.T, error)
	// SELECT * FROM @@table WHERE is_public = 1
	AllPublic() ([]gen.T, error)
	// SELECT * FROM @@table WHERE is_public = 0 AND created_by = @createdBy
	AllPrivateByAuthor(createdBy string) ([]gen.T, error)
	// UPDATE @@table SET is_public = @isPublic WHERE id = @id
	UpdateIsPublic(id int, isPublic bool) error
}

func main() {
	mysqlConf := conf.DefaultConf.MysqlConf
	gormdb, err := gorm.Open(mysql.Open(mysqlConf.GetDsn()))
	if err != nil {
		panic(err)
	}
	g := gen.NewGenerator(gen.Config{
		OutPath: "dao",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	// create table if table not exists
	if err = createApi(gormdb, g, model.AiModel{}, func(AiModelQuery) {}); err != nil {
		panic(err)
	}
	if err = createApi(gormdb, g, model.App{}, func(AppQuery) {}); err != nil {
		panic(err)
	}
}

func createApi(db *gorm.DB, gen *gen.Generator, model model.BaseTable, fc interface{}) error {
	// create api
	gen.UseDB(db)
	gen.ApplyBasic(model)
	gen.ApplyInterface(fc, model)
	gen.Execute()
	fmt.Printf("%s crud generated\n", model.TableName())
	return nil
}