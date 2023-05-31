package handler

import "github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/api/handler/interfaces"

type userHandler struct {
}

func NewUserHandler() interfaces.UserHandler {
	return &userHandler{}
}
