package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	productpb "github.com/ramiroschettino/Go-Market-Microservices/products-service/proto"
	"google.golang.org/grpc"
)

func main() {
	if len(os.Args) < 4 {
		log.Fatalf("Uso: go run cmd/client/main.go <nombre> <descripción> <precio>")
	}

	name := os.Args[1]
	description := os.Args[2]
	priceStr := os.Args[3]

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		log.Fatalf("❌ Precio inválido: %v", err)
	}

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("❌ No se pudo conectar a product service: %v", err)
	}
	defer conn.Close()

	client := productpb.NewProductServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &productpb.CreateProductRequest{
		Name:        name,
		Description: description,
		Price:       price,
	}

	res, err := client.CreateProduct(ctx, req)
	if err != nil {
		log.Fatalf("❌ Error al crear producto: %v", err)
	}

	log.Printf("🟢 Producto creado: %+v\n", res.GetProduct())
}
