package jwt

import (
	"os"
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