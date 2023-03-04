package main

import (
        "github.com/gin-gonic/gin"
        "gin-mongo-api/configs"
        "gin-mongo-api/routes"
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
        router.Run("localhost:6000") 
}