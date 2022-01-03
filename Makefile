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
kubectl-bootstrap: ## Update used helm repositories
	kubectl create namespace ingress-nginx || true
	kubectl create namespace ndsquared || true

.PHONY: bootstrap
bootstrap: asdf-bootstrap k8s-bootstrap kubectl-bootstrap ## Perform all bootstrapping to start your project

.PHONY: clean
clean: ## Delete local dev environment
	k3d cluster delete $(PROJECT_NAME) || echo "No cluster found"
	k3d registry delete $(PROJECT_NAME).localhost || echo "No registry found"

.PHONY: up
up: bootstrap ## Run a local development environment
	tilt up --context k3d-$(PROJECT_NAME) --hud

.PHONY: down
down: ## Shutdown local development and free those resources
	tilt down --context k3d-$(PROJECT_NAME)

.PHONY: generic-secret
generic-secret: ## Create a generic secret to be used in conjunction with make seal-secret
	kubectl create secret generic secret --namespace changeme --dry-run=client --from-literal="foo=bar" -o yaml > secret.yaml

.PHONY: seal-secret
seal-secret: ## Seal a generic secret
	kubeseal --controller-name core-sealed-secrets --controller-namespace core -o yaml <secret.yaml >sealedsecret.yaml
