package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ramiroschettino/Go-Market-Microservices/api-gateway/internal/handler"
)

func SetupRouter(orderHandler *handler.OrderHandler) *gin.Engine {
	r := gin.Default()

	r.POST("/orders", orderHandler.CreateOrderHandler)

	return r
}
