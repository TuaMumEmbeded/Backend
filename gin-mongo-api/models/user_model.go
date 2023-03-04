package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Patient struct {
    Id       primitive.ObjectID `bson:"_id,omitempty"`
    Patient_id    int             `bson:"patient_id,omitempty" validate:"required"`
    Name     string             `bson:"name,omitempty" validate:"required"`
}