package interfaces

import "github.com/gin-gonic/gin"

type ProductHandler interface {
	AddCategory(ctx *gin.Context)
	AddVariation(ctx *gin.Context)
	AddVariationOption(ctx *gin.Context)
}
