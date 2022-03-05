.DEFAULT_GOAL := help

.PHONY: help
help: ## View help information
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: docs ## Build go binary
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/goapi ndsquared/goapi/src

.PHONY: docs
docs: ## Generate Swagger docs
	swag init -d src/ -ot go

.PHONY: tests
tests: ## Run tests
	go test -v ndsquared/goapi/src

.PHONY: lint
lint: ## Run linter
	golangci-lint run --fast

.PHONY: format
format: ## Run formatter 
	go fmt ndsquared/goapi/src
	swag fmt -d src/

.PHONY: cq
cq: format tests lint ## Run code quality tools

.PHONY: cq-check
cq-check: lint tests ## Run code quality check

.PHONY: run-prod
run-prod: build ## Run prod binary
	GIN_MODE=release ./build/goapi

.PHONY: run-dev
run-dev: build ## Run dev binary
	GIN_MODE=debug ./build/goapi
