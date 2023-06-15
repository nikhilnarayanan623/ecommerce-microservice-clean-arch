package interfaces

import (
	"context"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/utils/request"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/utils/response"
)

type OrderClient interface {
	PlaceOrder(ctx context.Context, userID uint64) (shopOrderID uint64, err error)
	FindAllShopOrders(ctx context.Context, userID uint64, pagination request.Pagination) ([]response.ShopOrder, error)
}
