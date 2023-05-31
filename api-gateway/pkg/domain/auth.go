package domain

type UserSignupRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Age       uint64 `json:"age" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Phone     string `json:"phone" binding:"required,min=10,max=10"`
	Password  string `json:"password" binding:"required,min=6,max=30"`
}
