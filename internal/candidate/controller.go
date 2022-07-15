package candidate

import (
	"context"
	"net/http"

	"github.com/auxcube/ektimo-api/ent"
	"github.com/gin-gonic/gin"
)

type candidateFetcher interface {
	List(context.Context) ([]*ent.Candidate, error)
}

type candidateService interface {
	candidateFetcher
}

type Controller struct {
	service candidateService
}

func NewController(service candidateService) *Controller {
	return &Controller{service: service}
}

// RegisterRoutes registers candidate related http routes
func (c *Controller) RegisterRoutes(r gin.IRouter) {
	r.Group("candidates").
		GET("/", c.handleList)
}

// @Summary      List candidates
// @Description  Lists all candidates in the app
// @Tags         candidates
// @Success      200 {array} ent.Candidate
// @Router       /candidates [get]
func (c *Controller) handleList(ctx *gin.Context) {
	candidates, err := c.service.List(ctx.Request.Context())

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, candidates)
}
