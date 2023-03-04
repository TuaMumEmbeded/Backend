package routes


import (
    "github.com/gin-gonic/gin"
    "gin-mongo-api/controllers"
)

func UserRoute(router *gin.Engine)  {
    //All routes related to users comes here
    // router.POST("/user", controllers.CreateUser())
    router.GET("/patient/:patientId", controllers.GetPatient())
    // router.PUT("/user/:userId", controllers.EditAUser())
    // router.DELETE("/user/:userId", controllers.DeleteAUser())
    router.GET("/patients", controllers.GetAllPatients())
}