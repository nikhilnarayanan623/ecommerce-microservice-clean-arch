package domain

type Category struct {
	ID         uint      `json:"id" gorm:"primaryKey;not null"`
	CategoryID uint      `json:"category_id"`
	Category   *Category `json:"-"`
	Name       string    `json:"name" gorm:"unique;not null" binding:"required,min=1,max=30"`
}

type Variation struct {
	ID         uint     `json:"id" gorm:"primaryKey;not null"`
	CategoryID uint     `json:"category_id" gorm:"not null" binding:"required,numeric"`
	Category   Category `json:"-"`
	Name       string   `json:"name" gorm:"not null" binding:"required"`
}

type VariationOption struct {
	ID          uint      `json:"id" gorm:"primaryKey;not null"`
	VariationID uint      `json:"variation_id" gorm:"not null" binding:"required,numeric"`
	Variation   Variation `json:"-"`
	Value       string    `json:"value" gorm:"not null" binding:"required"`
}
