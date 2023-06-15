package token

import "time"

type TokenAuth interface {
	GenerateToken(req TokenRequest) (tokenString string, err error)
	VerifyToken(req TokenVerifyRequest) (TokenVerifyResponse, error)
}

type UserType string

const (
	Admin UserType = "admin"
	User  UserType = "user"
)

type TokenRequest struct {
	TokenID        string
	UserID         uint64
	UsedFor        UserType
	ExpirationDate time.Time
}

type TokenVerifyRequest struct {
	TokenString string
	UsedFor     UserType
}

type TokenVerifyResponse struct {
	TokenID string
	UserID  uint64
}
