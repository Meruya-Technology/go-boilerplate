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
		"eat": AddedTime,
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		logger.Fatal(err.Error())
	}
	return tokenString
}
