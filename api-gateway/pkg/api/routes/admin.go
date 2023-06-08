package routes

import (
	"github.com/gin-gonic/gin"
	handler "github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/api/handler/interfaces"
)

func SetupAdminRoutes(admin *gin.RouterGroup, productHandler handler.ProductHandler) {

	admin.POST("/category", productHandler.AddCategory)
	admin.GET("/category", productHandler.FindAllCategories)
	admin.POST("/variation", productHandler.AddVariation)
	admin.POST("/variation/option", productHandler.AddVariationOption)

	product := admin.Group("/product")
	{
		product.POST("/", productHandler.AddProduct)
		product.GET("/", productHandler.FindAllProducts)

		product.POST("/items", productHandler.AddProductItem)
		product.GET("/items/:product_id", productHandler.FindAllProductItems)
	}

}
