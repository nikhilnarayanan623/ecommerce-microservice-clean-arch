//go:build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/auth-service/pkg/api"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/auth-service/pkg/api/service"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/auth-service/pkg/client"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/auth-service/pkg/config"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/auth-service/pkg/db"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/auth-service/pkg/otp"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/auth-service/pkg/repository"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/auth-service/pkg/token"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/auth-service/pkg/usecase"
)

func InitializeServices(cfg *config.Config) (*api.ServiceServer, error) {

	wire.Build(
		db.ConnectDatabase,
		repository.NewAuthRepository,
		otp.NewTwiloOtpAuth,
		token.NewJWTAuth,
		client.NewUserClient, usecase.NewAuthUsecase,
		service.NewAuthServiceServer,
		api.ServerGRPC,
	)
	return &api.ServiceServer{}, nil
}
