package usecase

import (
	"context"
	"errors"
	"fmt"

	repo "github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/product-service/pkg/repository/interfaces"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/product-service/pkg/usecase/interfaces"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/product-service/pkg/utils"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/product-service/pkg/utils/request"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/product-service/pkg/utils/response"
)

type productUseCase struct {
	repo repo.ProductRepository
}

func NewProductUseCase(repo repo.ProductRepository) interfaces.ProductUseCase {
	return &productUseCase{
		repo: repo,
	}
}

var (
	ErrCategoryExist            = errors.New("a category already exist with the given name")
	ErrInvalidCategoryID        = errors.New("invalid category_id")
	ErrVariationExist           = errors.New("an variation already exist with given name")
	ErrInvalidVariationID       = errors.New("invalid variation_id")
	ErrVariationOptionExist     = errors.New("variation option already exist with given value")
	ErrProductExist             = errors.New("product already exist with given name")
	ErrInvalidProductID         = errors.New("invalid product_id")
	ErrProductItemExist         = errors.New("product item already exist with given configurations")
	ErrInvalidVariationOptionID = errors.New("invalid variation_option_id")

	ErrInvalidStockUpdateQty = errors.New("given stock decrease qty greater than product qty_in_stock")
)

func (c *productUseCase) AddCategory(ctx context.Context, category request.AddCategory) (uint64, error) {

	existCategory, err := c.repo.FindCategoryByName(ctx, category.CategoryName)
	if err != nil {
		return 0, fmt.Errorf("failed to check category already exist \nerror:%w", err)
	}
	if existCategory.ID != 0 {
		return 0, ErrCategoryExist
	}

	if category.MainCategoryID != 0 {
		category, err := c.repo.FindCategoryByID(ctx, category.MainCategoryID)
		if err != nil {
			return 0, fmt.Errorf("failed to verify category_id \nerror:%w", err)
		}
		if category.ID == 0 {
			return 0, ErrInvalidCategoryID
		}
	}

	categoryID, err := c.repo.SaveCategory(ctx, category)
	if err != nil {
		return 0, fmt.Errorf("failed to save category \nerror:%w", err)
	}
	return categoryID, nil
}
func (c *productUseCase) AddVariation(ctx context.Context, variation request.AddVariation) (uint64, error) {

	category, err := c.repo.FindCategoryByID(ctx, variation.CategoryID)
	if err != nil {
		return 0, fmt.Errorf("failed to verify category_id \nerror:%w", err)
	}
	if category.ID == 0 {
		return 0, ErrInvalidCategoryID
	}

	existVariation, err := c.repo.FindVariationByName(ctx, variation.VariationName)
	if err != nil {
		return 0, fmt.Errorf("failed to check variation already exist \nerror:%w", err)
	}
	if existVariation.ID != 0 {
		return 0, ErrVariationExist
	}

	variationID, err := c.repo.SaveVariation(ctx, variation)
	if err != nil {
		return 0, fmt.Errorf("failed to save variation \nerror:%w", err)
	}

	return variationID, nil
}
func (c *productUseCase) AddVariationOption(ctx context.Context, variationOption request.AddVariationOption) (uint64, error) {

	variation, err := c.repo.FindVariationByID(ctx, variationOption.VariationID)
	if err != nil {
		return 0, fmt.Errorf("failed to verify variation_id \nerror:%w", err)
	}
	if variation.ID == 0 {
		return 0, ErrInvalidVariationID
	}

	existVariationOp, err := c.repo.FindVariationOptionByValue(ctx, variationOption.VariationValue)
	if err != nil {
		return 0, fmt.Errorf("failed to check variation option already exist \nerror:%w", err)
	}
	if existVariationOp.ID != 0 {
		return 0, ErrVariationOptionExist
	}

	variationOptionID, err := c.repo.SaveVariationOption(ctx, variationOption)
	if err != nil {
		return 0, fmt.Errorf("failed to save variation_option \nerror:%w", err)
	}

	return variationOptionID, nil
}

func (c *productUseCase) FindAllCategories(ctx context.Context) ([]response.Category, error) {
	categories, err := c.repo.FindAllCategories(ctx)
	return categories, err
}

// Save Product
func (c *productUseCase) AddProduct(ctx context.Context, product request.AddProduct) (uint64, error) {

	productExist, err := c.repo.IsProductNameAlreadyExist(ctx, product.Name)
	if err != nil {
		return 0, fmt.Errorf("failed to check product name already exist \nerror:%w", err)
	}
	if productExist {
		return 0, ErrProductExist
	}

	category, err := c.repo.FindCategoryByID(ctx, product.CategoryID)

	if err != nil {
		return 0, fmt.Errorf("failed to verify category_id \nerror:%w", err)
	}
	if category.ID == 0 {
		return 0, ErrInvalidCategoryID
	}

	productID, err := c.repo.SaveProduct(ctx, product)
	if err != nil {
		return 0, fmt.Errorf("failed to save product \nerror:%w", err)
	}
	return productID, nil
}

func (c *productUseCase) FindAllProducts(ctx context.Context, pagination request.Pagination) ([]response.Product, error) {

	products, err := c.repo.FindAllProducts(ctx, pagination)

	return products, err
}

func (c *productUseCase) AddProductItem(ctx context.Context, productItems request.AddProductItem) (uint64, error) {
	// validate product_id
	valid, err := c.repo.IsValidProductID(ctx, productItems.ProductID)
	if err != nil {
		return 0, fmt.Errorf("failed to validate product_id \nerror:%w", err)
	}
	if !valid {
		return 0, ErrInvalidProductID
	}
	// validate variation_option_id
	valid, err = c.repo.IsValidVariationOptionID(ctx, productItems.VariationOptionID)
	if err != nil {
		return 0, fmt.Errorf("failed to validate variation_option_id \nerror:%w", err)
	}
	if !valid {
		return 0, ErrInvalidVariationOptionID
	}

	productItemExist, err := c.repo.IsProductItemAlreadyExist(ctx, productItems.ProductID, productItems.VariationOptionID)
	if err != nil {
		return 0, fmt.Errorf("failed to check product_item already exist \nerror:%w", err)
	}
	if productItemExist {
		return 0, ErrProductItemExist
	}
	productItems.SKU = utils.GenerateSKU()

	var productItemID uint64
	err = c.repo.Transactions(ctx, func(trxRepo repo.ProductRepository) error {
		productItemID, err = trxRepo.SaveProductItem(ctx, productItems)
		if err != nil {
			return fmt.Errorf("failed to save product_item \nerror:%w", err)
		}

		err = trxRepo.SaveProductConfiguration(ctx, productItemID, productItems.VariationOptionID)
		if err != nil {
			return fmt.Errorf("failed to save product_item configuration \nerror:%w", err)
		}
		return nil
	})
	if err != nil {
		return 0, err
	}

	return productItemID, nil
}

func (c *productUseCase) FindAllProductItems(ctx context.Context, productID uint64) ([]response.ProductItem, error) {

	valid, err := c.repo.IsValidProductID(ctx, productID)
	if err != nil {
		return nil, fmt.Errorf("failed to validate product_id \nerror:%w", err)
	}
	if !valid {
		return nil, ErrInvalidProductID
	}

	productItems, err := c.repo.FindProductItemsByProductID(ctx, productID)

	if err != nil {
		return nil, fmt.Errorf("failed to find all product_items \nerror:%w", err)
	}

	return productItems, nil
}

func (c *productUseCase) FindProductItemByID(ctx context.Context, productItemID uint64) (response.ProductItem, error) {
	productItem, err := c.repo.FindProductItemByID(ctx, productItemID)
	if err != nil {
		return response.ProductItem{}, fmt.Errorf("failed to find product_item \nerror:%w", err)
	}
	return productItem, nil
}

func (c *productUseCase) MultipleStockDecrease(ctx context.Context, stockDetails []request.StockDecrease) error {

	var updateTotalQty uint64
	for _, stock := range stockDetails {

		productItemStock, err := c.repo.FindProductItemsStockDetails(ctx, stock.SKU)
		if err != nil {
			return fmt.Errorf("failed to find product stock details \nerror:%w", err)
		}

		if stock.QtyToDecrease > productItemStock.QtyInStock {
			return ErrInvalidStockUpdateQty
		}

		updateTotalQty = productItemStock.QtyInStock - stock.QtyToDecrease
		err = c.repo.UpdateProductQty(ctx, productItemStock.SKU, updateTotalQty)
		if err != nil {
			return fmt.Errorf("failed to decrease product qty_in_stock \nerror:%w", err)
		}
	}

	return nil
}
