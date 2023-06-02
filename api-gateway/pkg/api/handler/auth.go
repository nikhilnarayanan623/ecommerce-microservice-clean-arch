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

	var body domain.UserLoginRequest

	if err := ctx.ShouldBindJSON(&body); err != nil {
		response := utils.ErrorResponse("failed bind inputs", err.Error(), body)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	tokenRes, err := c.client.UserLogin(ctx, body)

	if err != nil {
		response := utils.ErrorResponse("failed to login", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.SuccessResponse("successfully login completede ", tokenRes)
	ctx.JSON(http.StatusOK, response)
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

	response := utils.SuccessResponse("successfully otp send to user registered number", userID)
	ctx.JSON(http.StatusOK, response)
}

func (c *authHandler) UserSignupVerify(ctx *gin.Context) {

	var body utils.OtpVerify
	if err := ctx.ShouldBindJSON(&body); err != nil {
		response := utils.ErrorResponse("failed to bind inputs", err.Error(), body)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	tokenRes, err := c.client.UserSignupVerify(ctx, utils.OtpVerify{
		OtpID:   body.OtpID,
		OtpCode: body.OtpCode,
	})

	if err != nil {
		response := utils.ErrorResponse("failed to verify otp", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.SuccessResponse("successfully otp verified ", tokenRes)
	ctx.JSON(http.StatusOK, response)
}

func (c *authHandler) RefreshAccesstokenForUser(ctx *gin.Context) {

	var body utils.RefreshTokenRequest

	if err := ctx.ShouldBindJSON(&body); err != nil {
		response := utils.ErrorResponse("failed to bind inputs", err.Error(), body)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	accessToken, err := c.client.RefreshAccesstokenForUser(ctx, body.RefreshToken)
	if err != nil {
		response := utils.ErrorResponse("failed to refresh access token", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.SuccessResponse("successfully access_token generated", accessToken)
	ctx.JSON(http.StatusOK, response)
}
