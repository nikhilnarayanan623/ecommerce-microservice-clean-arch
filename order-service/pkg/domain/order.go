package domain

import "time"

type ShopOrder struct {
	ID              uint64    `json:"shop_order_id" gorm:"primaryKey;not null"`
	UserID          uint64    `json:"user_id" gorm:"not null"`
	OrderDate       time.Time `json:"order_date" gorm:"not null"`
	OrderTotalPrice float64   `json:"order_total_price" gorm:"not null"`
	Discount        float64   `json:"discount" gorm:"not null"`
}

type OrderLine struct {
	ID            uint64    `json:"id" gorm:"primaryKey;not null"`
	ProductItemID uint64    `json:"product_item_id" gorm:"not null"`
	ShopOrderID   uint64    `json:"shop_order_id" gorm:"not null"`
	ShopOrder     ShopOrder `json:"-"`
	Qty           uint64    `json:"qty" gorm:"not null"`
	Price         float64   `json:"price" gorm:"not null"`
}
