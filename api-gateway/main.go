package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ramiroschettino/Go-Market-Microservices/api-gateway/clients"
	"github.com/ramiroschettino/Go-Market-Microservices/api-gateway/handlers"
)

func main() {
	// Inicializar cliente gRPC para orders-service
	orderClient, err := clients.NewOrderClient("orders_service:50052") // Asegúrate de que el puerto coincida con el de tu orders-service
	if err != nil {
		log.Fatalf("could not create order client: %v", err)
	}

	// Inicializar el router de Gin
	r := gin.Default()

	// Inicializar handler para la creación de orden
	orderHandler := &handlers.OrderHandler{
		OrderClient: orderClient,
	}

	// Definir la ruta para crear órdenes
	r.POST("/orders", orderHandler.CreateOrderHandler)

	// Iniciar el servidor HTTP
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
