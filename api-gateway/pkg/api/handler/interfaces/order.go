package interfaces

import "github.com/gin-gonic/gin"

type OrderHandler interface {
	PlaceOrder(ctx *gin.Context)
	FindAllOrders(ctx *gin.Context)
}
