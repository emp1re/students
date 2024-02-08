package handler

import (
	"context"

	"github.com/emp1re/students/internal/service"
	"github.com/emp1re/students/pkg/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SvcHandler interface {
	CreateStudent(ctx context.Context, input models.InputStudent) (out models.OutStudent, err error)
	GetStudents(ctx context.Context) (out []models.OutStudent, err error)
	UpdateStudent(ctx context.Context, input models.InputUpdateStudent, id string) error
	DeleteStudent(ctx context.Context, id string) error
}

type Handler struct {
	*zap.Logger
	SvcHandler SvcHandler
	srv        *service.Repository
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
