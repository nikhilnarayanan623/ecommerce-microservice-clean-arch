package interfaces

import "github.com/gin-gonic/gin"

type CartHandler interface {
	AddToCart(ctx *gin.Context)
	FindCart(ctx *gin.Context)
}
