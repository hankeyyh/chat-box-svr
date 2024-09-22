dep: ## get dependices
	@echo "go dep..."
	@go mod tidy

table: dep # create table
	@echo "create table..."
	@go run script/create_table/main.go

table_force: dep # force create table
	@echo "force create table..."
	@go run script/create_table/main.go --force

crud: dep # create crud
	@echo "create crud..."
	@go run script/create_crud/main.go

init_model: dep # init model
	@echo "init model..."
	@go run script/init_model/main.go --table_name=all

build: dep # build server
	@echo "build server..."
	@go build -o cmd/server cmd/main.go

run: dep # run server
	@echo "run server..."
	@go run cmd/main.go