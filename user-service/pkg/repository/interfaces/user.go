package interfaces

import (
	"context"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/user-service/pkg/domain"
)

type UserRepository interface {
	FindUserByUserID(ctx context.Context, userID uint64) (user domain.User, err error)
	FindUserByEmail(ctx context.Context, email string) (domain.User, error)
	FindUserByPhone(ctx context.Context, phone string) (domain.User, error)

	SaveUser(ctx context.Context, user domain.User) (userID uint64, err error)
	UpdateUserVerified(ctx context.Context, userID uint64) error
}
