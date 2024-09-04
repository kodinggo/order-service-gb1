package utils

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"time"
)

func Dump(v interface{}) string {
	js, err := json.Marshal(v)
	if err != nil {
		return ""
	}

	return string(js)
}

func GenerateInvoice() string {
	prefix := "INV"

	currentTime := time.Now().Format("20060102150405")

	max := big.NewInt(1000)
	randomNumber, err := rand.Int(rand.Reader, max)
	if err != nil {
		return ""
	}

	invoiceNumber := fmt.Sprintf("%s-%s-%03d", prefix, currentTime, randomNumber)

	return invoiceNumber
}
