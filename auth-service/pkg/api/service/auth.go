package service

import (
	"context"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/auth-service/pkg/domain"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/auth-service/pkg/pb"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/auth-service/pkg/token"
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

// User Sginup
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

// User Signup Verify
func (c *authServiceServer) UserSignupVerify(ctx context.Context, req *pb.UserSignupVerifyRequest) (*pb.UserSignupVerifyResponse, error) {

	userID, err := c.usecase.OtpVerify(ctx, utils.OtpVerify{
		OtpID:   req.GetOtpId(),
		OtpCode: req.GetOtpCode(),
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}

	accessToken, err := c.usecase.GenerateAccessToken(ctx, userID, token.User)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}

	refreshToken, err := c.usecase.GenereateRefreshToken(ctx, userID, token.User)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}

	return &pb.UserSignupVerifyResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// Generate access token with refresh token
func (c *authServiceServer) RefreshAccessToken(ctx context.Context, req *pb.RefreshAccessTokenRequest) (*pb.RefreshAccessTokenResponse, error) {

	// check the token refreshing for user or admin
	tokenUser := token.User
	if req.UsedFor == pb.RefreshAccessTokenRequest_Admin {
		tokenUser = token.Admin
	}

	accessToken, err := c.usecase.RefreshAccessToken(ctx, req.GetRefreshToken(), tokenUser)
	if err != nil {
		return &pb.RefreshAccessTokenResponse{}, status.Errorf(codes.Internal, "%s", err.Error())
	}

	return &pb.RefreshAccessTokenResponse{
		AccessToken: accessToken,
	}, nil
}

func (c *authServiceServer) UserLogin(ctx context.Context, req *pb.UserLoginRequest) (*pb.UserLoginResponse, error) {

	userID, err := c.usecase.UserLogin(ctx, domain.UserLoginRequest{
		Email:    req.GetEmail(),
		Phone:    req.GetPhone(),
		Password: req.GetPassword(),
	})
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%s", err.Error())
	}

	accessToken, err := c.usecase.GenerateAccessToken(ctx, userID, token.User)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}

	refreshToken, err := c.usecase.GenereateRefreshToken(ctx, userID, token.User)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}

	return &pb.UserLoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
