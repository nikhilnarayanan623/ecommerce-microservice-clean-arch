package api

import (
	"github.com/gin-gonic/gin"
	handler "github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/api/handler/interfaces"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/api/routes"
)

type Server struct {
	engine *gin.Engine
}

// NewServerHTTP creates a new server with given handler functions
func NewServerHTTP(authHandler handler.AuthHandler, userHandler handler.UserHandler) *Server {
	engine := gin.New()
	engine.Use(gin.Logger())

	routes.SetupUserRoutes(engine, authHandler, userHandler)
	routes.SetupAdminRoutes(engine)

	return &Server{
		engine: engine,
	}
}
