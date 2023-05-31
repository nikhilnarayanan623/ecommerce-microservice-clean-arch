package client

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/domain"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/mock"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/pb"
	"github.com/stretchr/testify/assert"	
)

func TestUserSignup(t *testing.T) {

	testCases := map[string]struct {
		input          domain.UserSignupRequest
		buildStub      func(mClient *mock.MockAuthServiceClient)
		expectedOutput uint64
		expectedError  error
	}{
		"ServerErrorShouldReturnAsError": {
			input: domain.UserSignupRequest{
				FirstName: "Nikhil",
				LastName:  "N",
			},
			buildStub: func(mClient *mock.MockAuthServiceClient) {
				mClient.EXPECT().UserSignup(gomock.Any(), &pb.UserSignupRequest{
					FirstName: "Nikhil",
					LastName:  "N",
				}).Times(1).Return(&pb.UserSignupResponse{}, errors.New("failed by server"))
			},
			expectedOutput: 0,
			expectedError:  errors.New("failed by server"),
		},
		"SucessfulOnServerSideShouldReturnUserID": {
			input: domain.UserSignupRequest{
				FirstName: "Nikhil",
				LastName:  "N",
			},
			buildStub: func(mClient *mock.MockAuthServiceClient) {
				mClient.EXPECT().UserSignup(gomock.Any(), &pb.UserSignupRequest{
					FirstName: "Nikhil",
					LastName:  "N",
				}).Times(1).Return(&pb.UserSignupResponse{UserId: 8}, nil)
			},
			expectedOutput: 8,
			expectedError:  nil,
		},
	}

	for testName, test := range testCases {
		test := test
		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			ctl := gomock.NewController(t)
			defer ctl.Finish()

			client := mock.NewMockAuthServiceClient(ctl)
			test.buildStub(client)

			authClient := authClient{
				client: client,
			}

			output, err := authClient.UserSignup(context.Background(), test.input)
			assert.Equal(t, test.expectedError, err)
			assert.EqualValues(t, test.expectedOutput, output)

		})
	}

}
