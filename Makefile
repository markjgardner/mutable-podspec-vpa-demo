.PHONY: build run test docker-build docker-run clean help

BINARY_NAME=vpa-demo
IMAGE_NAME=vpa-demo:latest

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build the Go binary
	go build -o $(BINARY_NAME) main.go

run: build ## Run the application locally
	./$(BINARY_NAME)

test: build ## Test the application with custom settings
	MEMORY_ALLOC_MB=100 WAIT_SECONDS=5 timeout 15 ./$(BINARY_NAME) || true

docker-build: ## Build the Docker image
	docker build -t $(IMAGE_NAME) .

docker-run: docker-build ## Run the Docker container
	docker run --rm -e MEMORY_ALLOC_MB=100 -e WAIT_SECONDS=5 $(IMAGE_NAME)

clean: ## Clean up built artifacts
	rm -f $(BINARY_NAME)
	go clean

k8s-deploy: ## Deploy to Kubernetes
	kubectl apply -f k8s/pod.yaml
	kubectl apply -f k8s/vpa.yaml

k8s-delete: ## Delete from Kubernetes
	kubectl delete -f k8s/vpa.yaml --ignore-not-found
	kubectl delete -f k8s/pod.yaml --ignore-not-found

k8s-logs: ## View pod logs
	kubectl logs -f vpa-demo-app

k8s-status: ## Check pod and VPA status
	@echo "Pod status:"
	kubectl get pod vpa-demo-app
	@echo "\nVPA status:"
	kubectl describe vpa vpa-demo-vpa
