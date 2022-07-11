//go:generate mockgen -source=controller.go -destination=mocks/controller.go -package=mocks
package health

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	db database
}

type database interface {
	Ping(context.Context) error
}

func NewController(db database) *Controller {
	return &Controller{
		db,
	}
}

func (c *Controller) RegisterRoutes(r gin.IRouter) {
	r.GET("/healthz", c.handleHealth)
}

// @Summary      Health check
// @Description  checks if the app is healthy
// @Tags         health
// @Success      200  {object}  Response
// @Router       /healthz [get]
func (c *Controller) handleHealth(ctx *gin.Context) {
	err := c.db.Ping(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &Response{
			Status: StatusUnhealthy,
			Error:  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, &Response{
		Status: StatusHealthy,
	})
}
