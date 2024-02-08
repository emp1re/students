package handler

import (
	"fmt"
	"net/http"

	"github.com/emp1re/students/pkg/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (h *Handler) AddStudent(c *gin.Context) {

	bytes, err := c.GetRawData()
	if err != nil {
		h.Logger.Error("c.GetRawData", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	student := &models.InputStudent{}

	err = student.Decode(bytes)
	if err != nil {
		h.Logger.Error("student.Decode", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	validation := student.ValidateStudent(*student)
	if validation != nil {
		h.Logger.Error("student.ValidateStudent", zap.Error(validation))
		c.JSON(http.StatusInternalServerError, gin.H{"error": validation.Error()})
		return
	}
	stud, err := h.SvcHandler.CreateStudent(c, *student)
	if err != nil {
		h.Logger.Error("h.SvcHandler.CreateStudent", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, stud)

}
func (h *Handler) GetStudent(c *gin.Context) {

	stud, err := h.SvcHandler.GetStudents(c)
	if err != nil {
		h.Logger.Error("h.SvcHandler.GetStudent", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": zap.Error(err)})
		return
	}
	c.JSON(http.StatusOK, stud)
}
func (h *Handler) UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	bytes, err := c.GetRawData()
	if err != nil {
		h.Logger.Error("c.GetRawData", zap.Error(err))
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	student := &models.InputUpdateStudent{}
	if err := student.Decode(bytes); err != nil {
		h.Logger.Error("student.UpdateStudent.Decode", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(student)
	validation := student.ValidateUpdateStudent(*student)
	if validation != nil {
		h.Logger.Error("student.ValidateUpdateStudent", zap.Error(validation))
		c.JSON(http.StatusInternalServerError, gin.H{"error": validation.Error()})
		return
	}
	err = h.SvcHandler.UpdateStudent(c, *student, id)
	if err != nil {
		h.Logger.Error("h.SvcHandler.UpdateStudent", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Student updated successfully"})
}
func (h *Handler) DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	err := h.SvcHandler.DeleteStudent(c, id)
	if err != nil {
		h.Logger.Error("h.SvcHandler.DeleteStudent", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Student deleted successfully"})
}
