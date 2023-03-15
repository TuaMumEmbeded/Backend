package routes

import (
	"gin-mongo-api/controllers"

	"github.com/gin-gonic/gin"
)

func SensorRoute(router *gin.Engine) {
	router.GET("/sensor/patient/:patientId", controllers.GetSensor())
	router.PUT("/sensor/patient/:patientId", controllers.UpdateSensor())
}
