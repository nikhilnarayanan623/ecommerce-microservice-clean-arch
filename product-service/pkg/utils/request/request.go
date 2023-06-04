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

type Pagination struct {
	PageNumber uint64
	Count      uint64
}
