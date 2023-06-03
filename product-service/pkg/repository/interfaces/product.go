package interfaces

import (
	"context"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/product-service/pkg/domain"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/product-service/pkg/utils/request"
)

type ProductRepository interface {
	SaveCategory(ctx context.Context, category request.AddCategory) (categoryID uint64, err error)
	FindCategoryByName(ctx context.Context, categoryName string) (domain.Category, error)
	FindCategoryByID(ctx context.Context, categoryID uint64) (domain.Category, error)

	SaveVariation(ctx context.Context, variation request.AddVariation) (variationID uint64, err error)
	FindVariationByName(ctx context.Context, variationName string) (domain.Variation, error)
	FindVariationByID(ctx context.Context, variationID uint64) (domain.Variation, error)

	SaveVariationOption(ctx context.Context, variationOption request.AddVariationOption) (variationOptionID uint64, err error)
	FindVariationOptionByValue(ctx context.Context, variationValue string) (domain.VariationOption, error)
	FindVariationOptionByID(ctx context.Context, variationOptionID uint64) (domain.VariationOption, error)
}
