.DEFAULT_GOAL := help

.PHONY: help
help: ## View help information
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: shell
shell: ## Install dependencies and activate poetry virtual environment
	poetry install
	poetry shell

.PHONY: run
run: ## Run uvicorn web server
	API_ENVIRONMENT=development uvicorn app.main:app --host=0.0.0.0 --reload

.PHONY: format
format: ## Run formatter
	black .

.PHONY: format-check
format-check: ## Run formating check
	black --check .

.PHONY: lint
lint: ## Run linter
	flake8

.PHONY: tests
tests: ## Run tests
	API_ENVIRONMENT=testing pytest

.PHONY: cq
cq: format lint tests ## Run code quality tools

.PHONY: cq-check
cq-check: format-check lint tests ## Run code quality check
