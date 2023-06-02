package usecase

import (
	"context"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/user-service/pkg/domain"
	repo "github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/user-service/pkg/repository/interfaces"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/user-service/pkg/usecase/interfaces"
)

type userUsecase struct {
	repo repo.UserRepository
}

func NewUserUsecase(repo repo.UserRepository) interfaces.UserUsecase {
	return &userUsecase{
		repo: repo,
	}
}

func (c *userUsecase) SaveUser(ctx context.Context, user domain.User) (userID uint64, err error) {

	return c.repo.SaveUser(ctx, user)
}

func (c *userUsecase) FindUserByEmail(ctx context.Context, email string) (domain.User, error) {
	return c.repo.FindUserByEmail(ctx, email)
}
func (c *userUsecase) FindUserByID(ctx context.Context, userID uint64) (domain.User, error) {
	return c.repo.FindUserByUserID(ctx, userID)
}

func (c *userUsecase) UpdateUserVerified(ctx context.Context, userID uint64) error {
	return c.repo.UpdateUserVerified(ctx, userID)
}
