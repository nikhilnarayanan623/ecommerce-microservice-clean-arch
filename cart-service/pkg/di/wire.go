//go:build wireinject
package di

import (
	"github.com/google/wire"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/cart-service/pkg/api"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/cart-service/pkg/api/service"
	client "github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/cart-service/pkg/client"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/cart-service/pkg/config"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/cart-service/pkg/db"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/cart-service/pkg/repository"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/cart-service/pkg/usecase"
)

func InitializeService(cfg *config.Config) (*api.Server, error) {

	wire.Build(
		db.ConnectDatabase,
		repository.NewCartRepository,
		usecase.NewCartUseCase,
		client.NewProductServiceClient,
		service.NewCartServiceServer,
		api.NewServerGRPC,
	)
	return &api.Server{}, nil
}
