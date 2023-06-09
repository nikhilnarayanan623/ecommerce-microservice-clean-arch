package client

import (
	"context"
	"fmt"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/client/interfaces"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/config"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type orderClient struct {
	client pb.OrderServiceClient
}

func NewOrderClient(cfg *config.Config) (interfaces.OrderClient, error) {

	gcc, err := grpc.Dial(cfg.OrderServiceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to dial grpc client \nerror:%w", err)
	}

	client := pb.NewOrderServiceClient(gcc)

	return &orderClient{
		client: client,
	}, nil
}

func (c *orderClient) PlaceOrder(ctx context.Context, userID uint64) (shopOrderID uint64, err error) {

	res, err := c.client.PlaceOrder(ctx, &pb.PlaceOrderRequest{UserId: userID})

	if err != nil {
		return 0, err
	}

	return res.GetShopOrderId(), nil
}
