package interfaces

import (
	"context"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/order-service/pkg/utils/response"
)

type CartClient interface {
	FindCart(ctx context.Context, userID uint64) (response.Cart, error)
	RemoveAllCartItems(ctx context.Context, userID uint64) error
}
