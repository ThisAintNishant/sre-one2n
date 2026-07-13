package main

import (
	"log"

	"github.com/ThisAintNishant/sre-one2n/internal/config"
	"github.com/ThisAintNishant/sre-one2n/internal/database"
	"github.com/ThisAintNishant/sre-one2n/internal/routes"

	"github.com/swaggo/files"
    "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	db, err := database.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := gin.Default()

	routes.Register(router, db)

	log.Printf("Server listening on :%s", cfg.Port)

	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}