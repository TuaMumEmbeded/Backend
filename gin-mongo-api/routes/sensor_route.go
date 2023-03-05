package routes


import (
    "github.com/gin-gonic/gin"
    "gin-mongo-api/controllers"
)

func SensorRoute(router *gin.Engine)  {
    //All routes related to users comes here
    router.GET("/sensor/patient/:patientId", controllers.GetSensor())
    router.PUT("/sensor/patient/:patientId", controllers.UpdateSensor())
}