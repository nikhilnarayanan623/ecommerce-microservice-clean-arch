package usecase

import (
	"context"
	"fmt"
	"time"

	client "github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/auth-service/pkg/client/interfaces"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/auth-service/pkg/domain"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/auth-service/pkg/otp"
	repo "github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/auth-service/pkg/repository/interfaces"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/auth-service/pkg/token"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/auth-service/pkg/usecase/interfaces"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/auth-service/pkg/utils"
)

type authUsecase struct {
	repo       repo.AuthRepository
	userClient client.UserClient
	otpVerify  otp.OtpVerification
	tokenAuth  token.TokenAuth
}

const (
	accessTokenDuration  = time.Minute * 20
	refreshTokenDuration = time.Hour * 24 * 7
)

func NewAuthUsecase(repo repo.AuthRepository, userClient client.UserClient,
	otpVerify otp.OtpVerification, tokenAuth token.TokenAuth) interfaces.AuthUseCase {
	return &authUsecase{
		repo:       repo,
		userClient: userClient,
		otpVerify:  otpVerify,
		tokenAuth:  tokenAuth,
	}
}

// Signup for user
func (c *authUsecase) UserSignup(ctx context.Context, user domain.SaveUserRequest) (otpID string, err error) {

	existUser, err := c.userClient.FindUserByEmail(ctx, user.Email)
	if err != nil {
		return "", fmt.Errorf("failed to check user already exist \nerror:%s", err.Error())
	}
	// if user exist and verified also then return error
	if existUser.ID != 0 && existUser.Verified {
		return "", fmt.Errorf("user already exist with given details and verified")
	}

	userID := existUser.ID
	// if user not exist then save the user
	if userID == 0 {

		user.Password, err = utils.GenerateHashFromPassword(user.Password)
		if err != nil {
			return "", fmt.Errorf("failed to hash user password")
		}

		userID, err = c.userClient.SaveUser(ctx, user)
		if err != nil {
			return "", fmt.Errorf("failed to save user \nerror%s", err.Error())
		}
	}

	_, err = c.otpVerify.SentOtp("+91" + user.Phone)
	if err != nil {
		return "", fmt.Errorf("failed to send otp to given number \nerror:%s", err.Error())
	}

	otpID = utils.GenerateUniqueRandomString()
	otpExpireTime := time.Now().Add(time.Minute * 5)
	err = c.repo.SaveOtpSession(ctx, domain.OtpSession{
		OtpID:    otpID,
		UserID:   userID,
		Phone:    user.Phone,
		ExpireAt: otpExpireTime,
	})

	if err != nil {
		return "", fmt.Errorf("failed to save otp session \nerror:%s", err.Error())
	}

	return otpID, nil
}

// Verify OTP using otp id and otp
func (c *authUsecase) OtpVerify(ctx context.Context, otpDetails utils.OtpVerify) (uint64, error) {

	otpSession, err := c.repo.FindOtpSession(ctx, otpDetails.OtpID)
	if err != nil {
		return 0, fmt.Errorf("failed to get otp session \nerror:%s", err.Error())
	} else if otpSession.OtpID == "" {
		return 0, fmt.Errorf("invalid otp id")
	}

	if time.Since(otpSession.ExpireAt) > 0 {
		return 0, fmt.Errorf("otp validation time expired")
	}

	err = c.repo.Transactions(func(repo repo.AuthRepository) error {

		err = c.userClient.UpdateUserVerified(ctx, otpSession.UserID)
		if err != nil {
			return err
		}

		err = repo.DeleteAllOtpSessionsByUserID(ctx, otpSession.UserID)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return 0, err
	}

	return otpSession.UserID, nil
}

// Generate AccessToken
func (c *authUsecase) GenerateAccessToken(ctx context.Context, userID uint64, tokenUser token.UserType) (accessToken string, err error) {

	accessToken, err = c.tokenAuth.GenerateToken(token.TokenRequest{
		TokenID:        "no_id",
		UserID:         userID,
		UsedFor:        token.User,
		ExpirationDate: time.Now().Add(accessTokenDuration),
	})
	return
}

// Generate RefreshToken
func (c *authUsecase) GenerateRefreshToken(ctx context.Context, userID uint64, tokenUser token.UserType) (refreshToken string, err error) {

	tokenID := utils.GenerateUniqueRandomString()
	expireTime := time.Now().Add(refreshTokenDuration)
	refreshToken, err = c.tokenAuth.GenerateToken(token.TokenRequest{
		TokenID:        tokenID,
		UserID:         userID,
		UsedFor:        tokenUser,
		ExpirationDate: expireTime,
	})
	if err != nil {
		return "", err
	}

	err = c.repo.SaveRefreshSession(ctx, domain.RefreshSession{
		TokenID:      tokenID,
		UserID:       userID,
		RefreshToken: refreshToken,
		ExpireAt:     expireTime,
	})

	if err != nil {
		return "", nil
	}

	return refreshToken, nil
}

// UserLogin
func (c *authUsecase) UserLogin(ctx context.Context, loginDetail domain.UserLoginRequest) (uint64, error) {

	var (
		user domain.User
		err  error
	)

	if loginDetail.Email != "" {
		user, err = c.userClient.FindUserByEmail(ctx, loginDetail.Email)
	} else {
		user, err = c.userClient.FindUserByPhone(ctx, loginDetail.Phone)
	}

	if err != nil {
		return 0, fmt.Errorf("failed to find user \nerror:%s", err.Error())
	}
	if user.ID == 0 {
		return 0, fmt.Errorf("user not exist with given details")
	}

	if !user.Verified {
		return 0, fmt.Errorf("user not verified")
	} else if user.BlockStatus {
		return 0, fmt.Errorf("user blocked by admin")
	}

	if !utils.VerifyHashAndPassword(user.Password, loginDetail.Password) {
		return 0, fmt.Errorf("password doesn't match with given password")
	}

	return user.ID, nil
}

// Refresh AccessToken
func (c *authUsecase) RefreshAccessToken(ctx context.Context, refreshToken string, tokenUser token.UserType) (accessToken string, err error) {

	tokenRes, err := c.tokenAuth.VerifyToken(token.TokenVerifyRequest{
		UsedFor:     tokenUser,
		TokenString: refreshToken,
	})
	if err != nil {
		return "", fmt.Errorf("failed to verify refresh token \nerror:%s", err.Error())
	}

	session, err := c.repo.FindRefreshSessionByTokenID(ctx, tokenRes.TokenID)
	if err != nil {
		return "", fmt.Errorf("failed to get refresh session \nerror:%s", err.Error())
	}

	if session.TokenID == "" {
		return "", fmt.Errorf("session not found for the refresh token")
	}

	if session.IsBlocked {
		return "", fmt.Errorf("refresh token blocked in session")
	}
	fmt.Println(session.ExpireAt)
	if time.Since(session.ExpireAt) > 0 {
		return "", fmt.Errorf("refresh token expired")
	}

	accessToken, err = c.tokenAuth.GenerateToken(token.TokenRequest{
		TokenID:        "no_id",
		UserID:         session.UserID,
		UsedFor:        tokenUser,
		ExpirationDate: time.Now().Add(accessTokenDuration),
	})

	return
}

func (c *authUsecase) VerifyAccessToken(ctx context.Context, accessToken string, usedFor token.UserType) (uint64, error) {

	tokenRes, err := c.tokenAuth.VerifyToken(token.TokenVerifyRequest{
		TokenString: accessToken,
		UsedFor:     usedFor,
	})
	if err != nil {
		return 0, fmt.Errorf("failed to verify token %w", err)
	}

	return tokenRes.UserID, nil
}
