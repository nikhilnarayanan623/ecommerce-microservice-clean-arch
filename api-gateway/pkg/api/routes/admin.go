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

	engine.POST("/product", productHandler.AddProduct)
	engine.GET("/product", productHandler.FindAllProducts)

	engine.POST("/product/item", productHandler.AddProductItem)
	engine.GET("/product/item/:product_id", productHandler.FindAllProductItems)
}
