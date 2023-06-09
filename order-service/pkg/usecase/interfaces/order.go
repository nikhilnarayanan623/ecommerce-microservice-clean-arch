package interfaces

import "context"

type OrderUseCase interface {
	PlaceShopOrder(ctx context.Context, userID uint64) (shopOrderID uint64, err error)
	FindShopOrders(ctx context.Context, userID uint64)
}
