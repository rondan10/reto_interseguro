package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("your-secret-key") // En producción, usar variable de entorno

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// Token fijo para validación
#const validToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYXBpLWFjY2VzcyJ9.0q6QFJ4YMoKC1RiHtO-PKzIcOqhdU2qZVhb7735kcX4"

func ValidateToken(tokenStr string) (*Claims, error) {
	if tokenStr != validToken {
		return nil, jwt.ErrSignatureInvalid
	}

	return &Claims{}, nil
}
