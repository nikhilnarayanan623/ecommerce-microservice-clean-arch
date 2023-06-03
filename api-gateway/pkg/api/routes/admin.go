package routes

import (
	"github.com/gin-gonic/gin"
	handler "github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/api/handler/interfaces"
)

func SetupAdminRoutes(engine *gin.Engine, productHandler handler.ProductHandler) {

	engine.POST("/category", productHandler.AddCategory)
	engine.GET("/category", productHandler.FindAllCategories)
	engine.POST("/variation", productHandler.AddVariation)
	engine.POST("/variation/option", productHandler.AddVariationOption)
}
