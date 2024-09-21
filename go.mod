module github.com/hankeyyh/chat-box-svr

go 1.22.0

require (
	github.com/BurntSushi/toml v1.4.0
	github.com/sashabaranov/go-openai v1.30.3
	gorm.io/driver/mysql v1.5.7
	gorm.io/gorm v1.25.12
	gorm.io/plugin/dbresolver v1.5.0
)

require (
	golang.org/x/mod v0.17.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/tools v0.21.1-0.20240508182429-e35e4ccd0d2d // indirect
	gorm.io/datatypes v1.1.1-0.20230130040222-c43177d3cf8c // indirect
	gorm.io/hints v1.1.0 // indirect
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.8.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/text v0.18.0 // indirect
	gorm.io/gen v0.3.26
)

replace gorm.io/plugin/dbresolver v1.5.0 => github.com/go-gorm/dbresolver v1.5.3
