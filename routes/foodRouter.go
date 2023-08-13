package routes

import (
	"github.com/gin-gonic/gin"
	"resturant/controllers"
)

func FoodRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/foods", controllers.GetFoods())
	incomingRoutes.GET("/food/:food_id", controllers.GetFood())
	incomingRoutes.POST("/food", controllers.CreateFood())
	incomingRoutes.PATCH("/food/:food_id", controllers.UpdateFood())

}
