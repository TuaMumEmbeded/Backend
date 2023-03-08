package main

import (
        "github.com/gin-contrib/cors"
        "github.com/gin-gonic/gin"
        "gin-mongo-api/configs"
        "gin-mongo-api/routes"
)

func main() {
        gin.SetMode(gin.ReleaseMode)
        router := gin.Default()

        router.Use(cors.Default())

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