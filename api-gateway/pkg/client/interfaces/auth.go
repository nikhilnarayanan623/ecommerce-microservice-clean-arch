package interfaces

import (
	"context"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/domain"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/utils/request"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/utils/response"
)

type AuthClient interface {
	UserSignup(ctx context.Context, req domain.UserSignupRequest) (otpID string, err error)
	UserSignupVerify(ctx context.Context, otpVerify request.OtpVerify) (response.TokenResponse, error)
	UserLogin(ctx context.Context, loginDetails domain.UserLoginRequest) (response.TokenResponse, error)
	VerifyUserAccessToken(ctx context.Context, accessToken string) (userID uint64, err error)
	RefreshAccessTokenForUser(ctx context.Context, refreshToken string) (accessToken string, err error)
}
