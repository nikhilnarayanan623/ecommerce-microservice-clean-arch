package response

type Category struct {
	ID               uint64
	CategoryID       uint64
	MainCategoryName string
	Name             string
}

type Product struct {
	ID           uint64
	Name         string
	Description  string
	Price        float64
	Image        string
	CategoryID   uint64
	CategoryName string
}
