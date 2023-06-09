package domain

import "time"

type Product struct {
	ID            uint64    `json:"id" gorm:"primaryKey;not null"`
	Name          string    `json:"name" gorm:"not null" binding:"required,min=3,max=50"`
	Description   string    `json:"description" gorm:"not null" binding:"required,min=10,max=100"`
	CategoryID    uint64    `json:"category_id" binding:"omitempty,numeric"`
	Category      Category  `json:"-"`
	Price         uint64    `json:"price" gorm:"not null" binding:"required,numeric"`
	DiscountPrice uint64    `json:"discount_price"`
	Image         string    `json:"image" gorm:"not null"`
	CreatedAt     time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type ProductItem struct {
	ID            uint64 `json:"id" gorm:"primaryKey;not null"`
	ProductID     uint64 `json:"product_id" gorm:"not null" binding:"required,numeric"`
	Product       Product
	QtyInStock    uint64    `json:"qty_in_stock" gorm:"not null" binding:"required,numeric"`
	Price         uint64    `json:"price" gorm:"not null" binding:"required,numeric"`
	SKU           string    `json:"sku" gorm:"unique;not null"`
	DiscountPrice uint64    `json:"discount_price"`
	CreatedAt     time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type ProductConfiguration struct {
	ProductItemID     uint64          `json:"product_item_id" gorm:"not null"`
	ProductItem       ProductItem     `json:"-"`
	VariationOptionID uint64          `json:"variation_option_id" gorm:"not null"`
	VariationOption   VariationOption `json:"-"`
}
