package domain

import "time"

type User struct {
	ID          uint64    `json:"id" gorm:"primaryKey;unique"`
	FirstName   string    `json:"first_name" gorm:"not null" binding:"required,min=2,max=50"`
	LastName    string    `json:"last_name" gorm:"not null" binding:"required,min=1,max=50"`
	Age         uint64    `json:"age" binding:"required,numeric"`
	Email       string    `json:"email" gorm:"unique;not null" binding:"required,email"`
	Phone       string    `json:"phone" gorm:"unique" binding:"required,min=10,max=10"`
	Password    string    `json:"password" binding:"required"`
	Verified    bool      `json:"verified" gorm:"default:false"`
	BlockStatus bool      `json:"block_status" gorm:"not null;default:false"`
	CreatedAt   time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt   time.Time `json:"updated_at"`
}
