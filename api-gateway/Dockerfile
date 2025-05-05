FROM golang:1.23.6-alpine  

RUN apk add --no-cache netcat-openbsd && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/v0.4.24/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

RUN adduser -D -g '' appuser

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY internal/ internal/
COPY router/ router/
COPY utils/ utils/
COPY main.go .

RUN go build -ldflags="-s -w" -o api-gateway main.go

USER appuser

EXPOSE 8080

CMD ["./api-gateway"]