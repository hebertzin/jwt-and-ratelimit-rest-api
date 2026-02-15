package security

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("SECRET_JWT"))

type Jwt struct {
	expiration int64
	service    string
}

func NewTokenService(exp int64, service string) *Jwt {
	return &Jwt{
		expiration: exp,
		service:    service,
	}
}

func (t *Jwt) GenerateToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email":   email,
			"exp":     t.expiration,
			"service": t.service,
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (t *Jwt) VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
