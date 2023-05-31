package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/api/handler/interfaces"
	client "github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/client/interfaces"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/domain"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/utils"
)

type authHandler struct {
	client client.AuthClient
}

func NewAuthHandler(client client.AuthClient) interfaces.AuthHandler {
	return &authHandler{
		client: client,
	}
}

func (c *authHandler) UserLogin(ctx *gin.Context) {

}

func (c *authHandler) UserSignup(ctx *gin.Context) {

	var body domain.UserSignupRequest

	if err := ctx.ShouldBindJSON(&body); err != nil {
		response := utils.ErrorResponse("failed bind inputs", err.Error(), body)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	userID, err := c.client.UserSignup(context.Background(), body)

	if err != nil {
		response := utils.ErrorResponse("faild to signup user", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.SuccessResponse("successfully user signup completed", userID)
	ctx.JSON(http.StatusOK, response)
}
