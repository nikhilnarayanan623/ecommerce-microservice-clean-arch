package interfaces

import "context"

type OrderClient interface {
	PlaceOrder(ctx context.Context, userID uint64) (shopOrderID uint64, err error)
}
