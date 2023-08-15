package client

import (
	"context"
	"fmt"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/order-service/pkg/client/interfaces"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/order-service/pkg/config"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/order-service/pkg/pb"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/order-service/pkg/utils/response"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type cartClient struct {
	client pb.CartServiceClient
}

func NewCartServiceClient(cfg *config.Config) (interfaces.CartClient, error) {
	gcc, err := grpc.Dial(cfg.CartServiceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to dial grpc client \nerror:%w", err)
	}
	client := pb.NewCartServiceClient(gcc)
	return &cartClient{
		client: client,
	}, nil
}

func (c *cartClient) FindCart(ctx context.Context, userID uint64) (response.Cart, error) {

	res, err := c.client.FindCart(ctx, &pb.FindCartRequest{UserId: userID})
	if err != nil {
		return response.Cart{}, err
	}

	cartItems := make([]response.CartItem, len(res.GetCartItems()))

	for i := 0; i < len(res.GetCartItems()); i++ {

		cartItems[i].ProductItemID = res.GetCartItems()[i].GetProductItemId()
		cartItems[i].ProductName = res.GetCartItems()[i].GetProductName()
		cartItems[i].SKU = res.GetCartItems()[i].GetSku()
		cartItems[i].VariationValue = res.GetCartItems()[i].GetVariationValue()
		cartItems[i].Price = res.GetCartItems()[i].GetPrice()
		cartItems[i].SubTotal = res.GetCartItems()[i].GetSubTotal()
		cartItems[i].Qty = res.GetCartItems()[i].GetQty()
		cartItems[i].QtyInStock = res.GetCartItems()[i].GetQtyInStock()
		cartItems[i].DiscountPrice = res.GetCartItems()[i].GetDiscountPrice()

	}
	return response.Cart{
		TotalPrice: res.GetTotalPrice(),
		CartItems:  cartItems,
	}, nil
}

func (c *cartClient) RemoveAllCartItems(ctx context.Context, userID uint64) error {

	_, err := c.client.RemoveAllCartItems(ctx, &pb.RemoveAllCartItemsRequest{UserId: userID})

	return err
}
