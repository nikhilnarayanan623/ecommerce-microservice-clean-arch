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

	admin.POST("/product", productHandler.AddProduct)
	admin.GET("/product", productHandler.FindAllProducts)

	admin.POST("/product/item", productHandler.AddProductItem)
	admin.GET("/product/item/:product_id", productHandler.FindAllProductItems)
}
