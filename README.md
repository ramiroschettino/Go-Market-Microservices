# 🛒 Go Market Microservices

Un sistema de microservicios desarrollado en Go que simula un mercado de productos y órdenes. Este proyecto fue creado como práctica para aplicar conceptos modernos como arquitectura hexagonal, gRPC, REST, contenedores con Docker y persistencia en PostgreSQL.

## 🚀 Tecnologías Utilizadas

- **Go (Golang)**: Backend principal
- **gRPC**: Comunicación entre microservicios
- **REST (HTTP)**: Interfaz externa para creación de órdenes
- **PostgreSQL**: Almacenamiento de productos y órdenes
- **Docker & Docker Compose**: Orquestación y contenedores
- **Arquitectura Hexagonal (Ports & Adapters)**

## 📦 Estructura de Microservicios

- `products-service`: Gestión de productos
- `orders-service`: Creación de órdenes (HTTP + gRPC)
- `PostgreSQL`: Base de datos compartida

## 🛠️ Cómo levantar el proyecto

```bash
docker-compose down --volumes
docker-compose up --build -d

Cómo crear una orden:

POST a:

http://localhost:8080/orders
Con este body en JSON:

{
  "product_id": 1,
  "product_name": "Teclado Logitech",
  "product_description": "Teclado inalámbrico óptico",
  "product_price": 300.99,
  "quantity": 10
}