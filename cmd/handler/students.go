package handler

import (
	"fmt"

	"github.com/emp1re/students/pkg/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (h *Handler) AddStudent(c *gin.Context) {
	logger := h.Logger
	bytes, err := c.GetRawData()
	if err != nil {
		logger.Error("c.GetRawData", zap.Error(err))
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	student := &models.Student{}

	err = student.Decode(bytes)
	if err != nil {
		logger.Error("student.Decode", zap.Error(err))
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	validation := student.ValidateStudent(*student)
	if validation != nil {
		logger.Error("student.ValidateStudent", zap.Error(validation))
		c.JSON(400, gin.H{"error": validation.Error()})
		return
	}
	fmt.Println(student.Addresses)

}
func (h *Handler) GetStudent(c *gin.Context) {

}
func (h *Handler) UpdateStudent(c *gin.Context) {
}
func (h *Handler) DeleteStudent(c *gin.Context) {
}
