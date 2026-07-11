package routes

import (
	"net/http"

	"github.com/ThisAintNishant/sre-bootcamp/internal/handlers"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Register(router *gin.Engine, db *pgxpool.Pool) {
	health := handlers.NewHealthHandler(db)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to SRE Bootcamp API",
		})
	})

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	
	router.GET("/health/ready", health.Ready)
}