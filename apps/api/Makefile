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
	uvicorn app.main:app --host=0.0.0.0 --reload
