package api

import (
	"github.com/gin-gonic/gin"
	handler "github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/api/handler/interfaces"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/api/routes"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/config"
)

type Server struct {
	port   string
	engine *gin.Engine
}

// NewServerHTTP creates a new server with given handler functions
func NewServerHTTP(cfg *config.Config, authHandler handler.AuthHandler, userHandler handler.UserHandler,
	productHandler handler.ProductHandler) *Server {
	engine := gin.New()
	engine.Use(gin.Logger())

	routes.SetupUserRoutes(engine.Group("/"), authHandler, userHandler)
	routes.SetupAdminRoutes(engine.Group("/admin"), productHandler)

	return &Server{
		engine: engine,
		port:   cfg.Port,
	}
}

func (c *Server) Start() {
	c.engine.Run(c.port)
}
