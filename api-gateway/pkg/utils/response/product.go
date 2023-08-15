package response

type Category struct {
	ID               uint64 `json:"id"`
	CategoryID       uint64 `json:"category_id"`
	MainCategoryName string `json:"main_category_name"`
	Name             string `json:"name"`
}

type Product struct {
	ID           uint64  `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	Image        string  `json:"image"`
	CategoryID   uint64  `json:"category_id"`
	CategoryName string  `json:"category_name"`
}

type ProductItem struct {
	ID             uint64  `json:"id"`
	Name           string  `json:"name"`
	QtyInStock     uint64  `json:"qty_in_stock"`
	Price          float64 `json:"price"`
	SKU            string  `json:"sku"`
	DiscountPrice  uint64  `json:"discount_price"`
	VariationValue string  `json:"variation_value"`
}
