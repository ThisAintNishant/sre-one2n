package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type HealthHandler struct {
	DB *pgxpool.Pool
}

func NewHealthHandler(db *pgxpool.Pool) *HealthHandler {
	return &HealthHandler{
		DB: db,
	}
}

func (h *HealthHandler) Ready(c *gin.Context) {
	if err := h.DB.Ping(context.Background()); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status":   "DOWN",
			"database": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   "UP",
		"database": "UP",
	})
}