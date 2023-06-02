package interfaces

import (
	"context"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/domain"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/utils"
)

type AuthClient interface {
	UserSignup(ctx context.Context, req domain.UserSignupRequest) (otpID string, err error)
	UserSignupVerify(ctx context.Context, otpVerify utils.OtpVerify) (utils.TokenResponse, error)
	RefreshAccesstokenForUser(ctx context.Context, refreshToken string) (accessToken string, err error)
}
