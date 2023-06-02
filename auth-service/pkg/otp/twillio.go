package otp

import (
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/auth-service/pkg/config"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/verify/v2"
)

type twilioOtp struct {
	serviceID  string
	authToken  string
	accountSID string
}

func NewTwiloOtpAuth(cfg *config.Config) OtpVerification {
	return &twilioOtp{
		serviceID:  cfg.TwilioServiceID,
		authToken:  cfg.TwilioAuthToken,
		accountSID: cfg.TwilioAccountSID,
	}
}

func (c *twilioOtp) SentOtp(phoneNumber string) (string, error) {

	client := c.getNewTwiloClient()
	params := &twilioApi.CreateVerificationParams{}
	params.SetTo(phoneNumber)
	params.SetChannel("sms")

	resp, err := client.VerifyV2.CreateVerification(c.serviceID, params)
	if err != nil {
		return "", err
	}

	return *resp.Sid, nil
}

func (c *twilioOtp) VerifyOtp(phoneNumber string, code string) error {

	client := c.getNewTwiloClient()

	params := &twilioApi.CreateVerificationCheckParams{}
	params.SetTo(phoneNumber)
	params.SetCode(code)

	resp, err := client.VerifyV2.CreateVerificationCheck(c.serviceID, params)
	if err != nil {
		return err
	} else if *resp.Status == "approved" {
		return nil
	}

	return nil
}

func (c *twilioOtp) getNewTwiloClient() twilio.RestClient {

	return *twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: c.accountSID,
		Password: c.authToken,
	})
}
