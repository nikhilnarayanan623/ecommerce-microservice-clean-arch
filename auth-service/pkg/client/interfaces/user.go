package interfaces

import (
	"context"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/auth-service/pkg/domain"
)

type UserClient interface {
	FindUserByEmail(ctx context.Context, email string) (domain.User, error)
	FindUserByPhone(ctx context.Context, phone string) (domain.User, error)
	SaveUser(ctx context.Context, user domain.UserSignupRequest) (userID uint64, err error)
	UpdateUserVerified(ctx context.Context, userID uint64) error
}
