package repository

import (
	"context"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/auth-service/pkg/domain"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/auth-service/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type authDatabase struct {
	db *gorm.DB
}

// New AuthRepository with gorm database
func NewAuthRepository(db *gorm.DB) interfaces.AuthRepository {
	return &authDatabase{
		db: db,
	}
}

func (c *authDatabase) Transactions(trxFunc func(repo interfaces.AuthRepository) error) error {

	trx := c.db.Begin()

	repo := NewAuthRepository(trx)
	err := trxFunc(repo)

	if err != nil {
		return err
	}

	if err := trx.Commit().Error; err != nil {
		return err
	}
	return nil
}

func (c *authDatabase) SaveRefreshSession(ctx context.Context, refreshSession domain.RefreshSession) error {
	query := `INSERT INTO refresh_sessions (token_id, user_id, refresh_token, expire_at) 
	VALUES ($1, $2, $3, $4)`
	err := c.db.Exec(query, refreshSession.TokenID, refreshSession.UserID, refreshSession.RefreshToken, refreshSession.ExpireAt).Error

	return err
}

func (c *authDatabase) FindRefreshSessionByTokenID(ctx context.Context, tokenID string) (refreshSession domain.RefreshSession, err error) {
	query := `SELECT * FROM refresh_sessions WHERE token_id = $1`

	err = c.db.Raw(query, tokenID).Scan(&refreshSession).Error

	return
}

func (c *authDatabase) SaveOtpSession(ctx context.Context, otpSession domain.OtpSession) error {

	query := `INSERT INTO otp_sessions (otp_id, user_id, phone ,expire_at) 
	VALUES ($1, $2, $3, $4)`
	err := c.db.Exec(query, otpSession.OtpID, otpSession.UserID, otpSession.Phone, otpSession.ExpireAt).Error

	return err
}
func (c *authDatabase) FindOtpSession(ctx context.Context, otpID string) (otpSession domain.OtpSession, err error) {

	query := `SELECT * FROM otp_sessions WHERE otp_id = $1`
	err = c.db.Raw(query, otpID).Scan(&otpSession).Error

	return otpSession, err
}

func (c *authDatabase) DeleteAllOtpSessionsByUserID(ctx context.Context, userID uint64) error {

	query := `DELETE FROM otp_sessions WHERE user_id = $1`
	err := c.db.Exec(query, userID).Error

	return err
}
