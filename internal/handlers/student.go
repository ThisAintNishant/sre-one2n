package handlers

import (
	"net/http"

	"github.com/ThisAintNishant/sre-one2n/internal/models"
	"github.com/ThisAintNishant/sre-one2n/internal/service"

	"github.com/gin-gonic/gin"
)

type StudentHandler struct {
	service *service.StudentService
}

func NewStudentHandler(service *service.StudentService) *StudentHandler {
	return &StudentHandler{
		service: service,
	}
}

// CreateStudent godoc
//
// @Summary Create Student
// @Description Create a new student
// @Tags Students
// @Accept json
// @Produce json
// @Param student body models.Student true "Student"
// @Success 201 {object} models.Student
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /students [post]
func (h *StudentHandler) Create(c *gin.Context) {

	var student models.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.service.Create(c.Request.Context(), &student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, student)
}