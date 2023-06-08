package client

import (
	"context"
	"fmt"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/cart-service/pkg/client/interfaces"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/cart-service/pkg/config"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/cart-service/pkg/pb"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/cart-service/pkg/utils/response"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type productClient struct {
	client pb.ProductServiceClient
}

func NewProductServiceClient(cfg *config.Config) (interfaces.ProductClient, error) {

	gcc, err := grpc.Dial(cfg.ProductServiceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to dial grpc client \nerror:%w", err)
	}

	client := pb.NewProductServiceClient(gcc)
	return &productClient{
		client: client,
	}, nil
}

func (c *productClient) FindProductItemByID(ctx context.Context, productItemID uint64) (response.ProductItem, error) {

	res, err := c.client.FindProductItem(ctx, &pb.FindProductItemRequest{ProductId: productItemID})
	if err != nil {
		return response.ProductItem{}, err
	}
	return response.ProductItem{
		ID:             res.GetId(),
		Name:           res.GetName(),
		QtyInStock:     res.GetQtyInStock(),
		Price:          res.GetPrice(),
		SKU:            res.GetSku(),
		DiscountPrice:  res.GetDiscountPrice(),
		VariationValue: res.GetVariationValue(),
	}, nil
}
