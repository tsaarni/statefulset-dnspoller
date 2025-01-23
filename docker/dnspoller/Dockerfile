FROM golang:1.23-alpine3.21 AS builder
WORKDIR /app
COPY dnspoller.go .
RUN CGO_ENABLED=0 go build dnspoller.go

FROM alpine:3.21
COPY --from=builder /app/dnspoller .
