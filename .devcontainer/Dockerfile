# Use the official Go image as base
FROM mcr.microsoft.com/devcontainers/go:1-1.21-bullseye

# Install additional tools
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends \
    git \
    curl \
    ca-certificates \
    && apt-get autoremove -y && apt-get clean -y && rm -rf /var/lib/apt/lists/*

# The azure-cli and kubectl will be installed via devcontainer features
