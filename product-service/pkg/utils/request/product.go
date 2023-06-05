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
	Name        string
	Description string
	CategoryID  uint64
	Price       float64
	Image       string
}

type AddProductItem struct {
	ProductID         uint64
	QtyInStock        uint64
	Price             float64
	SKU               string
	VariationOptionID uint64
}
