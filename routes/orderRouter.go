package routes

import (
	"resturant/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/order", controllers.GetOrders())
	incomingRoutes.GET("/order/:order_id", controllers.GetOrder())
	incomingRoutes.POST("/order", controllers.CreateOrder())
	incomingRoutes.POST("/order/:order_id", controllers.UpdateOrder())
}
