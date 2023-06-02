package interfaces

import (
	"context"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/auth-service/pkg/domain"
)

type AuthRepository interface {
	SaveRefreshSession(ctx context.Context, refreshSession domain.RefreshSession) error
	FindRefreshSessionByTokenID(ctx context.Context, tokenID string) (refreshSession domain.RefreshSession, err error)
	SaveOtpSession(ctx context.Context, OTPSession domain.OTPSession) error
	FindOtpSession(ctx context.Context, OTPID string) (domain.OTPSession, error)
}
