package handlers

import (
	"errors"
	"net/http"

	"github.com/ThisAintNishant/sre-one2n/internal/models"
	"github.com/ThisAintNishant/sre-one2n/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
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

// GetStudents godoc
//
// @Summary Get all students
// @Description Returns all students
// @Tags Students
// @Produce json
// @Success 200 {array} models.Student
// @Failure 500 {object} map[string]string
// @Router /students [get]
func (h *StudentHandler) GetAll(c *gin.Context) {

	students, err := h.service.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, students)
}

// GetStudent godoc
//
// @Summary Get student by ID
// @Description Returns a single student by ID
// @Tags Students
// @Produce json
// @Param id path string true "Student ID"
// @Success 200 {object} models.Student
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /students/{id} [get]
func (h *StudentHandler) GetByID(c *gin.Context) {

	id := c.Param("id")

	student, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "student not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, student)
}

// UpdateStudent godoc
//
// @Summary Update Student
// @Description Update an existing student by ID
// @Tags Students
// @Accept json
// @Produce json
// @Param id path string true "Student ID"
// @Param student body models.Student true "Student"
// @Success 200 {object} models.Student
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /students/{id} [put]
func (h *StudentHandler) Update(c *gin.Context) {

	id := c.Param("id")

	var student models.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.service.Update(c.Request.Context(), id, &student); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "student not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, student)
}

// DeleteStudent godoc
//
// @Summary Delete Student
// @Description Delete a student by ID
// @Tags Students
// @Produce json
// @Param id path string true "Student ID"
// @Success 204 "No Content"
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /students/{id} [delete]
func (h *StudentHandler) Delete(c *gin.Context) {

	id := c.Param("id")

	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "student not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}