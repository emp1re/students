package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SvcHandler interface {
	// AddStudent(c *gin.Context, request *models.Student) (errCode int, err error)
}

type Handler struct {
	*zap.Logger
	SvcHandler SvcHandler
}

func NewHandler(SvcHandler SvcHandler, logger zap.Logger) *Handler {
	return &Handler{SvcHandler: SvcHandler, Logger: &logger}
}

func MakeRouter(h *Handler) *gin.Engine {
	router := gin.Default()
	router.Use(gin.Recovery())
	v1 := router.Group("api/v1")
	{
		v1.GET("/students", h.GetStudent)
		v1.POST("/students", h.AddStudent)
		v1.PATCH("/students/:id", h.UpdateStudent)
		v1.DELETE("/students/:id", h.DeleteStudent)

	}
	return router
}
