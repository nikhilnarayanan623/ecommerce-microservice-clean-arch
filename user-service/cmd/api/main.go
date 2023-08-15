package main

import (
	"log"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/user-service/pkg/config"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/user-service/pkg/di"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config error:%s", err.Error())
	}

	service, err := di.InitializeService(cfg)
	if err != nil {
		log.Fatalf("failed initialize service error:%s", err.Error())
	}

	service.Start()
}
