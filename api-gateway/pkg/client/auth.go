package client

import (
	"context"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/client/interfaces"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/config"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/domain"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/pb"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/utils"
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
		Passord:   req.Password,
	})
	if err != nil {
		return otpID, err
	}
	return res.GetOtpId(), nil
}

func (c *authClient) UserSignupVerify(ctx context.Context, otpVerify utils.OtpVerify) (utils.TokenResponse, error) {

	res, err := c.client.UserSignupVerify(ctx, &pb.UserSignupVerifyRequest{
		OtpId:   otpVerify.OtpID,
		OtpCode: otpVerify.OtpCode,
	})
	if err == nil {
		return utils.TokenResponse{}, err
	}

	return utils.TokenResponse{
		AccessToken:  res.GetAccesToken(),
		RefreshToken: res.GetRefreshToken(),
	}, nil
}
