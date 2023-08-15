//go:build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/user-service/pkg/api"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/user-service/pkg/api/service"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/user-service/pkg/config"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/user-service/pkg/db"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/user-service/pkg/repository"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/user-service/pkg/usecase"
)

func InitializeService(cfg *config.Config) (*api.ServiceServer, error) {

	wire.Build(
		db.ConnectDatabase,
		repository.NewUserRepository,
		usecase.NewUserUsecase,
		service.NewUserServiceServer,
		api.NewServerGRPC,
	)
	return &api.ServiceServer{}, nil
}
