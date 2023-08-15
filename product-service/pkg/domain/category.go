package domain

type Category struct {
	ID         uint64    `json:"id" gorm:"primaryKey;not null"`
	CategoryID uint64    `json:"category_id"`
	Category   *Category `json:"-"`
	Name       string    `json:"name" gorm:"unique;not null" binding:"required,min=1,max=30"`
}

type Variation struct {
	ID         uint64   `json:"id" gorm:"primaryKey;not null"`
	CategoryID uint64   `json:"category_id" gorm:"not null" binding:"required,numeric"`
	Category   Category `json:"-"`
	Name       string   `json:"name" gorm:"not null" binding:"required"`
}

type VariationOption struct {
	ID          uint64    `json:"id" gorm:"primaryKey;not null"`
	VariationID uint64    `json:"variation_id" gorm:"not null" binding:"required,numeric"`
	Variation   Variation `json:"-"`
	Value       string    `json:"value" gorm:"not null" binding:"required"`
}
