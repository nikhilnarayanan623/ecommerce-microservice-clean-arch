package service

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/user-service/pkg/domain"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/user-service/pkg/mock"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/user-service/pkg/pb"
	"github.com/stretchr/testify/assert"
)

func TestSaveUser(t *testing.T) {

	testCases := map[string]struct {
		input          *pb.SaveUserRequest
		buildStub      func(mUsecase *mock.MockUserUsecase)
		expectedOutput *pb.SaveUserResponse
		expectedError  error
	}{
		"UsecaseErrorShouldReturnError": {
			input: &pb.SaveUserRequest{
				FirstName: "Nikhil",
				LastName:  "N",
				Email:     "email@gmail.com",
			},
			buildStub: func(mUsecase *mock.MockUserUsecase) {
				mUsecase.EXPECT().SaveUser(gomock.Any(), domain.User{
					FirstName: "Nikhil",
					LastName:  "N",
					Email:     "email@gmail.com",
				}).Times(1).Return(uint64(0), errors.New("failed to save user from usecase"))
			},
			expectedOutput: nil,
			expectedError:  errors.New("failed to save user from usecase"),
		},
		"SuccessfulSaveShouldReturnUserID": {
			input: &pb.SaveUserRequest{
				FirstName: "Nikhil",
				LastName:  "N",
				Email:     "email@gmail.com",
			},
			buildStub: func(mUsecase *mock.MockUserUsecase) {
				mUsecase.EXPECT().SaveUser(gomock.Any(), domain.User{
					FirstName: "Nikhil",
					LastName:  "N",
					Email:     "email@gmail.com",
				}).Times(1).Return(uint64(6), nil)
			},
			expectedOutput: &pb.SaveUserResponse{UserId: 6},
			expectedError:  nil,
		},
	}

	for testName, test := range testCases {

		test := test

		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			ctl := gomock.NewController(t)
			usecase := mock.NewMockUserUsecase(ctl)

			test.buildStub(usecase)
			service := NewUserServiceServer(usecase)

			output, err := service.SaveUser(context.Background(), test.input)
			if test.expectedError == nil {
				assert.Nil(t, err)
			} else {
				assert.Contains(t, err.Error(), test.expectedError.Error())
			}
			assert.Equal(t, test.expectedOutput, output)

		})
	}

}
