package routes

import (
	"github.com/gin-gonic/gin"
	handler "github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/api/handler/interfaces"
)

func SetupUserRoutes(user *gin.RouterGroup, authHandler handler.AuthHandler,
	userHandler handler.UserHandler, productHandler handler.ProductHandler) {

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

	products := user.Group("/product")
	{
		products.GET("/", productHandler.FindAllProducts)
		products.GET("/items/:product_id", productHandler.FindAllProductItems)
	}
}
