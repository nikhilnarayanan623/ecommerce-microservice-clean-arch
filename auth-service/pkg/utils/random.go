package utils

import "github.com/google/uuid"

func GenerateUniqueRandomString() string {
	
	return uuid.NewString()
}
