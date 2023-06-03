package interfaces

import (
	"context"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/utils/request"
)

type ProductClient interface {
	AddCategory(ctx context.Context, category request.AddCategory) (categoryID uint64, err error)
	AddVariation(ctx context.Context, variation request.AddVariation) (variationID uint64, err error)
	AddVariationOption(ctx context.Context, variationOption request.AddVariationOption) (variationOptionID uint64, err error)
}
