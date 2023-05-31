package interfaces

import (
	"context"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/domain"
)

type AuthClient interface {
	UserSignup(ctx context.Context, req domain.UserSignupRequest) (userId uint64, err error)
}
