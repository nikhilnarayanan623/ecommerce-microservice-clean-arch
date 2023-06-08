package domain

type Cart struct {
	ID              uint64  `gorm:"primaryKey;not null"`
	UserID          uint64  `gorm:"not null"`
	TotalPrice      float64 `gorm:"not null"`
	AppliedCouponID uint64
	DiscountAmount  float64
}

type CartItem struct {
	ID            uint64 `gorm:"primaryKey;not null"`
	CartID        uint64
	Cart          Cart
	ProductItemID uint64 `gorm:"not null"`
	Qty           uint64 `gorm:"not null"`
}
