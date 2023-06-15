package interfaces

import (
	"context"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/cart-service/pkg/domain"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/cart-service/pkg/utils/response"
)

type CartRepository interface {
	FindCartByUserID(ctx context.Context, userID uint64) (domain.Cart, error)
	SaveCart(ctx context.Context, userID uint64) (cartID uint64, err error)

	SaveCartItem(ctx context.Context, cartID, productItemID uint64) error
	IsProductItemAlreadyExistInCart(ctx context.Context, cartID, productItemID uint64) (exist bool, err error)
	FindCartItemsByCartID(ctx context.Context, cartID uint64) ([]response.CartItem, error)
	RemoveAllCartItems(ctx context.Context, cartID uint64) error
}
