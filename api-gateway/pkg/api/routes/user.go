package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	handler "github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/api/handler/interfaces"
)

func SetupUserRoutes(user *gin.RouterGroup, authHandler handler.AuthHandler, userHandler handler.UserHandler) {

	auth := user.Group("/auth")
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

	user.Use(authHandler.AuthenticateUser)
)

}
