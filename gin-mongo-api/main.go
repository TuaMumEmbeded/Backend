package main

import (
	"gin-mongo-api/configs"
	"gin-mongo-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "Hello from Gin-gonic & mongoDB",
		})
	})
	// run database
	configs.ConnectDB()
	// routes
	routes.UserRoute(router)
	routes.SensorRoute(router)
	router.Run("localhost:6000")

}
