package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/auxcube/ektimo-api/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware(ctx *gin.Context) {
	tokenHeader := ctx.Request.Header.Get("Authorization")
	if tokenHeader == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	tokenStr := ""
	tokenSplit := strings.Split(tokenHeader, " ")
	if len(tokenSplit) == 2 && tokenSplit[0] == "Bearer" {
		tokenStr = tokenSplit[1]
	}
	if tokenStr == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	secret := config.Global.Auth.Secret
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		ctx.Set("user", claims["user"])
		ctx.Next()
	} else {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

}
