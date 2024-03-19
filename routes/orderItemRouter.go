package routes

import (
	"github.com/gin-gonic/gin"
	"restaurant/controllers"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItems", controllers.GetOrderItems())
	incomingRoutes.GET("/orderItems/:orderItemId", controllers.GetOrderItem())
	incomingRoutes.GET("/orderItems-order/:orderId", controllers.GetOrderItemsByOrder())
	incomingRoutes.POST("/orderItems", controllers.CreateOrderItem())
	incomingRoutes.PATCH("/orderItems", controllers.UpdateOrderItem())
}
