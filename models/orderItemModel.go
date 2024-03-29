package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type OrderItem struct {
	ID          primitive.ObjectID `bson:"_id"`
	Quantity    *string            `json:"quantity" validate:"required,eq=S|eq=M|eq=L"`
	UnitPrice   *float64           `json:"unit_price" validate:"required"`
	FoodID      *string            `json:"food_id" validate:"required"`
	OrderId     string             `json:"order_id" validate:"required"`
	OrderItemId string             `json:"order_item_id"`
	Created_at  time.Time          `json:"created_at"`
	Updated_at  time.Time          `json:"updated_at"`
}
