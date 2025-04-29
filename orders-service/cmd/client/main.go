package main

import (
	"context"
	"log"
	"time"

	orderpb "github.com/ramiroschettino/Go-Market-Microservices/orders-service/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar: %v", err)
	}
	defer conn.Close()

	client := orderpb.NewOrderServiceClient(conn)

	req := &orderpb.CreateOrderRequest{
		ProductId: 1,
		Quantity:  2,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := client.CreateOrder(ctx, req)
	if err != nil {
		log.Fatalf("Error al crear orden: %v", err)
	}

	log.Printf("🧾 Orden creada: %+v\n", res.GetOrder())
}
