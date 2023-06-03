package interfaces

import (
	"context"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/product-service/pkg/utils/request"
)

type ProductUseCase interface {
	AddCategory(ctx context.Context, category request.AddCategory) (categoryID uint64, err error)
	AddVariation(ctx context.Context, variation request.AddVariation) (variationID uint64, err error)
	AddVariationOption(ctx context.Context, variationOption request.AddVariationOption) (variationOptionID uint64, err error)
}
