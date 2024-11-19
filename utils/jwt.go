package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"customer-orders/config"
)

// JWT Secret loaded from the .env file
var jwtKey = []byte(config.GetEnv("JWT_SECRET"))

// Claims structure for JWT
type Claims struct {
	ID    uint   `json:"id"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

// GenerateJWT creates a new JWT token
func GenerateJWT(id uint, role string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		ID:   id,
		Role: role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "go-oauth-jwt-app",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// ValidateJWT validates the token and returns the claims
func ValidateJWT(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
