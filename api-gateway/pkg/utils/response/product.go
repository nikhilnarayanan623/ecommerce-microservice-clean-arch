package response

type Category struct {
	ID               uint64 `json:"id"`
	CategoryID       uint64 `json:"category_id"`
	MainCategoryName string `json:"main_category_name"`
	Name             string `json:"name"`
}
