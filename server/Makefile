dep: ## get dependices
	@echo "go dep..."
	@go mod tidy

model: dep # create model
	@echo "create model..."
	@go run script/create_model/main.go

init_model: dep # init model
	@echo "init model..."
	@go run script/init_model/main.go --table_name=all

build: dep # build server
	@echo "build server..."
	@go build -o cmd/server cmd/main.go

run: dep # run server
	@echo "run server..."
	@go run cmd/main.go