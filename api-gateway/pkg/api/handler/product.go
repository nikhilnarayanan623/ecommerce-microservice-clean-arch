package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/api/handler/interfaces"
	client "github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/client/interfaces"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/utils/request"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/utils/response"
)

type productHandler struct {
	client client.ProductClient
}

func NewProductHandler(client client.ProductClient) interfaces.ProductHandler {
	return &productHandler{
		client: client,
	}
}

func (c *productHandler) AddCategory(ctx *gin.Context) {

	var body request.AddCategory
	if err := ctx.ShouldBindJSON(&body); err != nil {
		response.ErrorResponse(ctx, "failed to bind inputs", err, body)
		return
	}

	categoryID, err := c.client.AddCategory(ctx, body)
	if err != nil {
		response.ErrorResponse(ctx, "failed to add category", err, nil)
		return
	}
	response.SuccessResponse(ctx, "successfully category added", categoryID)
}
func (c *productHandler) AddVariation(ctx *gin.Context) {

	var body request.AddVariation
	if err := ctx.ShouldBindJSON(&body); err != nil {
		response.ErrorResponse(ctx, "failed to bind inputs", err, body)
		return
	}

	variationID, err := c.client.AddVariation(ctx, body)
	if err != nil {
		response.ErrorResponse(ctx, "failed to add variation", err, nil)
		return
	}
	response.SuccessResponse(ctx, "successfully variation added", variationID)
}
func (c *productHandler) AddVariationOption(ctx *gin.Context) {

	var body request.AddVariationOption
	if err := ctx.ShouldBindJSON(&body); err != nil {
		response.ErrorResponse(ctx, "failed to bind inputs", err, body)
		return
	}

	variationOptionID, err := c.client.AddVariationOption(ctx, body)

	if err != nil {
		response.ErrorResponse(ctx, "failed to add variation_option", err, nil)
		return
	}
	response.SuccessResponse(ctx, "successfully variation_option added", variationOptionID)
}
