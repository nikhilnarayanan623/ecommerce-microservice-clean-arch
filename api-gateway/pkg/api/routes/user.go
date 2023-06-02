package routes

import (
	"github.com/gin-gonic/gin"
	handler "github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/api/handler/interfaces"
)

func SetupUserRoutes(router *gin.Engine, authHandler handler.AuthHandler, userHandler handler.UserHandler) {

	auth := router.Group("/auth")
	{
		signup := auth.Group("/signup")
		{
			signup.POST("/", authHandler.UserSignup)
			signup.POST("/verify", authHandler.UserSignupVerify)

		}

		login := auth.Group("/login")
		{
			login.POST("/", authHandler.UserLogin)
		}

		auth.POST("/refresh-token", authHandler.RefreshAccessTokenForUser)
	}

}
