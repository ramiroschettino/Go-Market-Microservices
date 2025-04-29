package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	orderpb "github.com/ramiroschettino/Go-Market-Microservices/orders-service/proto"
	"google.golang.org/grpc"

	productpb "github.com/ramiroschettino/Go-Market-Microservices/orders-service/proto/product"
)

type server struct {
	orderpb.UnimplementedOrderServiceServer
	mu             sync.Mutex
	orders         []*orderpb.Order
	nextID         int64
	productService productpb.ProductServiceClient
}

func (s *server) CreateOrder(ctx context.Context, req *orderpb.CreateOrderRequest) (*orderpb.CreateOrderResponse, error) {
	// Verificamos que exista el producto
	productReq := &productpb.GetProductRequest{Id: req.GetProductId()}
	_, err := s.productService.GetProduct(ctx, productReq)
	if err != nil {
		log.Printf("❌ Producto no encontrado: %v", err)
		return nil, fmt.Errorf("producto con ID %d no existe", req.GetProductId())
	}

	// Si existe, seguimos
	s.mu.Lock()
	defer s.mu.Unlock()

	s.nextID++
	newOrder := &orderpb.Order{
		Id:        s.nextID,
		ProductId: req.GetProductId(),
		Quantity:  req.GetQuantity(),
	}

	s.orders = append(s.orders, newOrder)

	log.Printf("🟢 Orden creada: %+v\n", newOrder)

	return &orderpb.CreateOrderResponse{Order: newOrder}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Error al escuchar: %v", err)
	}

	// ✅ Primero nos conectamos al products-service
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar con products-service: %v", err)
	}
	defer conn.Close()

	productClient := productpb.NewProductServiceClient(conn)

	// ✅ Luego creamos el servidor con ese cliente
	orderServer := &server{
		orders:         []*orderpb.Order{},
		nextID:         0,
		productService: productClient,
	}

	grpcServer := grpc.NewServer()
	orderpb.RegisterOrderServiceServer(grpcServer, orderServer)

	log.Println("📦 Order Service escuchando en puerto :50052...")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error al servir: %v", err)
	}
}
