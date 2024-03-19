package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Order struct {
	ID         primitive.ObjectID `bson:"_id"`
	OrderId    string             `json:"order_id"`
	TableId    *string            `json:"table_id" validate:"required"`
	OrderDate  time.Time          `json:"order_date" validate:"required"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
}
