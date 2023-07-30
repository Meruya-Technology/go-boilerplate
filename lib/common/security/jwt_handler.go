package security

import (
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
)

type JwtHandler struct{}

func (j JwtHandler) Generate(Secret string, ExpiredAt string) string {
	logger, _ := zap.NewProduction()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": ExpiredAt,
	})

	signatureKey := []byte(Secret)
	tokenString, err := token.SignedString(signatureKey)
	if err != nil {
		logger.Fatal(err.Error())
	}
	return tokenString
}
