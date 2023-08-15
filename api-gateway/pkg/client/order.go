package client

import (
	"context"
	"fmt"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/client/interfaces"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/config"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/pb"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/utils/request"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/utils/response"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
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

func (c *orderClient) FindAllShopOrders(ctx context.Context, userID uint64, pagination request.Pagination) ([]response.ShopOrder, error) {

	request := &pb.FindAllOrderRequest{
		UserId:     userID,
		PageNumber: pagination.PageNumber,
		Count:      pagination.Count,
	}
	res, err := c.client.FindAllOrder(ctx, request)

	if err != nil {
		return nil, err
	}

	shopOrders := make([]response.ShopOrder, len(res.GetOrders()))

	for i, order := range res.GetOrders() {
		ts := timestamppb.New(order.OrderDate.AsTime())
		shopOrders[i].ID = order.ShopOrderId
		shopOrders[i].OrderDate = ts.AsTime()
		shopOrders[i].OrderTotalPrice = order.OrderTotalPrice
		shopOrders[i].Discount = order.Discount
	}

	return shopOrders, nil
}
