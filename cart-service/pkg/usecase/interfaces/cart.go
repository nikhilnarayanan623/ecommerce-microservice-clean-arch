package interfaces

import (
	"context"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/cart-service/pkg/domain"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/cart-service/pkg/utils/response"
)

type CartUseCase interface {
	AddToCart(ctx context.Context, userID, productItemID uint64) error
	FindCart(ctx context.Context, userID uint64) (domain.Cart, error)
	FindCartItem(ctx context.Context, cartID uint64) ([]response.CartItem, error)
}
