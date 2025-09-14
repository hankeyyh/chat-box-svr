module github.com/hankeyyh/chat-box-svr

go 1.24

require (
	github.com/BurntSushi/toml v1.4.0
	github.com/joho/godotenv v1.5.1
	github.com/sashabaranov/go-openai v1.30.3
	google.golang.org/genai v1.24.0
	google.golang.org/grpc v1.66.2
	google.golang.org/protobuf v1.34.2
	gorm.io/driver/mysql v1.5.7
	gorm.io/gorm v1.25.12
	gorm.io/plugin/dbresolver v1.5.0
)

require (
	cloud.google.com/go v0.116.0 // indirect
	cloud.google.com/go/auth v0.9.3 // indirect
	cloud.google.com/go/compute/metadata v0.5.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/s2a-go v0.1.8 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.4 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	go.opencensus.io v0.24.0 // indirect
	golang.org/x/crypto v0.27.0 // indirect
	golang.org/x/mod v0.17.0 // indirect
	golang.org/x/net v0.29.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/sys v0.25.0 // indirect
	golang.org/x/tools v0.21.1-0.20240508182429-e35e4ccd0d2d // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240903143218-8af14fe29dc1 // indirect
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
