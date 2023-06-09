package client

import (
	"context"
	"fmt"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/order-service/pkg/client/interfaces"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/order-service/pkg/config"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/order-service/pkg/pb"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/order-service/pkg/utils/request"
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

func (c *productClient) DecreaseMultipleStocks(ctx context.Context, stocksToDecrease []request.StockDecrease) error {

	inputStocksToDecrease := make([]*pb.MultipleStockDecreaseRequest_StockDecreases, len(stocksToDecrease))

	for i, stock := range stocksToDecrease {
		inputStocksToDecrease[i] = &pb.MultipleStockDecreaseRequest_StockDecreases{
			Sku:           stock.SKU,
			QtyToDecrease: stock.QtyToDecrease,
		}
	}

	_, err := c.client.MultipleStockDecrease(ctx, &pb.MultipleStockDecreaseRequest{
		StockDecreases: inputStocksToDecrease,
	})

	return err
}
