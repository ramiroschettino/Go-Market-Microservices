package clients

import (
	"context"
	"log"
	"time"

	orderpb "github.com/ramiroschettino/Go-Market-Microservices/orders-service/proto"
	"google.golang.org/grpc"
)

type OrderClient struct {
	Client orderpb.OrderServiceClient
}

func NewOrderClient(address string) *OrderClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(5*time.Second))
	if err != nil {
		log.Fatalf("Could not connect to order service: %v", err)
	}

	client := orderpb.NewOrderServiceClient(conn)

	return &OrderClient{
		Client: client,
	}
}

// CreateOrder is an example function you might call from a handler
func (oc *OrderClient) CreateOrder(ctx context.Context, req *orderpb.CreateOrderRequest) (*orderpb.CreateOrderResponse, error) {
	return oc.Client.CreateOrder(ctx, req)
}
