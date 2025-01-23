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
$ kubectl apply -f manifest/dnspoller.yaml

# Wait for the image to be loaded and pods to be running.

# Read logs
$ kubectl logs dnspoller-0
time=2025-01-23T09:26:02.067Z level=INFO msg="Own addresses" addresses="[127.0.0.1/8 10.244.0.5/24 ::1/128 fe80::a851:f0ff:fe4d:819d/64]"
time=2025-01-23T09:26:02.067Z level=INFO msg=Lookup hostname=dnspoller-headless
time=2025-01-23T09:26:02.075Z level=ERROR msg="Error looking up" hostname=dnspoller-headless error="lookup dnspoller-headless on 10.96.0.10:53: server misbehaving"
time=2025-01-23T09:26:03.086Z level=ERROR msg="Error looking up" hostname=dnspoller-headless error="lookup dnspoller-headless on 10.96.0.10:53: server misbehaving"
time=2025-01-23T09:26:04.097Z level=ERROR msg="Error looking up" hostname=dnspoller-headless error="lookup dnspoller-headless on 10.96.0.10:53: server misbehaving"
time=2025-01-23T09:26:05.108Z level=ERROR msg="Error looking up" hostname=dnspoller-headless error="lookup dnspoller-headless on 10.96.0.10:53: server misbehaving"
time=2025-01-23T09:26:06.116Z level=ERROR msg="Error looking up" hostname=dnspoller-headless error="lookup dnspoller-headless on 10.96.0.10:53: server misbehaving"
time=2025-01-23T09:26:07.125Z level=ERROR msg="Error looking up" hostname=dnspoller-headless error="lookup dnspoller-headless on 10.96.0.10:53: server misbehaving"
time=2025-01-23T09:26:08.137Z level=ERROR msg="Error looking up" hostname=dnspoller-headless error="lookup dnspoller-headless on 10.96.0.10:53: server misbehaving"
time=2025-01-23T09:26:09.147Z level=ERROR msg="Error looking up" hostname=dnspoller-headless error="lookup dnspoller-headless on 10.96.0.10:53: server misbehaving"
time=2025-01-23T09:26:10.157Z level=ERROR msg="Error looking up" hostname=dnspoller-headless error="lookup dnspoller-headless on 10.96.0.10:53: server misbehaving"
time=2025-01-23T09:26:11.165Z level=ERROR msg="Error looking up" hostname=dnspoller-headless error="lookup dnspoller-headless on 10.96.0.10:53: server misbehaving"
time=2025-01-23T09:26:12.176Z level=ERROR msg="Error looking up" hostname=dnspoller-headless error="lookup dnspoller-headless on 10.96.0.10:53: server misbehaving"
time=2025-01-23T09:26:13.188Z level=ERROR msg="Error looking up" hostname=dnspoller-headless error="lookup dnspoller-headless on 10.96.0.10:53: server misbehaving"
time=2025-01-23T09:26:14.190Z level=INFO msg=Added address=10.244.0.5
time=2025-01-23T09:26:26.204Z level=INFO msg=Added address=10.244.0.6
time=2025-01-23T09:26:38.219Z level=INFO msg=Added address=10.244.0.7

$ kubectl logs dnspoller-1
time=2025-01-23T09:26:14.731Z level=INFO msg="Own addresses" addresses="[127.0.0.1/8 10.244.0.6/24 ::1/128 fe80::985b:d7ff:fe06:89b3/64]"
time=2025-01-23T09:26:14.732Z level=INFO msg=Lookup hostname=dnspoller-headless
time=2025-01-23T09:26:14.732Z level=INFO msg=Added address=10.244.0.5
time=2025-01-23T09:26:26.749Z level=INFO msg=Added address=10.244.0.6
time=2025-01-23T09:26:38.770Z level=INFO msg=Added address=10.244.0.7

$ kubectl logs dnspoller-2
time=2025-01-23T09:26:26.693Z level=INFO msg="Own addresses" addresses="[127.0.0.1/8 10.244.0.7/24 ::1/128 fe80::b0f7:91ff:fe5b:9513/64]"
time=2025-01-23T09:26:26.693Z level=INFO msg=Lookup hostname=dnspoller-headless
time=2025-01-23T09:26:26.694Z level=INFO msg=Added address=10.244.0.5
time=2025-01-23T09:26:26.694Z level=INFO msg=Added address=10.244.0.6
time=2025-01-23T09:26:38.709Z level=INFO msg=Added address=10.244.0.7
```
