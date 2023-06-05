package repository

import (
	"context"
	"time"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/product-service/pkg/domain"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/product-service/pkg/repository/interfaces"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/product-service/pkg/utils/request"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/product-service/pkg/utils/response"
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

func (c *productDatabase) Transactions(ctx context.Context, trxFn func(repo interfaces.ProductRepository) error) error {

	trx := c.db.Begin()

	repo := NewProductRepository(trx)

	if err := trxFn(repo); err != nil {
		trx.Rollback()
		return err
	}

	if err := trx.Commit().Error; err != nil {
		trx.Rollback()
		return err
	}
	return nil
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

func (c *productDatabase) FindAllCategories(ctx context.Context) (categories []response.Category, err error) {

	query := `SELECT c.id, c.name, mc.id AS category_id, mc.name AS main_category_name  FROM categories c 
	LEFT JOIN categories mc ON c.category_id = mc.id`
	err = c.db.Raw(query).Scan(&categories).Error

	return
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

// func (c *productDatabase) FindVariationOptionByID(ctx context.Context, variationOptionID uint64) (domain.VariationOption, error) {

// 	var variationOption domain.VariationOption
// 	query := `SELECT id, variation_id, value FROM variation_options WHERE id = $1`
// 	err := c.db.Raw(query, variationOptionID).Scan(&variationOption).Error

// 	return variationOption, err
// }

func (c *productDatabase) IsValidVariationOptionID(ctx context.Context, variationOptionID uint64) (valid bool, err error) {

	query := `SELECT EXISTS(SELECT 1) AS valid FROM variation_options WHERE id = $1`
	err = c.db.Raw(query, variationOptionID).Scan(&valid).Error

	return
}

func (c *productDatabase) SaveProduct(ctx context.Context, product request.AddProduct) (productID uint64, err error) {

	query := `INSERT INTO products (name, description, category_id, price, image,created_at) 
	 VALUES ($1, $2, $3, $4, $5, $6) RETURNING id AS product_id`
	createdAt := time.Now()
	err = c.db.Raw(query, product.Name, product.Description, product.CategoryID,
		product.Price, product.Image, createdAt).Scan(&productID).Error

	return
}

func (c *productDatabase) IsValidProductID(ctx context.Context, productID uint64) (valid bool, err error) {

	query := `SELECT EXISTS(SELECT 1) AS valid FROM products WHERE id = $1`
	err = c.db.Raw(query, productID).Scan(&valid).Error

	return
}

func (c *productDatabase) FindAllProducts(ctx context.Context, pagination request.Pagination) (products []response.Product, err error) {

	limit := pagination.Count
	offset := (pagination.PageNumber - 1) * limit

	query := `SELECT p.id, p.name, p.description, p.price, p.image, p.category_id, c.name AS category_name 
	FROM products p INNER JOIN categories c ON c.id = p.category_id 
	LIMIT $1 OFFSET  $2`

	err = c.db.Raw(query, limit, offset).Scan(&products).Error

	return
}
func (c *productDatabase) IsProductNameAlreadyExist(ctx context.Context, productName string) (exist bool, err error) {

	query := `SELECT DISTINCT EXISTS(SELECT 1 FROM products WHERE name = $1) AS exist FROM products`
	err = c.db.Raw(query, productName).Scan(&exist).Error

	return
}

func (c *productDatabase) SaveProductItem(ctx context.Context, productItem request.AddProductItem) (productItemID uint64, err error) {

	query := `INSERT INTO product_items (product_id, qty_in_stock, price, sku,created_at) VALUES($1, $2, $3, $4, $5) 
	 RETURNING id AS product_item_id`
	createdAt := time.Now()
	err = c.db.Raw(query, productItem.ProductID, productItem.QtyInStock, productItem.Price, productItem.SKU, createdAt).
		Scan(&productItemID).Error

	return
}

func (c *productDatabase) SaveProductConfiguration(ctx context.Context, productItemID, variationOptionID uint64) error {

	query := `INSERT INTO product_configurations (product_item_id, variation_option_id) VALUES ($1, $2)`
	err := c.db.Exec(query, productItemID, variationOptionID).Error

	return err
}

func (c *productDatabase) FindProductItemsByProductID(ctx context.Context, productID uint64) (productItems []response.ProductItem, err error) {

	query := `SELECT pi.id, pi.price, pi.qty_in_stock, pi.sku, 
	pi.discount_price, p.name, vo.value AS variation_value  FROM product_items pi 
	INNER JOIN products p ON pi.product_id = p.id 
	INNER JOIN product_configurations pc ON pi.id = pc.product_item_id  
	INNER JOIN variation_options vo ON pc.variation_option_id = vo.id 
	WHERE pi.product_id = $1`
	err = c.db.Raw(query, productID).Scan(&productItems).Error

	return
}

func (c *productDatabase) IsProductItemAlreadyExist(ctx context.Context, productID, variationOptionID uint64) (exist bool, err error) {

	query := `SELECT CASE WHEN id != 0 THEN 'T' ELSE 'F' END AS exist FROM product_items pi 
	WHERE id = ( SELECT product_item_id FROM product_configurations WHERE variation_option_id = $1 )                        
	AND pi.product_id = $2`
	err = c.db.Raw(query, variationOptionID, productID).Scan(&exist).Error

	return
}
