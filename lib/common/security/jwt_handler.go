package security

import (
	"time"

	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
)

type JwtHandler struct{}

func (j JwtHandler) Generate(secret string) string {
	logger, _ := zap.NewProduction()
	AddedTime := time.Now().AddDate(2, 0, 0)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": AddedTime,
	})

	signatureKey := []byte(secret)
	tokenString, err := token.SignedString(signatureKey)
	if err != nil {
		logger.Fatal(err.Error())
	}
	return tokenString
}
