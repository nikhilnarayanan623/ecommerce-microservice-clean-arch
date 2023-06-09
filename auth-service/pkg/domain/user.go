package domain

type UserSignupRequest struct {
	FirstName string
	LastName  string
	Age       uint64
	Email     string
	Phone     string
	Password  string
}

type UserLoginRequest struct {
	Email    string
	Phone    string
	Password string
}

type User struct {
	ID          uint64
	FirstName   string
	LastName    string
	Age         uint64
	Email       string
	Phone       string
	Password    string
	Verified    bool
	BlockStatus bool
}
