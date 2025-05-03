FROM golang:1.23.6-alpine AS builder

WORKDIR /app

# 1. Instalar dependencias
RUN apk add --no-cache git make

# 2. Copiar solo los archivos de módulos primero
COPY orders-service/go.mod orders-service/go.sum ./
COPY common/go.mod common/go.sum ./common/

# 3. Configurar el reemplazo
RUN go mod edit -replace github.com/ramiroschettino/Go-Market-Microservices/common=/app/common && \
    go mod download

# 4. Copiar todo el código
COPY . .

# 5. Construir los binarios
WORKDIR /app/orders-service
RUN CGO_ENABLED=0 go build -o /grpcserver ./cmd/server/main.go && \
    CGO_ENABLED=0 go build -o /httpserver ./cmd/http/main.go

# Etapa final
FROM alpine:latest
WORKDIR /app
COPY --from=builder /grpcserver .
COPY --from=builder /httpserver .
EXPOSE 8080 50052
CMD ["sh", "-c", "./httpserver & ./grpcserver"]