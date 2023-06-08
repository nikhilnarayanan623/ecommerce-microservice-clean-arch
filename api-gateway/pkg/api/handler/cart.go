package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/api/handler/interfaces"
	client "github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/client/interfaces"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/utils"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/utils/response"
)

type cartHandler struct {
	client client.CartClient
}

func NewCartHandler(client client.CartClient) interfaces.CartHandler {
	return &cartHandler{
		client: client,
	}
}

func (c *cartHandler) AddToCart(ctx *gin.Context) {
	productItemID, err := utils.StringToUint64(ctx.Param("product_item_id"))
	if err != nil {
		response.ErrorResponse(ctx, "failed to parse params", err, nil)
		return
	}

	userID := utils.GetUserIDFromContext(ctx)

	err = c.client.AddToCart(ctx, userID, productItemID)
	if err != nil {
		response.ErrorResponse(ctx, "failed to add product to cart", err, nil)
		return
	}

	response.SuccessResponse(ctx, "successfully product_item added to cart")
}
func (c *cartHandler) FindCart(ctx *gin.Context) {

	userID := utils.GetUserIDFromContext(ctx)

	cart, err := c.client.FindCart(ctx, userID)

	if err != nil {
		response.ErrorResponse(ctx, "failed to get user cart", err, nil)
		return
	}
	response.SuccessResponse(ctx, "successfully found user cart", cart)
}
