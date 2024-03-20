package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"restaurant/database"
	"restaurant/models"
	"time"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")
var menuCollection *mongo.Collection = database.OpenCollection(database.Client, "menu")
var validate = validator.New()

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		foodID := c.Param("food_id")
		var food models.Food

		err := foodCollection.FindOne(ctx, bson.M{"food_id": foodID}).Decode(&food)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while fetching the food item"})
		}

		c.JSON(http.StatusOK, food)
		defer cancel()
	}
}
func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		result, err := foodCollection.Find(context.TODO(), bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var allFoods []bson.M
		if err := result.All(ctx, &allFoods); err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, allFoods)
	}
}
func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var food *models.Food
		var menu *models.Menu

		if err := c.BindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := validate.Struct(food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := menuCollection.FindOne(ctx, bson.M{"menu_id": food.MenuId}).Decode(&menu); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "menu was not found"})
			return
		}

		food.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.ID = primitive.NewObjectID()
		food.FoodId = food.ID.Hex()
		num := toFixed(*food.Price, 2)
		food.Price = &num

		result, err := foodCollection.InsertOne(ctx, food)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "food item was not created"})
			return
		}

		c.JSON(http.StatusOK, result)
	}
}

func round(num float64) int {

}

func toFixed(num float64, precision int) float64 {

}
