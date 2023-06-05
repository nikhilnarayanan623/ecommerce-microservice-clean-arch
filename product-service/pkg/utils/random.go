package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateSKU() string {

	sku := make([]byte, 10)

	rand.Read(sku)

	return hex.EncodeToString(sku)
}
