package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/api/handler/interfaces"
	client "github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/client/interfaces"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/utils"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/utils/request"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/utils/response"
)

type orderHandler struct {
	client client.OrderClient
}

func NewOrderHandler(client client.OrderClient) interfaces.OrderHandler {

	return &orderHandler{
		client: client,
	}
}

func (c *orderHandler) PlaceOrder(ctx *gin.Context) {

	userID := utils.GetUserIDFromContext(ctx)

	shopOrderId, err := c.client.PlaceOrder(ctx, userID)

	if err != nil {
		response.ErrorResponse(ctx, "failed to place order", err, nil)
		return
	}

	response.SuccessResponse(ctx, "successfully order placed", gin.H{
		"shop_order_Id": shopOrderId,
	})
}

func (c *orderHandler) FindAllOrders(ctx *gin.Context) {

	userID := utils.GetUserIDFromContext(ctx)

	pagination := request.GetPagination(ctx)

	shopOrders, err := c.client.FindAllShopOrders(ctx, userID, pagination)

	if err != nil {
		response.ErrorResponse(ctx, "failed to find all shop orders", err, nil)
		return
	}

	if shopOrders == nil {
		response.SuccessResponse(ctx, "there is no shop orders to show")
	}

	response.SuccessResponse(ctx, "successfully found all shop orders of user", shopOrders)
}
