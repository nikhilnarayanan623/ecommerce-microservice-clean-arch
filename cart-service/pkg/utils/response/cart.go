package response

type CartItem struct {
	ProductItemID  uint64
	ProductName    string
	SKU            string
	VariationValue string
	Price          float64
	SubTotal       float64
	Qty            uint64
	QtyInStock     uint64
	DiscountPrice  uint64
}
