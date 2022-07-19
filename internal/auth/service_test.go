package auth

import (
	"context"
	"testing"

	"github.com/auxcube/ektimo-api/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestService_Login_generatesToken(t *testing.T) {
	config.Initialize("test", "../../config")
	testCreds := CredentialsDto{
		Username: "user",
		Password: "pass1",
	}
	service := NewService()
	token, err := service.Login(context.Background(), testCreds)
	assert.Nil(t, err)
	assert.NotNil(t, token)
}
