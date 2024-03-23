package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"restaurant/models"
	"time"
)

func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		menuId := c.Param("menu_id")

		var menu *models.Food
		if err := c.BindJSON(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		if err := menuCollection.FindOne(ctx, bson.M{"menu_id": menuId}).Decode(menu); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, menu)
	}
}
func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		result, err := menuCollection.Find(context.TODO(), bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while listing the menu items"})
		}

		var allMenus []bson.M
		if err = result.All(ctx, &allMenus); err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, allMenus)

	}
}
func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var menu *models.Menu
		if err := c.BindJSON(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		menuId := c.Param("menu_id")
		var updateObj primitive.D

		if menu.StartDate != nil && menu.EndDate != nil {
			if !inTimeSpan(*menu.StartDate, *menu.EndDate, time.Now()) {
				msg := "retype the time"
				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
				defer cancel()
				return
			}
			updateObj = append(updateObj, bson.E{Key: "start_date", Value: menu.StartDate})
			updateObj = append(updateObj, bson.E{Key: "end_date", Value: menu.EndDate})

			if menu.Name != "" {
				updateObj = append(updateObj, bson.E{Key: "name", Value: menu.Name})
			}
			if menu.Category != "" {
				updateObj = append(updateObj, bson.E{Key: "category", Value: menu.Category})
			}
			menu.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			updateObj = append(updateObj, bson.E{Key: "updated_at", Value: menu.Updated_at})

			upsert := true
			opt := options.UpdateOptions{
				Upsert: &upsert,
			}
			result, err := menuCollection.UpdateOne(ctx, bson.M{"menu_id": menuId}, bson.D{
				{"$set", updateObj},
			}, &opt)

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "menu update failed"})

			}

			c.JSON(http.StatusOK, result)

		}

	}
}
func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var menu *models.Menu
		if err := c.BindJSON(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		if err := validate.Struct(menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		menu.ID = primitive.NewObjectID()
		menu.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.MenuId = menu.ID.Hex()

		insertResult, insertError := menuCollection.InsertOne(ctx, menu)
		if insertError != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "menu item was not created"})
			return
		}

		c.JSON(http.StatusOK, insertResult)
	}
}

func inTimeSpan(start , end , check time.Time) bool {
return start.After(time.Now()) && end.After(start)
}
