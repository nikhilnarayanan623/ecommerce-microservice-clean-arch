package repository

import (
	"context"
	"time"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/user-service/pkg/domain"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/user-service/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type userDatabase struct {
	db *gorm.DB
}

// New UserRepository using gorm postgres databse
func NewUserRepository(db *gorm.DB) interfaces.UserRepository {
	return &userDatabase{
		db: db,
	}
}

func (c *userDatabase) FindUserByUserID(ctx context.Context, userID uint64) (user domain.User, err error) {

	query := `SELECT * FROM users WHERE id = $1`
	err = c.db.Raw(query, userID).Scan(&user).Error

	return user, err
}

func (c *userDatabase) FindUserByEmail(ctx context.Context, email string) (user domain.User, err error) {

	query := `SELECT * FROM users WHERE email = $1`
	err = c.db.Raw(query, email).Scan(&user).Error

	return user, err
}

// Save a new user
func (c *userDatabase) SaveUser(ctx context.Context, user domain.User) (uint64, error) {
	query := `INSERT INTO users ( first_name, last_name, age, email, phone, password,created_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	createdAt := time.Now()
	err := c.db.Raw(query, user.FirstName, user.LastName,
		user.Age, user.Email, user.Phone, user.Password, createdAt).Scan(&user).Error

	return user.ID, err
}

func (c *userDatabase) UpdateUserVerified(ctx context.Context, userID uint64) error {
	query := `UPDATE users SET verified = 'T' WHERE id = $1`
	err := c.db.Exec(query, userID).Error

	return err
}
