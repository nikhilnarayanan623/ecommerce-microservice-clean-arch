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

	signupRequest := domain.UserSignupRequest{
		FirstName: req.GetFirstName(),
		LastName:  req.GetLastName(),
		Age:       req.GetAge(),
		Email:     req.GetEmail(),
		Phone:     req.GetPhone(),
		Password:  req.GetPassword(),
	}

	otpID, err := c.usecase.UserSignup(ctx, signupRequest)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}

	return &pb.UserSignupResponse{OtpId: otpID}, nil
}

// User Signup Verify
func (c *authServiceServer) UserSignupVerify(ctx context.Context, req *pb.UserSignupVerifyRequest) (*pb.UserSignupVerifyResponse, error) {

	optRequest := utils.OtpVerify{
		OtpID:   req.GetOtpId(),
		OtpCode: req.GetOtpCode(),
	}

	userID, err := c.usecase.OtpVerify(ctx, optRequest)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}

	// two token(string) channel and one error channel of size 2(two go routines)
	accessTokenChan := make(chan string)
	refreshTokenChan := make(chan string)
	errChan := make(chan error, 2)

	go c.getAccessToken(ctx, userID, token.User, accessTokenChan, errChan)
	go c.getRefreshToken(ctx, userID, token.User, refreshTokenChan, errChan)

	var accessToken, refreshToken string
	// wait for two channel input one is access token and refresh token with error or not
	for i := 1; i <= 2; i++ {
		select {
		case token := <-accessTokenChan:
			accessToken = token
		case token := <-refreshTokenChan:
			refreshToken = token
		case err := <-errChan:
			return nil, status.Errorf(codes.Internal, "%s", err.Error())
		}
	}

	return &pb.UserSignupVerifyResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// for concurrently getting access token
func (c *authServiceServer) getAccessToken(ctx context.Context, userID uint64, tokenUser token.UserType, ch chan string, errChan chan error) {
	accessToken, err := c.usecase.GenerateAccessToken(ctx, userID, tokenUser)
	if err != nil {
		errChan <- err
	}
	ch <- accessToken
}

// for concurrently getting refresh token
func (c *authServiceServer) getRefreshToken(ctx context.Context, userID uint64, tokenUser token.UserType, ch chan string, errChan chan error) {
	refreshToken, err := c.usecase.GenerateRefreshToken(ctx, userID, tokenUser)
	if err != nil {
		errChan <- err
	}
	ch <- refreshToken
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

	loginRequest := domain.UserLoginRequest{
		Email:    req.GetEmail(),
		Phone:    req.GetPhone(),
		Password: req.GetPassword(),
	}

	userID, err := c.usecase.UserLogin(ctx, loginRequest)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%s", err.Error())
	}

	// two token(string) channel and one error channel of size 2(two go routines)
	accessTokenChan := make(chan string)
	refreshTokenChan := make(chan string)
	errChan := make(chan error, 2)

	go c.getAccessToken(ctx, userID, token.User, accessTokenChan, errChan)
	go c.getRefreshToken(ctx, userID, token.User, refreshTokenChan, errChan)

	var accessToken, refreshToken string
	// wait for two channel input one is access token and refresh token with error or not
	for i := 1; i <= 2; i++ {
		select {
		case token := <-accessTokenChan:
			accessToken = token
		case token := <-refreshTokenChan:
			refreshToken = token
		case err := <-errChan:
			return nil, status.Errorf(codes.Internal, "%s", err.Error())
		}
	}

	return &pb.UserLoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (c *authServiceServer) VerifyUserAccessToken(ctx context.Context, req *pb.VerifyUserAccessTokenRequest) (*pb.VerifyUserAccessTokenResponse, error) {

	userID, err := c.usecase.VerifyAccessToken(ctx, req.GetAccessToken(), token.User)
	if err != nil {
		return &pb.VerifyUserAccessTokenResponse{}, status.Error(codes.Unauthenticated, err.Error())
	}

	return &pb.VerifyUserAccessTokenResponse{UserId: userID}, nil
}
