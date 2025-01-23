build:
	go build -o dnspoller .

container:
	docker build --tag quay.io/tsaarni/dnspoller:latest --file docker/dnspoller/Dockerfile .
