package interfaces

import "github.com/gin-gonic/gin"

type ProductHandler interface {
	AddCategory(ctx *gin.Context)
	AddVariation(ctx *gin.Context)
	AddVariationOption(ctx *gin.Context)
	FindAllCategories(ctx *gin.Context)

	AddProduct(ctx *gin.Context)
	FindAllProducts(ctx *gin.Context)

	AddProductItem(ctx *gin.Context)
	FindAllProductItems(ctx *gin.Context)
}
