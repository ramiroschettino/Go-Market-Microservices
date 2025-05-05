🛒 Go Market Microservices
Sistema de microservicios en Go para gestión de productos y pedidos

📌 Descripción
Este proyecto es un e-commerce modular basado en microservicios, desarrollado en Go (Golang), que utiliza:

API Gateway (REST) como punto único de entrada.

gRPC para comunicación interna entre servicios.

PostgreSQL como base de datos principal.

Docker para containerización y despliegue.

Ideal para aprender:
✅ Arquitectura hexagonal (Ports & Adapters)
✅ Comunicación entre servicios (gRPC + REST)
✅ Manejo de contenedores con Docker Compose

🚀 Instalación
Clonar el repositorio:

bash
git clone https://github.com/ramiroschettino/Go-Market-Microservices.git
cd go-market-microservices
Iniciar los servicios con Docker:

bash
docker-compose up --build -d
Verificar que todo esté funcionando:

bash
curl http://localhost:8080/health
# Respuesta esperada: {"status":"ok"}
🔍 Uso
1. Crear una nueva orden (vía API Gateway)
Endpoint:

http
POST http://localhost:8080/orders
Ejemplo de solicitud (JSON):

json
{
  "product_id": 1,
  "product_name": "Mouse Logitech MX Master 3",
  "product_description": "Mouse inalámbrico ergonómico",
  "product_price": 99.99,
  "quantity": 1
}
Respuesta exitosa:

json
{
  "order_id": 101,
  "product_id": 1,
  "quantity": 1,
  "total_price": 99.99
}
2. Consultar la base de datos (PostgreSQL)
Acceder a la DB desde Docker:

bash
docker-compose exec postgres psql -U postgres -d products_db
Comandos útiles dentro de PostgreSQL:

sql
-- Ver tablas disponibles
\dt

-- Consultar órdenes recientes
SELECT * FROM orders LIMIT 5;

-- Salir
\q
⚙️ Estructura del Proyecto
bash
.
├── api-gateway/           # Punto de entrada (REST)
├── orders-service/        # Procesamiento de pedidos (gRPC)
├── products-service/      # Catálogo de productos (gRPC + DB)
├── proto/                 # Definiciones de Protocol Buffers
└── docker-compose.yml     # Configuración de contenedores
📜 Comandos Útiles
Comando	Descripción
docker-compose logs -f	Ver logs en tiempo real
docker-compose restart api-gateway	Reiniciar solo el API Gateway
docker-compose down --volumes	Detener y eliminar todo (incluyendo datos)
📌 Tecnologías
Lenguaje: Go (Golang)

Comunicación: gRPC (interno) + REST (público)

Base de datos: PostgreSQL

Contenedores: Docker + Docker Compose