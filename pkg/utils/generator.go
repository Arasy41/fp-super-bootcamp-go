package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateUid() (string, error) {
	b := make([]byte, 12)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
