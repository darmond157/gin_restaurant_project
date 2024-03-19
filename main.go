package main

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"restaurant/database"
	"restaurant/middleware"
	"restaurant/routes"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(middleware.Authentication())

	routes.UserRoutes(router)
	routes.MenuRoutes(router)
	routes.FoodRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)

	router.Run(":" + port)

}
