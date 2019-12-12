package service

import (
	"encoding/json"

	"github.com/dgrijalva/jwt-go"
)

var (
	secretKey = "58oEIPpcj1quAqjuXx27NPBqWHfbkkRyKjWNQVuS9nUDkG1LwbDEbGcmuxACL5vYH5wsAREqHeTeGVjLiORie2mkndLP2G5RD9Eu9rTGh27UBwkJJxSAPaZ1g0O5sFd4cJbgcjoq5U9FXsjeDN5O7nCMX5weE9rZdqn-uqhHIaM"
)

// TokenGenerate generate token
func TokenGenerate(claim map[string]interface{}) (string, error) {

	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set token claims
	tokeClaims := token.Claims.(jwt.MapClaims)
	buff, err := json.Marshal(claim)
	if err != nil {
		return "", err
	}
	if err := json.Unmarshal(buff, &tokeClaims); err != nil {
		return "", err
	}

	// Sign the token with secret key
	tokenString, err := token.SignedString([]byte(secretKey))
	return tokenString, err
}
