package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/api/handler/interfaces"
	client "github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/client/interfaces"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/utils"
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

func (c *productHandler) FindAllCategories(ctx *gin.Context) {

	categories, err := c.client.FindAllCategories(ctx)
	if err != nil {
		response.ErrorResponse(ctx, "failed to find all categories", err, nil)
		return
	}

	response.SuccessResponse(ctx, "successfully found all categories", categories)
}

func (c *productHandler) AddProduct(ctx *gin.Context) {

	var body request.AddProduct
	if err := ctx.ShouldBindJSON(&body); err != nil {
		response.ErrorResponse(ctx, "failed to bind inputs", err, body)
		return
	}

	productID, err := c.client.AddProduct(ctx, body)
	if err != nil {
		response.ErrorResponse(ctx, "failed to add product", err, nil)
		return
	}

	response.SuccessResponse(ctx, "successfully product added", productID)
}

func (c *productHandler) FindAllProducts(ctx *gin.Context) {

	pagination := utils.GetPagination(ctx)

	products, err := c.client.FindAllProducts(ctx, pagination)
	if err != nil {
		response.ErrorResponse(ctx, "failed to find all products", err, nil)
		return
	}

	response.SuccessResponse(ctx, "successfully found all products", products)
}

func (c *productHandler) AddProductItem(ctx *gin.Context) {

	var body request.AddProductItem
	if err := ctx.ShouldBindJSON(&body); err != nil {
		response.ErrorResponse(ctx, "failed to bind inputs", err, body)
		return
	}

	productItemID, err := c.client.AddProductItem(ctx, body)
	if err != nil {
		response.ErrorResponse(ctx, "failed to add product_items", err, nil)
		return
	}

	response.SuccessResponse(ctx, "successfully product_item added", productItemID)
}

func (c *productHandler) FindAllProductItems(ctx *gin.Context) {

	productIDStr := ctx.Param("product_id")

	productID, err := utils.StringToUint64(productIDStr)
	if err != nil {
		response.ErrorResponse(ctx, "failed to parse params", err, nil)
		return
	}

	productItems, err := c.client.FindAllProductItems(ctx, productID)

	if err != nil {
		response.ErrorResponse(ctx, "failed to find product items", err, nil)
		return
	}

	if len(productItems) == 0 {
		response.SuccessResponse(ctx, "there is no product items for given product_id")
		return
	}

	response.SuccessResponse(ctx, "successfully found product items", productItems)
}
