# VPA InPlaceOrRecreate Demo

Demonstrating in-place resizing of a pod using VPA (Vertical Pod Autoscaler) on Kubernetes 1.33+

## Overview

This demo application showcases the `InPlaceOrRecreate` update policy of the Vertical Pod Autoscaler. The application allocates a large chunk of memory at startup, holds it for a configurable duration, then releases it to demonstrate how VPA can adjust pod resources in-place without recreating the pod.

## Application Behavior

The Go application:
1. Allocates a configurable amount of memory at startup (default: 512 MB)
2. Waits for a configurable duration (default: 30 seconds)
3. Releases the memory and forces garbage collection
4. Continues running with minimal resource usage
5. Prints memory statistics periodically

## Configuration

The application is configured via environment variables:

- `MEMORY_ALLOC_MB`: Amount of memory to allocate in megabytes (default: 512)
- `WAIT_SECONDS`: Time to wait before releasing memory in seconds (default: 30)

## Building the Application

### Build the Docker Image

```bash
docker build -t vpa-demo:latest .
```

### Run Locally

```bash
# Run with defaults
go run main.go

# Run with custom configuration
MEMORY_ALLOC_MB=1024 WAIT_SECONDS=60 go run main.go
```

## Kubernetes Deployment

### Prerequisites

- Kubernetes 1.33 or later (with in-place pod resize feature enabled)
- VPA (Vertical Pod Autoscaler) installed in the cluster
- The `InPlacePodVerticalScaling` feature gate must be enabled

### Deploy the Application

1. Load the Docker image into your cluster (e.g., for kind or minikube):
   ```bash
   docker build -t vpa-demo:latest .
   # For kind:
   kind load docker-image vpa-demo:latest
   # For minikube:
   minikube image load vpa-demo:latest
   ```

2. Deploy the pod:
   ```bash
   kubectl apply -f k8s/pod.yaml
   ```

3. Deploy the VPA:
   ```bash
   kubectl apply -f k8s/vpa.yaml
   ```

### Monitor the Demo

Watch the pod's resource allocation:
```bash
kubectl get pod vpa-demo-app -w
```

View the application logs:
```bash
kubectl logs -f vpa-demo-app
```

Check VPA recommendations:
```bash
kubectl describe vpa vpa-demo-vpa
```

### Expected Behavior

1. The pod starts with low resource requests (128Mi memory, 100m CPU)
2. The application allocates 800MB of memory (configured in pod.yaml)
3. VPA observes the high memory usage and increases the pod's memory request in-place
4. After 60 seconds, the application releases the memory
5. VPA observes the decreased memory usage and may reduce the pod's memory request in-place

## VPA Configuration

The VPA is configured with:
- **Update Mode**: `InPlaceOrRecreate` - Attempts in-place resize first, falls back to pod recreation if needed
- **Resource Policies**:
  - Min memory: 64Mi, Max memory: 2Gi
  - Min CPU: 50m, Max CPU: 1000m
- **Controlled Resources**: Both memory and CPU

## Pod Resize Policy

The pod is configured with `resizePolicy` set to `NotRequired` for both memory and CPU, allowing the resources to be resized without restarting the container.

## Cleanup

```bash
kubectl delete -f k8s/vpa.yaml
kubectl delete -f k8s/pod.yaml
```

## Notes

- The `InPlaceOrRecreate` update mode is a feature of VPA that attempts to resize pods in-place when possible
- In-place resizing requires Kubernetes 1.27+ with the feature gate enabled
- Not all resource changes can be done in-place; some may still require pod recreation
- The demo application intentionally creates a memory usage pattern that demonstrates VPA's ability to adjust resources dynamically
