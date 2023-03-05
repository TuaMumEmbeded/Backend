package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Playtime struct {
	StartTime time.Time `bson:"startTime,omitempty"`
    EndTime time.Time `bson:"endTime,omitempty"`
}
type Patient struct {
    Id       primitive.ObjectID `bson:"_id,omitempty"`
    Patient_id    int             `bson:"patient_id,omitempty" validate:"required"`
    Firstname     string             `bson:"firstname,omitempty" validate:"required"`
    Lastname     string             `bson:"lastname,omitempty" validate:"required"`
    Age     int             `bson:"age,omitempty" validate:"required"`
    Gender     string             `bson:"gender,omitempty" validate:"required"`
    Playtimes []Playtime `bson:"playtimes,omitempty"`
}