package token

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/auth-service/pkg/config"
)

type jwtClaims struct {
	TokenID string
	UserID  uint64
	jwt.RegisteredClaims
}

type jwtAuth struct {
	adminSecretKey string
	userSecretKey  string
}

// New TokenAuth With JWT Implimentation
func NewJWTAuth(cfg *config.Config) TokenAuth {

	return &jwtAuth{
		adminSecretKey: cfg.AdminSecretKey,
		userSecretKey:  cfg.UserSecretKey,
	}
}

// Generate a new JWT token string from token request
func (c *jwtAuth) GenerateToken(req TokenRequest) (string, error) {

	if req.UsedFor != Admin && req.UsedFor != User {
		return "", fmt.Errorf("invalid user type")
	}

	claims := &jwtClaims{
		TokenID: req.TokenID,
		UserID:  req.UserID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(req.ExpirationDate),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	if req.UsedFor == Admin {
		return token.SignedString([]byte(c.adminSecretKey))
	}
	return token.SignedString([]byte(c.userSecretKey))
}

// Verify JWT token string and return TokenResponse
func (c *jwtAuth) VerifyToken(req TokenVerifyRequest) (TokenVerifyResponse, error) {

	if req.UsedFor != Admin && req.UsedFor != User {
		return TokenVerifyResponse{}, fmt.Errorf("invalid user type")
	}

	token, err := jwt.ParseWithClaims(req.TokenString, &jwtClaims{}, func(t *jwt.Token) (interface{}, error) {

		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method on token")
		}
		if req.UsedFor == Admin {
			return []byte(c.adminSecretKey), nil
		}
		return []byte(c.userSecretKey), nil
	})

	if err != nil {
		return TokenVerifyResponse{}, err
	}

	claims, ok := token.Claims.(*jwtClaims)
	if !ok {
		return TokenVerifyResponse{}, fmt.Errorf("failed to convert token claims")
	}

	return TokenVerifyResponse{
		TokenID: claims.TokenID,
		UserID:  claims.UserID,
	}, nil
}
