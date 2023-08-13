package routes

import (
	"resturant/controllers"

	"github.com/gin-gonic/gin"
)

func MenuRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/orders", controllers.GetOrders())
	incomingRoutes.GET("/orders/:order_id", controllers.GetOrder())
	incomingRoutes.POST("/orders", controllers.CreateOrder())
	incomingRoutes.GET("/orders/:order_id", controllers.UpdateOrder())
}
