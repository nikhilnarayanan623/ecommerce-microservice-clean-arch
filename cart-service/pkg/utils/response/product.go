package response

type ProductItem struct {
	ID             uint64
	Name           string
	QtyInStock     uint64
	Price          float64
	SKU            string
	DiscountPrice  uint64
	VariationValue string
}
