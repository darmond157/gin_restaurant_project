package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Table struct {
	ID             primitive.ObjectID `bson:"_id"`
	NumberOfGuests *int               `json:"number_of_guests" validate:"required"`
	TableNumber    *int               `json:"table_number" validate:"required"`
	TableId        string             `json:"table_id"`
	Created_at     time.Time          `json:"created_at"`
	Updated_at     time.Time          `json:"updated_at"`
}
