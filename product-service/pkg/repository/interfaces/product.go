package interfaces

import (
	"context"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/product-service/pkg/domain"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/product-service/pkg/utils/request"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/product-service/pkg/utils/response"
)

type ProductRepository interface {
	Transactions(ctx context.Context, trxFn func(repo ProductRepository) error) error

	SaveCategory(ctx context.Context, category request.AddCategory) (categoryID uint64, err error)
	FindCategoryByName(ctx context.Context, categoryName string) (domain.Category, error)
	FindCategoryByID(ctx context.Context, categoryID uint64) (domain.Category, error)
	FindAllCategories(ctx context.Context) ([]response.Category, error)

	SaveVariation(ctx context.Context, variation request.AddVariation) (variationID uint64, err error)
	FindVariationByName(ctx context.Context, variationName string) (domain.Variation, error)
	FindVariationByID(ctx context.Context, variationID uint64) (domain.Variation, error)

	SaveVariationOption(ctx context.Context, variationOption request.AddVariationOption) (variationOptionID uint64, err error)
	FindVariationOptionByValue(ctx context.Context, variationValue string) (domain.VariationOption, error)
	//FindVariationOptionByID(ctx context.Context, variationOptionID uint64) (domain.VariationOption, error)
	IsValidVariationOptionID(ctx context.Context, variationOptionID uint64) (valid bool, err error)

	SaveProduct(ctx context.Context, product request.AddProduct) (productID uint64, err error)
	IsProductNameAlreadyExist(ctx context.Context, productName string) (exist bool, err error)
	IsValidProductID(ctx context.Context, productID uint64) (valid bool, err error)
	FindAllProducts(ctx context.Context, pagination request.Pagination) ([]response.Product, error)

	SaveProductItem(ctx context.Context, productItem request.AddProductItem) (productItemID uint64, err error)
	FindProductItemsByProductID(ctx context.Context, productID uint64) (productItems []response.ProductItem, err error)
	FindProductItemByID(ctx context.Context, productItemID uint64) (response.ProductItem, error)
	IsProductItemAlreadyExist(ctx context.Context, productID, variationOptionID uint64) (exist bool, err error)
	SaveProductConfiguration(ctx context.Context, productItemID, variationOptionID uint64) error

	FindProductItemsStockDetails(ctx context.Context, sku string) (response.StockDetails, error)
	UpdateProductQty(ctx context.Context, sku string, updateTotalQty uint64) error
}
