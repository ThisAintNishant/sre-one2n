package routes

import "github.com/gin-gonic/gin"

func Register(router *gin.Engine) {

	router.GET("/", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "Welcome to SRE Bootcamp API",
		})

	})

}