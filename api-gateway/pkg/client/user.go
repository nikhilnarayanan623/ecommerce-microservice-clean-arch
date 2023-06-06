package client

import (
	"context"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/client/interfaces"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/config"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/pb"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/utils/response"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type userClient struct {
	client pb.UserServiceClient
}

func NewUserClient(cfg *config.Config) (interfaces.UserClient, error) {
	gcc, err := grpc.Dial(cfg.UserServiceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := pb.NewUserServiceClient(gcc)

	return &userClient{
		client: client,
	}, nil
}

func (c *userClient) GetUserProfile(ctx context.Context, userID uint64) (response.User, error) {

	res, err := c.client.GetUserProfile(ctx, &pb.GetUserProfileRequest{UserId: userID})
	if err != nil {
		return response.User{}, err
	}

	return response.User{
		FirstName: res.GetFirstName(),
		LastName:  res.GetLastName(),
		Age:       res.GetAge(),
		Email:     res.GetEmail(),
		Phone:     res.GetPhone(),
	}, nil
}
