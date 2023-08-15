package interfaces

import (
	"context"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/order-service/pkg/utils/request"
)

type ProductClient interface {
	DecreaseMultipleStocks(ctx context.Context, stocksDecrease []request.StockDecrease) error
}
