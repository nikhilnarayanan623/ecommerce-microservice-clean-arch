package interfaces

import (
	"context"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/utils/response"
)

type UserClient interface {
	GetUserProfile(ctx context.Context, userID uint64) (response.User, error)
}
