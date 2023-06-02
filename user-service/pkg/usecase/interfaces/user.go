package interfaces

import (
	"context"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/user-service/pkg/domain"
)

type UserUsecase interface {
	SaveUser(ctx context.Context, user domain.User) (userID uint64, err error)
	FindUserByEmail(ctx context.Context, email string) (domain.User, error)
	FindUserByID(ctx context.Context, userID uint64) (domain.User, error)
	UpdateUserVerified(ctx context.Context, userID uint64) error
}
