FROM golang:1.23.6-alpine AS builder

WORKDIR /app
COPY . .

RUN apk add --no-cache git make

RUN CGO_ENABLED=0 go build -o /grpcserver ./cmd/server/main.go && \
    CGO_ENABLED=0 go build -o /httpserver ./cmd/http/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /grpcserver .
COPY --from=builder /httpserver .
COPY entrypoint.sh .

RUN apk add --no-cache netcat-openbsd && \
    chmod +x /app/entrypoint.sh

EXPOSE 8080 50052

ENTRYPOINT ["/app/entrypoint.sh"]