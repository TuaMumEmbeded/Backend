package routes

import (
	"gin-mongo-api/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("/patient", controllers.CreatePatient())
	router.GET("/patient/:patientId", controllers.GetPatient())
	router.GET("/patients", controllers.GetAllPatients())
}
