package main

import (
	"context"
	"log"
	"net"

	productpb "github.com/ramiroschettino/Go-Market-Microservices/products-service/proto"
	"google.golang.org/grpc"
)

// Server estructura vacía
type server struct {
	productpb.UnimplementedProductServiceServer
}

// GetProduct implementa el método gRPC
func (s *server) GetProduct(ctx context.Context, req *productpb.GetProductRequest) (*productpb.GetProductResponse, error) {
	log.Printf("Recibida solicitud para producto con ID: %s", req.GetId())

	// Simular que buscamos el producto (hardcodeado)
	product := &productpb.Product{
		Id:          req.GetId(),
		Name:        "Producto de prueba",
		Description: "Descripción del producto de prueba",
		Price:       99.99,
	}

	return &productpb.GetProductResponse{
		Product: product,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error al escuchar: %v", err)
	}

	grpcServer := grpc.NewServer()
	productpb.RegisterProductServiceServer(grpcServer, &server{})

	log.Println("Product Service escuchando en puerto :50051...")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error al servir: %v", err)
	}
}
