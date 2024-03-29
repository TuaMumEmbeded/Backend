package controllers

import (
	"context"
	"gin-mongo-api/models"
	"gin-mongo-api/responses"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetSensor() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		patientId, err := strconv.Atoi(c.Param("patientId"))
		if err != nil {
			panic(err)
		}

		var sensor models.Sensor
		defer cancel()

		err = sensorCollection.FindOne(ctx, bson.M{"patient_id": patientId}).Decode(&sensor)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.DataResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"result": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.DataResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"result": sensor}})

	}
}
func UpdateSensor() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
		patientId, err := strconv.Atoi(c.Param("patientId"))

		if err != nil {
			panic(err)
		}

		var sensor models.Sensor
		var old_sensor models.Sensor
		sensorCollection.FindOne(ctx, bson.M{"patient_id": patientId}).Decode(&old_sensor)

		defer cancel()

		err = c.BindJSON(&sensor)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.DataResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    map[string]interface{}{"result": err.Error()}})
			return
		}

		var old_data models.Patient

		userCollection.FindOne(ctx, bson.M{"patient_id": patientId}).Decode(&old_data)

		new_data := bson.M{"emergency": sensor.Emergency, "bed": sensor.Bed, "restroom": sensor.Restroom, "hungry": sensor.Hungry, "game": sensor.Game}

		result, err := sensorCollection.UpdateOne(ctx, bson.M{"patient_id": patientId}, bson.M{"$set": new_data})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.DataResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"result": err.Error()}})
			return
		}

		var updatedData models.Sensor
		if result.MatchedCount == 1 {
			err := sensorCollection.FindOne(ctx, bson.M{"patient_id": patientId}).Decode(&updatedData)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.DataResponse{
					Status:  http.StatusInternalServerError,
					Message: "error",
					Data:    map[string]interface{}{"result": err.Error()}})
				return
			}
		}

		c.JSON(http.StatusOK, responses.DataResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"result": updatedData}})
	}
}
