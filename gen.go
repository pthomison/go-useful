package usefulgo

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateRandomB64(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	state := base64.StdEncoding.EncodeToString(b)

	return state, nil
}
