package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/ramiroschettino/Go-Market-Microservices/orders-service/internal/domain"
	"github.com/ramiroschettino/Go-Market-Microservices/orders-service/internal/ports/db"
	orderpb "github.com/ramiroschettino/Go-Market-Microservices/orders-service/proto"
	productpb "github.com/ramiroschettino/Go-Market-Microservices/orders-service/proto/product"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type server struct {
	orderpb.UnimplementedOrderServiceServer
	productService productpb.ProductServiceClient
}

func (s *server) CreateOrder(ctx context.Context, req *orderpb.CreateOrderRequest) (*orderpb.CreateOrderResponse, error) {
	productReq := &productpb.GetProductRequest{Id: req.GetProductId()}
	product, err := s.productService.GetProduct(ctx, productReq)

	if err != nil {
		log.Printf("❌ Producto no encontrado, creando nuevo producto...")

		createProductReq := &productpb.CreateProductRequest{
			Name:        req.GetProductName(),
			Description: req.GetProductDescription(),
			Price:       req.GetProductPrice(),
		}
		createdProduct, err := s.productService.CreateProduct(ctx, createProductReq)
		if err != nil {
			log.Printf("❌ Error al crear el producto: %v", err)
			return nil, fmt.Errorf("no se pudo crear el producto")
		}
		log.Printf("🟢 Producto creado: %v", createdProduct)

		product = &productpb.GetProductResponse{
			Product: createdProduct.Product,
		}
	}

	order := domain.Order{
		ProductID: product.Product.Id,
		Quantity:  req.GetQuantity(),
	}

	if err := db.DB.Create(&order).Error; err != nil {
		log.Printf("❌ Error al guardar orden: %v", err)
		return nil, err
	}

	log.Printf("🟢 Orden persistida: %+v\n", order)

	return &orderpb.CreateOrderResponse{
		Order: &orderpb.Order{
			Id:        int64(order.ID),
			ProductId: order.ProductID,
			Quantity:  order.Quantity,
		},
	}, nil
}

func main() {
	db.Connect()

	var conn *grpc.ClientConn
	var err error
	maxRetries := 3
	for i := 0; i < maxRetries; i++ {
		conn, err = grpc.Dial("products-service:50051",
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithBlock(),
			grpc.WithTimeout(5*time.Second))

		if err == nil {
			break
		}

		log.Printf("Intento %d: No se pudo conectar a products-service: %v", i+1, err)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatalf("No se pudo conectar con products-service después de %d intentos: %v", maxRetries, err)
	}
	defer conn.Close()

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Error al crear el listener: %v", err)
	}

	productClient := productpb.NewProductServiceClient(conn)

	orderServer := &server{
		productService: productClient,
	}

	grpcServer := grpc.NewServer()
	orderpb.RegisterOrderServiceServer(grpcServer, orderServer)

	log.Println("📦 Order Service escuchando en puerto :50052...")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error al iniciar el servidor gRPC: %v", err)
	}
}
