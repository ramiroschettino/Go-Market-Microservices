FROM golang:1.23.6-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git make && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/v0.4.24/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

COPY . .

WORKDIR /app/products-service

RUN go mod edit -replace github.com/ramiroschettino/Go-Market-Microservices/common=../common

RUN go mod download

RUN CGO_ENABLED=0 go build -o /server ./cmd/server/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /server .
COPY --from=builder /bin/grpc_health_probe /bin/grpc_health_probe
EXPOSE 50051
CMD ["./server"]