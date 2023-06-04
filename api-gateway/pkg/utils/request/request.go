package request

type AddCategory struct {
	MainCategoryID uint64
	CategoryName   string
}

type AddVariation struct {
	CategoryID    uint64
	VariationName string
}

type AddVariationOption struct {
	VariationID    uint64
	VariationValue string
}

type AddProduct struct {
	Name        string  `json:"product_name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	CategoryID  uint64  `json:"category_id" binding:"required,min=1"`
	Price       float64 `json:"price" binding:"required,min=1"`
	Image       string  `json:"image" binding:"required"`
}

type Pagination struct {
	PageNumber uint64
	Count      uint64
}
