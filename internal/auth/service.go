package auth

import (
	"context"
	"errors"

	"github.com/auxcube/ektimo-api/pkg/config"
	"github.com/golang-jwt/jwt"
)

var (
	ErrNoSecretDefined = errors.New("no secret defined")
	ErrGeneratingToken = errors.New("error generating auth token")
	ErrIncorrectCreds  = errors.New("incorrect username or password")
)

type Service struct {
}

func NewService() Service {
	return Service{}
}

func (s *Service) Login(ctx context.Context, creds CredentialsDto) (string, error) {
	if config.Global.Auth.AdminUsername == creds.Username && config.Global.Auth.AdminPassword == creds.Password {
		tokenString, err := generateToken(creds.Username)
		if err != nil {
			return "", ErrGeneratingToken
		}
		return tokenString, nil
	}
	return "", ErrIncorrectCreds
}

func generateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": username,
	})
	if config.Global.Auth.Secret == "" {
		return "", ErrNoSecretDefined
	}
	secret := []byte(config.Global.Auth.Secret)
	return token.SignedString(secret)
}
