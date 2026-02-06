# Development Container

This directory contains the configuration for a VS Code development container that provides a consistent development environment for this project.

## Features

The devcontainer includes:

- **Go 1.21**: The Go programming language and toolchain
- **Azure CLI**: Command-line interface for managing Azure resources
- **kubectl**: Kubernetes command-line tool for cluster management
- **Docker-in-Docker**: Ability to build and run Docker containers from within the devcontainer
- **Helm**: Kubernetes package manager (included with kubectl feature)

## VS Code Extensions

The following VS Code extensions are automatically installed:

- **Go**: Official Go language support
- **Kubernetes Tools**: Enhanced Kubernetes support for VS Code
- **Azure Resource Groups**: Azure resource management tools

## Using the Devcontainer

### Prerequisites

- [Visual Studio Code](https://code.visualstudio.com/)
- [Docker Desktop](https://www.docker.com/products/docker-desktop) (or Docker Engine on Linux)
- [Dev Containers extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)

### Opening the Project

1. Clone the repository
2. Open the folder in VS Code
3. When prompted, click "Reopen in Container" (or use Command Palette: "Dev Containers: Reopen in Container")
4. Wait for the container to build (first time will take a few minutes)

### Post-Creation Setup

The devcontainer automatically runs `go mod download` after creation to download Go dependencies.

## Available Tools

Once inside the devcontainer, you can use:

```bash
# Go development
go run main.go
go build
go test ./...

# Kubernetes
kubectl version
kubectl get pods
kubectl apply -f k8s/

# Azure CLI
az --version
az login
az account list

# Docker
docker build -t vpa-demo:latest .
docker run vpa-demo:latest
```

## Customization

To customize the devcontainer:

- Edit `devcontainer.json` to add features, extensions, or change settings
- Edit `Dockerfile` to install additional system packages or tools
- Rebuild the container: Command Palette → "Dev Containers: Rebuild Container"
