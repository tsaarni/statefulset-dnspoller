# Simple dnspoller to check how the DNS entries work for StatefulSets

## Usage

```console
kubectl apply -f manifest/dnspoller.yaml

# Read the logs.
kubectl logs dnspoller-0
kubectl logs dnspoller-1
kubectl logs dnspoller-2

# To test scale down and up.
kubectl scale statefulset dnspoller --replicas 0
kubectl scale statefulset dnspoller --replicas 3
```

To build locally

```console
make container
```

To run on Kind

```console
kind create cluster --name dnspoller
kind load docker-image quay.io/tsaarni/dnspoller:latest --name dnspoller
```

## Example run

```console
