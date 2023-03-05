package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Sensor struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	Patient_id int `bson:"patient_id,omitempty"`
	Emergency bool `bson:"emergency"`
	Bed bool `bson:"bed"`
	Restroom bool `bson:"restroom"`
	Hungry bool `bson:"hungry"`
	Game bool `bson:"game"`
}