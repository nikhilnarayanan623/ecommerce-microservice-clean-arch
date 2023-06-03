package repository

import (
	"context"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/product-service/pkg/domain"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/product-service/pkg/repository/interfaces"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/product-service/pkg/utils/request"
	"gorm.io/gorm"
)

type productDatabase struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) interfaces.ProductRepository {
	return &productDatabase{
		db: db,
	}
}

func (c *productDatabase) SaveCategory(ctx context.Context, category request.AddCategory) (categoryID uint64, err error) {

	if category.MainCategoryID == 0 {
		query := `INSERT INTO categories (name) VALUES($1) RETURNING id AS category_id`
		err = c.db.Raw(query, category.CategoryName).Scan(&categoryID).Error
	} else {
		//otherwise add its with main category
		query := `INSERT INTO categories (category_id, name) VALUES($1, $2) RETURNING id AS category_id`
		err = c.db.Raw(query, category.MainCategoryID, category.CategoryName).Scan(&categoryID).Error
	}

	return
}

func (c *productDatabase) FindCategoryByName(ctx context.Context, categoryName string) (domain.Category, error) {

	var category domain.Category
	query := `SELECT id, category_id, name FROM categories WHERE name = $1`
	err := c.db.Raw(query, categoryName).Scan(&category).Error

	return category, err
}
func (c *productDatabase) FindCategoryByID(ctx context.Context, categoryID uint64) (domain.Category, error) {

	var category domain.Category
	query := `SELECT id, category_id, name FROM categories WHERE id = $1`
	err := c.db.Raw(query, categoryID).Scan(&category).Error

	return category, err
}
func (c *productDatabase) SaveVariation(ctx context.Context, variation request.AddVariation) (variationID uint64, err error) {

	query := `INSERT INTO variations (category_id, name) 
	VALUES($1, $2) RETURNING id AS variation_id`

	err = c.db.Raw(query, variation.CategoryID, variation.VariationName).Scan(&variationID).Error

	return
}

func (c *productDatabase) FindVariationByName(ctx context.Context, variationName string) (domain.Variation, error) {

	var variation domain.Variation
	query := `SELECT id, category_id, name FROM variations WHERE name = $1`
	err := c.db.Raw(query, variationName).Scan(&variation).Error

	return variation, err
}
func (c *productDatabase) FindVariationByID(ctx context.Context, variationID uint64) (domain.Variation, error) {

	var variation domain.Variation
	query := `SELECT id, category_id, name FROM variations WHERE id = $1`
	err := c.db.Raw(query, variationID).Scan(&variation).Error

	return variation, err
}
func (c *productDatabase) SaveVariationOption(ctx context.Context, variationOption request.AddVariationOption) (variationOptionID uint64, err error) {

	query := `INSERT INTO variation_options (variation_id,value) 
	VALUES($1, $2) RETURNING id AS variation_option_id`
	err = c.db.Raw(query, variationOption.VariationID, variationOption.VariationValue).Scan(&variationOptionID).Error

	return
}

func (c *productDatabase) FindVariationOptionByValue(ctx context.Context, variationValue string) (domain.VariationOption, error) {

	var variationOption domain.VariationOption
	query := `SELECT id, variation_id, value FROM variation_options WHERE value = $1`
	err := c.db.Raw(query, variationValue).Scan(&variationOption).Error

	return variationOption, err
}
func (c *productDatabase) FindVariationOptionByID(ctx context.Context, variationOptionID uint64) (domain.VariationOption, error) {

	var variationOption domain.VariationOption
	query := `SELECT id, variation_id, value FROM variation_options WHERE id = $1`
	err := c.db.Raw(query, variationOptionID).Scan(&variationOption).Error

	return variationOption, err
}
