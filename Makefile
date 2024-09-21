dep: ## get dependices
	@echo "go dep..."
	@go mod tidy

db: dep # create db
	@echo "create db..."
	@go run model/gen/main.go
