package response

import "time"

type ShopOrder struct {
	ID              uint64
	OrderDate       time.Time
	OrderTotalPrice float64
	Discount        float64
}
