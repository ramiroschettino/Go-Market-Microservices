package client

import (
	"context"
	"time"

	orderpb "github.com/ramiroschettino/Go-Market-Microservices/orders-service/proto"
	"google.golang.org/grpc"
)

type OrderClient struct {
	Client orderpb.OrderServiceClient
}

func NewOrderClient(address string) (*OrderClient, error) {
	conn, err := grpc.Dial(address,
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithTimeout(5*time.Second))
	if err != nil {
		return nil, err
	}

	client := orderpb.NewOrderServiceClient(conn)

	return &OrderClient{
		Client: client,
	}, nil
}

func (oc *OrderClient) CreateOrder(ctx context.Context, req *orderpb.CreateOrderRequest) (*orderpb.CreateOrderResponse, error) {
	return oc.Client.CreateOrder(ctx, req)
}
