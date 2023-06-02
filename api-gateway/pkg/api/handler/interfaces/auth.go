package interfaces

import "github.com/gin-gonic/gin"

type AuthHandler interface {
	UserSignup(ctx *gin.Context)
	UserSignupVerify(ctx *gin.Context)
	UserLogin(ctx *gin.Context)

	RefreshAccessTokenForUser(ctx *gin.Context)
}
