//go:build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/product-service/pkg/api"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/product-service/pkg/api/service"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/product-service/pkg/config"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/product-service/pkg/db"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/product-service/pkg/repository"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/product-service/pkg/usecase"
)

func InitializeServices(cfg *config.Config) (*api.Server, error) {

	wire.Build(
		db.ConnectDatabase,
		repository.NewProductRepository,
		usecase.NewProductUseCase,
		service.NewProductServiceServer,
		api.NewServerGRPC,
	)
	return &api.Server{}, nil
}
