package handler

import (
	"context"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/api/handler/interfaces"
	client "github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/client/interfaces"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/domain"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/utils/request"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/utils/response"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		response.ErrorResponse(ctx, "failed bind inputs", err, body)
		return
	}

	tokenRes, err := c.client.UserLogin(ctx, body)

	if err != nil {
		response.ErrorResponse(ctx, "failed to login", err, nil)
		return
	}

	response.SuccessResponse(ctx, "successfully login completed", tokenRes)
}

func (c *authHandler) UserSignup(ctx *gin.Context) {

	var body domain.UserSignupRequest

	if err := ctx.ShouldBindJSON(&body); err != nil {
		response.ErrorResponse(ctx, "failed bind inputs", err, body)
		return
	}

	userID, err := c.client.UserSignup(context.Background(), body)

	if err != nil {
		response.ErrorResponse(ctx, "failed to signup user", err, nil)
		return
	}

	response.SuccessResponse(ctx, "successfully otp send to user registered number", userID)
}

func (c *authHandler) UserSignupVerify(ctx *gin.Context) {

	var body request.OtpVerify
	if err := ctx.ShouldBindJSON(&body); err != nil {
		response.ErrorResponse(ctx, "failed bind inputs", err, body)
		return
	}

	tokenRes, err := c.client.UserSignupVerify(ctx, request.OtpVerify{
		OtpID:   body.OtpID,
		OtpCode: body.OtpCode,
	})

	if err != nil {
		response.ErrorResponse(ctx, "failed to verify otp", err, nil)
		return
	}

	response.SuccessResponse(ctx, "successfully otp verified ", tokenRes)
}

func (c *authHandler) AuthenticateUser(ctx *gin.Context) {

	authHeader := ctx.GetHeader("authorization")

	authFields := strings.Fields(authHeader)
	if len(authFields) < 2 {
		response.ErrorResponse(ctx, "unauthorized user", status.Error(codes.Unauthenticated, "token not found on header"), nil)
		ctx.Abort()
		return
	}
	accessToken := authFields[1]
	userID, err := c.client.VerifyUserAccessToken(ctx, accessToken)
	if err != nil {
		response.ErrorResponse(ctx, "failed authenticate user", err, nil)
		ctx.Abort()
		return
	}

	ctx.Set("userId", userID)
}

func (c *authHandler) RefreshAccessTokenForUser(ctx *gin.Context) {

	var body request.RefreshTokenRequest

	if err := ctx.ShouldBindJSON(&body); err != nil {
		response.ErrorResponse(ctx, "failed bind inputs", err, body)
		return
	}

	accessToken, err := c.client.RefreshAccessTokenForUser(ctx, body.RefreshToken)
	if err != nil {
		response.ErrorResponse(ctx, "failed to refresh access token", err, nil)
		return
	}

	response.SuccessResponse(ctx, "successfully access_token generated", accessToken)
}
