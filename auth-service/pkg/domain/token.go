package domain

import (
	"time"
)

type RefreshSession struct {
	TokenID      string    `json:"token_id" gorm:"not null"`
	UserID       uint64    `json:"user_id" gorm:"not null"`
	RefreshToken string    `json:"refresh_token" gorm:"not null"`
	ExpireAt     time.Time `json:"expire_at" gorm:"not null"`
	IsBlocked    bool      `json:"is_blocked" gorm:"not null;default:false"`
}
