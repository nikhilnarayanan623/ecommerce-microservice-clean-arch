package service

import (
	"context"
	"log"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/user-service/pkg/domain"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/user-service/pkg/pb"
	usecase "github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/user-service/pkg/usecase/interfaces"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServiceServer struct {
	usecase usecase.UserUsecase
	pb.UnimplementedUserServiceServer
}

func NewUserServiceServer(usecase usecase.UserUsecase) pb.UserServiceServer {
	return &UserServiceServer{
		usecase: usecase,
	}
}
func (c *UserServiceServer) SaveUser(ctx context.Context, req *pb.SaveUserRequest) (*pb.SaveUserResponse, error) {
	log.Println("save user called")
	userID, err := c.usecase.SaveUser(context.Background(), domain.User{
		FirstName: req.GetFirstName(),
		LastName:  req.GetLastName(),
		Age:       req.GetAge(),
		UserName:  req.GetUserName(),
		Phone:     req.GetPhone(),
		Email:     req.GetEmail(),
		Password:  req.GetPassword(),
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}

	return &pb.SaveUserResponse{UserId: userID}, nil
}
func (c *UserServiceServer) FindUserByEmail(ctx context.Context, req *pb.FindUserByEmailRequest) (*pb.FindUserByEmailResponse, error) {
	log.Println("find user by email called")
	user, err := c.usecase.FindUserByEmail(ctx, req.GetEmail())
	if err != nil {
		log.Println("failed to find user by email")
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}
	log.Println("successfully find user by email")
	return &pb.FindUserByEmailResponse{
		UserId:    user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Age:       user.Age,
		UserName:  user.UserName,
		Phone:     user.Phone,
		Email:     user.Email,
		Password:  user.Password,
	}, nil
}
