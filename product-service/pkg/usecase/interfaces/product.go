package interfaces

import (
	"context"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/product-service/pkg/utils/request"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/product-service/pkg/utils/response"
)

type ProductUseCase interface {
	AddCategory(ctx context.Context, category request.AddCategory) (categoryID uint64, err error)
	AddVariation(ctx context.Context, variation request.AddVariation) (variationID uint64, err error)
	AddVariationOption(ctx context.Context, variationOption request.AddVariationOption) (variationOptionID uint64, err error)

	FindAllCategories(ctx context.Context) ([]response.Category, error)

	AddProduct(ctx context.Context, product request.AddProduct) (productID uint64, err error)
	FindAllProducts(ctx context.Context, pagination request.Pagination) ([]response.Product, error)

	AddProductItem(ctx context.Context, productItems request.AddProductItem) (productItemID uint64, err error)
	FindAllProductItems(ctx context.Context, productID uint64) ([]response.ProductItem, error)

	FindProductItemByID(ctx context.Context, productItemID uint64) (response.ProductItem, error)
}
