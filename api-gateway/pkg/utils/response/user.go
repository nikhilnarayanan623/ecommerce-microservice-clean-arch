package response

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       uint64 `json:"age"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}
