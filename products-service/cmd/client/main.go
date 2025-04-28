package main

import (
	"context"
	"log"
	"time"

	productpb "github.com/ramiroschettino/Go-Market-Microservices/products-service/proto"
	"google.golang.org/grpc"
)

func main() {
	// Conectamos al servidor
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar: %v", err)
	}
	defer conn.Close()

	client := productpb.NewProductServiceClient(conn)

	// Contexto con timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Hacemos la llamada a GetProduct
	req := &productpb.GetProductRequest{Id: 1234}
	res, err := client.GetProduct(ctx, req)
	if err != nil {
		log.Fatalf("Error llamando a GetProduct: %v", err)
	}

	log.Printf("Producto recibido: %+v", res.GetProduct())
}
