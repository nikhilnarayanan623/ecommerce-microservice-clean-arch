package client

import (
	"context"
	"fmt"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/auth-service/pkg/client/interfaces"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/auth-service/pkg/config"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/auth-service/pkg/domain"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/auth-service/pkg/pb"
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

func (c *userClient) FindUserByEmail(ctx context.Context, email string) (domain.User, error) {

	res, err := c.client.FindUserByEmail(ctx, &pb.FindUserByEmailRequest{
		Email: email,
	})

	if err != nil {
		return domain.User{}, err
	}

	return domain.User{
		ID:          res.GetUserId(),
		FirstName:   res.GetFirstName(),
		LastName:    res.GetLastName(),
		Age:         res.GetAge(),
		Email:       res.GetEmail(),
		Phone:       res.GetPhone(),
		Password:    res.GetPassword(),
		Verified:    res.GetVerified(),
		BlockStatus: res.GetBlockStatus(),
	}, nil
}

func (c *userClient) FindUserByPhone(ctx context.Context, phone string) (domain.User, error) {

	res, err := c.client.FindUserByPhone(ctx, &pb.FindUserByPhoneRequest{
		Phone: phone,
	})
	if err != nil {
		return domain.User{}, err
	}
	
	return domain.User{
		ID:          res.GetUserId(),
		FirstName:   res.GetFirstName(),
		LastName:    res.GetLastName(),
		Age:         res.GetAge(),
		Email:       res.GetEmail(),
		Phone:       res.GetPhone(),
		Password:    res.GetPassword(),
		Verified:    res.GetVerified(),
		BlockStatus: res.GetBlockStatus(),
	}, nil
}

func (c *userClient) SaveUser(ctx context.Context, user domain.SaveUserRequest) (userID uint64, err error) {
	res, err := c.client.SaveUser(ctx, &pb.SaveUserRequest{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Age:       user.Age,
		Email:     user.Email,
		Phone:     user.Phone,
		Password:  user.Password,
	})
	if err != nil {
		return 0, err
	}
	fmt.Println("user client res", res)
	return res.GetUserId(), nil
}

func (c *userClient) UpdateUserVerified(ctx context.Context, userID uint64) error {

	_, err := c.client.UpdateUserVerified(ctx, &pb.UpdateUserVerifyRequest{UserId: userID})

	return err
}
