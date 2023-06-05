package interfaces

import (
	"context"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/utils/request"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/utils/response"
)

type ProductClient interface {
	AddCategory(ctx context.Context, category request.AddCategory) (categoryID uint64, err error)
	AddVariation(ctx context.Context, variation request.AddVariation) (variationID uint64, err error)
	AddVariationOption(ctx context.Context, variationOption request.AddVariationOption) (variationOptionID uint64, err error)

	FindAllCategories(ctx context.Context) ([]response.Category, error)

	AddProduct(ctx context.Context, product request.AddProduct) (productID uint64, err error)
	FindAllProducts(ctx context.Context, pagination request.Pagination) ([]response.Product, error)

	AddProductItem(ctx context.Context, productItem request.AddProductItem) (productItemID uint64, err error)
	FindAllProductItems(ctx context.Context, productID uint64) ([]response.ProductItem, error)
}
