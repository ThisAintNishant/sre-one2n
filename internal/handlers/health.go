package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type HealthHandler struct {
	DB *pgxpool.Pool
	StartedAt time.Time
}

func NewHealthHandler(db *pgxpool.Pool) *HealthHandler {
	return &HealthHandler{
		DB: db,
	}
}

func (h *HealthHandler) Ready(c *gin.Context) {

	err := h.DB.Ping(context.Background())

	if err != nil {

		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "DOWN",
			"database": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   "UP",
		"database": "UP",
		"version":  version.Version,
		"uptime":   time.Since(h.StartedAt).Round(time.Second).String(),
	})

}