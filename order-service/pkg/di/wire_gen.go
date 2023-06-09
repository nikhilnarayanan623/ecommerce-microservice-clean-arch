// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/order-service/pkg/api"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/order-service/pkg/api/service"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/order-service/pkg/client"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/order-service/pkg/config"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/order-service/pkg/db"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/order-service/pkg/repository"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/order-service/pkg/usecase"
)

// Injectors from wire.go:

func InitializeServices(cfg *config.Config) (*api.Server, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
	orderRepository := repository.NewOrderRepository(gormDB)
	cartClient, err := client.NewCartServiceClient(cfg)
	if err != nil {
		return nil, err
	}
	productClient, err := client.NewProductServiceClient(cfg)
	if err != nil {
		return nil, err
	}
	orderUseCase := usecase.NewOrderUseCase(orderRepository, cartClient, productClient)
	orderServiceServer := service.NewOrderServiceServer(orderUseCase)
	server, err := api.NewServerGRPC(cfg, orderServiceServer)
	if err != nil {
		return nil, err
	}
	return server, nil
}
