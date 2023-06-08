package interfaces

import (
	"context"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/cart-service/pkg/utils/response"
)

type ProductClient interface {
	FindProductItemByID(ctx context.Context, productItemID uint64) (response.ProductItem, error)
}
