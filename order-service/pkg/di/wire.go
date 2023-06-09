//go:build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/order-service/pkg/api"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/order-service/pkg/api/service"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/order-service/pkg/client"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/order-service/pkg/config"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/order-service/pkg/db"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/order-service/pkg/repository"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/order-service/pkg/usecase"
)

func InitializeServices(cfg *config.Config) (*api.Server, error) {

	wire.Build(
		db.ConnectDatabase,
		repository.NewOrderRepository,
		client.NewCartServiceClient,
		client.NewProductServiceClient,
		usecase.NewOrderUseCase,
		service.NewOrderServiceServer,
		api.NewServerGRPC,
	)
	return &api.Server{}, nil
}
