package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/user-service/pkg/domain"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/user-service/pkg/pb"
	usecase "github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/user-service/pkg/usecase/interfaces"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/user-service/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
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
	utils.LogMessage(utils.Cyan, "SaveUser Invoked")
	userID, err := c.usecase.SaveUser(context.Background(), domain.User{
		FirstName: req.GetFirstName(),
		LastName:  req.GetLastName(),
		Age:       req.GetAge(),
		Phone:     req.GetPhone(),
		Email:     req.GetEmail(),
		Password:  req.GetPassword(),
	})

	if err != nil {
		utils.LogMessage(utils.Red, err.Error())
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}
	utils.LogMessage(utils.Green, "Successfully User Details Saved")
	return &pb.SaveUserResponse{UserId: userID}, nil
}

func (c *UserServiceServer) FindUserByEmail(ctx context.Context, req *pb.FindUserByEmailRequest) (*pb.FindUserByEmailResponse, error) {
	utils.LogMessage(utils.Cyan, "FindUserByEmail Invoked")
	user, err := c.usecase.FindUserByEmail(ctx, req.GetEmail())
	if err != nil {
		utils.LogMessage(utils.Red, err.Error())
		return nil, status.Errorf(codes.InvalidArgument, "%s", err.Error())
	}
	utils.LogMessage(utils.Green, "Successfully Found User By Email")
	return &pb.FindUserByEmailResponse{
		UserId:      user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Age:         user.Age,
		Phone:       user.Phone,
		Email:       user.Email,
		Password:    user.Password,
		Verified:    user.Verified,
		BlockStatus: user.BlockStatus,
	}, nil
}

func (c *UserServiceServer) FindUserByPhone(ctx context.Context, req *pb.FindUserByPhoneRequest) (*pb.FindUserByPhoneResponse, error) {
	utils.LogMessage(utils.Cyan, "FindUserByPhone Invoked")

	user, err := c.usecase.FindUserByPhone(ctx, req.GetPhone())
	if err != nil {
		utils.LogMessage(utils.Red, err.Error())
		return nil, status.Errorf(codes.InvalidArgument, "%s", err.Error())
	}
	utils.LogMessage(utils.Green, "Successfully Found User By Email")
	return &pb.FindUserByPhoneResponse{
		UserId:      user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Age:         user.Age,
		Phone:       user.Phone,
		Email:       user.Email,
		Password:    user.Password,
		Verified:    user.Verified,
		BlockStatus: user.BlockStatus,
	}, nil
}

func (c *UserServiceServer) UpdateUserVerified(ctx context.Context, req *pb.UpdateUserVerifyRequest) (*empty.Empty, error) {
	utils.LogMessage(utils.Cyan, "UpdateUserVerified Invoked")
	err := c.usecase.UpdateUserVerified(ctx, req.GetUserId())
	if err != nil {
		utils.LogMessage(utils.Red, err.Error())
		return &emptypb.Empty{}, status.Errorf(codes.Internal, "%s", err.Error())
	}
	utils.LogMessage(utils.Green, "Successfully User Verified")
	return &emptypb.Empty{}, nil
}

func (c *UserServiceServer) GetUserProfile(ctx context.Context, req *pb.GetUserProfileRequest) (*pb.GetUserProfileResponse, error) {
	utils.LogMessage(utils.Cyan, "GetUserProfile Invoked")

	user, err := c.usecase.FindUserByID(ctx, req.GetUserId())
	if err != nil {
		utils.LogMessage(utils.Red, err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	utils.LogMessage(utils.Green, "Successfully found user details")
	return &pb.GetUserProfileResponse{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Age:       user.Age,
		Email:     user.Email,
		Phone:     user.Phone,
	}, nil
}
