package routes

import (
	"resturant/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/orderItems", controllers.GetOrderItems())
	incomingRoutes.GET("/orderItems/:orderItem_id", controllers.GetOrderItem())
	incomingRoutes.GET("/orderItems-order/:order_id", controllers.GetOrderItemsByOrder())
	incomingRoutes.POST("/orderItems", controllers.CreateOrderItem())
	incomingRoutes.POST("/orderItems/:orderItem_id", controllers.UpdateOrderItem())
}
