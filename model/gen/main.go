package main

import (
	"flag"
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
	// SELECT * FROM @@table WHERE id = @id LIMIT 1
	GetByID(id int) ([]gen.T, error)
	// SEECT * FROM @@table WHERE created_by = @createdBy
	GetByAuthor(createdBy string) ([]gen.T, error)
	// SELECT * FROM @@table WHERE is_public = 1
	AllPublic() ([]gen.T, error)
	// SELECT * FROM @@table WHERE is_public = 0 AND created_by = @createdBy
	AllPrivateByAuthor(createdBy string) ([]gen.T, error)
}

var initModelData = flag.Bool("init_model", false, "init model data")
var forceCreate = flag.Bool("force", false, "force create table")

func main() {
	flag.Parse()

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
	if err = createModelAndApi(gormdb, g, model.AiModel{}, func(AiModelQuery) {}); err != nil {
		panic(err)
	}
	if err = createModelAndApi(gormdb, g, model.App{}, func(AppQuery) {}); err != nil {
		panic(err)
	}

	// fill data
	if *initModelData {
		if err = initAiModel(gormdb); err != nil {
			fmt.Printf("Fill ai_model failed: %v\n", err)
		}
		fmt.Println("ai_model initialized")

		if err = initApp(gormdb); err != nil {
			fmt.Printf("Fill app failed: %v\n", err)
		}
		fmt.Println("app initialized")
	}
}

func createModelAndApi(db *gorm.DB, gen *gen.Generator, model model.BaseTable, fc interface{}) error {
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

	// create api
	gen.UseDB(db)
	gen.ApplyBasic(model)
	gen.ApplyInterface(fc, model)
	gen.Execute()
	return nil
}

func initAiModel(db *gorm.DB) error {
	// fill data
	return db.Exec(`
		INSERT INTO ai_model (name, enabled, max_output_token)
		VALUES ('gpt-4o-mini', 1, 100), ('gpt-4o', 1, 100);
	`).Error
}

func initApp(db *gorm.DB) error {
	return db.Exec(`
		INSERT INTO app (model_id, name, created_by, introduction, prompt, is_public)
		VALUES (1, '变量命名专家', 'yuhanyang', '擅长生成变量名和函数名', '角色\n你是一个英语纯熟的计算机程序员。你的主要特长是根据功能描述为用户产生变量名或函数名。\n\n技能\n技能 1: 生成变量名\n细读用户提供的功描述。\n根据描述选取关键词，转化成英文（如果用户提供的是非英文描述）\n基于这些关键词，构建符合命名规范的变量名。示例格式：\n=====\n<!---->\n变量名: <variable name>\n====\n\n技能 2: 生成函数名\n细读用户提供的功描述。\n取出描述中的动作或动词部分，转化成英文（如果用户提供的是非英文描述）\n根据这些关键词，构建符合规范的函数名。示例格式：\n=====\n<!---->\n函数名: <function name>\n=====\n\n限制\n只解答与变量命名和函数命名相关的问题。如果用户提问其他问题，不进行回答。\n使用与原始提示一致的语言进行回答。\n使用用户使用的语言进行回答。\n直接以优化的提示开始你的回答。', 1)
	`).Error
}
