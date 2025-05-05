FROM golang:1.23.6-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git make

COPY orders-service/go.mod orders-service/go.sum ./
COPY common/go.mod common/go.sum ./common/

RUN go mod edit -replace github.com/ramiroschettino/Go-Market-Microservices/common=/app/common && \
    go mod download

COPY . .

WORKDIR /app/orders-service
RUN CGO_ENABLED=0 go build -o /grpcserver ./cmd/server/main.go && \
    CGO_ENABLED=0 go build -o /httpserver ./cmd/http/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /grpcserver .
COPY --from=builder /httpserver .
EXPOSE 8080 50052
CMD ["sh", "-c", "./httpserver & ./grpcserver"]