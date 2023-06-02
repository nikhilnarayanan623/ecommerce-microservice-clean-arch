package service

import (
	"context"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/auth-service/pkg/domain"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/auth-service/pkg/pb"
	usecase "github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/auth-service/pkg/usecase/interfaces"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/auth-service/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type authServiceServer struct {
	usecase usecase.AuthUseCase
	pb.UnimplementedAuthServiceServer
}

func NewAuthServiceServer(usecase usecase.AuthUseCase) pb.AuthServiceServer {
	return &authServiceServer{
		usecase: usecase,
	}
}

func (c *authServiceServer) UserSignup(ctx context.Context, req *pb.UserSignupRequest) (*pb.UserSignupResponse, error) {

	otpID, err := c.usecase.UserSignup(ctx, domain.SaveUserRequest{
		FirstName: req.GetFirstName(),
		LastName:  req.GetLastName(),
		Age:       req.GetAge(),
		Email:     req.GetEmail(),
		Phone:     req.GetPhone(),
		Password:  req.GetPassord(),
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}

	return &pb.UserSignupResponse{OtpId: otpID}, nil
}
func (c *authServiceServer) UserSignupVerify(ctx context.Context, req *pb.UserSignupVerifyRequest) (*pb.UserSignupVerifyResponse, error) {

	usreID, err := c.usecase.OtpVerify(ctx, utils.OtpVerify{
		OtpID:   req.GetOtpId(),
		OtpCode: req.GetOtpCode(),
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}

	accessToken, err := c.usecase.GenerateAccessToken(ctx, usreID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}

	refreshToken, err := c.usecase.GenereateRefreshToken(ctx, usreID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}

	return &pb.UserSignupVerifyResponse{
		AccesToken:   accessToken,
		RefreshToken: refreshToken,
	}, nil
}
