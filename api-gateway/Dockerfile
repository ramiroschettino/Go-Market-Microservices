# Usa una imagen base de Go
FROM golang:1.23.6-alpine

# Establece el directorio de trabajo
WORKDIR /app

# Instala las dependencias necesarias (como grpc_health_probe si es necesario)
RUN apk add --no-cache netcat-openbsd && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/v0.4.24/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

# Copia el go.mod y go.sum para instalar las dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copia todo el código fuente del API Gateway
COPY . .

# Compila el código
RUN go build -o api-gateway main.go

# Expón el puerto donde el servicio HTTP de la API Gateway estará corriendo
EXPOSE 8080

# Ejecuta el servidor HTTP
CMD ["./api-gateway"]
