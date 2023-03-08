package controllers

import (
	"context"
	"gin-mongo-api/configs"
	"gin-mongo-api/models"
	"gin-mongo-api/responses"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "Info")
var sensorCollection *mongo.Collection = configs.GetCollection(configs.DB, "Sensor")
var validate = validator.New()

func CreatePatient() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var patient models.Patient
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&patient); err != nil {
			c.JSON(http.StatusBadRequest, responses.DataResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&patient); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.DataResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newPatient := models.Patient{
			Id:         primitive.NewObjectID(),
			Patient_id: patient.Patient_id,
			Firstname:  patient.Firstname,
			Lastname:   patient.Lastname,
			Age:        patient.Age,
			Gender:     patient.Gender,
			Playtimes:  []models.Playtime{},
		}

		newSensor := models.Sensor{
			Id:         primitive.NewObjectID(),
			Patient_id: patient.Patient_id,
			Emergency:  false,
			Bed:        false,
			Restroom:   false,
			Hungry:     false,
			Game:       false,
		}

		count, err := userCollection.CountDocuments(ctx, bson.M{"patient_id": newPatient.Patient_id})

		if err != nil {
			panic(err)
		}

		if count != 0 {
			c.JSON(http.StatusInternalServerError, responses.DataResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": "patient_id already exists"}})
			return
		}

		patientID, err := userCollection.InsertOne(ctx, newPatient)
		sensorID, err := sensorCollection.InsertOne(ctx, newSensor)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.DataResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.DataResponse{
			Status:  http.StatusCreated,
			Message: "success",
			Data:    map[string]interface{}{"patient_data": patientID, "sensor_data": sensorID}})
	}
}

func GetPatient() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		patientId, err := strconv.Atoi(c.Param("patientId"))
		if err != nil {
			panic(err)
		}

		var patient models.Patient
		defer cancel()

		err = userCollection.FindOne(ctx, bson.M{"patient_id": patientId}).Decode(&patient)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.DataResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.DataResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"data": patient}})
	}
}

// func EditAUser() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 			userId := c.Param("userId")
// 			var user models.User
// 			defer cancel()
// 			objId, _ := primitive.ObjectIDFromHex(userId)

// 			//validate the request body
// 			if err := c.BindJSON(&user); err != nil {
// 					c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
// 					return
// 			}

// 			//use the validator library to validate required fields
// 			if validationErr := validate.Struct(&user); validationErr != nil {
// 					c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
// 					return
// 			}

// 			update := bson.M{"name": user.Name, "location": user.Location, "title": user.Title}
// 			result, err := userCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
// 			if err != nil {
// 					c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
// 					return
// 			}

// 			//get updated user details
// 			var updatedUser models.User
// 			if result.MatchedCount == 1 {
// 					err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedUser)
// 					if err != nil {
// 							c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
// 							return
// 					}
// 			}

// 			c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedUser}})
// 	}
// }

// func DeleteAUser() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 			userId := c.Param("userId")
// 			defer cancel()

// 			objId, _ := primitive.ObjectIDFromHex(userId)

// 			result, err := userCollection.DeleteOne(ctx, bson.M{"id": objId})
// 			if err != nil {
// 					c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
// 					return
// 			}

// 			if result.DeletedCount < 1 {
// 					c.JSON(http.StatusNotFound,
// 							responses.UserResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "User with specified ID not found!"}},
// 					)
// 					return
// 			}

// 			c.JSON(http.StatusOK,
// 					responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "User successfully deleted!"}},
// 			)
// 	}
// }

func GetAllPatients() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var patients []models.Patient
		defer cancel()

		results, err := userCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.DataResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()}})
			return
		}

		defer results.Close(ctx)
		for results.Next(ctx) {
			var patient models.Patient
			if err = results.Decode(&patient); err != nil {
				c.JSON(http.StatusInternalServerError, responses.DataResponse{
					Status:  http.StatusInternalServerError,
					Message: "error",
					Data:    map[string]interface{}{"data": err.Error()}})
			}

			patients = append(patients, patient)
		}

		c.JSON(http.StatusOK, responses.DataResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"data": patients}})
	}
}
