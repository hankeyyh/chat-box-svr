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
	GetByID(id uint64) ([]gen.T, error)
	// SELECT * FROM @@table
	All() ([]gen.T, error)
}

type AppQuery interface {
	// SELECT * FROM @@table WHERE name = @name
	GetByName(name string) ([]gen.T, error)
	// SELECT * FROM @@table WHERE model_id = @modelId
	GetByModelID(modelId uint64) ([]gen.T, error)
	// SELECT * FROM @@table WHERE id = @id LIMIT 1
	GetByID(id uint64) ([]gen.T, error)
	// SELECT * FROM @@table WHERE created_by = @createdBy and id = @id LIMIT 1
	GetByAuthorAndId(createdBy uint64, id uint64) (gen.T, error)
	// SELECT * FROM @@table WHERE is_public = 1
	AllPublic() ([]gen.T, error)
	// SELECT * FROM @@table WHERE is_public = 0 AND created_by = @createdBy
	AllPrivateByAuthor(createdBy uint64) ([]gen.T, error)
	// UPDATE @@table SET is_public = @isPublic WHERE id = @id
	UpdateIsPublic(id uint64, isPublic bool) error
}

type ChatHistoryQuery interface {
	// SELECT * FROM @@table WHERE id = @id LIMIT 1
	GetByID(id uint64) ([]gen.T, error)
	// SELECT * FROM @@table WHERE parent_id = @parentId
	GetByParentID(parentId uint64) ([]gen.T, error)
	// SELECT * FROM @@table WHERE 
	// {{if lastId != 0}}
	// 	id < @lastId AND 
	// {{end}}
	// app_id = @appId AND user_id = @userId ORDER BY id DESC LIMIT @offset, @limit
	BatchGetRecentByUserID(appId, userId uint64, lastId uint64, offset int, limit int) ([]gen.T, error)
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
	if err = createApi(gormdb, g, model.ChatHistory{}, func(ChatHistoryQuery) {}); err != nil {
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