// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/product-service/pkg/api"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/product-service/pkg/api/service"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/product-service/pkg/config"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/product-service/pkg/db"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/product-service/pkg/repository"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/product-service/pkg/usecase"
)

// Injectors from wire.go:

func InitializeServices(cfg *config.Config) (*api.Server, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
	productRepository := repository.NewProductRepository(gormDB)
	productUseCase := usecase.NewProductUseCase(productRepository)
	productServiceServer := service.NewProductServiceServer(productUseCase)
	server, err := api.NewServerGRPC(cfg, productServiceServer)
	if err != nil {
		return nil, err
	}
	return server, nil
}
