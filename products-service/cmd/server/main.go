package main

import (
	"context"
	"log"
	"net"

	productpb "github.com/ramiroschettino/Go-Market-Microservices/common/proto"
	"google.golang.org/grpc"
	"gorm.io/gorm"

	"github.com/ramiroschettino/Go-Market-Microservices/products-service/internal/domain"
	"github.com/ramiroschettino/Go-Market-Microservices/products-service/internal/ports/db"
)

type server struct {
	productpb.UnimplementedProductServiceServer
	db *gorm.DB
}

func (s *server) GetProduct(ctx context.Context, req *productpb.GetProductRequest) (*productpb.GetProductResponse, error) {
	log.Printf("🔎 Buscando producto con ID: %s", req.GetId())

	var product domain.Product
	if err := s.db.First(&product, req.GetId()).Error; err != nil {
		log.Printf("❌ Producto no encontrado: %v", err)
		return nil, err
	}

	return &productpb.GetProductResponse{
		Product: &productpb.Product{
			Id:          int64(product.ID),
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
		},
	}, nil
}

func (s *server) CreateProduct(ctx context.Context, req *productpb.CreateProductRequest) (*productpb.CreateProductResponse, error) {
	log.Printf("🆕 Creando producto: %s", req.GetName())

	product := domain.Product{
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Price:       req.GetPrice(),
	}

	if err := s.db.Create(&product).Error; err != nil {
		log.Printf("❌ Error al crear producto: %v", err)
		return nil, err
	}

	return &productpb.CreateProductResponse{
		Product: &productpb.Product{
			Id:          int64(product.ID),
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
		},
	}, nil
}

func main() {
	db.Connect()
	db.DB.AutoMigrate(&domain.Product{})

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("❌ Error al escuchar: %v", err)
	}

	s := &server{db: db.DB}

	grpcServer := grpc.NewServer()
	productpb.RegisterProductServiceServer(grpcServer, s)

	log.Println("🛒 Product Service escuchando en puerto :50051...")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("❌ Error al servir: %v", err)
	}
}
