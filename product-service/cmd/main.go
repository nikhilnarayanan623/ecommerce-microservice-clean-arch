package main

import (
	"log"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/product-service/pkg/config"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/product-service/pkg/di"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config error:%s", err.Error())
	}

	service, err := di.InitializeServices(cfg)
	if err != nil {
		log.Fatalf("failed initialize service error:%s", err.Error())
	}

	service.Start()
}
