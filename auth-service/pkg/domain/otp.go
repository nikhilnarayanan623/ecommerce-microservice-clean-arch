package domain

import (
	"time"
)

type OtpSession struct {
	OtpID    string    `json:"otp_id" gorm:"unique;not null"`
	UserID   uint64    `json:"user_id" gorm:"not null"`
	Phone    string    `json:"phone" gorm:"not null"`
	ExpireAt time.Time `json:"expire_at" gorm:"not null"`
}
