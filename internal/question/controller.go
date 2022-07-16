package question

import (
	"context"
	"net/http"

	"github.com/auxcube/ektimo-api/pkg/httputil"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	service questionService
}

type questionService interface {
	FindAll(ctx context.Context) (Questions, error)
}

func NewController(service questionService) *Controller {
	return &Controller{
		service,
	}
}

func (c *Controller) RegisterRoutes(r gin.IRouter) {
	r.GET("/question", c.handleFindAll)
	r.GET("/question/:id", c.handleFindOne)
	r.POST("/question/text", c.handleCreateTextQuestion)
	r.POST("/question/mcq", c.handleCreateMCQ)
	r.POST("/question/coding", c.handleCreateCodingQuestion)
}

func (c *Controller) handleFindAll(ctx *gin.Context) {
	questions, err := c.service.FindAll(ctx.Request.Context())
	if err != nil {
		httputil.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, questions)
}

func (c *Controller) handleFindOne(ctx *gin.Context) {
	panic("unimplemented")
}

func (c *Controller) handleCreateTextQuestion(ctx *gin.Context) {
	panic("unimplemented")
}

func (c *Controller) handleCreateMCQ(ctx *gin.Context) {
	panic("unimplemented")
}

func (c *Controller) handleCreateCodingQuestion(ctx *gin.Context) {
	panic("unimplemented")
}
