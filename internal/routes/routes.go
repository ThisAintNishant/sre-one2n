package routes

import (
	"net/http"

	"github.com/ThisAintNishant/sre-one2n/internal/handlers"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/prometheus/client_golang/prometheus/promhttp"

    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/ThisAintNishant/sre-one2n/internal/repository"
	"github.com/ThisAintNishant/sre-one2n/internal/service"
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

	router.GET("/swagger/*any",
    ginSwagger.WrapHandler(swaggerFiles.Handler))

	repo := repository.NewPostgresStudentRepository(db)
	svc := service.NewStudentService(repo)
	studentHandler := handlers.NewStudentHandler(svc)
	
	router.POST("/students", studentHandler.Create)
	router.GET("/students", studentHandler.GetAll)
	router.GET("/students/:id", studentHandler.GetByID)
	router.PUT("/students/:id", studentHandler.Update)
	router.DELETE("/students/:id", studentHandler.Delete)
}