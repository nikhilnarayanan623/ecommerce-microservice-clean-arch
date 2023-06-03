// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/api"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/api/handler"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/client"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/config"
)

// Injectors from wire.go:

func InitializeApi(cfg *config.Config) (*api.Server, error) {
	authClient, err := client.NewAuthClient(cfg)
	if err != nil {
		return nil, err
	}
	authHandler := handler.NewAuthHandler(authClient)
	userHandler := handler.NewUserHandler()
	productClient, err := client.NewProductClient(cfg)
	if err != nil {
		return nil, err
	}
	productHandler := handler.NewProductHandler(productClient)
	server := api.NewServerHTTP(cfg, authHandler, userHandler, productHandler)
	return server, nil
}
