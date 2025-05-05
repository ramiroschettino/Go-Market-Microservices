package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ramiroschettino/Go-Market-Microservices/api-gateway/internal/client"
	"github.com/ramiroschettino/Go-Market-Microservices/api-gateway/internal/handler"
)

func main() {
	serviceAddress := getServiceAddress()

	orderClient, err := client.NewOrderClient(serviceAddress)
	if err != nil {
		log.Fatalf("Could not create order client: %v", err)
	}

	r := setupRouter(orderClient)

	log.Println("Starting API Gateway on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

func getServiceAddress() string {
	address := os.Getenv("ORDER_SERVICE_ADDRESS")
	if address == "" {
		address = "orders-service:50052"
	}
	return address
}

func setupRouter(orderClient *client.OrderClient) *gin.Engine {
	r := gin.Default()

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Order routes
	orderHandler := &handler.OrderHandler{
		OrderClient: orderClient,
	}
	r.POST("/orders", orderHandler.CreateOrderHandler)

	return r
}
