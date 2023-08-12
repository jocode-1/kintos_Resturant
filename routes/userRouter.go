package routes

import (
	"github.com/gin-gonic/gin"
	"resturant/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id", controllers.GetUser())
	incomingRoutes.POST("/users/signup", controllers.RegisterUser())
	incomingRoutes.POST("/users/login", controllers.Login())

}
