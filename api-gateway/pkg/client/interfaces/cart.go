package interfaces

import (
	"context"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/utils/response"
)

type CartClient interface {
	AddToCart(ctx context.Context, userID, productItemId uint64) error
	FindCart(ctx context.Context, userID uint64) (response.Cart, error)
}
