package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/api/handler/interfaces"
	client "github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/client/interfaces"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/utils"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/utils/response"
)

type userHandler struct {
	client client.UserClient
}

func NewUserHandler(client client.UserClient) interfaces.UserHandler {
	return &userHandler{
		client: client,
	}
}

func (c *userHandler) GetProfile(ctx *gin.Context) {

	userID := utils.GetUserIDFromContext(ctx)

	user, err := c.client.GetUserProfile(ctx, userID)

	if err != nil {
		response.ErrorResponse(ctx, "failed get user details", err, nil)
		return

	}

	response.SuccessResponse(ctx, "successfully got user details", user)
}
