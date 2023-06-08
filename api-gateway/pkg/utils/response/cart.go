package response

type Cart struct {
	TotalPrice float64 `json:"total_price"`
	CartItems  []CartItem
}

type CartItem struct {
	ProductItemID  uint64  `json:"product_item_id"`
	ProductName    string  `json:"product_name"`
	SKU            string  `json:"sku"`
	VariationValue string  `json:"variation_value"`
	Price          float64 `json:"price"`
	SubTotal       float64 `json:"sub_total"`
	Qty            uint64  `json:"qty"`
	QtyInStock     uint64  `json:"qty_in_stock"`
	DiscountPrice  uint64  `json:"discount_price"`
}
