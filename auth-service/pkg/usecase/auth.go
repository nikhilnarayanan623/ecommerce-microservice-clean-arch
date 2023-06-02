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
		return "", fmt.Errorf("user alredy exist with given details and verified")
	}

	userID := existUser.ID
	// if user not exist then save the user
	if userID == 0 {
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
	err = c.repo.SaveOtpSession(ctx, domain.OTPSession{
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

	err = c.userClient.UpdateUserVerified(ctx, otpSession.UserID)
	if err != nil {
		return 0, err
	}

	return otpSession.UserID, nil
}

// Generate AccessToken
func (c *authUsecase) GenerateAccessToken(ctx context.Context, userID uint64) (accessToken string, err error) {

	accessToken, err = c.tokenAuth.GenerateToken(token.TokenRequest{
		TokenID:        "no_id",
		UserID:         userID,
		UsedFor:        token.User,
		ExpirationDate: time.Now().Add(time.Minute * 20),
	})
	return
}

// Generate RefreshToken
func (c *authUsecase) GenereateRefreshToken(ctx context.Context, userID uint64) (refreshToken string, err error) {

	tokenID := utils.GenerateUniqueRandomString()
	expireTime := time.Now().Add(time.Hour * 24 * 7)
	refreshToken, err = c.tokenAuth.GenerateToken(token.TokenRequest{
		TokenID:        tokenID,
		UserID:         userID,
		UsedFor:        token.User,
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
