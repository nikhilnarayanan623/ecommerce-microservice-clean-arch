package interfaces

import "github.com/gin-gonic/gin"

type UserHandler interface {
	GetProfile(ctx *gin.Context)
}
