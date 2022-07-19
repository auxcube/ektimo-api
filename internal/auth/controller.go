package auth

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service AuthService
}

type AuthService interface {
	Login(ctx context.Context, creds CredentialsDto) (string, error)
}

func NewController(service AuthService) *Controller {
	return &Controller{service: service}
}

func (c *Controller) RegisterRoutes(r gin.IRouter) {
	r.POST("/login", c.handleLogin)
}

func (c *Controller) handleLogin(ctx *gin.Context) {
	var creds CredentialsDto
	err := ctx.Bind(&creds)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	token, err := c.service.Login(ctx.Request.Context(), creds)
	if errors.Is(err, ErrIncorrectCreds) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	resp := AuthTokenDto{Token: token}
	ctx.JSON(http.StatusOK, resp)
}
