package interfaces

import (
	"context"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/auth-service/pkg/domain"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/auth-service/pkg/utils"
)

type AuthUseCase interface {
	UserSignup(ctx context.Context, user domain.SaveUserRequest) (otpID string, err error)
	OtpVerify(ctx context.Context, otpDetails utils.OtpVerify) (userID uint64, err error)
	GenerateAccessToken(ctx context.Context, userID uint64) (accessToken string, err error)
	GenereateRefreshToken(ctx context.Context, userID uint64) (refreshToken string, err error)
}