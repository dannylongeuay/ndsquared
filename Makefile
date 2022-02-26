.DEFAULT_GOAL := help

PROJECT_NAME = ndsquared
REGISTRY_PORT = 37893

.PHONY: help
help: ## View help information
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: asdf-bootstrap
asdf-bootstrap: ## Install all tools through asdf-vm
	asdf plugin-add helm    || asdf install helm
	asdf plugin-add k3d     || asdf install k3d
	asdf plugin-add kubectl || asdf install kubectl
	asdf plugin-add tilt    || asdf install tilt

.PHONY: k8s-bootstrap
k8s-bootstrap: ## Create a Kubernetes cluster for local development
	k3d registry create $(PROJECT_NAME).localhost --port=$(REGISTRY_PORT) || echo "Registry already exists"
	k3d cluster create $(PROJECT_NAME) --registry-use k3d-$(PROJECT_NAME).localhost:$(REGISTRY_PORT) -p "8000:80@loadbalancer" || echo "Cluster already exists"

.PHONY: kubectl-bootstrap
kubectl-bootstrap: ## Create namespaces in k3d cluster
	kubectl create namespace ingress-nginx || true
	kubectl create namespace local || true

.PHONY: bootstrap
bootstrap: asdf-bootstrap k8s-bootstrap kubectl-bootstrap ## Perform all bootstrapping required for local development

.PHONY: clean
clean: ## Destroy local development environment
	k3d cluster delete $(PROJECT_NAME) || echo "No cluster found"
	k3d registry delete $(PROJECT_NAME).localhost || echo "No registry found"

.PHONY: up
up: bootstrap ## Run local development environment
	tilt up --context k3d-$(PROJECT_NAME) --hud

.PHONY: down
down: ## Stop local development environment
	tilt down --context k3d-$(PROJECT_NAME)

.PHONY: gitlab-secret
gitlab-secret: ## Create a gitlab secret to be used with the external secrets store
	cp gitlab-secret.example.yaml gitlab-secret.yaml
