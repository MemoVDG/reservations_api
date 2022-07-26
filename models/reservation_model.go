package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Reservation struct {
	Id        primitive.ObjectID `json:"id"`
	User      string             `json:"user_id" validate:"required"`
	Property  string             `json:"property_id" validate:"required"`
	StartDate string             `json:"start_date" validate:"required"`
	EndDate   string             `json:"end_date" validate:"required"`
}
