package client

import (
	"context"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/client/interfaces"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/config"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/domain"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/pb"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/utils/request"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/utils/response"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type authClient struct {
	client pb.AuthServiceClient
}

func NewAuthClient(cfg *config.Config) (interfaces.AuthClient, error) {

	gcc, err := grpc.Dial(cfg.AuthServiceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := pb.NewAuthServiceClient(gcc)

	return &authClient{
		client: client,
	}, nil
}

func (c *authClient) UserSignup(ctx context.Context, req domain.UserSignupRequest) (otpID string, err error) {

	res, err := c.client.UserSignup(ctx, &pb.UserSignupRequest{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Age:       req.Age,
		Email:     req.Email,
		Phone:     req.Phone,
		Password:  req.Password,
	})
	if err != nil {
		return otpID, err
	}
	return res.GetOtpId(), nil
}

func (c *authClient) UserSignupVerify(ctx context.Context, otpVerify request.OtpVerify) (response.TokenResponse, error) {

	res, err := c.client.UserSignupVerify(ctx, &pb.UserSignupVerifyRequest{
		OtpId:   otpVerify.OtpID,
		OtpCode: otpVerify.OtpCode,
	})
	if err != nil {
		return response.TokenResponse{}, err
	}

	return response.TokenResponse{
		AccessToken:  res.GetAccessToken(),
		RefreshToken: res.GetRefreshToken(),
	}, nil
}

// User Login
func (c *authClient) UserLogin(ctx context.Context, loginDetails domain.UserLoginRequest) (response.TokenResponse, error) {

	res, err := c.client.UserLogin(ctx, &pb.UserLoginRequest{
		Email:    loginDetails.Email,
		Phone:    loginDetails.Phone,
		Password: loginDetails.Password,
	})
	if err != nil {
		return response.TokenResponse{}, err
	}
	return response.TokenResponse{
		AccessToken:  res.GetAccessToken(),
		RefreshToken: res.GetRefreshToken(),
	}, nil
}

func (c *authClient) VerifyUserAccessToken(ctx context.Context, accessToken string) (userID uint64, err error) {

	res, err := c.client.VerifyUserAccessToken(ctx, &pb.VerifyUserAccessTokenRequest{
		AccessToken: accessToken,
	})

	if err != nil {
		return 0, err
	}

	return res.GetUserId(), nil
}

// Refresh access token for user using refresh token
func (c *authClient) RefreshAccessTokenForUser(ctx context.Context, refreshToken string) (accessToken string, err error) {

	res, err := c.client.RefreshAccessToken(ctx, &pb.RefreshAccessTokenRequest{
		UsedFor:      pb.RefreshAccessTokenRequest_User,
		RefreshToken: refreshToken,
	})
	if err != nil {
		return "", err
	}
	return res.GetAccessToken(), nil
}
