//go:build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/api"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/api/handler"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/client"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/config"
)

func InitializeApi(cfg *config.Config) (*api.Server, error) {

	wire.Build(
		client.NewAuthClient,
		client.NewUserClient,
		client.NewProductClient,
		client.NewCartClient,
		client.NewOrderClient,

		handler.NewAuthHandler,
		handler.NewUserHandler,
		handler.NewProductHandler,
		handler.NewCartHandler,
		handler.NewOrderHandler,
		api.NewServerHTTP,
	)
	return &api.Server{}, nil
}
