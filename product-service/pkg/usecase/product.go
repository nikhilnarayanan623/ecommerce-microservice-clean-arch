package usecase

import (
	"context"
	"errors"
	"fmt"

	repo "github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/product-service/pkg/repository/interfaces"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/product-service/pkg/usecase/interfaces"
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
	ErrCategoryExist        = errors.New("a category already exist with the given name")
	ErrInvalidCategoryID    = errors.New("invalid category_id")
	ErrVariationExist       = errors.New("an variation already exist with given name")
	ErrInvalidVariationID   = errors.New("invalid variation_id")
	ErrVariationOptionExist = errors.New("variation option already exist with given value")
	ErrProductExist         = errors.New("product already exist with given name")
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

	products, err := c.repo.FindAllProducts(ctx,pagination )

	return products, err
}
