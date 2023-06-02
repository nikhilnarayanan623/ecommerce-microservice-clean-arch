package interfaces

import (
	"context"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/auth-service/pkg/domain"
)

type AuthRepository interface {
	Transactions(trxFunc func(repo AuthRepository) error) error

	SaveRefreshSession(ctx context.Context, refreshSession domain.RefreshSession) error
	FindRefreshSessionByTokenID(ctx context.Context, tokenID string) (refreshSession domain.RefreshSession, err error)
	SaveOtpSession(ctx context.Context, otpSession domain.OtpSession) error
	FindOtpSession(ctx context.Context, Otp string) (domain.OtpSession, error)
	DeleteAllOtpSessionsByUserID(ctx context.Context, userID uint64) error
}
