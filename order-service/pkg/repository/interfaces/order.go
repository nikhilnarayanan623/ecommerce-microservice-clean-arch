package interfaces

import (
	"context"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/order-service/pkg/domain"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/order-service/pkg/utils/request"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/order-service/pkg/utils/response"
)

type OrderRepository interface {
	Transaction(trxFun func(trxRepo OrderRepository) error) error
	SaveShopOrder(ctx context.Context, shopOrder domain.ShopOrder) (shopOrderID uint64, err error)
	FindShopOrderByShopOrderID(ctx context.Context, shopOrderID uint64) (domain.ShopOrder, error)
	SaveOrderLine(ctx context.Context, orderLine domain.OrderLine) error
	FindOrderLinesByShopOrderID(ctx context.Context, shopOrderID uint64) ([]domain.OrderLine, error)

	FindAllShopOrdersByUserID(ctx context.Context, userID uint64, pagination request.Pagination) ([]response.ShopOrder, error)
}
