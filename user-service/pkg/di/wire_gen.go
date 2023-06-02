// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/user-service/pkg/api"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/user-service/pkg/api/service"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/user-service/pkg/config"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/user-service/pkg/db"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/user-service/pkg/repository"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/user-service/pkg/usecase"
)

// Injectors from wire.go:

func InitializeService(cfg *config.Config) (*api.ServiceServer, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
	userRepository := repository.NewUserRepository(gormDB)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userServiceServer := service.NewUserServiceServer(userUsecase)
	serviceServer, err := api.NewServerGRPC(cfg, userServiceServer)
	if err != nil {
		return nil, err
	}
	return serviceServer, nil
}