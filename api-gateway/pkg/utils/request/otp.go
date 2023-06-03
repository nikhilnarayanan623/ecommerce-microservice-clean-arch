package request

type OtpVerify struct {
	OtpID   string `json:"otp_id" binding:"required"`
	OtpCode string `json:"otp_code" binding:"required"`
}
