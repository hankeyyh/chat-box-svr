dep: ## get dependices
	@echo "go dep..."
	@go mod tidy

db: dep # create db
	@echo "create db..."
	@go run model/gen/main.go

init_model: dep # init model
	@echo "init model..."
	@go run model/gen/main.go --init_model

build: dep # build server
	@echo "build server..."
	@go build -o cmd/server cmd/main.go

run: dep # run server
	@echo "run server..."
	@go run cmd/main.go