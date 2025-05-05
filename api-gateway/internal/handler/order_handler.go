package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ramiroschettino/Go-Market-Microservices/api-gateway/clients"
	orderpb "github.com/ramiroschettino/Go-Market-Microservices/orders-service/proto"
)

type OrderHandler struct {
	OrderClient *clients.OrderClient
}

type CreateOrderInput struct {
	ProductName        string  `json:"product_name" binding:"required"`
	ProductDescription string  `json:"product_description" binding:"required"`
	ProductPrice       float32 `json:"product_price" binding:"required"`
	Quantity           int32   `json:"quantity" binding:"required"`
}

func (h *OrderHandler) CreateOrderHandler(c *gin.Context) {
	var input CreateOrderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &orderpb.CreateOrderRequest{
		ProductName:        input.ProductName,
		ProductDescription: input.ProductDescription,
		ProductPrice:       input.ProductPrice,
		Quantity:           input.Quantity,
	}

	resp, err := h.OrderClient.CreateOrder(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"order_id": resp.OrderId,
	})
}
