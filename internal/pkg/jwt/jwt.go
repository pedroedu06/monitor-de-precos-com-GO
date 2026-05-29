package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GetSecretKey() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		panic("chave nao encontrada!")
	}

	return []byte(secret)
}

type Claims struct {
	UserID string `json:"userId"`
	jwt.RegisteredClaims
}

func GenerateToken(userID string) (string, error) {
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString(GetSecretKey())
}

