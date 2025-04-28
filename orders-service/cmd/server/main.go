package main

import (
	"log"
	"net"

	orderpb "github.com/ramiroschettino/Go-Market-Microservices/orders-service/proto"
	"google.golang.org/grpc"
)

// Server vacío por ahora
type server struct {
	orderpb.UnimplementedOrderServiceServer
}

func main() {
	lis, err := net.Listen("tcp", ":50052") // ⚡ Escuchamos en 50052 (no mismo puerto que products-service)
	if err != nil {
		log.Fatalf("Error al escuchar: %v", err)
	}

	grpcServer := grpc.NewServer()

	orderpb.RegisterOrderServiceServer(grpcServer, &server{})

	log.Println("Order Service escuchando en puerto :50052...")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error al servir: %v", err)
	}
}
