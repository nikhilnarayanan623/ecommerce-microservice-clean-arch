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

type AddProductItem struct {
	ProductID         uint64  `json:"product_id" binding:"required"`
	QtyInStock        uint64  `json:"qty_in_stock"  binding:"required"`
	Price             float64 `json:"price"  binding:"required,min=1"`
	VariationOptionID uint64  `json:"variation_option_id"  binding:"required"`
}
