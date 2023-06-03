package response

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Error   interface{} `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func SuccessResponse(ctx *gin.Context, message string, data ...interface{}) {

	response := Response{
		Success: true,
		Message: message,
		Error:   nil,
		Data:    data,
	}
	ctx.JSON(http.StatusOK, response)
}

func ErrorResponse(ctx *gin.Context, message string, err error, data interface{}) {
	errFields := strings.Split(err.Error(), "\n")
	response := Response{
		Success: false,
		Message: message,
		Error:   errFields,
		Data:    data,
	}

	ctx.JSON(getErrorCode(err), response)
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func getErrorCode(err error) int {

	var httpCode int

	switch status.Code(err) {
	case codes.InvalidArgument:
		httpCode = http.StatusBadRequest
	case codes.AlreadyExists:
		httpCode = http.StatusConflict
	case codes.Internal:
		httpCode = http.StatusInternalServerError
	default:
		httpCode = http.StatusServiceUnavailable
	}
	return httpCode
}
