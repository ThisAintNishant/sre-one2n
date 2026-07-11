package main

import (
	"log"

	"github.com/ThisAintNishant/sre-bootcamp/internal/config"
	"github.com/ThisAintNishant/sre-bootcamp/internal/database"
	"github.com/ThisAintNishant/sre-bootcamp/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	startedAt := time.Now()
	cfg := config.Load()

	db, err := database.New(cfg)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	router := gin.Default()

	routes.Register(router, db)
	logger := logger.New()
	logger.Info("Server listening on :%s", cfg.Port)
	// log.Printf("Server listening on :%s", cfg.Port)

	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}