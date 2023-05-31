package utils

import (
	"strings"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Error   interface{} `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func SuccessResponse(message string, data ...interface{}) Response {

	return Response{
		Success: true,
		Message: message,
		Error:   nil,
		Data:    data,
	}
}

func ErrorResponse(message, errString string, data interface{}) Response {
	errFields := strings.Split(errString, "\n")
	return Response{
		Success: false,
		Message: message,
		Error:   errFields,
		Data:    data,
	}
}
